package types

type VirtualSvc struct {
	Id   string          `yaml:"id,omitempty"`
	Name string          `yaml:"name,omitempty"`
	Opts *VirtualSvcOpts `yaml:"opts,omitempty"`
}

type VirtualSvcOpts struct {
	Auth     *AuthOpts     `yaml:"auth,omitempty"`
	Mastodon *MastodonOpts `yaml:"mastodon,omitempty"`
	Sqlx     *SqlxOpts     `yaml:"sqlx,omitempty"`
}

type AuthOpts struct {
	PrivateKey string `yaml:"privateKey,omitempty"`
	Salt       string `yaml:"salt,omitempty"`
}

type MastodonOpts struct {
	Host         string `yaml:"host,omitempty"`
	ClientId     string `yaml:"clientId,omitempty"`
	ClientSecret string `yaml:"clientSecret,omitempty"`
}

type SqlxOpts struct {
	Log        bool     `yaml:"log,omitempty"`
	Driver     string   `yaml:"driver,omitempty"`
	Connection string   `yaml:"connection,omitempty"`
	Types      []string `yaml:"types,omitempty"`
}
