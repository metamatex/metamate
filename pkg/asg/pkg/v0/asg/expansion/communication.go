package expansion

import (
	"github.com/metamatex/asg/pkg/v0/asg/graph"
	"github.com/metamatex/asg/pkg/v0/asg/graph/typeflags"
)

func generateCommunications(root *graph.RootNode) {
	root.
		Types.
		FilterFunc(func(tn *graph.TypeNode) bool {
			return tn.Flags().Is(typeflags.GetEndpoints, true)
		}).
		Each(func(tn *graph.TypeNode) {
			//generateGetCommunication(root, tn)
			//generateCreateCommunication(root, tn)
			//generateUpdateCommunication(root, tn)
			//generateDeleteCommunication(root, tn)
			//generateStreamCommunication(root, tn)

			if tn.Flags().Is(typeflags.IsInList, true) {
				//generateModifyListByIdCommunication(root, tn)
			}
		})

	root.
		Types.
		FilterFunc(func(tn *graph.TypeNode) bool {
			return tn.Flags().Is(typeflags.GetPassEndpoint, true)
		}).
		Each(func(tn *graph.TypeNode) {
			//generatePassCommunication(root, tn)
		})
}
