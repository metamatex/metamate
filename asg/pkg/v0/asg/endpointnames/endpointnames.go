package endpointnames

import (
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/utils"
)

const (
	AuthenticateClientAccount = "AuthenticateClientAccount"
	VerifyToken               = "VerifyToken"
	LookupService             = "LookupService"
)

func Get(s string) (string) {
	return "Get" + utils.Plural(s)
}

func Pipe(s string) (string) {
	return "Pipe" + utils.Plural(s)
}

func Post(s string) (string) {
	return "Post" + utils.Plural(s)
}

func Put(s string) (string) {
	return "Put" + utils.Plural(s)
}

func Delete(s string) (string) {
	return "Delete" + utils.Plural(s)
}
