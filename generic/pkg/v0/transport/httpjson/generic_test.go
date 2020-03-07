package httpjson

import (
	"context"
	"github.com/metamatex/metamatemono/generic/pkg/v0/generic"
	"net/http"
	"testing"

	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/expansion"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph"
	"github.com/stretchr/testify/assert"

	"github.com/metamatex/metamatemono/gen/v0/sdk"
	"github.com/metamatex/metamatemono/gen/v0/sdk/utils/ptr"
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
			StringField:  ptr.String(""),
			Int32Field:   ptr.Int32(0),
			Float64Field: ptr.Float64(0),
			BoolField:    ptr.Bool(false),
			EnumField:    ptr.String(sdk.WhateverKind_RED),
			UnionField: &sdk.WhateverUnion{
				Kind:        ptr.String(sdk.WhateverUnionKind_stringField),
				StringField: ptr.String(""),
			},
			StringList:  []string{""},
			Int32List:   []int32{0},
			Float64List: []float64{0},
			BoolList:    []bool{false},
			EnumList:    []string{sdk.WhateverKind_RED},
			UnionList: []sdk.WhateverUnion{
				{
					Kind:        ptr.String(sdk.WhateverUnionKind_stringField),
					StringField: ptr.String(""),
				},
			},
		},
	},
}

var req = sdk.GetWhateversRequest{
	Select: &sdk.GetWhateversResponseSelect{
		Whatevers: &sdk.WhateverSelect{
			StringField:  ptr.Bool(true),
			Int32Field:   ptr.Bool(true),
			Float64Field: ptr.Bool(true),
			BoolField:    ptr.Bool(true),
			EnumField:    ptr.Bool(true),
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
