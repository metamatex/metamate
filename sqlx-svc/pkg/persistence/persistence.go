package persistence

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamatemono/generic/pkg/v0/generic"
	"github.com/metamatex/metamatemono/gen/v0/sdk"
	"github.com/metamatex/metamatemono/gen/v0/sdk/utils/ptr"
	"github.com/metamatex/metamatemono/sqlx-svc/pkg/persistence/sql"
)

func Insert(db sqlx.Ext, g generic.Generic) (id string, err error) {
	id, ok := g.String(fieldnames.Id, fieldnames.Value)
	if !ok {
		var u uuid.UUID
		u, err = uuid.NewUUID()
		if err != nil {
			return
		}

		id = u.String()

		g.MustSetString([]string{fieldnames.Id, fieldnames.Value}, id)
	}

	q, values, err := sql.InsertOne(g)
	if err != nil {
		return
	}

	_, err = db.Exec(q, values...)
	if err != nil {
		return
	}

	gSlice, ok := g.GenericSlice(fieldnames.AlternativeIds)
	if ok {
		for _, gId := range gSlice.Get() {
			var alternativeId = sdk.Id{}
			err = gId.ToStruct(&alternativeId)
			if err != nil {
				return
			}

			var q string
			var values []interface{}
			q, values, err = sql.InsertAlternativeId(g.Type().Name(), id, alternativeId)
			if err != nil {
				return
			}

			_, err = db.Exec(q, values...)
			if err != nil {
				return
			}
		}
	}

	return
}

func InsertSlice(db sqlx.Ext, gSlice generic.Slice) (id int64, err error) {
	for _, g := range gSlice.Get() {
		_, err = Insert(db, g)
		if err != nil {
			return
		}
	}

	return
}

func Get(db sqlx.Ext, f generic.Factory, tn *graph.TypeNode, gFilter generic.Generic) (gSlice generic.Slice, err error) {
	q, values, err := sql.Get(tn.Name(), gFilter)
	if err != nil {
		return
	}

	rows, err := db.Queryx(q, values...)
	if err != nil {
		return
	}
	defer rows.Close()

	ms := []map[string]interface{}{}
	for rows.Next() {
		m := map[string]interface{}{}
		err = rows.MapScan(m)
		if err != nil {
			return
		}

		ms = append(ms, m)
	}

	gSlice, err = f.UnflattenSlice(tn, "_", ms)
	if err != nil {
		return
	}

	return
}

func GetById(supportedIdKinds map[string]bool, db sqlx.Ext, f generic.Factory, gEntitySelect generic.Generic, id sdk.Id) (g generic.Generic, err error) {
	tn := gEntitySelect.Type().Edges.Type.For()
	typeName := tn.Name()

	var id0 string
	switch *id.Kind {
	case sdk.IdKind.ServiceId:
		id0 = *id.ServiceId.Value
	default:
		var q string
		var values []interface{}
		q, values, err = sql.ResolveAlternativeId(supportedIdKinds, typeName, id)
		if err != nil {
			return
		}

		err = sqlx.Get(db, &id0, q, values...)
		if err != nil {
			return
		}
	}

	q, values := sql.GetById(typeName, id0)

	r := db.QueryRowx(q, values...)
	if r.Err() != nil {
		return
	}

	m := map[string]interface{}{}
	err = r.MapScan(m)
	if err != nil {
		err = nil
		return
	}

	g, err = f.Unflatten(tn, "_", m)
	if err != nil {
		return
	}

	gIdSelect, ok := gEntitySelect.Generic(fieldnames.AlternativeIds)
	if ok {
		var alternativeIds []sdk.Id
		alternativeIds, err = GetAlternativeIds(supportedIdKinds, db, typeName, gIdSelect, id0)
		if err != nil {
			return
		}

		if len(alternativeIds) != 0 {
			g.MustSetGenericSlice([]string{fieldnames.AlternativeIds}, f.MustFromStructs(alternativeIds))
		}
	}

	return
}

func GetAlternativeIds(supportedIdKinds map[string]bool, db sqlx.Ext, typeName string, gIdSelect generic.Generic, id string) (ids []sdk.Id, err error) {
	kinds := []string{}
	gIdSelect.EachBool(func(fn *graph.FieldNode, b bool) {
		b, ok := supportedIdKinds[fn.Name()]
		if ok && b {
			kinds = append(kinds, fn.Name())
		}
	})

	gIdSelect.EachGeneric(func(fn *graph.FieldNode, b generic.Generic) {
		_, ok := supportedIdKinds[fn.Name()]
		if ok {
			kinds = append(kinds, fn.Name())
		}
	})

	for _, k := range kinds {
		var q string
		var values []interface{}
		q, values, err = sql.GetAlternativeId(supportedIdKinds, typeName, k, id)
		if err != nil {
			return
		}

		var rows *sqlx.Rows
		rows, err = db.Queryx(q, values...)
		if err != nil {
			return
		}
		defer rows.Close()

		ss := []string{}
		for rows.Next() {
			var s string
			err = rows.Scan(&s)
			if err != nil {
				return
			}
			ss = append(ss, s)
		}

		for _, s := range ss {
			var id sdk.Id
			switch k {
			case sdk.IdKind.Email:
				id = sdk.Id{
					Kind: ptr.String(k),
					Email: &sdk.Email{
						Value: ptr.String(s),
					},
				}
			case sdk.IdKind.Name:
				id = sdk.Id{
					Kind: ptr.String(k),
					Name: ptr.String(s),
				}
			case sdk.IdKind.Username:
				id = sdk.Id{
					Kind:     ptr.String(k),
					Username: ptr.String(s),
				}
			case sdk.IdKind.Ean:
				id = sdk.Id{
					Kind: ptr.String(k),
					Ean:  ptr.String(s),
				}
			case sdk.IdKind.Url:
				id = sdk.Id{
					Kind: ptr.String(k),
					Url: &sdk.Url{
						Value: ptr.String(s),
					},
				}
			default:
				err = errors.New(fmt.Sprintf("id.kind %v not supported", k))

				return
			}

			ids = append(ids, id)
		}
	}

	return
}

func Migrate(supoortedIdKinds map[string]bool, db sqlx.Ext, tnm graph.TypeNodeMap, rnm graph.RelationNodeMap) (err error) {
	q, err := sql.Create(supoortedIdKinds, tnm, rnm)
	if err != nil {
		return
	}

	_, err = db.Exec(q)
	if err != nil {
		return
	}

	return
}

func GetRelations(db sqlx.Ext, f generic.Factory, tn *graph.TypeNode, relationName string, isActive bool, id sdk.ServiceId) (gSlice generic.Slice, err error) {
	q, values := sql.GetRelation(tn.Name(), relationName, isActive, *id.Value)

	rows, err := db.Queryx(q, values...)
	if err != nil {
		return
	}
	defer rows.Close()

	ms := []map[string]interface{}{}
	for rows.Next() {
		m := map[string]interface{}{}
		err = rows.MapScan(m)
		if err != nil {
			return
		}

		ms = append(ms, m)
	}

	gSlice, err = f.UnflattenSlice(tn, "_", ms)
	if err != nil {
		return
	}

	return
}
