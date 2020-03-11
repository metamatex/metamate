package sql

import (
	"bytes"
	"fmt"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/fieldflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/typenames"
)

func Create(supportedIdKinds map[string]bool, tnm graph.TypeNodeMap, rnm graph.RelationNodeMap) (q string, err error) {
	b := bytes.Buffer{}

	generateRelations(&b, rnm)

	tnm.Each(func(tn *graph.TypeNode) {
		generateCreateAlternativeIdsTables(supportedIdKinds, &b, tn)

		generateTable(&b, tn)
	})

	q = b.String()

	return
}

func generateCreateAlternativeIdsTables(supportedIdKinds map[string]bool, b *bytes.Buffer, tn *graph.TypeNode) {
	for a, _ := range supportedIdKinds {
		b.WriteString("CREATE TABLE ")
		b.WriteString(tn.Name())
		b.WriteString(a)
		b.WriteString(" (\n\tid_value VARCHAR NOT NULL,\n\tvalue TEXT NOT NULL\n);\n\n")
	}

	return
}

func generateTable(b *bytes.Buffer, tn *graph.TypeNode) {
	b.WriteString("CREATE TABLE ")
	b.WriteString(tn.Name())
	b.WriteString(" (\n")

	columns := []string{
		"id_value VARCHAR PRIMARY KEY,",
	}

	columns = append(columns, getColumns(tn)...)
	columns[len(columns)-1] = columns[len(columns)-1][:len(columns[len(columns)-1])-1]

	for _, column := range columns {
		b.WriteString("\t")
		b.WriteString(column)
		b.WriteString("\n")
	}

	b.WriteString(");\n\n")

	return
}

func generateRelations(b *bytes.Buffer, rnm graph.RelationNodeMap) {
	rnm.Each(func(rn *graph.RelationNode) {
		b.WriteString("CREATE TABLE ")
		b.WriteString(rn.Name())
		b.WriteString(" (\n\tactive_id_value VARCHAR NOT NULL,\n\tpassive_id_value VARCHAR NOT NULL\n);\n\n")
	})

	return
}

func getColumns(tn *graph.TypeNode) (columns []string) {
	flat := FlattenTypeNode(tn, []string{typenames.Service, typenames.ServiceId}, []string{typeflags.IsRelations, typeflags.IsRelationships})

	for k, fn := range flat {
		if fn.Flags().Is(fieldflags.IsHash, true) {
			continue
		}

		if k == "id_value" {
			continue
		}

		switch fn.Kind() {
		case graph.FieldKindEnum:
			columns = append(columns, fmt.Sprintf("%s TEXT NULL,", k))

			break
		case graph.FieldKindString:
			columns = append(columns, fmt.Sprintf("%s TEXT NULL,", k))

			break
		case graph.FieldKindBool:
			columns = append(columns, fmt.Sprintf("%s TINYINT NULL,", k))

			break
		case graph.FieldKindInt32:
			columns = append(columns, fmt.Sprintf("%s INT NULL,", k))

			break
		case graph.FieldKindFloat64:
			columns = append(columns, fmt.Sprintf("%s DOUBLE NULL,", k))

			break
		}
	}

	return
}
