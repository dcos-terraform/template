package template

type Provider struct {
	Name       string            `hcl:",key" hcle:"omitempty" json:"name"`
	Version    string            `hcl:"version" hcle:"omitempty" json:"version,omitempty"`
	Region     string            `hcl:"region" hcle:"omitempty" json:"region,omitempty"`
	Alias      string            `hcl:"alias" hcle:"omitempty" json:"alias,omitempty"`
	AssumeRole map[string]string `hcl:"assume_role" hcle:"omitempty" json:"assume_role,omitempty"`
}
