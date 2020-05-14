package _go

import (
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/metamatex/metamate/metactl/pkg/v0/utils/ptr"
)

const (
	TaskVersion = "TaskVersion"
)

func init() {
	tasks[TaskVersion] = types.RenderTask{
		TemplateData: &versionTpl,
		Out:          ptr.String("version_.go"),
	}
}

var versionTpl = `package mql

const Version = "{{ .Version.Version }}"
`