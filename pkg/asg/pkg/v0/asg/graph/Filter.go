package graph

type Filter struct {
	Names *NamesSubset `yaml:"names,omitempty" json:"names,omitempty"`
	Flags *FlagsSubset `yaml:"flags,omitempty" json:"flags,omitempty"`
}

type FlagsSubset struct {
	Or  []string `yaml:"or,omitempty" json:"or,omitempty"`
	Nor []string `yaml:"nor,omitempty" json:"nor,omitempty"`
	And []string `yaml:"and,omitempty" json:"and,omitempty"`
}

type NamesSubset struct {
	ContainsOr []string `yaml:"containsOr,omitempty" json:"containsOr,omitempty"`
	Or         []string `yaml:"or,omitempty" json:"or,omitempty"`
	Nor        []string `yaml:"nor,omitempty" json:"nor,omitempty"`
}
