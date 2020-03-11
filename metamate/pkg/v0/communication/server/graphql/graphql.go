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

func composeResolve(f generic.Factory, serveFunc types.ServeFunc, en *graph.EndpointNode) func(params graphql.ResolveParams) (m interface{}, err error) {
	return func(params graphql.ResolveParams) (m interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("panic: ", string(debug.Stack()))
				panic(r)
			}
		}()

		selectedMap, err := getSelectedFields(params)
		if err != nil {
			return
		}

		gCliReq := f.New(en.Edges.Type.Request())
		for _, fn := range en.Edges.Type.Request().Edges.Fields.Holds() {
			_, ok := params.Args[fn.Name()]
			if ok {
				gCliReq.MustSetGeneric([]string{fn.Name()}, f.MustFromStringInterfaceMap(fn.Edges.Type.Holds(), params.Args[fn.Name()].(map[string]interface{})))
			}
		}

		gSelect := f.MustFromStringInterfaceMap(en.Edges.Type.Response().Edges.Type.SelectedBy(), selectedMap)

		gCliReq.MustSetGeneric([]string{fieldnames.Select}, gSelect)

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

func getSelectedFields(params graphql.ResolveParams) (selected map[string]interface{}, err error) {
	fieldASTs := params.Info.FieldASTs
	if len(fieldASTs) == 0 {
		err = fmt.Errorf("getSelectedFields: ResolveParams has no fields")

		return
	}

	selected, err = selectedFieldsFromSelections(params, fieldASTs[0].SelectionSet.Selections)
	if err != nil {
		return
	}

	return
}

//func getRelationsFromSelections(selections []ast.Selection) (relations map[string]interface{}, err error) {
//	relations = map[string]interface{}{}
//
//	for _, s := range selections {
//		switch sth := s.(type) {
//		case *ast.Field:
//			if len(sth.Arguments) != 0 {
//				relations[sth.Name.Value] = argumentsToMap(sth.Arguments)
//			}
//
//			if sth.SelectionSet != nil {
//				selected[s.Name.Value], err = getRelationsFromSelections(sth.SelectionSet.Selections)
//				if err != nil {
//					return
//				}
//			}
//		case *ast.FragmentSpread:
//		default:
//			err = fmt.Errorf("getRelationsFromSelections: found unexpected selection type %v", sth)
//
//			return
//		}
//	}
//
//	return
//}

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
