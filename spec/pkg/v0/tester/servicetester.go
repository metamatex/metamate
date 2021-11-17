package tester

import (
	"context"
	"errors"
	"fmt"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/spec/pkg/v0/mtesting"
	"github.com/stretchr/testify/assert"
)

type BusTester struct {
	print bool
	t     mtesting.TB
	ctx   context.Context
	rn    *graph.RootNode
	f     generic.Factory
	h     func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)
}

func NewBusTester(print bool, t mtesting.TB, ctx context.Context, f generic.Factory, rn *graph.RootNode, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) BusTester {
	return BusTester{
		print: print,
		t:     t,
		ctx:   ctx,
		rn:    rn,
		f:     f,
		h:     h,
	}
}

func (tester BusTester) NewSubTester(print bool, name string, f func(tester BusTester)) {
	tester.t.Run(name, func(t mtesting.TB) {
		tester.t = t
		tester.print = print
		f(tester)
	})
}

func (tester BusTester) send(req interface{}) (gRsp generic.Generic, err error) {
	return tester.h(tester.ctx, tester.f.MustFromStruct(req))
}

func (tester BusTester) sendG(gReq generic.Generic) (gRsp generic.Generic, err error) {
	return tester.h(tester.ctx, gReq)
}

func (tester BusTester) TestGetSearch(print bool, typeName string, term string, svcFilter *mql.ServiceFilter, relations []string, pages int) {
	mode := mql.GetMode{
		Kind: &mql.GetModeKind.Search,
		Search: &mql.SearchGetMode{
			Term: &term,
		},
	}

	tester.Test(print, "TestGet%vSearch", typeName, mode, svcFilter, relations, pages)
}

func (tester BusTester) TestGetRelation(print bool, typeName string, id mql.Id, relation string, svcFilter *mql.ServiceFilter, relations []string, pages int) {
	mode := mql.GetMode{
		Kind: &mql.GetModeKind.Relation,
		Relation: &mql.RelationGetMode{
			Id:       &id,
			Relation: &relation,
		},
	}

	tester.Test(print, "TestGet%vRelation", typeName, mode, svcFilter, relations, pages)
}

func (tester BusTester) TestGetCollection(print bool, typeName string, svcFilter *mql.ServiceFilter, relations []string, pages int) {
	mode := mql.GetMode{
		Kind:       &mql.GetModeKind.Collection,
		Collection: &mql.CollectionGetMode{},
	}

	tester.Test(print, "TestGet%vCollection", typeName, mode, svcFilter, relations, pages)
}

func (tester BusTester) TestGetId(print bool, typeName string, id mql.Id, svcFilter *mql.ServiceFilter, relations []string) {
	mode := mql.GetMode{
		Kind: &mql.GetModeKind.Id,
		Id:   &id,
	}

	tester.Test(print, "TestGet%vId", typeName, mode, svcFilter, relations, 1)
}

func (tester BusTester) Test(print bool, testNameFormat string, typeName string, mode mql.GetMode, svcFilter *mql.ServiceFilter, relations []string, pages int) {
	tn, err := tester.rn.Types.ByName(typeName)
	if err != nil {
		tester.t.Error(err)

		return
	}

	name := fmt.Sprintf(testNameFormat, tn.PluralName())
	tester.t.Run(name, func(t mtesting.TB) {
		t.Parallel()

		err := func() (err error) {
			var next []mql.ServicePage
			serviceIdMap := map[string]bool{}
			for i := 0; i < pages; i++ {
				gReq := tester.f.New(tn.Edges.Type.GetClientRequest())
				gReq.MustSetGeneric([]string{mql.FieldNames.Mode}, tester.f.MustFromStruct(mode))

				if svcFilter != nil {
					gReq.MustSetGeneric([]string{mql.FieldNames.ServiceFilter}, tester.f.MustFromStruct(*svcFilter))
				}

				relationsTn := tn.Edges.Type.GetRelations()

				if len(relations) != 0 && relationsTn == nil {
					err = errors.New(fmt.Sprintf("relations %v set but no relations type node", relations))

					return
				}

				for _, r := range relations {
					gReq.MustSetGeneric([]string{mql.FieldNames.Relations, r}, tester.f.New(relationsTn.Edges.Fields.Holds().MustByName(r).Edges.Type.Holds()))
				}

				if len(next) != 0 {
					gReq.MustSetGenericSlice([]string{mql.FieldNames.Pages}, tester.f.MustFromStructs(next))
				}

				var gGetRsp generic.Generic
				gGetRsp, err = tester.sendG(gReq)
				if err != nil {
					return
				}
				if !assert.True(t, gGetRsp.Type().Flags().Is(typeflags.IsBusResponse, true), "response %v is not a BusResponse", gGetRsp.Type().Name()) {
					return
				}

				if tester.print || print {
					gGetRsp.Print()
				}

				err = getErrors(gGetRsp)
				if err != nil {
					return
				}

				gEntities, ok := gGetRsp.GenericSlice(tn.PluralFieldName())
				if !assert.True(t, ok) {
					return
				}

				if !assert.True(t, len(gEntities.Get()) != 0) {
					return
				}

				for _, gEntity := range gEntities.Get() {
					for _, r := range relations {
						gRelationEntities, ok := gEntity.GenericSlice(mql.FieldNames.Relations, r, relationsTn.Edges.Fields.Holds().MustByName(r).Edges.Type.Holds().Edges.Type.For().PluralFieldName())
						if !assert.True(t, ok, fmt.Sprintf("relation %v is empty[0]", r)) {
							continue
						}

						if !assert.NotEmpty(t, gRelationEntities.Get(), fmt.Sprintf("relation %v is empty[1]", r)) {
							continue
						}
					}
				}

				if pages > 1 {
					currentServiceIdMap := getServiceIdMap(gEntities)
					serviceIdMap, err = mergeUniqueMaps(serviceIdMap, currentServiceIdMap)
					if err != nil {
						return
					}

					if i+1 < pages {
						gServicePages, ok := gGetRsp.GenericSlice(mql.FieldNames.Pagination, mql.FieldNames.Next)
						if !ok {
							err = errors.New(fmt.Sprintf("expected pagination.next"))

							return
						}

						gServicePages.MustToStructs(&next)
					}
				}
			}
			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func getServiceIdMap(gSlice generic.Slice) (m map[string]bool) {
	m = map[string]bool{}

	for _, g := range gSlice.Get() {
		var svcId mql.ServiceId
		g.MustGeneric(mql.FieldNames.Id).MustToStruct(&svcId)

		m[*svcId.ServiceName+"/"+*svcId.Value] = true
	}

	return
}

func mergeUniqueMaps(a, b map[string]bool) (map[string]bool, error) {
	for k, _ := range b {
		_, ok := a[k]
		if ok {
			return nil, errors.New(fmt.Sprintf("duplicate key: %v", k))
		}

		a[k] = true
	}

	return a, nil
}

func getErrors(gRsp generic.Generic) (err error) {
	gErrs, ok := gRsp.GenericSlice(mql.FieldNames.Errors)
	if ok {
		var errs []mql.Error
		gErrs.MustToStructs(&errs)

		var msg string
		for _, e := range errs {
			var svc string
			if e.Service != nil && e.Service.Name != nil {
				svc = *e.Service.Name + ": "
			}

			var kind string
			if e.Kind != nil {
				kind = *e.Kind
			}
			var message string
			if e.Message != nil {
				message = *e.Message
			}
			msg = msg + fmt.Sprintf("%v%v %v\n", svc, kind, message)
		}

		err = errors.New(msg)

		return
	}

	return
}
