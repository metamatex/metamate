package get

import (
	"github.com/metamatex/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamatemono/pkg/generic/pkg/v0/generic"
	"github.com/metamatex/metamatemono/pkg/generic/pkg/v0/transport/httpjson"
	"github.com/metamatex/metamatemono/gen/v0/sdk"
	"github.com/metamatex/metamatemono/pkg/metactl/pkg/v0/types"
	"net/http"
)

func Get(rn *graph.RootNode, f generic.Factory, addr string, token string, typenodeId string) (o types.Output, err error) {
	c := httpjson.NewClient(f, &http.Client{}, addr, token)

	tn, err := rn.Types.ById(graph.NodeId(typenodeId))
	if err != nil {
		return
	}

	gReq := f.New(tn.Edges.Type.GetRequest())


	mode := sdk.GetMode{
		Kind: &sdk.GetModeKind.Collection,
		Collection: &sdk.CollectionGetMode{},
	}

	gReq.MustSetGeneric([]string{fieldnames.Mode}, f.MustFromStruct(mode))

	gRsp, err := c.Send(gReq)
	if err != nil {
	    return
	}

	o.Text = gRsp.Sprint()
	o.Data = gRsp.ToStringInterfaceMap()

	return
}
