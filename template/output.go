package template

type Output struct {
	Name  string `hcl:",key"`
	Value string `hcl:"value"`
}
