module github.com/metamatex/metamate

go 1.13

replace github.com/metamatex/metamate/asg => ./asg

replace github.com/metamatex/metamate/auth-svc => ./auth-svc

replace github.com/metamatex/metamate/gen => ./gen

replace github.com/metamatex/metamate/generic => ./generic

replace github.com/metamatex/metamate/mastodon-svc => ./mastodon-svc

replace github.com/metamatex/metamate/kubernetes-svc => ./kubernetes-svc

replace github.com/metamatex/metamate/metactl => ./metactl

replace github.com/metamatex/metamate/metamate => ./metamate

replace github.com/metamatex/metamate/spec => ./spec

replace github.com/metamatex/metamate/sqlx-svc => ./sqlx-svc

require (
	github.com/metamatex/metamate/metactl v0.0.0-00010101000000-000000000000 // indirect
	github.com/metamatex/metamate/metamate v0.0.0-00010101000000-000000000000 // indirect
)
