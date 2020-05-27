module github.com/metamatex/metamate/reddit-svc

go 1.13

replace github.com/metamatex/metamate/gen => ../gen

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/metamatex/metamate/gen v0.0.0-00010101000000-000000000000
	github.com/thecsw/mira v4.0.0+incompatible
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
