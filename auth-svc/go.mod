module github.com/metamatex/metamatemono/auth-svc

go 1.13

replace github.com/metamatex/metamatemono/gen => ../gen

require (
	github.com/metamatex/metamatemono/gen v0.0.0-00010101000000-000000000000
	github.com/square/go-jose v2.4.1+incompatible
	github.com/stretchr/testify v1.5.1 // indirect
	golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073 // indirect
	gopkg.in/square/go-jose.v2 v2.4.1 // indirect
)
