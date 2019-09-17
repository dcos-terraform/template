package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/google/go-github/github"
	"github.com/hashicorp/hcl/hcl/scanner"
	"github.com/iancoleman/strcase"
	"golang.org/x/oauth2"
)

func parse(provider string) {
	var structName string
	var moduleName string
	var githubOwner string
	var contentstrs []string

	switch provider {
	case "aws":
		structName = "AWS"
	case "azurerm":
		structName = "AzureRM"
	case "gcp":
		structName = "GCP"
	}

	moduleName = "terraform-" + provider + "-dcos"
	githubOwner = "dcos-terraform"

	githubToken, ok := os.LookupEnv("GITHUB_TOKEN")
	if ok != true {
		fmt.Println("Make sure you set GITHUB_TOKEN as variable, to not force the low GitHub API Limit of 60.")
		os.Exit(1)
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// grab GitHub dirContent from latest release
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), githubOwner, moduleName)
	fmt.Printf("Found latest release %v for %v module\n", release.GetTagName(), moduleName)
	_, dirContent, _, err := client.Repositories.GetContents(context.Background(), githubOwner, moduleName, "/", &github.RepositoryContentGetOptions{Ref: release.GetTagName()})
	if err != nil {
		fmt.Println(err)
		return
	}

	// only regex check "variables" files
	r, _ := regexp.Compile("variables")
	for _, file := range dirContent {
		if r.MatchString(*file.Name) {
			raw, _, err := client.Git.GetBlobRaw(context.Background(), githubOwner, moduleName, file.GetSHA())
			if err != nil {
				fmt.Println(err)
				return
			}
			// hcl scan all the things
			var fieldname string
			var fieldtype string
			var jsonfieldname string
			var hclfieldname string
			var hcle string
			hclScanner := scanner.New(raw)
			for {
				out := hclScanner.Scan()
				if out.Text == "variable" {
					out = hclScanner.Scan()
					fieldname = strcase.ToCamel(out.Text[1 : len(out.Text)-1])
					fieldtype = strings.ToLower(out.Type.String())
					jsonfieldname = strings.ToLower(out.Text[1 : len(out.Text)-1])
					hclfieldname = (out.Text[1 : len(out.Text)-1])
					for {
						out = hclScanner.Scan()
						if out.Text == "default" {
							out = hclScanner.Scan()
							jsonfieldname = jsonfieldname + ",omitempty"
							hcle = " hcle:\"omitempty\""
						}
						if out.Text == "}" {
							break
						}
					}
					contentstrs = append(contentstrs, fmt.Sprintf("%v %v `hcl:\"%v\"%v json:\"%v\"`", fieldname, fieldtype, hclfieldname, hcle, jsonfieldname))
					fieldname = ""
					fieldtype = ""
					jsonfieldname = ""
					hclfieldname = ""
					hcle = ""
				}
				if out.Text == "" {
					break
				}
			}
		}
	}

	// Write the file
	prefixstr := []byte("package template\n\n\ntype " + structName + " struct {\n")
	suffixstr := []byte("\n}")
	err = ioutil.WriteFile("template/"+moduleName+".go", prefixstr, 0644)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	file, err := os.OpenFile("template/"+moduleName+".go", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	// ModuleSource and ModuleVersion
	var sourceandversion []string
	sourceandversion = append(sourceandversion, fmt.Sprintf("ModuleSource string `hcl:\"source\" default:\"dcos-terraform/dcos/%v\" json:\"source\"`\n", provider))
	sourceandversion = append(sourceandversion, fmt.Sprintf("ModuleVersion string `hcl:\"version\" default:\"~> %v\" json:\"version\"`\n", release.GetTagName()))
	_, err = file.WriteString(strings.Join(sourceandversion, ""))
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	// All extracted variables
	_, err = file.WriteString(strings.Join(contentstrs, "\n"))
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	// Closing curly bracket
	_, err = file.Write(suffixstr)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
}

func main() {
	parse(os.Args[1])
}
