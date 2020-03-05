package expansion

import (
	"fmt"
	"github.com/metamatex/asg/pkg/v0/asg/endpointnames"
	"github.com/metamatex/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/asg/pkg/v0/asg/graph"
	"github.com/metamatex/asg/pkg/v0/asg/graph/fieldflags"
	"github.com/metamatex/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/asg/pkg/v0/asg/typenames"
)

func generatePostEndpointType(root *graph.RootNode, tn *graph.TypeNode) {
	endpointName := endpointnames.Post(tn.Name())

	endpointTn := root.AddTypeNode(typenames.Endpoint(endpointName), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Filter, typenames.Filter(typenames.Request(endpointName)), graph.Flags{
			fieldflags.Filter: false,
		}),
	}, graph.Flags{
		typeflags.IsEndpoint: true,
	})

	tn.Edges.Type.Resolver.SetPostEndpoint(endpointTn.Id())
}

func generateGetEndpointType(root *graph.RootNode, tn *graph.TypeNode) {
	endpointName := endpointnames.Get(tn.Name())

	endpointTn := root.AddTypeNode(typenames.Endpoint(endpointName), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Filter, typenames.Filter(typenames.Request(endpointName)), graph.Flags{
			fieldflags.Filter: false,
		}),
	}, graph.Flags{
		typeflags.IsEndpoint: true,
	})

	tn.Edges.Type.Resolver.SetGetEndpoint(endpointTn.Id())
}

func generatePutEndpointType(root *graph.RootNode, tn *graph.TypeNode) {
	endpointName := endpointnames.Put(tn.Name())

	endpointTn := root.AddTypeNode(typenames.Endpoint(endpointName), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Filter, typenames.Filter(typenames.Request(endpointName)), graph.Flags{
			fieldflags.Filter: false,
		}),
	}, graph.Flags{
		typeflags.IsEndpoint: true,
	})

	tn.Edges.Type.Resolver.SetPutEndpoint(endpointTn.Id())
}

func generateDeleteEndpointType(root *graph.RootNode, tn *graph.TypeNode) {
	endpointName := endpointnames.Delete(tn.Name())

	endpointTn := root.AddTypeNode(typenames.Endpoint(endpointName), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Filter, typenames.Filter(typenames.Request(endpointName)), graph.Flags{
			fieldflags.Filter: false,
		}),
	}, graph.Flags{
		typeflags.IsEndpoint: true,
	})

	tn.Edges.Type.Resolver.SetDeleteEndpoint(endpointTn.Id())
}

func generatePipeEndpointType(root *graph.RootNode, tn *graph.TypeNode) {
	endpointName := endpointnames.Pipe(tn.Name())

	endpointTn := root.AddTypeNode(typenames.Endpoint(endpointName), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Filter, typenames.Filter(typenames.Request(endpointName)), graph.Flags{
			fieldflags.Filter: false,
		}),
	}, graph.Flags{
		typeflags.IsEndpoint: true,
	})

	tn.Edges.Type.Resolver.SetPipeEndpoint(endpointTn.Id())
}

func generateActionEndpointType(root *graph.RootNode, en *graph.EndpointNode) {
	root.AddTypeNode(typenames.Endpoint(en.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Filter, typenames.Filter(typenames.Request(en.Name())), graph.Flags{
			fieldflags.Filter: false,
		}),
	}, graph.Flags{
		typeflags.IsEndpoint: true,
	})
}

func generateEndpointType(root *graph.RootNode, en *graph.EndpointNode) {
	tn := root.AddTypeNode(typenames.Endpoint(en.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Filter, typenames.Filter(typenames.Request(en.Name())), graph.Flags{
			fieldflags.Filter: false,
		}),
	}, graph.Flags{
		typeflags.IsEndpoint: true,
	})

	switch en.Data.Method {
	case graph.MethodPost:
		en.Edges.Type.For().Edges.Type.Resolver.SetPostEndpoint(tn.Id())
		tn.Edges.Endpoint.Resolver.SetPost(en.Id())
	case graph.MethodGet:
		en.Edges.Type.For().Edges.Type.Resolver.SetGetEndpoint(tn.Id())
		tn.Edges.Endpoint.Resolver.SetGet(en.Id())
	case graph.MethodPut:
		en.Edges.Type.For().Edges.Type.Resolver.SetPutEndpoint(tn.Id())
		tn.Edges.Endpoint.Resolver.SetPut(en.Id())
	case graph.MethodDelete:
		en.Edges.Type.For().Edges.Type.Resolver.SetDeleteEndpoint(tn.Id())
		tn.Edges.Endpoint.Resolver.SetDelete(en.Id())
	case graph.MethodAction:
	default:
		panic(fmt.Sprintf("unsupported method %v", en.Data.Method))
	}
}

//func generateEndpointUnion(root *graph.RootNode) {
//	names := []interface{}{}
//	root.Types.Flagged(typeflags.IsEndpoint, true).Each(func(tn *graph.TypeNode) {
//		names = append(names, tn.Name())
//	})
//
//	endpointNode, _ := root.AddUnion(typenames.EndpointNode, names)
//
//	endpointNode.Flags().
//		Set(typeflags.IsOptionalValueUnion, true)
//
//	root.Endpoints.Each(func(en *graph.EndpointNode) {
//		n := strcase.ToLowerCamel(typenames.EndpointNode(en.Name()))
//
//		en.Edges.Type.Request().Data.EndpointNode = &n
//	})
//}

func generateEndpoints(root *graph.RootNode) {
	fns := graph.FieldNodeSlice{}

	root.Endpoints.Each(func(en *graph.EndpointNode) {
		fns = append(fns, graph.TypeField(en.FieldName(), typenames.Endpoint(en.Name())))
	})

	root.AddTypeNode(typenames.Endpoints, fns)
}

func addEndpointKindToRequests(root *graph.RootNode) {
	root.Endpoints.Each(func(en *graph.EndpointNode) {
		en.Edges.Type.Response().Edges.Endpoint.Resolver.SetBelongsTo(en.Id())

		en.Edges.Type.Request().Edges.Endpoint.Resolver.SetBelongsTo(en.Id())
	})
}