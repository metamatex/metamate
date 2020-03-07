module github.com/metamatex/metamatemono/metamate

go 1.13

replace github.com/metamatex/metamatemono/asg => ../asg

replace github.com/metamatex/metamatemono/auth-svc => ../auth-svc

replace github.com/metamatex/metamatemono/mastodon-svc => ../mastodon-svc

replace github.com/metamatex/metamatemono/sqlx-svc => ../sqlx-svc

replace github.com/metamatex/metamatemono/gen => ../gen

replace github.com/metamatex/metamatemono/generic => ../generic

replace github.com/metamatex/metamatemono/spec => ../spec

require (
	github.com/asaskevich/govalidator v0.0.0-20200108200545-475eaeb16496 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/go-chi/chi v4.0.3+incompatible
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/google/uuid v1.1.1
	github.com/graphql-go/graphql v0.7.9
	github.com/graphql-go/handler v0.2.3
	github.com/metamatex/metamatemono/asg v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamatemono/auth-svc v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamatemono/gen v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamatemono/generic v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamatemono/mastodon-svc v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamatemono/spec v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamatemono/sqlx-svc v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.5.0
	github.com/prometheus/common v0.9.1
	github.com/rs/cors v1.7.0
	gopkg.in/yaml.v2 v2.2.8
)
