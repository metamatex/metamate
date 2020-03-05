package _go

import (
	"github.com/metamatex/metamatemono/pkg/metactl/pkg/v0/utils/ptr"
	"github.com/metamatex/metamatemono/pkg/metactl/pkg/v0/types"
)

const (
	TaskHttpJson                 = "TaskHttpJson"
	TaskHttpJsonTypedGenericTest = "TaskHttpJsonTypedGenericTest"
)

func init() {
	tasks[TaskHttpJson] = types.RenderTask{
		TemplateData: &goHttpJsonTpl,
		Out:          ptr.String("transport/httpjson_.go"),
	}

	tasks[TaskHttpJsonTypedGenericTest] = types.RenderTask{
		TemplateData: &goHttpJsonTypedGenericTestTpl,
		Out:          ptr.String("transport/typed_generic_test.go"),
	}
}

var goHttpJsonTpl = `package transport
const (
	METAMATE_TYPE_HEADER = "X-MetaMate-Type"
	CONTENT_TYPE_JSON = "application/json; charset=utf-8"
	CONTENT_TYPE_HEADER = "Content-type"
	AUTHORIZATION_HEADER = "Authorization"
)
`

var goHttpJsonTypedGenericTestTpl = `package httpjson
{{ $package := index .Data "package" }}
import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"{{ $package }}/gen/v0/sdk"
)

func TestTypedClientGenericServer(t *testing.T) {
	t.Parallel()

	err := func() (err error) {
		addr := "127.0.0.1:57002"
		s := NewGenericServer(addr, root, NewHandler(t))

		go func() {
			err := s.Listen()
			if err != nil {
				t.Error(err)
			}
		}()

		c := NewTypedClient(&http.Client{}, "http://"+addr)

		rsp0, err := c.GetWhatevers(ctx, req)
		if err != nil {
			return
		}

		assert.Equal(t, rsp, *rsp0)

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

func TestGenericClientTypedServer(t *testing.T) {
	t.Parallel()

	err := func() (err error) {
		addr := "127.0.0.1:57003"
		s := NewTypedServer(addr)

		s.SetGetWhateversEndpoint(NewEndpoint(t))

		go func() {
			err := s.Listen()
			if err != nil {
				t.Error(err)
			}
		}()

		c := NewClient(&http.Client{}, "http://"+addr, "")

		gRsp, err := c.Send(mapper.StructToGeneric(req))
		if err != nil {
			return
		}

		var rsp0 types.GetWhateversResponse
		err = types.GenericToStruct(gRsp, &rsp0)
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
}`
