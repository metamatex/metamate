package expansion

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph"
	"testing"
)

func Test(t *testing.T) {
	root := graph.NewRoot()

	root.Wire()

	errs := root.Validate()
	if len(errs) != 0 {
		spew.Dump(errs)

	    return
	}

	Expand(0, root)

	//spew.Dump(root.Interfaces.GetIds())

	//unused := graph.GetUnused(root, graph.TYPE)
	//spew.Dump(unused)

	//root.types.Flagged(typeflags.IsRelationships, true).Each(func(n *graph.TypeNode) {
	//	n.Print()
	//})

	//root.types.ByName(graph.Person).Print()

	errs = root.Validate()
	if len(errs) != 0 {
		spew.Dump(errs)
	}

	//root.Endpoints.Filter(graph.Filter{
	//	Names: &graph.NamesSubset{
	//		Or: []string{"CreateWhatevers"},
	//	},
	//}).BroadcastPrint()


	//root.Types.MustById("getattachmentsrequest").Print()

	root.Types.MustByName("ClientAccountsCollectionSelect").Print()

	//s, m := graph.GetEdgeMaps(root.types.ById(graph.ToNodeId(typenames.Service)).Edges)
	//
	//spew.Dump(s)
	//spew.Dump(m)

	//root.Interfaces.ById("requestselect").Print()
}
