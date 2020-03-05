package communication

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/metamatex/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamatemono/pkg/generic/pkg/v0/generic"
	"github.com/metamatex/metamatemono/pkg/services/sqlx-svc/gen/v0/sdk"
	"github.com/metamatex/metamatemono/pkg/services/sqlx-svc/pkg/persistence"
)

func handleGet(ctx context.Context, supportedIdKinds map[string]bool, rn *graph.RootNode, db sqlx.Ext, f generic.Factory, gReq generic.Generic) (gRsp generic.Generic, errs []error) {
	gMode := gReq.MustGeneric(fieldnames.Mode)

	mode := sdk.GetMode{}
	gMode.MustToStruct(&mode)

	gRsp = f.New(gReq.Type().Edges.Type.Response())

	switch *mode.Kind {
	case sdk.GetModeKind.Collection:
		gRsp, errs = handleGetCollection(ctx, db, f, gReq)
		break
	case sdk.GetModeKind.Id:
		gRsp, errs = handleGetId(ctx, supportedIdKinds, db, f, gReq)
		break
	case sdk.GetModeKind.Relation:
		gRsp, errs = handleGetRelation(ctx, rn, db, f, gReq)
		break
	default:
		errs = append(errs, errors.New(fmt.Sprintf("mode.kind %v not supported", *mode.Kind)))
	}

	return
}

func handleGetCollection(ctx context.Context, db sqlx.Ext, f generic.Factory, gReq generic.Generic) (gRsp generic.Generic, errs []error) {
	gRsp = f.New(gReq.Type().Edges.Type.Response())

	gFilter, _ := gReq.Generic(fieldnames.Filter)

	gSlice, err := persistence.Get(db, f, gReq.Type().Edges.Type.For(), gFilter)
	if err != nil {
		errs = append(errs, err)

		return
	}

	gRsp.MustSetGenericSlice([]string{gReq.Type().Edges.Type.For().PluralFieldName()}, gSlice)

	return
}

func handleGetId(ctx context.Context, supportedIdKinds map[string]bool, db sqlx.Ext, f generic.Factory, gReq generic.Generic) (gRsp generic.Generic, errs []error) {
	gRsp = f.New(gReq.Type().Edges.Type.Response())

	gEntitySelect := gReq.MustGeneric(fieldnames.Select, gReq.Type().Edges.Type.For().PluralFieldName())

	id := sdk.Id{}
	gReq.MustGeneric(fieldnames.Mode, fieldnames.Id).MustToStruct(&id)

	g, err := persistence.GetById(supportedIdKinds, db, f, gEntitySelect, id)
	if err != nil {
		errs = append(errs, err)

		return
	}

	gSlice := f.NewSlice(gReq.Type().Edges.Type.For())

	if g != nil {
		gSlice.Append(g)
	}

	gRsp.MustSetGenericSlice([]string{gReq.Type().Edges.Type.For().PluralFieldName()}, gSlice)

	return
}

func handleGetRelation(ctx context.Context, rn *graph.RootNode, db sqlx.Ext, f generic.Factory, gReq generic.Generic) (gRsp generic.Generic, errs []error) {
	gRsp = f.New(gReq.Type().Edges.Type.Response())

	mode := &sdk.RelationGetMode{}
	gReq.MustGeneric(fieldnames.Mode, fieldnames.Relation).MustToStruct(&mode)

	pn, err := rn.Paths.ByName(*mode.Relation)
	if err != nil {
		errs = append(errs, err)

		return
	}

	rn0 := pn.Edges.Relation.BelongsTo()

	gSlice, err := persistence.GetRelations(db, f, gReq.Type().Edges.Type.For(), rn0.Name(), pn.Data.IsActive, *mode.Id)
	if err != nil {
		errs = append(errs, err)

		return
	}

	gRsp.MustSetGenericSlice([]string{gReq.Type().Edges.Type.For().PluralFieldName()}, gSlice)

	return
}
