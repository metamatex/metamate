module github.com/metamatex/metamate/metamate

go 1.13

replace github.com/metamatex/metamate/asg => ../asg

replace github.com/metamatex/metamate/auth-svc => ../auth-svc

replace github.com/metamatex/metamate/mastodon-svc => ../mastodon-svc

replace github.com/metamatex/metamate/sqlx-svc => ../sqlx-svc

replace github.com/metamatex/metamate/hackernews-svc => ../hackernews-svc

replace github.com/metamatex/metamate/kubernetes-svc => ../kubernetes-svc

replace github.com/metamatex/metamate/generic => ../generic

replace github.com/metamatex/metamate/spec => ../spec

replace github.com/metamatex/metamate/gen => ../gen

require (
	github.com/asaskevich/govalidator v0.0.0-20200108200545-475eaeb16496 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/go-chi/chi v4.0.3+incompatible
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/google/uuid v1.1.1
	github.com/graphql-go/graphql v0.7.9
	github.com/graphql-go/handler v0.2.3
	github.com/metamatex/metamate/asg v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/auth-svc v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/gen v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/generic v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/hackernews-svc v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/kubernetes-svc v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/mastodon-svc v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/spec v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/sqlx-svc v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.5.0
	github.com/prometheus/common v0.9.1
	github.com/rs/cors v1.7.0
	github.com/spf13/afero v1.1.2
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.6.2
	gopkg.in/yaml.v2 v2.2.8
)
