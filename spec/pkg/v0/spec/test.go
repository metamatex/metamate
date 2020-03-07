package spec

import (
	"context"
	"fmt"
	"github.com/metamatex/metamatemono/generic/pkg/v0/generic"
	"github.com/metamatex/metamatemono/gen/v0/sdk"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamatemono/gen/v0/sdk/utils/ptr"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type T struct {
	Errors []error
	Failed bool
}

func (t *T) Errorf(format string, args ...interface{}) {
	t.Failed = true
	t.Errors = append(t.Errors, errors.New(fmt.Sprintf(format, args)))
}

func requirePostRsp(t *testing.T, gRsp generic.Generic) {
	require.NotNil(t, gRsp)
	requireNoRspMetaErrs(t, gRsp)
	requireHasGenerics(t, gRsp)
}

func requireGetRsp(t *testing.T, gRsp generic.Generic) {
	require.NotNil(t, gRsp)
	requireNoRspMetaErrs(t, gRsp)
	requireHasGenerics(t, gRsp)
}

func requireSvcIdValues(t *testing.T, gSlice generic.Slice, expectedValues []string) {
	actualValues := map[string]bool{}

	for _, g := range gSlice.Get() {
		v := g.MustString(fieldnames.Id, fieldnames.Value)

		actualValues[v] = true
	}

	for _, v := range expectedValues {
		_, ok := actualValues[v]

		assert.True(t, ok, fmt.Sprintf("expected id.value %v to be present", v))
	}
}

func requireEmptyGetRsp(t *testing.T, gRsp generic.Generic) {
	require.NotNil(t, gRsp)
	requireNoRspMetaErrs(t, gRsp)
	requireNoGenerics(t, gRsp)
}

func GetCollectionMetaSelect() (*sdk.CollectionMetaSelect) {
	return &sdk.CollectionMetaSelect{
		Errors: &sdk.ErrorSelect{
			Kind: ptr.Bool(true),
			Service: &sdk.ServiceSelect{
				Name: ptr.Bool(true),
			},
			Message: &sdk.TextSelect{
				Formatting: ptr.Bool(true),
				Value: ptr.Bool(true),
			},
			Id: &sdk.IdSelect{
				Kind: ptr.Bool(true),
			},
		},
	}
}

func GetResponseMetaSelect() (*sdk.ResponseMetaSelect) {
	return &sdk.ResponseMetaSelect{
		Errors: &sdk.ErrorSelect{
			Kind: ptr.Bool(true),
			Service: &sdk.ServiceSelect{
				Name: ptr.Bool(true),
			},
			Message: &sdk.TextSelect{
				Formatting: ptr.Bool(true),
				Value: ptr.Bool(true),
			},
			Id: &sdk.IdSelect{
				Kind: ptr.Bool(true),
			},
		},
	}
}

func requireHasGenerics(t *testing.T, gRsp generic.Generic) {
	gSlice, _ := gRsp.GenericSlice(gRsp.Type().Edges.Type.For().PluralFieldName())
	require.NotNil(t, gSlice, "requireHasGenerics: gSlice is nil")

	require.NotEqual(t, 0, len(gSlice.Get()), "requireHasGenerics: len(gSlice.Get) is 0")
}

func requireNoGenerics(t *testing.T, gRsp generic.Generic) {
	gSlice, _ := gRsp.GenericSlice(gRsp.Type().Edges.Type.For().PluralFieldName())
	if gSlice == nil {
		require.Nil(t, gSlice, "requireHasGenerics: gSlice is nil")
	} else {
		require.Equal(t, 0, len(gSlice.Get()), "requireHasGenerics: len(gSlice.Get) is 0")
	}
}

func requireNoCollectionMetaErrors(t *testing.T, gRsp generic.Generic) {
	gRspMeta, ok := gRsp.Generic(fieldnames.Meta)
	if !ok {
		return
	}

	rspMeta := sdk.CollectionMeta{}
	gRspMeta.MustToStruct(&rspMeta)

	require.Len(t, rspMeta.Errors, 0)
}

func assertNRspMetaErrs(t *testing.T, gRsp generic.Generic, n int) (bool) {
	gErrs, ok := gRsp.GenericSlice(fieldnames.Meta, fieldnames.Errors)
	if !ok {
		return n == 0
	}

	return assert.Len(t, gErrs.Get(), n, gErrs.Sprint())
}

func requireNoRspMetaErrs(t *testing.T, gRsp generic.Generic) {
	gErrs, ok := gRsp.GenericSlice(fieldnames.Meta, fieldnames.Errors)
	if !ok {
	    return
	}

	require.Len(t, gErrs.Get(), 0, gErrs.Sprint())
}

func requireSvcHasSvcId(t *testing.T, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), f generic.Factory, ctx context.Context, svcName string, svcIdValue string) (err error) {
	getReq := sdk.GetWhateversRequest{
		ServiceFilter: &sdk.ServiceFilter{
			Id: &sdk.ServiceIdFilter{
				Value: &sdk.StringFilter{
					Is: ptr.String(svcName),
				},
			},
		},
		Mode: &sdk.GetMode{
			Kind: &sdk.GetModeKind.Id,
			Id: &sdk.Id{
				Kind: &sdk.IdKind.ServiceId,
				ServiceId: &sdk.ServiceId{
					Value: ptr.String(svcIdValue),
				},
			},
		},
		Select: &sdk.GetWhateversResponseSelect{
			Meta: GetCollectionMetaSelect(),
			Whatevers: &sdk.WhateverSelect{
				Id: &sdk.ServiceIdSelect{
					Value: ptr.Bool(true),
				},
			},
		},
	}

	gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
	if err != nil {
	    return
	}

	requireGetRsp(t, gGetRsp)

	getRsp := sdk.GetWhateversResponse{}
	gGetRsp.MustToStruct(&getRsp)

	require.Len(t, getRsp.Whatevers, 1)
	assert.Equal(t, svcIdValue, *getRsp.Whatevers[0].Id.Value)

	return
}

func requireSvcHasNotSvcId(t *testing.T, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), f generic.Factory, ctx context.Context, svcName string, svcIdValue string) (err error) {
	getReq := sdk.GetWhateversRequest{
		ServiceFilter: &sdk.ServiceFilter{
			Id: &sdk.ServiceIdFilter{
				Value: &sdk.StringFilter{
					Is: ptr.String(svcName),
				},
			},
		},
		Mode: &sdk.GetMode{
			Kind: &sdk.GetModeKind.Id,
			Id: &sdk.Id{
				Kind: &sdk.IdKind.ServiceId,
				ServiceId: &sdk.ServiceId{
					Value: ptr.String(svcIdValue),
				},
			},
		},
		Select: &sdk.GetWhateversResponseSelect{
			Meta: GetCollectionMetaSelect(),
			Whatevers: &sdk.WhateverSelect{
				Id: &sdk.ServiceIdSelect{
					Value: ptr.Bool(true),
				},
			},
		},
	}

	gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
	if err != nil {
	    return
	}

	requireEmptyGetRsp(t, gGetRsp)

	return
}
