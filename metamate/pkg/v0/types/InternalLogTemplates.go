package types

import (
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"text/template"
)

type InternalLogData struct {
	Subject generic.Generic
	Ctx     ReqCtx
}

type InternalLogTemplates map[string]map[string]*template.Template
