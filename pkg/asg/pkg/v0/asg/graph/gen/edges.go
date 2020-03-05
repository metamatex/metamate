package main

import (
	"fmt"
	"os"
	"reflect"
	"text/template"
	"unicode"
	"unicode/utf8"
)

type StaticDynamic struct {
	Static  []string
	Dynamic []string
}

type StaticDynamicContext struct {
	StaticDynamic
	From string
	To   string
}

type FromToContext struct {
	From string
	To   []string
}

type Edges struct {
	BasicType *StaticDynamic
	Endpoint  *StaticDynamic
	Enum      *StaticDynamic
	Field     *StaticDynamic
	Relation  *StaticDynamic
	Type      *StaticDynamic
	Path      *StaticDynamic
}

type NodeEdge struct {
	Many *Edges
	One  *Edges
}

type NodeEdges struct {
	BasicType *NodeEdge
	Endpoint  *NodeEdge
	Enum      *NodeEdge
	Field     *NodeEdge
	Relation  *NodeEdge
	Type      *NodeEdge
	Path      *NodeEdge
}

const (
	BasicType = "BasicType"
	Endpoint  = "Endpoint"
	Enum      = "Enum"
	Field     = "Field"
	Relation  = "Relation"
	Type      = "Type"
	Path      = "Path"
)

const (
	Request  = "Request"
	Response = "Response"
)

var nodeEdges = NodeEdges{
	Endpoint: &NodeEdge{
		One: &Edges{
			Type: &StaticDynamic{Static: []string{"For", "Request", "Response"}},
		},
		Many: &Edges{
			Type: &StaticDynamic{Dynamic: []string{"Dependencies"}},
			Enum: &StaticDynamic{Dynamic: []string{"Dependencies"}},
		},
	},
	Enum: &NodeEdge{
		One: &Edges{
			Type: &StaticDynamic{Static: []string{"FilteredBy"},},
		},
	},
	Field: &NodeEdge{
		One: &Edges{
			BasicType: &StaticDynamic{Static: []string{"Holds"},},
			Enum:      &StaticDynamic{Static: []string{"Holds"}},
			Field:     &StaticDynamic{Static: []string{"For", "RelatedTo"}},
			Relation:  &StaticDynamic{Static: []string{"RelatedThrough"}},
			Type:      &StaticDynamic{Static: []string{"Holds", "HeldBy"}},
			Path:      &StaticDynamic{Static: []string{"BelongsTo"}},
		},
	},
	Relation: &NodeEdge{
		One: &Edges{
			Type: &StaticDynamic{Static: []string{"NodeA", "NodeB"},},
			Path: &StaticDynamic{Static: []string{"Active", "Passive"},},
		},
		// todo remove many Holds Fields egde
		Many: &Edges{
			Field: &StaticDynamic{Static: []string{"Holds"}},
		},
	},
	Type: &NodeEdge{
		One: &Edges{
			Type: &StaticDynamic{Static: []string{"For", "FilteredBy", "SortedBy", "SelectedBy", "Collection", "Request", "Response", "GetRequest", "GetCollection", "GetRelations", "GetResponse", "PostRequest", "PostResponse", "PutRequest", "PutResponse", "DeleteRequest", "DeleteResponse", "PipeRequest", "PipeResponse", "PostEndpoint", "GetEndpoint", "PutEndpoint", "DeleteEndpoint", "PipeEndpoint"}},
			Enum: &StaticDynamic{Static: []string{"For", "ListKind"}},
			Endpoint: &StaticDynamic{Static: []string{"BelongsTo", "Post", "Get", "Pipe", "Put", "Delete"}},
		},
		Many: &Edges{
			Field:    &StaticDynamic{Static: []string{"Holds", "EdgedByFields", "EdgedByListFields"}},
			Relation: &StaticDynamic{Static: []string{"Holds"}},
			Type:     &StaticDynamic{Static: []string{"ToOneRelations", "ToManyRelations"}, Dynamic: []string{"Misses", "Dependencies"}},
			Enum:     &StaticDynamic{Dynamic: []string{"Dependencies"}},
		},
	},
	Path: &NodeEdge{
		One: &Edges{
			Type:     &StaticDynamic{Static: []string{"From", "To"}},
			Relation: &StaticDynamic{Static: []string{"BelongsTo"}},
		},
	},
}

func main() {
	errs := processNodeEdges(nodeEdges)
	if len(errs) != 0 {
		fmt.Printf("%v", errs)
	}
}

func processNodeEdges(nodeEdges NodeEdges) (errs []error) {
	v := reflect.ValueOf(nodeEdges)
	t := v.Type()

	one := []StaticDynamicContext{}
	many := []StaticDynamicContext{}
	for i := 0; i < t.NumField(); i++ {
		v0 := v.Field(i)
		if v0.IsNil() {
			continue
		}

		nodeEdge := v0.Elem().Interface().(NodeEdge)

		one0, many0 := getRenderContextsFromNodeEdge(t.Field(i).Name, nodeEdge)
		one = append(one, one0...)
		many = append(many, many0...)
	}

	err := renderOne(one)
	if err != nil {
		errs = append(errs, err)
	}

	err = renderMany(many)
	if err != nil {
		errs = append(errs, err)
	}

	fromTo := map[string][]string{}
	for _, c := range one {
		fromTo[c.From] = append(fromTo[c.From], c.To)
	}

	for _, c := range many {
		fromTo[c.From] = append(fromTo[c.From], c.To+"s")
	}

	fromToCtxs := []FromToContext{}
	for k, v := range fromTo {
		fromToCtxs = append(fromToCtxs, FromToContext{From: k, To: v})
	}

	err = renderEdges(fromToCtxs)
	if err != nil {
		errs = append(errs, err)
	}

	return
}

func getRenderContextsFromNodeEdge(from string, edge NodeEdge) (one []StaticDynamicContext, many []StaticDynamicContext) {
	if edge.One != nil {
		one = getRenderContextsFromEdges(from, *edge.One)
	}

	if edge.Many != nil {
		many = getRenderContextsFromEdges(from, *edge.Many)
	}

	return
}

func getRenderContextsFromEdges(from string, edges Edges) (contexts []StaticDynamicContext) {
	v := reflect.ValueOf(edges)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		v0 := v.Field(i)
		if v0.IsNil() {
			continue
		}

		staticDynamic := v0.Elem().Interface().(StaticDynamic)

		contexts = append(contexts, StaticDynamicContext{
			From:          from,
			To:            t.Field(i).Name,
			StaticDynamic: staticDynamic,
		})
	}

	return
}

func renderEdges(contexts []FromToContext) (err error) {
	err = render("Edges_", contexts, edgesTemplate)
	if err != nil {
		return
	}

	return
}

func renderOne(contexts []StaticDynamicContext) (err error) {
	err = render("OneIdResolver_", contexts, oneIdResolverTemplate)
	if err != nil {
		return
	}

	err = render("OneEdges_", contexts, oneEdgesTemplate)
	if err != nil {
		return
	}

	return
}

func renderMany(contexts []StaticDynamicContext) (err error) {
	err = render("ManyIdResolver_", contexts, manyIdResolverTemplate)
	if err != nil {
		return
	}

	err = render("ManyEdges_", contexts, manyEdgesTemplate)
	if err != nil {
		return
	}

	return
}

func lower(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

func render(name string, contexts interface{}, templateData string) (err error) {
	f, err := os.Create(name + ".go")
	if err != nil {
		return
	}
	defer f.Close()

	err = template.Must(template.New("").Funcs(template.FuncMap{"lower": lower}).Parse(templateData)).Execute(f, contexts)
	if err != nil {
		return
	}

	return
}

var oneIdResolverTemplate = `// generated by go:generate go run gen/edges.go
package graph

import "fmt"
{{ range $i, $context := . }}
{{- $name := printf "%vTo%vIdResolver" $context.From $context.To -}}
{{ range $i, $s := $context.Dynamic }}
var {{ $name }}_{{ $s }} func({{ $name }}) (NodeId) 
{{- end }}
type {{ $name }} struct {
	n *{{ $context.From }}Node
	d  map[string]NodeId
}

func New{{ $name }}(n *{{ $context.From }}Node) ({{ $name }}) {
	return {{ $name }}{
		n: n,
		d: map[string]NodeId{},
	}
}

func (r *{{ $name }}) set(name string, id NodeId) {
	id0, ok := r.d[name]
	if ok {
		panic(fmt.Sprintf("%v {{ $context.From }} edge %v already set to %v", r.n.Name(), name, id0))
	}

	r.d[name] = id
}
{{ range $i, $s := .Static }}
func (r *{{ $name }}) Set{{ $s }}(id NodeId) (*{{ $name }}) {
	r.set("{{ $s }}", id)

	return r
}

func (r {{ $name }}) {{ $s }}()(NodeId) {
	return r.d["{{ $s }}"]
}
{{ end }}
{{ range $i, $s := .Dynamic }}
func (r *{{ $name }}) Set{{ $s }}(id NodeId) {
	panic(fmt.Sprintf("dynamic egde can not be set to %v", id))
}

func (r {{ $name }}) {{ $s }}()(NodeId) {
	return {{ $name }}_{{ $s }}(r)
}
{{ end }}
{{- end }}
`

var oneEdgesTemplate = `// generated by go:generate go run gen/edges.go
package graph
{{ range $i, $context := . }}
const (
{{- range $i, $s := $context.Static }}
	{{ $context.From }}{{ $s }}{{ $context.To }} = "{{ $context.From }},one,{{ $context.To }},{{ $s }}" 
{{- end }}
{{- range $i, $s := $context.Dynamic }}
	{{ $context.From }}{{ $s }}{{ $context.To }} = "{{ $context.From }},one,{{ $context.To }},{{ $s }}"
{{- end }}
)
{{ $name := printf "%vTo%vEdges" $context.From $context.To }}
type {{ $name }} struct {
	Resolver {{ $context.From }}To{{ $context.To }}IdResolver ` + "`" + `resolves:"{{ $context.From }},one,{{ $context.To }}"` + "`" + `
}

{{ range $i, $s := $context.Static }}
func (e {{ $name }}) {{ $s }}() (*{{ $context.To }}Node) {
	if e.Resolver.{{ $s }}() == "" {
		return nil
	}

	return e.Resolver.n.root.{{ $context.To }}s.MustById(e.Resolver.{{ $s }}())
}
{{- end }}
{{ range $i, $s := $context.Dynamic }}
func (e {{ $name }}) {{ $s }}() (*{{ $context.To }}Node) {
	return e.Resolver.n.root.{{ $context.To }}s.MustById(e.Resolver.{{ $s }}())
}
{{ end }}
{{- end }}
`

var manyIdResolverTemplate = `// generated by go:generate go run gen/edges.go
package graph

import (
	"fmt"
	"reflect"
	"context"
)
{{ range $i, $context := . }}
{{- $name := printf "%vTo%vsIdResolver" $context.From $context.To -}}
{{ range $i, $s := $context.Dynamic }}
var {{ $name }}_{{ $s }} func(context.Context, {{ $name }}) ([]NodeId) 
{{- end }}
type {{ $name }} struct {
	n *{{ $context.From }}Node
	d  map[string]map[NodeId]bool
}

func New{{ $name }}(n *{{ $context.From }}Node) ({{ $name }}) {
	return {{ $name }}{
		n: n,
		d: map[string]map[NodeId]bool{},
	}
}

func (r *{{ $name }}) add(name string, ids []NodeId) {
	_, ok := r.d[name]
	if !ok {
		r.d[name] = map[NodeId]bool{}
	}

	for _, id := range ids {
		id0, ok := r.d[name][id]
		if ok {
			panic(fmt.Sprintf("%v %v %v already set %v", r.n.Name(), reflect.TypeOf(r).Name(), name, id0))
		}

		r.d[name][id] = true
	}
}

func (r {{ $name }}) get(name string) (ids []NodeId) {
	for id, _ := range r.d[name] {
		ids = append(ids, id)
	}

	return
}
{{ range $i, $s := .Static }}
func (r *{{ $name }}) Add{{ $s }}(ids ...NodeId) (*{{ $name }}) {
	r.add("{{ $s }}", ids)

	return r
}

func (r {{ $name }}) {{ $s }}()([]NodeId) {
	return r.get("{{ $s }}")
}
{{ end }}
{{ range $i, $s := .Dynamic }}
func (r *{{ $name }}) Add{{ $s }}(ids ...NodeId) {
	panic(fmt.Sprintf("dynamic egde can not be set to %v", ids))
}

func (r {{ $name }}) {{ $s }}()([]NodeId) {
	return r.{{ lower $s }}(context.Background())
}

func (r {{ $name }}) {{ lower $s }}(ctx context.Context)([]NodeId) {
	return {{ $name }}_{{ $s }}(ctx, r)
}
{{ end }}
{{- end }}
`

var manyEdgesTemplate = `// generated by go:generate go run gen/edges.go
package graph

{{ range $i, $context := .}}
const (
{{- range $i, $s := $context.Static }}
	{{ $context.From }}{{ $s }}{{ $context.To }}s = "{{ $context.From }},many,{{ $context.To }},{{ $s }}" 
{{- end }}
{{- range $i, $s := $context.Dynamic }}
	{{ $context.From }}{{ $s }}{{ $context.To }}s = "{{ $context.From }},many,{{ $context.To }},{{ $s }}"
{{- end }}
)
{{- $name := printf "%vTo%vsEdges" $context.From $context.To }}
type {{ $name }} struct {
	Resolver {{ $context.From }}To{{ $context.To }}sIdResolver ` + "`" + `resolves:"{{ $context.From }},many,{{ $context.To }}"` + "`" + `
}

{{ range $i, $s := .Static }}
func (e {{ $name }}) {{ $s }}() ({{ $context.To }}NodeMap) {
	return e.Resolver.n.root.{{ $context.To }}s.ByIds(e.Resolver.{{ $s }}()...)
}
{{- end }}
{{ range $i, $s := .Dynamic }}
func (e {{ $name }}) {{ $s }}() ({{ $context.To }}NodeMap) {
	return e.Resolver.n.root.{{ $context.To }}s.ByIds(e.Resolver.{{ $s }}()...)
}
{{- end }}
{{- end }}
`

var edgesTemplate = `// generated by go:generate go run gen/edges.go
package graph

{{ range $i, $context := .}}
type {{ $context.From }}NodeEdges struct {
{{- range $i, $to := $context.To }}
	{{ $to }} {{ $context.From }}To{{ $to }}Edges
{{- end }}
}

func New{{ $context.From }}NodeEdges(n *{{ $context.From }}Node) ({{ $context.From }}NodeEdges) {
	return {{ $context.From }}NodeEdges{
{{- range $i, $to := $context.To }}
		{{ $to }}: {{ $context.From }}To{{ $to }}Edges{
			Resolver: New{{ $context.From }}To{{ $to }}IdResolver(n),
		},
{{- end }}
	}
}
{{ end }}
`
