package sql

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/metamatex/metamatemono/pkg/generic/pkg/v0/generic"
	"github.com/metamatex/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/asg/pkg/v0/asg/graph"
	"github.com/metamatex/asg/pkg/v0/asg/graph/fieldflags"
	"github.com/metamatex/metamatemono/pkg/services/sqlx-svc/gen/v0/sdk"
	"strings"
)

func InsertOne(g generic.Generic) (q string, values []interface{}, err error) {
	g = g.Copy()

	g.MustDelete(fieldnames.AlternativeIds)
	g.MustDelete(fieldnames.Meta, fieldnames.Service)

	g.WalkDelete(func(fn *graph.FieldNode) bool {
		return fn.Flags().Is(fieldflags.IsHash, true)
	})

	_, ok := g.Generic(fieldnames.Relations)
	if ok {
		g.MustDelete(fieldnames.Relations)
	}

	_, ok = g.Generic(fieldnames.Relationships)
	if ok {
		g.MustDelete(fieldnames.Relationships)
	}

	f, err := g.Flatten("_")
	if err != nil {
		return
	}

	b := bytes.Buffer{}
	b.WriteString("INSERT INTO ")
	b.WriteString(g.Type().Name())
	b.WriteString(" (")

	ks := []string{}
	values = []interface{}{}
	for k, v := range f {
		ks = append(ks, k)
		values = append(values, v)
	}

	for i := 0; i < len(ks)-1; i++ {
		b.WriteString(ks[i])
		b.WriteString(",")
	}

	b.WriteString(ks[len(ks)-1])
	b.WriteString(") VALUES (")
	b.WriteString(strings.Repeat("?,", len(ks)-1))
	b.WriteString("?)")

	q = b.String()

	return
}

func InsertAlternativeId(typeName string, id string, alternativeId sdk.Id) (q string, values []interface{}, err error) {
	var kind string
	var value string
	switch *alternativeId.Kind {
	case sdk.IdKind.Email:
		kind = sdk.IdKind.Email
		value = *alternativeId.Email.Value
	case sdk.IdKind.Name:
		kind = sdk.IdKind.Name
		value = *alternativeId.Name
	case sdk.IdKind.Username:
		kind = sdk.IdKind.Username
		value = *alternativeId.Username
	case sdk.IdKind.Ean:
		kind = sdk.IdKind.Ean
		value = *alternativeId.Ean
	case sdk.IdKind.Url:
		kind = sdk.IdKind.Url
		value = *alternativeId.Url.Value
	default:
		err = errors.New(fmt.Sprintf("kind %v not supported", *alternativeId.Kind))

		return
	}

	q = `INSERT INTO ` + typeName + kind + ` (id_value, value) VALUES (?, ?)`

	values = append(values, id, value)

	return
}