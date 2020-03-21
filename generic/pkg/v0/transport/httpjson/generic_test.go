package httpjson

import (
	"context"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"net/http"
	"testing"

	"github.com/metamatex/metamate/asg/pkg/v0/asg/expansion"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/stretchr/testify/assert"

	"github.com/metamatex/metamate/gen/v0/sdk"
)

var root *graph.RootNode
var f generic.Factory

func init() {
	root = graph.NewRoot()

	err := expansion.Expand(0, root)
	if err != nil {
		panic(err)
	}

	f = generic.NewFactory(root)
}

var rsp = sdk.GetWhateversResponse{
	Whatevers: []sdk.Whatever{
		{
			StringField:  sdk.String(""),
			Int32Field:   sdk.Int32(0),
			Float64Field: sdk.Float64(0),
			BoolField:    sdk.Bool(false),
			EnumField:    sdk.String(sdk.WhateverKind.Red),
			UnionField: &sdk.WhateverUnion{
				Kind:        sdk.String(sdk.WhateverUnionKind.StringField),
				StringField: sdk.String(""),
			},
			StringList:  []string{""},
			Int32List:   []int32{0},
			Float64List: []float64{0},
			BoolList:    []bool{false},
			EnumList:    []string{sdk.WhateverKind.Red},
			UnionList: []sdk.WhateverUnion{
				{
					Kind:        sdk.String(sdk.WhateverUnionKind.StringField),
					StringField: sdk.String(""),
				},
			},
		},
	},
}

var req = sdk.GetWhateversRequest{
	Select: &sdk.GetWhateversResponseSelect{
		Whatevers: &sdk.WhateverSelect{
			StringField:  sdk.Bool(true),
			Int32Field:   sdk.Bool(true),
			Float64Field: sdk.Bool(true),
			BoolField:    sdk.Bool(true),
			EnumField:    sdk.Bool(true),
		},
	},
}

var ctx = context.Background()

func NewHandler(t *testing.T, f generic.Factory) (func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic)) {
	return func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic) {
		var req0 sdk.GetWhateversRequest
		err := gReq.ToStruct(&req0)
		if err != nil {
			t.Error(err)

			return
		}

		assert.Equal(t, req0, req)

		return f.MustFromStruct(rsp)
	}
}

func TestGenericClientGenericServer(t *testing.T) {
	t.Parallel()

	err := func() (err error) {
		addr := "127.0.0.1:57004"
		s := NewServer(root, f, NewHandler(t, f), addr)

		go func() {
			err := s.Listen()
			if err != nil {
				t.Error(err)
			}
		}()

		c := NewClient(f, &http.Client{}, "http://"+addr, "")

		gRsp, err := c.Send(f.MustFromStruct(req))
		if err != nil {
			return
		}

		var rsp0 sdk.GetWhateversResponse
		err = gRsp.ToStruct(&rsp0)
		if err != nil {
			return
		}

		assert.Equal(t, rsp, rsp0)

		err = s.Close(ctx)
		if err != nil {
			return
		}

		return
	}()
	if err != nil {
		t.Error(err)
	}
}
