package graphql

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/endpointflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
)

type SchemaContext struct {
	enums        map[string]*graphql.Enum
	objects      map[string]*graphql.Object
	inputObjects map[string]*graphql.InputObject
}

func (c SchemaContext) MustGetEnum(name string) *graphql.Enum {
	e, ok := c.enums[name]
	if !ok {
		panic(fmt.Sprintf("enum %v not found", name))
	}

	return e
}

func (c SchemaContext) MustGetInputObject(name string) *graphql.InputObject {
	o, ok := c.inputObjects[name]
	if !ok {
		panic(fmt.Sprintf("inputObject %v not found", name))
	}

	return o
}

func (c SchemaContext) MustGetObject(name string) *graphql.Object {
	o, ok := c.objects[name]
	if !ok {
		panic(fmt.Sprintf("object %v not found", name))
	}

	return o
}

func GetSchema(f generic.Factory, serveFunc types.ServeFunc, rn *graph.RootNode) (schema graphql.Schema, err error) {
	sCtx := SchemaContext{
		enums:        map[string]*graphql.Enum{},
		objects:      map[string]*graphql.Object{},
		inputObjects: map[string]*graphql.InputObject{},
	}

	rn.Types.Each(func(tn *graph.TypeNode) {
		if tn.Flags().Is(typeflags.RequestScope, true) {
			sCtx.inputObjects[tn.Name()+"Input"] = &graphql.InputObject{}
		}

		if tn.Flags().Is(typeflags.ResponseScope, true) {
			sCtx.objects[tn.Name()] = &graphql.Object{}
		}
	})

	rn.Enums.Each(func(en *graph.EnumNode) {
		sCtx.enums[en.Name()] = &graphql.Enum{}
	})

	rn.Types.Each(func(tn *graph.TypeNode) {
		if tn.Flags().Is(typeflags.RequestScope, true) {
			*sCtx.inputObjects[tn.Name()+"Input"] = *TypeToInputObject(sCtx, tn)
		}

		if tn.Flags().Is(typeflags.ResponseScope, true) {
			*sCtx.objects[tn.Name()] = *TypeToObject(sCtx, tn)
		}
	})

	rn.Enums.Each(func(en *graph.EnumNode) {
		*sCtx.enums[en.Name()] = *TypeToEnum(en)
	})

	schema, err = graphql.NewSchema(graphql.SchemaConfig{Query: getQueryObject(f, serveFunc, sCtx, rn), Mutation: getMutationObject(f, serveFunc, sCtx, rn)})
	if err != nil {
		return
	}

	return
}

func getEndpointField(f generic.Factory, serveFunc types.ServeFunc, en *graph.EndpointNode, sCtx SchemaContext) *graphql.Field {
	args := graphql.FieldConfigArgument{}
	for _, fn := range en.Edges.Type.Request().Edges.Fields.Holds() {
		switch fn.Name() {
		case fieldnames.Relations:
		case fieldnames.Select:
		default:
			if fn.Name() == "pages" {
				println("a")
			}
			if fn.IsTypeList() {
				args[fn.Name()] = &graphql.ArgumentConfig{
					Type: graphql.NewList(sCtx.MustGetInputObject(fn.Edges.Type.Holds().Name() + "Input")),
				}
			} else {
				args[fn.Name()] = &graphql.ArgumentConfig{
					Type: sCtx.MustGetInputObject(fn.Edges.Type.Holds().Name() + "Input"),
				}
			}
		}
	}

	return &graphql.Field{
		Type:        sCtx.MustGetObject(en.Edges.Type.Response().Name()),
		Description: "",
		Args:        args,
		Resolve:     composeResolve(f, serveFunc, en),
	}
}

func getQueryObject(f generic.Factory, serveFunc types.ServeFunc, sCtx SchemaContext, rn *graph.RootNode) (q *graphql.Object) {
	fs := graphql.Fields{}

	enm := rn.Endpoints.Filter(graph.Filter{
		Flags: &graph.FlagsSubset{
			Or: []string{endpointflags.IsGetEndpoint},
		},
	})

	enm.Each(func(en *graph.EndpointNode) {
		fs[en.FieldName()] = getEndpointField(f, serveFunc, en, sCtx)
	})

	q = graphql.NewObject(graphql.ObjectConfig{
		Name:   "Query",
		Fields: fs,
	})

	return
}

func getMutationObject(f generic.Factory, serveFunc types.ServeFunc, sCtx SchemaContext, rn *graph.RootNode) (q *graphql.Object) {
	fs := graphql.Fields{}

	enm := rn.Endpoints.Filter(graph.Filter{
		Flags: &graph.FlagsSubset{
			Or: []string{endpointflags.IsDeleteEndpoint, endpointflags.IsPutEndpoint, endpointflags.IsPostEndpoint, endpointflags.IsActionEndpoint},
		},
	})

	enm.Each(func(en *graph.EndpointNode) {
		fs[en.FieldName()] = getEndpointField(f, serveFunc, en, sCtx)
	})

	q = graphql.NewObject(graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: fs,
	})

	return
}

func TypeToInputObject(sCtx SchemaContext, tn *graph.TypeNode) (i *graphql.InputObject) {
	fields := graphql.InputObjectConfigFieldMap{}

	tn.Edges.Fields.Holds().Each(func(fn *graph.FieldNode) {
		fields[fn.Name()] = &graphql.InputObjectFieldConfig{
			Type: getObjectFieldType(sCtx, true, fn),
		}
	})

	if len(fields) == 0 {
		fields["dummy"] = &graphql.InputObjectFieldConfig{Type: graphql.Boolean}
	}

	i = graphql.NewInputObject(graphql.InputObjectConfig{
		Name:   tn.Name() + "Input",
		Fields: fields,
	})

	return
}

func getObjectFieldType(sCtx SchemaContext, asInput bool, fn *graph.FieldNode) (o graphql.Output) {
	suffix := ""
	if asInput {
		suffix = "Input"
	}

	switch fn.Kind() {
	case graph.FieldKindType:
		switch asInput {
		case true:
			o = sCtx.MustGetInputObject(fn.Edges.Type.Holds().Name() + suffix)
		case false:
			o = sCtx.MustGetObject(fn.Edges.Type.Holds().Name() + suffix)
		}
	case graph.FieldKindEnum:
		o = graphql.String
	case graph.FieldKindString:
		o = graphql.String
	case graph.FieldKindInt32:
		o = graphql.Int
	case graph.FieldKindFloat64:
		o = graphql.Float
	case graph.FieldKindBool:
		o = graphql.Boolean
	case graph.FieldKindTypeList:
		switch asInput {
		case true:
			o = graphql.NewList(sCtx.MustGetInputObject(fn.Edges.Type.Holds().Name() + suffix))
		case false:
			o = graphql.NewList(sCtx.MustGetObject(fn.Edges.Type.Holds().Name() + suffix))
		}
	case graph.FieldKindEnumList:
		o = graphql.NewList(graphql.String)
	case graph.FieldKindStringList:
		o = graphql.NewList(graphql.String)
	case graph.FieldKindInt32List:
		o = graphql.NewList(graphql.Int)
	case graph.FieldKindFloat64List:
		o = graphql.NewList(graphql.Float)
	case graph.FieldKindBoolList:
		o = graphql.NewList(graphql.Boolean)
	default:
		panic(fmt.Sprintf("unexpected field kind %v", fn.Kind()))
	}

	return
}

func TypeToObject(sCtx SchemaContext, tn *graph.TypeNode) (i *graphql.Object) {
	fields := graphql.Fields{}

	isRelations := tn.Flags().Is(typeflags.IsRelations, true)

	tn.Edges.Fields.Holds().Each(func(fn *graph.FieldNode) {
		var args graphql.FieldConfigArgument

		if isRelations && fn.Edges.Type.Holds().Flags().Is(typeflags.IsCollection, true) {
			args = graphql.FieldConfigArgument{}

			for _, fn0 := range fn.Edges.Type.Holds().Edges.Type.For().Edges.Type.GetCollection().Edges.Fields.Holds() {
				switch fn0.Name() {
				case fieldnames.Relations:
				case fieldnames.Select:
				default:
					args[fn0.Name()] = &graphql.ArgumentConfig{
						Type: sCtx.MustGetInputObject(fn0.Edges.Type.Holds().Name() + "Input"),
					}
				}
			}
		}

		fields[fn.Name()] = &graphql.Field{
			Type: getObjectFieldType(sCtx, false, fn),
			Args: args,
		}
	})

	if len(fields) == 0 {
		fields["dummy"] = &graphql.Field{
			Type: graphql.Boolean,
		}
	}

	i = graphql.NewObject(graphql.ObjectConfig{
		Name:   tn.Name(),
		Fields: fields,
	})

	return
}

func TypeToEnum(en *graph.EnumNode) (e *graphql.Enum) {
	values := graphql.EnumValueConfigMap{}

	for i, v := range en.Data.Values {
		values[v] = &graphql.EnumValueConfig{
			Value: i,
		}
	}

	e = graphql.NewEnum(graphql.EnumConfig{
		Name:   en.Name(),
		Values: values,
	})

	return
}
