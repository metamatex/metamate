package graphql

import (
	"context"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/handler"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"log"
	"net/http"
	"runtime/debug"
)

func MustGetHandler(rn *graph.RootNode, f generic.Factory, serverFunc types.ServeFunc) (h http.Handler) {
	err := func() (err error) {
		schema, err := GetSchema(f, serverFunc, rn)
		if err != nil {
			return
		}

		h = handler.New(&handler.Config{
			Schema:   &schema,
			Pretty:   true,
			GraphiQL: true,
		})

		return
	}()
	if err != nil {
		panic(err)
	}

	return
}

func ExecuteQuery(schema graphql.Schema, query string) (r *graphql.Result, err error) {
	r = graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(r.Errors) > 0 {
		err = errors.New(fmt.Sprintf("wrong result, unexpected errors: %v", r.Errors))

		return
	}

	return
}

func getField(path []string, f *ast.Field) (*ast.Field) {
	if len(path) == 0 {
		panic("len(path) == 0")
	}
	for _, s := range f.SelectionSet.Selections {
		s, ok := s.(*ast.Field)
		if !ok {
			continue
		}

		if s.Name.Value == path[0] {
			if len(path) == 1 {
				return s
			} else {
				return getField(path[1:], s)
			}
		}
	}

	return nil
}

func fillGetNode(f generic.Factory, params graphql.ResolveParams, g generic.Generic, field *ast.Field) (g0 generic.Generic, err error) {
	argsM := argumentsToMap(field.Arguments)

	selectedMap, err := selectedFieldsFromSelections(params, field.SelectionSet.Selections)
	if err != nil {
		return
	}

	for _, fn := range g.Type().Edges.Fields.Holds() {
		_, ok := argsM[fn.Name()]
		if ok {
			g.MustSetGeneric([]string{fn.Name()}, f.MustFromStringInterfaceMap(fn.Edges.Type.Holds(), argsM[fn.Name()].(map[string]interface{})))
		}

		if fn.Edges.Type.Holds().Flags().Is(typeflags.IsSelect, true) {
			gSelect := f.MustFromStringInterfaceMap(fn.Edges.Type.Holds(), selectedMap)

			g.MustSetGeneric([]string{fieldnames.Select}, gSelect)
		}
	}

	getRelationsField := getField([]string{g.Type().Edges.Type.For().PluralFieldName(), fieldnames.Relations}, params.Info.FieldASTs[0])
	if getRelationsField != nil {
		gGetRelations := f.New(g.Type().Edges.Type.For().Edges.Type.GetRelations())

		for _, fn := range gGetRelations.Type().Edges.Fields.Holds() {
			getCollectionField := getField([]string{fn.Name()}, getRelationsField)
			if getCollectionField != nil {
				gGetCollection := f.New(fn.Edges.Type.Holds())
				gGetCollection, err = fillGetNode(f, params, gGetCollection, getCollectionField)
				if err != nil {
				    return
				}

				gGetRelations.MustSetGeneric([]string{fn.Name()}, gGetCollection)
			}
		}

		g.MustSetGeneric([]string{fieldnames.Relations}, gGetRelations)
	}

	g0 = g

	return
}

func composeResolve(f generic.Factory, serveFunc types.ServeFunc, en *graph.EndpointNode) func(params graphql.ResolveParams) (m interface{}, err error) {
	return func(params graphql.ResolveParams) (m interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("panic: ", string(debug.Stack()))
				panic(r)
			}
		}()

		gCliReq := f.New(en.Edges.Type.ClientRequest())

		gCliReq, err = fillGetNode(f, params, gCliReq, params.Info.FieldASTs[0])
		if err != nil {
		    return
		}

		var ctx context.Context
		if params.Context != nil {
			ctx = params.Context
		} else {
			ctx = context.Background()
		}

		gResponse := serveFunc(ctx, gCliReq)

		m = gResponse.ToStringInterfaceMap()

		return
	}
}

func selectedFieldsFromSelections(params graphql.ResolveParams, selections []ast.Selection) (selected map[string]interface{}, err error) {
	selected = map[string]interface{}{}

	for _, s := range selections {
		switch s := s.(type) {
		case *ast.Field:
			if s.SelectionSet == nil {
				selected[s.Name.Value] = true
			} else {
				selected[s.Name.Value], err = selectedFieldsFromSelections(params, s.SelectionSet.Selections)
				if err != nil {
					return
				}
			}
		case *ast.FragmentSpread:
			n := s.Name.Value
			frag, ok := params.Info.Fragments[n]
			if !ok {
				err = fmt.Errorf("getSelectedFields: no fragment found with name %v", n)

				return
			}

			selected[s.Name.Value], err = selectedFieldsFromSelections(params, frag.GetSelectionSet().Selections)
			if err != nil {
				return
			}
		default:
			err = fmt.Errorf("getSelectedFields: found unexpected selection type %v", s)

			return
		}
	}

	return
}

func argumentsToMap(args []*ast.Argument) (m map[string]interface{}) {
	m = map[string]interface{}{}

	for _, arg := range args {
		name := arg.Name.Value

		switch sth := arg.Value.(type) {
		case *ast.ObjectValue:
			m[name] = objectFieldsToMap(sth.Fields)
		default:
			spew.Dump(sth)
			panic("unhandled argument value")
		}
	}

	return
}

func objectFieldsToMap(fields []*ast.ObjectField) (m map[string]interface{}) {
	m = map[string]interface{}{}

	for _, f := range fields {
		name := f.Name.Value
		switch sth := f.GetValue().(type) {
		case *ast.ObjectValue:
			m[name] = objectFieldsToMap(sth.Fields)
		case *ast.StringValue:
			m[name] = sth.Value
		case *ast.BooleanValue:
			m[name] = sth.Value
		case *ast.FloatValue:
			m[name] = sth.Value
		case *ast.IntValue:
			m[name] = sth.Value
		case *ast.EnumValue:
			m[name] = sth.Value
		default:
			spew.Dump(sth)
			panic("unhandled field value")
		}
	}

	return
}
