module github.com/metamatex/metamate/spec

go 1.13

replace github.com/metamatex/metamate/gen => ../gen

replace github.com/metamatex/metamate/generic => ../generic

replace github.com/metamatex/metamate/asg => ../asg

require (
	github.com/metamatex/metamate/asg v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/gen v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/generic v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.5.1
)
