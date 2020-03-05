package sql

import (
	"github.com/metamatex/metamatemono/pkg/generic/pkg/v0/generic"
	"github.com/metamatex/asg/pkg/v0/asg/graph"
	"github.com/metamatex/asg/pkg/v0/asg/typenames"
)

func Filter(gFilter generic.Generic) (q string, values []interface{}) {
	return ComposeFilter(GetFilterConditions(gFilter, ""))
}

func GetFilterConditions(gFilter generic.Generic, prefix string) (c Conditions) {
	gFilter.EachGeneric(func(fn *graph.FieldNode, v generic.Generic) {
		switch v.Type().Name() {
		case typenames.StringFilter:
			c.And = append(c.And, GetStringConditions(v, prefix+fn.Name()))
		case typenames.Int32Filter:
			c.And = append(c.And, GetInt32Conditions(v, prefix+fn.Name()))
		case typenames.Float64Filter:
			c.And = append(c.And, GetFloat64Conditions(v, prefix+fn.Name()))
		case typenames.BoolFilter:
			c.And = append(c.And, GetBoolConditions(v, prefix+fn.Name()))
		case typenames.EnumFilter:
			c.And = append(c.And, GetEnumConditions(v, prefix+fn.Name()))
		default:
			c.And = append(c.And, GetFilterConditions(v, prefix+fn.Name()+"_"))
		}
	})

	return
}
