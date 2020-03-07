package types

type AuthOpts struct {
	PrivateKey string
	Salt       string
}

type MastodonOpts struct {
	Host         string
	ClientId     string
	ClientSecret string
}

type SqlxOpts struct {
	Log        bool
	Driver     string
	Connection string
	Types      []string
}

type VirtualSvcOpts struct {
	Name     string
	Auth     *AuthOpts
	Mastodon *MastodonOpts
	Sqlx     *SqlxOpts
}
