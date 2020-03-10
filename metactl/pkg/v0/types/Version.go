package types

type Version struct {
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
	Commit  string `json:"commit,omitempty" yaml:"commit,omitempty"`
	Date    string `json:"date,omitempty" yaml:"date,omitempty"`
}
