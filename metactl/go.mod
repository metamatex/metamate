module github.com/metamatex/metamatemono/metactl

go 1.13

replace github.com/metamatex/metamatemono/asg => ../asg

replace github.com/metamatex/metamatemono/gen => ../gen

replace github.com/metamatex/metamatemono/generic => ../generic

require (
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible
	github.com/blang/semver v3.5.1+incompatible
	github.com/davecgh/go-spew v1.1.1
	github.com/google/uuid v1.1.1 // indirect
	github.com/huandu/xstrings v1.3.0 // indirect
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/jinzhu/inflection v1.0.0
	github.com/metamatex/metamatemono/asg v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamatemono/gen v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamatemono/generic v0.0.0-00010101000000-000000000000
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/olekukonko/tablewriter v0.0.4
	github.com/pkg/errors v0.9.1
	github.com/rhysd/go-github-selfupdate v1.2.1
	github.com/spf13/afero v1.2.2
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.6.2
	gopkg.in/yaml.v2 v2.2.8
)
