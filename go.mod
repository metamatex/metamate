module github.com/metamatex/metamatemono

go 1.13

replace github.com/metamatex/metamatemono/asg => ./asg

replace github.com/metamatex/metamatemono/auth-svc => ./auth-svc

replace github.com/metamatex/metamatemono/gen => ./gen

replace github.com/metamatex/metamatemono/generic => ./generic

replace github.com/metamatex/metamatemono/mastodon-svc => ./mastodon-svc

replace github.com/metamatex/metamatemono/metactl => ./metactl

replace github.com/metamatex/metamatemono/metamate => ./metamate

replace github.com/metamatex/metamatemono/spec => ./spec

replace github.com/metamatex/metamatemono/sqlx-svc => ./sqlx-svc

require (
	github.com/metamatex/metamatemono/metactl v0.0.0-00010101000000-000000000000 // indirect
	github.com/metamatex/metamatemono/metamate v0.0.0-00010101000000-000000000000 // indirect
)
