package expansion

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/endpointnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/endpointflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/fieldflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/typenames"
)

func generatePipeRequest(root *graph.RootNode, tn *graph.TypeNode) {
	requestNode := root.AddTypeNode(typenames.PipeRequest(tn.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Meta, typenames.RequestMeta),
		graph.TypeField(fieldnames.Auth, typenames.Auth, graph.Flags{
			fieldflags.Filter:        false,
			fieldflags.ValidateIsSet: false,
		}),
		graph.TypeField(fieldnames.Mode, typenames.PipeMode, graph.Flags{
			fieldflags.ValidateIsSet: true,
		}),
		graph.TypeField(fieldnames.Context, typenames.PipeContext(tn.Name())),
	}, graph.Flags{
		typeflags.IsRequest:       true,
		typeflags.IsPipeRequest:   true,
	})
	requestNode.Edges.Type.Resolver.
		SetFor(tn.Id()).
		SetResponse(graph.ToNodeId(typenames.PipeResponse(tn.Name())))

	tn.Edges.Type.Resolver.
		SetPipeRequest(requestNode.Id())
}

func generatePipeResponse(root *graph.RootNode, tn *graph.TypeNode) {
	responseNode := root.AddTypeNode(typenames.PipeResponse(tn.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Meta, typenames.ResponseMeta),
		graph.TypeField(fieldnames.Context, typenames.PipeContext(tn.Name())),
	}, graph.Flags{
		typeflags.IsResponse: true,
	})
	responseNode.Edges.Type.Resolver.
		SetFor(tn.Id()).
		SetRequest(graph.ToNodeId(typenames.PipeRequest(tn.Name())))

	tn.Edges.Type.Resolver.
		SetPipeResponse(responseNode.Id())
}

func generatePipeEndpoint(root *graph.RootNode, tn *graph.TypeNode) {
	en := root.AddEndpointNode(endpointnames.Pipe(tn.Name()), graph.MethodPipe, tn.Name(), typenames.PipeRequest(tn.Name()), typenames.PipeResponse(tn.Name()), graph.Flags{
		endpointflags.IsPipeEndpoint: true,
	})

	tn.Edges.Endpoint.Resolver.SetPipe(en.Id())
}

func generatePipeContext(root *graph.RootNode, tn *graph.TypeNode) {
	root.AddTypeNode(typenames.PipeContext(tn.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Post, typenames.PipePostContext(tn.Name())),
		graph.TypeField(fieldnames.Get, typenames.PipeGetContext(tn.Name())),
		graph.TypeField(fieldnames.Put, typenames.PipePutContext(tn.Name())),
		graph.TypeField(fieldnames.Delete, typenames.PipeDeleteContext(tn.Name())),
	})
}

func generatePipePostContext(root *graph.RootNode, tn *graph.TypeNode) {
	root.AddTypeNode(typenames.PipePostContext(tn.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.ClientRequest, typenames.PostRequest(tn.Name())),
		graph.TypeField(fieldnames.ServiceRequest, typenames.PostRequest(tn.Name())),
		graph.TypeField(fieldnames.ServiceResponse, typenames.PostResponse(tn.Name())),
		graph.TypeField(fieldnames.ClientResponse, typenames.PostResponse(tn.Name())),
	})
}

func generatePipeGetContext(root *graph.RootNode, tn *graph.TypeNode) {
	root.AddTypeNode(typenames.PipeGetContext(tn.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.ClientRequest, typenames.GetRequest(tn.Name())),
		graph.TypeField(fieldnames.ServiceRequest, typenames.GetRequest(tn.Name())),
		graph.TypeField(fieldnames.ServiceResponse, typenames.GetResponse(tn.Name())),
		graph.TypeField(fieldnames.ClientResponse, typenames.GetResponse(tn.Name())),
	})
}

func generatePipePutContext(root *graph.RootNode, tn *graph.TypeNode) {
	root.AddTypeNode(typenames.PipePutContext(tn.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.ClientRequest, typenames.PutRequest(tn.Name())),
		graph.TypeField(fieldnames.ServiceRequest, typenames.PutRequest(tn.Name())),
		graph.TypeField(fieldnames.ServiceResponse, typenames.PutResponse(tn.Name())),
		graph.TypeField(fieldnames.ClientResponse, typenames.PutResponse(tn.Name())),
	})
}

func generatePipeDeleteContext(root *graph.RootNode, tn *graph.TypeNode) {
	root.AddTypeNode(typenames.PipeDeleteContext(tn.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.ClientRequest, typenames.DeleteRequest(tn.Name())),
		graph.TypeField(fieldnames.ServiceRequest, typenames.DeleteRequest(tn.Name())),
		graph.TypeField(fieldnames.ServiceResponse, typenames.DeleteResponse(tn.Name())),
		graph.TypeField(fieldnames.ClientResponse, typenames.DeleteResponse(tn.Name())),
	})
}
