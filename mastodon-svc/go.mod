module github.com/metamatex/metamate/mastodon-svc

go 1.13

replace github.com/metamatex/metamate/gen => ../gen

require (
	github.com/mattn/go-mastodon v0.0.4
	github.com/metamatex/metamate/gen v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
