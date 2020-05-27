package boot

import (
	"context"
	"errors"
	"fmt"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Tester struct {
	print bool
	t     *testing.T
	ctx   context.Context
	rn    *graph.RootNode
	f     generic.Factory
	h     func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)
}

func NewTester(print bool, t *testing.T, ctx context.Context, f generic.Factory, rn *graph.RootNode, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) Tester {
	return Tester{
		print: print,
		t:     t,
		ctx:   ctx,
		rn:    rn,
		f:     f,
		h:     h,
	}
}

func (tester Tester) NewSubTester(print bool, name string, f func(tester Tester)) {
	tester.t.Run(name, func(t *testing.T) {
		tester.t = t
		tester.print = print
		f(tester)
	})
}

func (tester Tester) send(req interface{}) (gRsp generic.Generic, err error) {
	return tester.h(tester.ctx, tester.f.MustFromStruct(req))
}

func (tester Tester) sendG(gReq generic.Generic) (gRsp generic.Generic, err error) {
	return tester.h(tester.ctx, gReq)
}

func (tester Tester) TestGetId(print bool, typeName string, id mql.Id, relations []string) {
	tn, err := tester.rn.Types.ByName(typeName)
	if err != nil {
		tester.t.Error(err)

		return
	}

	name := fmt.Sprintf("TestGet%vId", tn.PluralName())
	tester.t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			tn, err := tester.rn.Types.ByName(typeName)
			if err != nil {
				return
			}

			gReq := tester.f.New(tn.Edges.Type.GetRequest())
			gReq.MustSetGeneric([]string{mql.FieldNames.Mode}, tester.f.MustFromStruct(mql.GetMode{
				Kind: &mql.GetModeKind.Id,
				Id:   &id,
			}))

			relationsTn := tn.Edges.Type.GetRelations()

			if len(relations) != 0 && relationsTn == nil {
				err = errors.New(fmt.Sprintf("relations %v set but no relations type node", relations))

				return
			}

			for _, r := range relations {
				gReq.MustSetGeneric([]string{mql.FieldNames.Relations, r}, tester.f.New(relationsTn.Edges.Fields.Holds().MustByName(r).Edges.Type.Holds()))
			}

			gGetRsp, err := tester.sendG(gReq)
			if err != nil {
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

			if id.ServiceId != nil {
				if !assert.True(t, len(gEntities.Get()) == 1) {
					return
				}
			} else {
				if !assert.True(t, len(gEntities.Get()) != 0) {
					return
				}
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

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func (tester Tester) TestGetRelation(print bool, typeName string, id mql.Id, relation string, relations []string, flipPages int) {
	tn, err := tester.rn.Types.ByName(typeName)
	if err != nil {
		tester.t.Error(err)

		return
	}

	name := fmt.Sprintf("TestGet%vRelation", tn.PluralName())
	tester.t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			var next []mql.ServicePage
			serviceIdMap := map[string]bool{}
			for i := 0; i < flipPages; i++ {
				gReq := tester.f.New(tn.Edges.Type.GetRequest())
				gReq.MustSetGeneric([]string{mql.FieldNames.Mode}, tester.f.MustFromStruct(mql.GetMode{
					Kind: &mql.GetModeKind.Relation,
					Relation: &mql.RelationGetMode{
						Id:       &id,
						Relation: &relation,
					},
				}))

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

				if flipPages > 1 {
					currentServiceIdMap := getServiceIdMap(gEntities)
					serviceIdMap, err = mergeUniqueMaps(serviceIdMap, currentServiceIdMap)
					if err != nil {
						return
					}

					gServicePages, ok := gGetRsp.GenericSlice(mql.FieldNames.Pagination, mql.FieldNames.Next)
					if !ok {
						err = errors.New(fmt.Sprintf("expected pagination.next"))

						return
					}

					gServicePages.MustToStructs(&next)
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

func (tester Tester) TestGetSearch(print bool, typeName string, term string, svcFilter *mql.ServiceFilter, relations []string) {
	tn, err := tester.rn.Types.ByName(typeName)
	if err != nil {
		tester.t.Error(err)

		return
	}

	name := fmt.Sprintf("TestGet%vSearch%v", tn.PluralName(), term)
	tester.t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			tn, err := tester.rn.Types.ByName(typeName)
			if err != nil {
				return
			}

			gReq := tester.f.New(tn.Edges.Type.GetRequest())
			gReq.MustSetGeneric([]string{mql.FieldNames.Mode}, tester.f.MustFromStruct(mql.GetMode{
				Kind: &mql.GetModeKind.Search,
				Search: &mql.SearchGetMode{
					Term: &term,
				},
			}))

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

			gGetRsp, err := tester.sendG(gReq)
			if err != nil {
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

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func (tester Tester) TestGetCollection(print bool, typeName string, svcFilter *mql.ServiceFilter, relations []string) {
	tn, err := tester.rn.Types.ByName(typeName)
	if err != nil {
		tester.t.Error(err)

		return
	}

	name := fmt.Sprintf("TestGet%vCollection", tn.PluralName())
	tester.t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			gReq := tester.f.New(tn.Edges.Type.GetRequest())
			gReq.MustSetGeneric([]string{mql.FieldNames.Mode}, tester.f.MustFromStruct(mql.GetMode{
				Kind:       &mql.GetModeKind.Collection,
				Collection: &mql.CollectionGetMode{},
			}))

			if svcFilter != nil {
				gReq.MustSetGeneric([]string{mql.FieldNames.ServiceFilter}, tester.f.MustFromStruct(*svcFilter))
			}

			relationsTn := tn.Edges.Type.GetRelations()
			if relationsTn != nil {
				for _, r := range relations {
					gReq.MustSetGeneric([]string{mql.FieldNames.Relations, r}, tester.f.New(tn.Edges.Type.GetCollection()))
				}
			}

			gGetRsp, err := tester.sendG(gReq)
			if err != nil {
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
					gRelationEntities, ok := gEntity.GenericSlice(mql.FieldNames.Relations, r, tn.PluralFieldName())
					if !assert.True(t, ok) {
						continue
					}

					if !assert.NotEmpty(t, gRelationEntities.Get()) {
						continue
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
