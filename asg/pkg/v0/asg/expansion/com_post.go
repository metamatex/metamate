package expansion

import (
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/endpointnames"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph/endpointflags"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph/fieldflags"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/typenames"
)

func generatePostRequest(root *graph.RootNode, tn *graph.TypeNode) {
	requestNode := root.AddTypeNode(typenames.PostRequest(tn.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Meta, typenames.RequestMeta),
		graph.TypeField(fieldnames.Mode, typenames.PostMode, graph.Flags{
			fieldflags.ValidateIsSet: true,
		}),
		graph.TypeField(fieldnames.Auth, typenames.Auth, graph.Flags{
			fieldflags.Filter:        false,
			fieldflags.ValidateIsSet: false,
		}),
		graph.TypeField(fieldnames.ServiceFilter, typenames.Filter(typenames.Service), graph.Flags{fieldflags.Filter: false}),
		graph.TypeField(fieldnames.Select, typenames.Select(typenames.PostResponse(tn.Name())), graph.Flags{
			fieldflags.Filter:        false,
			fieldflags.ValidateIsSet: true,
		}),
		graph.ListField(graph.TypeField(tn.PluralFieldName(), tn.Name(), graph.Flags{
			fieldflags.Filter:        true,
			fieldflags.ValidateIsSet: true,
		})),
	}, graph.Flags{
		typeflags.IsRequest:       true,
		typeflags.IsPostRequest:   true,
		typeflags.GetPassEndpoint: true,
	})
	requestNode.Edges.Type.Resolver.
		SetFor(tn.Id()).
		SetResponse(graph.ToNodeId(typenames.PostResponse(tn.Name())))

	tn.Edges.Type.Resolver.
		SetPostRequest(requestNode.Id())
}

func generatePostResponse(root *graph.RootNode, tn *graph.TypeNode) {
	responseNode := root.AddTypeNode(typenames.PostResponse(tn.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Meta, typenames.ResponseMeta),
		graph.ListField(graph.TypeField(tn.PluralFieldName(), tn.Name())),
	}, graph.Flags{
		typeflags.IsResponse: true,
	})
	responseNode.Edges.Type.Resolver.
		SetFor(tn.Id()).
		SetRequest(graph.ToNodeId(typenames.PostRequest(tn.Name())))

	tn.Edges.Type.Resolver.
		SetPostResponse(responseNode.Id())
}

func generatePostEndpoint(root *graph.RootNode, tn *graph.TypeNode) {
	en := root.AddEndpointNode(endpointnames.Post(tn.Name()), graph.MethodPost, tn.Name(), typenames.PostRequest(tn.Name()), typenames.PostResponse(tn.Name()), graph.Flags{
		endpointflags.IsPostEndpoint: true,
	})

	tn.Edges.Endpoint.Resolver.SetPost(en.Id())
}
