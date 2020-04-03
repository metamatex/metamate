package types

type VirtualSvc struct {
	Id   string          `yaml:"id,omitempty"`
	Name string          `yaml:"name,omitempty"`
	Opts *VirtualSvcOpts `yaml:"opts,omitempty"`
}

type VirtualSvcOpts struct {
	Mastodon *MastodonOpts `yaml:"mastodon,omitempty"`
}

type MastodonOpts struct {
	Host         string `yaml:"host,omitempty"`
	ClientId     string `yaml:"clientId,omitempty"`
	ClientSecret string `yaml:"clientSecret,omitempty"`
}
