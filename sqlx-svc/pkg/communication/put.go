package communication

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/gen/v0/sdk"
)

func handlePut(ctx context.Context, rn *graph.RootNode, db sqlx.Ext, f generic.Factory, gReq generic.Generic) (gRsp generic.Generic, errs []error) {
	gMode := gReq.MustGeneric(fieldnames.Mode)

	mode := sdk.GetMode{}
	gMode.MustToStruct(&mode)

	gRsp = f.New(gReq.Type().Edges.Type.Response())

	switch *mode.Kind {
	case sdk.PutModeKind.Relation:
		gRsp, errs = handlePutRelation(ctx, rn, db, f, gReq)
		break
	default:
		errs = append(errs, errors.New(fmt.Sprintf("mode.kind %v not supported", *mode.Kind)))
	}

	return
}

func handlePutRelation(ctx context.Context, rn *graph.RootNode, db sqlx.Ext, f generic.Factory, gReq generic.Generic) (gRsp generic.Generic, errs []error) {
	gRsp = f.New(gReq.Type().Edges.Type.Response())

	putMode := sdk.PutMode{}
	gReq.MustGeneric(fieldnames.Mode).MustToStruct(&putMode)

	pn, err := rn.Paths.ByName(*putMode.Relation.Relation)
	if err != nil {
		errs = append(errs, err)

		return
	}

	rn0 := pn.Edges.Relation.BelongsTo()

	errs0 := putRelations(db, rn0.Name(), *putMode.Relation.Operation, pn.Data.IsActive, *putMode.Relation.Id, putMode.Relation.Ids)
	if len(errs0) != 0 {
		errs = append(errs, errs0...)

		return
	}

	return
}

func putRelations(db sqlx.Ext, relationName string, operation string, active bool, id sdk.ServiceId, ids []sdk.ServiceId) (errs []error) {
	for _, id0 := range ids {
		var activeId sdk.ServiceId
		var passiveId sdk.ServiceId
		if active {
			activeId = id
			passiveId = id0
		} else {
			activeId = id0
			passiveId = id
		}

		q, values, err := generatePutSql(relationName, operation, *activeId.Value, *passiveId.Value)
		if err != nil {
			errs = append(errs, err)

			return
		}

		_, err = db.Exec(q, values...)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return
}

func generatePutSql(relationName string, operation string, activeId, passiveId string) (sql string, values []interface{}, err error) {
	switch operation {
	case sdk.RelationOperation.Add:
		sql = "DELETE FROM " + relationName + " WHERE active_id_value=? AND passive_id_value=?"
	case sdk.RelationOperation.Add:
		sql = "INSERT INTO " + relationName + " (active_id_value, passive_id_value) VALUES (?, ?)"
	default:
		err = errors.New(fmt.Sprintf("relation operation %v unknown", operation))
		return
	}

	values = []interface{}{activeId, passiveId}

	return
}
