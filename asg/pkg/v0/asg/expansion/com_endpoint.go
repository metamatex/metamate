package expansion

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/endpointnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/fieldflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/typenames"
)

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
