package graph

type PathNodeData struct {
	Cardinality string `yaml:",omitempty" json:"cardinality,omitempty"`
	Verb        string `yaml:",omitempty" json:"verb,omitempty"`
	IsActive    bool   `yaml:",omitempty" json:"is_active,omitempty"`
}
