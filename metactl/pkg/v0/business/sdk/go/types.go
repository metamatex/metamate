package _go

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/metamatex/metamate/metactl/pkg/v0/utils/ptr"
)

const (
	TaskServiceInterface = "TaskServiceInterface"
	TaskTypes            = "TaskTypes"
	TaskEnums            = "TaskEnums"
	TaskFieldNames       = "TaskFieldNames"
	TaskTypeNames        = "TaskTypeNames"
	TaskPathNames        = "TaskPathNames"
)

func init() {
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

	tasks[TaskFieldNames] = types.RenderTask{
		TemplateData: &goFieldNamesTpl,
		Out:          ptr.String("fieldnames_.go"),
	}

	tasks[TaskTypeNames] = types.RenderTask{
		TemplateData: &goTypeNamesTpl,
		Out:          ptr.String("typenames_.go"),
	}

	tasks[TaskPathNames] = types.RenderTask{
		TemplateData: &goPathNamesTpl,
		Out:          ptr.String("pathnames_.go"),
	}
}

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

var goFieldNamesTpl = `package mql

var FieldNames = struct{
{{- range $i, $n := sortAlpha .Fields.UniqueNames }}
	{{ camel $n }} string
{{- end }}
}{
{{- range $i, $n := sortAlpha .Fields.UniqueNames }}
	{{ camel $n }}: "{{ $n }}",
{{- end }}
}`

var goTypeNamesTpl = `package mql

var TypeNames = struct{
{{- range $i, $tn := .Types.Slice.Sort }}
	{{ camel $tn.Name }} string
{{- end }}
}{
{{- range $i, $tn := .Types.Slice.Sort }}
	{{ camel $tn.Name }}: "{{ $tn.Name }}",
{{- end }}
}`

var goPathNamesTpl = `package mql

var PathNames = struct{
{{- range $i, $pn := .Paths.Slice.Sort }}
	{{ camel $pn.Name }} string
{{- end }}
}{
{{- range $i, $pn := .Paths.Slice.Sort }}
	{{ camel $pn.Name }}: "{{ $pn.Name }}",
{{- end }}
}`
