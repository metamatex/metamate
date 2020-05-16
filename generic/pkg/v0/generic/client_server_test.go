package generic

import (
	"context"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/expansion"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var rsp = mql.GetDummiesResponse{
	Dummies: []mql.Dummy{
		{
			StringField:  mql.String(""),
			Int32Field:   mql.Int32(0),
			Float64Field: mql.Float64(0),
			BoolField:    mql.Bool(false),
			EnumField:    mql.String(mql.DummyKind.Red),
			UnionField: &mql.DummyUnion{
				Kind:        mql.String(mql.DummyUnionKind.StringField),
				StringField: mql.String(""),
			},
			StringList:  []string{""},
			Int32List:   []int32{0},
			Float64List: []float64{0},
			BoolList:    []bool{false},
			EnumList:    []string{mql.DummyKind.Red},
			UnionList: []mql.DummyUnion{
				{
					Kind:        mql.String(mql.DummyUnionKind.StringField),
					StringField: mql.String(""),
				},
			},
		},
	},
}

var req = mql.GetDummiesRequest{
	Select: &mql.GetDummiesResponseSelect{
		Dummies: &mql.DummySelect{
			StringField:  mql.Bool(true),
			Int32Field:   mql.Bool(true),
			Float64Field: mql.Bool(true),
			BoolField:    mql.Bool(true),
			EnumField:    mql.Bool(true),
		},
	},
}

var ctx = context.Background()

func NewHandler(t *testing.T, f Factory) func(ctx context.Context, gReq Generic) (gRsp Generic) {
	return func(ctx context.Context, gReq Generic) (gRsp Generic) {
		var req0 mql.GetDummiesRequest
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
		root := graph.NewRoot()

		err = expansion.Expand(0, root)
		if err != nil {
			return
		}

		f := NewFactory(root)

		addr := "127.0.0.1:57004"
		s := NewServer(ServerOpts{Root: root, Factory: f, Handler: NewHandler(t, f)})

		go func() {
			err = http.ListenAndServe(addr, s)
			if err != nil {
				t.Error(err)
			}
		}()

		c := NewClient(f, &http.Client{}, "http://"+addr, "")

		gRsp, err := c.Send(f.MustFromStruct(req))
		if err != nil {
			return
		}

		var rsp0 mql.GetDummiesResponse
		err = gRsp.ToStruct(&rsp0)
		if err != nil {
			return
		}

		assert.Equal(t, rsp, rsp0)

		return
	}()
	if err != nil {
		t.Error(err)
	}
}
