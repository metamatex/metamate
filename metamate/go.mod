module github.com/metamatex/metamate/metamate

go 1.13

replace github.com/metamatex/metamate/asg => ../asg

replace github.com/metamatex/metamate/mastodon-svc => ../mastodon-svc

replace github.com/metamatex/metamate/reddit-svc => ../reddit-svc

replace github.com/metamatex/metamate/hackernews-svc => ../hackernews-svc

replace github.com/metamatex/metamate/kubernetes-svc => ../kubernetes-svc

replace github.com/metamatex/metamate/generic => ../generic

replace github.com/metamatex/metamate/spec => ../spec

replace github.com/metamatex/metamate/gen => ../gen

require (
	github.com/asaskevich/govalidator v0.0.0-20200108200545-475eaeb16496 // indirect
	github.com/blang/semver v3.5.1+incompatible
	github.com/davecgh/go-spew v1.1.1
	github.com/go-chi/chi v4.1.1+incompatible
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/golang/groupcache v0.0.0-20190129154638-5b532d6fd5ef
	github.com/google/uuid v1.1.1
	github.com/graphql-go/graphql v0.7.9
	github.com/graphql-go/handler v0.2.3
	github.com/hashicorp/golang-lru v0.5.4
	github.com/metamatex/metamate/asg v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/gen v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/generic v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/hackernews-svc v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/kubernetes-svc v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/mastodon-svc v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/reddit-svc v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/spec v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.5.0
	github.com/rhysd/go-github-selfupdate v1.2.1
	github.com/rs/cors v1.7.0
	github.com/spf13/afero v1.1.2
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.5.1
	golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073
	gopkg.in/yaml.v2 v2.3.0
)
