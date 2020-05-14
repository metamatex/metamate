package _go

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/metamatex/metamate/metactl/pkg/v0/utils/ptr"
)

const (
	TaskClientInterface  = "TaskClientInterface"
	TaskServiceInterface = "TaskServiceInterface"
	TaskTypes            = "TaskTypes"
	TaskEnums            = "TaskEnums"
)

func init() {
	tasks[TaskClientInterface] = types.RenderTask{
		TemplateData: &goClientInterfaceTpl,
		Out:          ptr.String("client_.go"),
	}

	tasks[TaskServiceInterface] = types.RenderTask{
		Name:         ptr.String(TaskServiceInterface),
		TemplateData: &goServiceInterfaceTpl,
		Out:          ptr.String("{{ index .Data \"name\" }}_service_.go"),
	}

	tasks[TaskTypes] = types.RenderTask{
		TemplateData: &goTypesTpl,
		Out:          ptr.String("{{ .Type.Name }}_.go"),
		Iterate:      ptr.String(graph.TYPE),
	}

	tasks[TaskEnums] = types.RenderTask{
		TemplateData: &goEnumsTpl,
		Out:          ptr.String("{{ .Enum.Name }}_.go"),
		Iterate:      ptr.String(graph.ENUM),
	}
}

var goClientInterfaceTpl = `package mql

import (
    "context"
)

type Client interface {
{{- range $i, $endpoint := .Endpoints.Slice.Sort }}
	{{ $endpoint.Name }}(context.Context, {{ $endpoint.Edges.Type.Request.Name }}) (*{{ $endpoint.Edges.Type.Response.Name }}, error)
{{- end }}
}`

var goServiceInterfaceTpl = `package mql

import (
    "context"
)

type {{ index .Data "name" | title }}Service interface {
	Name() (string)
{{- range $i, $endpoint := .Endpoints.Slice.Sort }}
	Get{{ $endpoint.Name }}Endpoint() ({{ $endpoint.Name }}Endpoint)
    {{ $endpoint.Name }}(ctx context.Context, req {{ $endpoint.Edges.Type.Request.Name }}) (rsp {{ $endpoint.Edges.Type.Response.Name }})
{{- end }}
}`

var goTypesTpl = `package mql
{{ define "fields" }}
{{- range $fi, $field := . }}
{{- if $field.IsBool }}
    {{ camel $field.Name }} *bool ` + "`" + `json:"{{ $field.Name }},omitempty" yaml:"{{ $field.Name }},omitempty"` + "`" + `
{{- end }}
{{- if $field.IsFloat64 }}
    {{ camel $field.Name }} *float64 ` + "`" + `json:"{{ $field.Name }},omitempty" yaml:"{{ $field.Name }},omitempty"` + "`" + `
{{- end }}
{{- if $field.IsString }}
    {{ camel $field.Name }} *string ` + "`" + `json:"{{ $field.Name }},omitempty" yaml:"{{ $field.Name }},omitempty"{{- if $field.Flags.Is "hash" false -}},hash:"ignore"{{- end -}}` + "`" + `
{{- end }}
{{- if $field.IsInt32 }}
    {{ camel $field.Name }} *int32 ` + "`" + `json:"{{ $field.Name }},omitempty" yaml:"{{ $field.Name }},omitempty"` + "`" + `
{{- end }}
{{- if $field.IsType }}
    {{ camel $field.Name }} *{{ $field.Edges.Type.Holds.Name }} ` + "`" + `json:"{{ $field.Name }},omitempty" yaml:"{{ $field.Name }},omitempty"` + "`" + `
{{- end }}
{{- if $field.IsEnum }}
    {{ camel $field.Name }} *string ` + "`" + `json:"{{ $field.Name }},omitempty" yaml:"{{ $field.Name }},omitempty"` + "`" + `
{{- end }}
{{- if $field.IsStringList }}
    {{ camel $field.Name }} []string ` + "`" + `json:"{{ $field.Name }},omitempty" yaml:"{{ $field.Name }},omitempty"` + "`" + `
{{- end }}
{{- if $field.IsInt32List }}
    {{ camel $field.Name }} []int32 ` + "`" + `json:"{{ $field.Name }},omitempty" yaml:"{{ $field.Name }},omitempty"` + "`" + `
{{- end }}
{{- if $field.IsFloat64List }}
    {{ camel $field.Name }} []float64 ` + "`" + `json:"{{ $field.Name }},omitempty" yaml:"{{ $field.Name }},omitempty"` + "`" + `
{{- end }}
{{- if $field.IsBoolList }}
    {{ camel $field.Name }} []bool ` + "`" + `json:"{{ $field.Name }},omitempty" yaml:"{{ $field.Name }},omitempty"` + "`" + `
{{- end }}
{{- if $field.IsTypeList }}
    {{ camel $field.Name }} []{{ $field.Edges.Type.Holds.Name }} ` + "`" + `json:"{{ $field.Name }},omitempty" yaml:"{{ $field.Name }},omitempty"` + "`" + `
{{- end }}
{{- if $field.IsEnumList }}
    {{ camel $field.Name }} []string ` + "`" + `json:"{{ $field.Name }},omitempty" yaml:"{{ $field.Name }},omitempty"` + "`" + `
{{- end }}
{{- end }}
{{- end }}
const (
	{{ .Type.Name }}Name = "{{ .Type.Name }}"
)

type {{ .Type.Name }} struct {
{{- template "fields" .Type.Edges.Fields.Holds.Slice.Sort }}
}`

var goEnumsTpl = `package mql
{{ $enum := .Enum }}
const (
	{{ $enum.Name }}EnumName = "{{ $enum.Name }}"
)

var {{ $enum.Name }} = struct{
{{- range $vi, $value := sortAlpha $enum.Data.Values }}
    {{ camel $value }} string
{{- end }}
}{
{{- range $vi, $value := sortAlpha $enum.Data.Values }}
    {{ camel $value }}: "{{ $value }}",
{{- end }}
}
`
