module github.com/metamatex/metamatemono/sqlx-svc

go 1.13

replace github.com/metamatex/metamatemono/asg => ../asg

replace github.com/metamatex/metamatemono/gen => ../gen

replace github.com/metamatex/metamatemono/generic => ../generic

replace github.com/metamatex/metamatemono/spec => ../spec

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/google/uuid v1.1.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/metamatex/metamatemono/asg v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamatemono/gen v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamatemono/generic v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamatemono/spec v0.0.0-00010101000000-000000000000
	google.golang.org/appengine v1.6.5 // indirect
)
