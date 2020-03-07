module github.com/metamatex/metamatemono/spec

go 1.13

replace github.com/metamatex/metamatemono/gen => ../gen

replace github.com/metamatex/metamatemono/generic => ../generic

replace github.com/metamatex/metamatemono/asg => ../asg

require (
	github.com/metamatex/metamatemono/asg v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamatemono/gen v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamatemono/generic v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.5.1
)
