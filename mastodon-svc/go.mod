module github.com/metamatex/metamatemono/mastodon-svc

go 1.13

replace github.com/metamatex/metamatemono/gen => ../gen

require (
	github.com/mattn/go-mastodon v0.0.4
	github.com/metamatex/metamatemono/gen v0.0.0-00010101000000-000000000000
)
