package _go

import (
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/metamatex/metamate/metactl/pkg/v0/utils/ptr"
)

const (
	TaskHttpJson                 = "TaskHttpJson"
	TaskHttpJsonTypedGenericTest = "TaskHttpJsonTypedGenericTest"
)

func init() {
	tasks[TaskHttpJson] = types.RenderTask{
		TemplateData: &goHttpJsonTpl,
		Out:          ptr.String("httpjson_.go"),
	}

	tasks[TaskHttpJsonTypedGenericTest] = types.RenderTask{
		TemplateData: &goHttpJsonTypedGenericTestTpl,
		Out:          ptr.String("typed_generic_test.go"),
	}
}

var goHttpJsonTpl = `package mql
const (
	MetamateTypeHeader = "X-MetaMate-Type"
	ContentTypeJson = "application/json; charset=utf-8"
	ContentTypeHeader = "Content-type"
	AuthorizationHeader = "Authorization"
)
`

var goHttpJsonTypedGenericTestTpl = `package mql

import (
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
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

		rsp0, err := c.GetDummies(ctx, req)
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

		s.SetGetDummiesEndpoint(NewEndpoint(t))

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

		var rsp0 types.GetDummiesResponse
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
