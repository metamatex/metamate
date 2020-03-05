package types

import "net/http"

type Route struct {
	Methods     []string
	Path        string
	HandlerFunc http.HandlerFunc
	Handler     http.Handler
}
