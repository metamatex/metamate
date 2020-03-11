module github.com/metamatex/metamate/sqlx-svc

go 1.13

replace github.com/metamatex/metamate/asg => ../asg

replace github.com/metamatex/metamate/gen => ../gen

replace github.com/metamatex/metamate/generic => ../generic

replace github.com/metamatex/metamate/spec => ../spec

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/google/uuid v1.1.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/metamatex/metamate/asg v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/gen v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/generic v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/spec v0.0.0-00010101000000-000000000000
	google.golang.org/appengine v1.6.5 // indirect
)
