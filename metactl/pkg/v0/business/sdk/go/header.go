package _go

import (
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/metamatex/metamate/metactl/pkg/v0/utils/ptr"
)

const (
	TaskHeader = "TaskHeader"
)

func init() {
	tasks[TaskHeader] = types.RenderTask{
		TemplateData: &goHeaderTpl,
		Out:          ptr.String("header_.go"),
	}
}

var goHeaderTpl = `package mql
const (
	AsgTypeHeader = "X-Asg-type"
	ContentTypeJson = "application/json; charset=utf-8"
	ContentTypeHeader = "Content-type"
	AuthorizationHeader = "Authorization"
)
`
