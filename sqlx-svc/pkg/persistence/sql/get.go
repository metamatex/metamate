package sql

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/gen/v0/sdk"
)

func ResolveAlternativeId(supportedIdKinds map[string]bool, typeName string, alternativeId sdk.Id) (q string, values []interface{}, err error) {
	var value string
	switch *alternativeId.Kind {
	case sdk.IdKind.Email:
		value = *alternativeId.Email.Value
	case sdk.IdKind.Name:
		value = *alternativeId.Name
	case sdk.IdKind.Username:
		value = *alternativeId.Username
	case sdk.IdKind.Ean:
		value = *alternativeId.Ean
	case sdk.IdKind.Url:
		value = *alternativeId.Url.Value
	default:
		err = errors.New(fmt.Sprintf("kind %v not supported", *alternativeId.Kind))
	}

	_, ok := supportedIdKinds[*alternativeId.Kind]
	if !ok {
		err = errors.New(fmt.Sprintf("kind %v not supported", *alternativeId.Kind))

		return
	}

	q = `SELECT id_value FROM ` + typeName + *alternativeId.Kind + ` WHERE value = ?`

	values = append(values, value)

	return
}

func GetAlternativeId(supportedIdKinds map[string]bool, typeName string, idKind string, id string) (q string, values []interface{}, err error) {
	_, ok := supportedIdKinds[idKind]
	if !ok {
		err = errors.New(fmt.Sprintf("kind %v not supported", idKind))

	    return
	}

	q = `SELECT value FROM ` + typeName + idKind + ` WHERE id_value = ?`

	values = append(values, id)

	return
}

func GetById(typeName string, id string) (q string, values []interface{}) {
	q = `SELECT * FROM ` + typeName + ` WHERE id_value = ?`
	values = []interface{}{id}

	return
}

func Get(tableName string, gFilter generic.Generic) (q string, values []interface{}, err error) {
	q = `SELECT * FROM ` + tableName + ` `

	if gFilter != nil {
		q0, values0 := Filter(gFilter)
		q += " WHERE" + q0
		values = append(values, values0...)
	}

	q, values, err = sqlx.In(q, values...)
	if err != nil {
		return
	}

	return
}
