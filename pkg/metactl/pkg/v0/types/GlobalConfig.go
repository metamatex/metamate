package types

type GlobalConfig struct {
	V0 V0Global
}

type V0Global struct {
	Index          Index          `json:"index,omitempty",yaml:"index,omitempty"`
	Buses          []NamedBus     `json:"buses,omitempty",yaml:"buses,omitempty"`
	Users          []NamedUser    `json:"users,omitempty",yaml:"users,omitempty"`
	Contexts       []NamedContext `json:"contexts,omitempty",yaml:"contexts,omitempty"`
	CurrentContext string         `json:"currentContext,omitempty",yaml:"currentContext,omitempty"`
}

type Index struct {
	Path         string `json:"path,omitempty",yaml:"path,omitempty"`
	NGramMinimum int    `json:"nGramMinimum,omitempty",yaml:"nGramMinimum,omitempty"`
	NGramMaximum int    `json:"nGramMinimum,omitempty",yaml:"nGramMinimum,omitempty"`
}

type NamedBus struct {
	Name string `json:"name,omitempty",yaml:"name,omitempty"`
	Bus  Bus    `json:"bus,omitempty",yaml:"bus,omitempty"`
}

type Bus struct {
	CertificateAuthorityData string `json:"certificateAuthorityData,omitempty",yaml:"certificateAuthorityData,omitempty"`
	CertificateAuthority     string `json:"certificateAuthority,omitempty",yaml:"certificateAuthority,omitempty"`
	Addr                     string `json:"addr,omitempty",yaml:"addr,omitempty"`
}

type NamedContext struct {
	Name    string  `json:"name,omitempty",yaml:"name,omitempty"`
	Context Context `json:"context,omitempty",yaml:"context,omitempty"`
}

type Context struct {
	Bus  string `json:"bus,omitempty",yaml:"bus,omitempty"`
	User string `json:"user,omitempty",yaml:"user,omitempty"`
}

type NamedUser struct {
	Name string `json:"name,omitempty",yaml:"name,omitempty"`
	User User   `json:"user,omitempty",yaml:"user,omitempty"`
}

type User struct {
	ClientCertificate     string `json:"clientCertificate,omitempty",yaml:"clientCertificate,omitempty"`
	ClientCertificateData string `json:"clientCertificateData,omitempty",yaml:"clientCertificateData,omitempty"`
	ClientToken           string `json:"clientToken,omitempty",yaml:"clientToken,omitempty"`
	ClientTokenData       string `json:"clientTokenData,omitempty",yaml:"clientTokenData,omitempty"`
}
