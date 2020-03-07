package sql

import (
	"github.com/metamatex/metamatemono/generic/pkg/v0/generic"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph"
	"strings"

	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/fieldnames"
)

type Conditions struct {
	Conditions []string
	Values     []interface{}
	Or         []Conditions
	And        []Conditions
}

func ComposeFilter(c Conditions) (q string, values []interface{}) {
	var conditionsQ string
	var conditionsValues []interface{}
	firstConditions := true
	for _, c0 := range c.Conditions {
		if firstConditions {
			firstConditions = false
		} else {
			conditionsQ += " AND "
		}

		conditionsQ += c0
	}
	conditionsValues = append(conditionsValues, c.Values...)

	var orQ string
	var orValues []interface{}
	firstOr := true
	for _, c0 := range c.Or {
		if firstOr {
			firstOr = false
		} else {
			orQ += " OR "
		}

		q0, values0 := ComposeFilter(c0)
		orQ += q0
		orValues = append(orValues, values0...)
	}

	var andQ string
	var andValues []interface{}
	firstAnd := true
	for _, c0 := range c.And {
		if firstAnd {
			firstAnd = false
		} else {
			andQ += " AND "
		}

		q0, values0 := ComposeFilter(c0)

		andQ += q0
		andValues = append(andValues, values0...)
	}

	if conditionsQ != "" {
		q += conditionsQ
		values = append(values, conditionsValues...)
	}

	if orQ != "" {
		if q != "" {
			q += " AND "
		}

		q += orQ
		values = append(values, orValues...)
	}

	if andQ != "" {
		if q != "" {
			q += " AND "
		}

		q += andQ
		values = append(values, andValues...)
	}

	if q != "" {
		q = "(" + q + ")"
	}

	return
}

func GetStringConditions(gFilter generic.Generic, fieldName string) (c Conditions) {
	var fieldName0 string
	caseSensitive, _ := gFilter.Bool(fieldnames.CaseSensitive)
	if !caseSensitive {
		fieldName0 = "lower(" + fieldName + ")"
	} else {
		fieldName0 = fieldName
	}

	gFilter.EachString(func(fn *graph.FieldNode, v string) {
		if caseSensitive {
			v = strings.ToLower(v)
		}

		switch fn.Name() {
		case fieldnames.Contains:
			c.Conditions = append(c.Conditions, "("+fieldName0+" LIKE ?)")
			v = "%" + v + "%"
		case fieldnames.EndsWith:
			c.Conditions = append(c.Conditions, "("+fieldName0+" LIKE ?)")
			v = "%" + v
		case fieldnames.Is:
			c.Conditions = append(c.Conditions, "("+fieldName0+" LIKE ?)")
		case fieldnames.Not:
			c.Conditions = append(c.Conditions, "(NOT ("+fieldName0+" LIKE ?))")
		case fieldnames.NotContains:
			c.Conditions = append(c.Conditions, "(NOT ("+fieldName0+" LIKE ?))")
			v = "%" + v + "%"
		case fieldnames.NotEndsWith:
			c.Conditions = append(c.Conditions, "(NOT ("+fieldName0+" LIKE ?))")
			v = "%" + v
		case fieldnames.NotStartsWith:
			c.Conditions = append(c.Conditions, "(NOT ("+fieldName0+" LIKE ?))")
			v = v + "%"
		case fieldnames.StartsWith:
			c.Conditions = append(c.Conditions, "("+fieldName0+" LIKE ?)")
			v = v + "%"
		default:
			return
		}

		c.Values = append(c.Values, v)
	})

	gFilter.EachBool(func(fn *graph.FieldNode , v bool) {
		switch fn.Name() {
		case fieldnames.NotSet:
			if v {
				c.Conditions = append(c.Conditions, "("+fieldName0+" IS NULL)")
			} else {
				c.Conditions = append(c.Conditions, "("+fieldName0+" IS NOT NULL)")
			}
		case fieldnames.Set:
			if v {
				c.Conditions = append(c.Conditions, "("+fieldName0+" IS NOT NULL)")
			} else {
				c.Conditions = append(c.Conditions, "("+fieldName0+" IS NULL)")
			}
		default:
			return
		}
	})

	gFilter.EachStringSlice(func(fn *graph.FieldNode, v []string) {
		if caseSensitive {
			for i, _ := range v {
				v[i] = strings.ToLower(v[i])
			}
		}

		switch fn.Name() {
		case fieldnames.In:
			c.Conditions = append(c.Conditions, "("+fieldName0+" IN (?))")
		case fieldnames.NotIn:
			c.Conditions = append(c.Conditions, "(NOT ("+fieldName0+" IN (?)))")
		default:
			return
		}

		c.Values = append(c.Values, v)
	})

	gFilter.EachGenericSlice(func(fn *graph.FieldNode, v generic.Slice) {
		switch fn.Name() {
		case fieldnames.Or:
			for _, g := range v.Get() {
				c.Or = append(c.Or, GetStringConditions(g, fieldName))
			}
		case fieldnames.And:
			for _, g := range v.Get() {
				c.And = append(c.And, GetStringConditions(g, fieldName))
			}
		default:
			return
		}
	})

	return
}

func GetInt32Conditions(gFilter generic.Generic, fieldName string) (c Conditions) {
	gFilter.EachInt32(func(fn *graph.FieldNode, v int32) {
		switch fn.Name() {
		case fieldnames.Gt:
			c.Conditions = append(c.Conditions, "("+fieldName+" > ?)")
		case fieldnames.Gte:
			c.Conditions = append(c.Conditions, "("+fieldName+" >= ?)")
		case fieldnames.Is:
			c.Conditions = append(c.Conditions, "("+fieldName+" = ?)")
		case fieldnames.Lt:
			c.Conditions = append(c.Conditions, "("+fieldName+" < ?)")
		case fieldnames.Lte:
			c.Conditions = append(c.Conditions, "("+fieldName+" <= ?)")
		case fieldnames.Not:
			c.Conditions = append(c.Conditions, "(NOT ("+fieldName+" = ?))")
		default:
			return
		}
	})

	gFilter.EachBool(func(fn *graph.FieldNode, v bool) {
		switch fn.Name() {
		case fieldnames.NotSet:
			if v {
				c.Conditions = append(c.Conditions, "("+fieldName+" IS NULL)")
			} else {
				c.Conditions = append(c.Conditions, "("+fieldName+" IS NOT NULL)")
			}
		case fieldnames.Set:
			if v {
				c.Conditions = append(c.Conditions, "("+fieldName+" IS NOT NULL)")
			} else {
				c.Conditions = append(c.Conditions, "("+fieldName+" IS NULL)")
			}
		default:
			return
		}
	})

	gFilter.EachInt32Slice(func(fn *graph.FieldNode, v []int32) {
		switch fn.Name() {
		case fieldnames.In:
			c.Conditions = append(c.Conditions, "("+fieldName+" IN (?))")
		case fieldnames.NotIn:
			c.Conditions = append(c.Conditions, "(NOT ("+fieldName+" IN (?)))")
		default:
			return
		}

		c.Values = append(c.Values, v)
	})

	gFilter.EachGenericSlice(func(fn *graph.FieldNode, v generic.Slice) {
		switch fn.Name() {
		case fieldnames.Or:
			for _, g := range v.Get() {
				c.Or = append(c.Or, GetStringConditions(g, fieldName))
			}
		case fieldnames.And:
			for _, g := range v.Get() {
				c.And = append(c.And, GetStringConditions(g, fieldName))
			}
		default:
			return
		}
	})

	return
}

func GetFloat64Conditions(gFilter generic.Generic, fieldName string) (c Conditions) {
	gFilter.EachFloat64(func(fn *graph.FieldNode, v float64) {
		switch fn.Name() {
		case fieldnames.Gt:
			c.Conditions = append(c.Conditions, "("+fieldName+" > ?)")
		case fieldnames.Gte:
			c.Conditions = append(c.Conditions, "("+fieldName+" >= ?)")
		case fieldnames.Is:
			c.Conditions = append(c.Conditions, "("+fieldName+" = ?)")
		case fieldnames.Lt:
			c.Conditions = append(c.Conditions, "("+fieldName+" < ?)")
		case fieldnames.Lte:
			c.Conditions = append(c.Conditions, "("+fieldName+" <= ?)")
		case fieldnames.Not:
			c.Conditions = append(c.Conditions, "(NOT ("+fieldName+" = ?))")
		default:
			return
		}

		c.Values = append(c.Values, v)
	})

	gFilter.EachBool(func(fn *graph.FieldNode, v bool) {
		switch fn.Name() {
		case fieldnames.NotSet:
			if v {
				c.Conditions = append(c.Conditions, "("+fieldName+" IS NULL)")
			} else {
				c.Conditions = append(c.Conditions, "("+fieldName+" IS NOT NULL)")
			}
		case fieldnames.Set:
			if v {
				c.Conditions = append(c.Conditions, "("+fieldName+" IS NOT NULL)")
			} else {
				c.Conditions = append(c.Conditions, "("+fieldName+" IS NULL)")
			}
		default:
			return
		}
	})

	gFilter.EachFloat64Slice(func(fn *graph.FieldNode, v []float64) {
		switch fn.Name() {
		case fieldnames.In:
			c.Conditions = append(c.Conditions, "("+fieldName+" IN (?))")
		case fieldnames.NotIn:
			c.Conditions = append(c.Conditions, "(NOT ("+fieldName+" IN (?)))")
		default:
			return
		}

		c.Values = append(c.Values, v)
	})

	gFilter.EachGenericSlice(func(fn *graph.FieldNode, v generic.Slice) {
		switch fn.Name() {
		case fieldnames.Or:
			for _, g := range v.Get() {
				c.Or = append(c.Or, GetStringConditions(g, fieldName))
			}
		case fieldnames.And:
			for _, g := range v.Get() {
				c.And = append(c.And, GetStringConditions(g, fieldName))
			}
		default:
			return
		}
	})

	return
}

func GetBoolConditions(gFilter generic.Generic, fieldName string) (c Conditions) {
	gFilter.EachBool(func(fn *graph.FieldNode, v bool) {
		switch fn.Name() {
		case fieldnames.Is:
			c.Conditions = append(c.Conditions, "("+fieldName+" = ?)")
		case fieldnames.Not:
			c.Conditions = append(c.Conditions, "(NOT ("+fieldName+" = ?))")
		default:
			return
		}

		c.Values = append(c.Values, v)
	})

	gFilter.EachBool(func(fn *graph.FieldNode, v bool) {
		switch fn.Name() {
		case fieldnames.NotSet:
			if v {
				c.Conditions = append(c.Conditions, "("+fieldName+" IS NULL)")
			} else {
				c.Conditions = append(c.Conditions, "("+fieldName+" IS NOT NULL)")
			}
		case fieldnames.Set:
			if v {
				c.Conditions = append(c.Conditions, "("+fieldName+" IS NOT NULL)")
			} else {
				c.Conditions = append(c.Conditions, "("+fieldName+" IS NULL)")
			}
		default:
			return
		}
	})

	gFilter.EachGenericSlice(func(fn *graph.FieldNode, v generic.Slice) {
		switch fn.Name() {
		case fieldnames.Or:
			for _, g := range v.Get() {
				c.Or = append(c.Or, GetStringConditions(g, fieldName))
			}
		case fieldnames.And:
			for _, g := range v.Get() {
				c.And = append(c.And, GetStringConditions(g, fieldName))
			}
		default:
			return
		}
	})

	return
}

func GetEnumConditions(gFilter generic.Generic, fieldName string) (c Conditions) {
	gFilter.EachString(func(fn *graph.FieldNode, v string) {
		switch fn.Name() {
		case fieldnames.Is:
			c.Conditions = append(c.Conditions, "("+fieldName+" = ?)")
		case fieldnames.Not:
			c.Conditions = append(c.Conditions, "(NOT ("+fieldName+" = ?))")
		default:
			return
		}

		c.Values = append(c.Values, v)
	})

	gFilter.EachBool(func(fn *graph.FieldNode, v bool) {
		switch fn.Name() {
		case fieldnames.NotSet:
			if v {
				c.Conditions = append(c.Conditions, "("+fieldName+" IS NULL)")
			} else {
				c.Conditions = append(c.Conditions, "("+fieldName+" IS NOT NULL)")
			}
		case fieldnames.Set:
			if v {
				c.Conditions = append(c.Conditions, "("+fieldName+" IS NOT NULL)")
			} else {
				c.Conditions = append(c.Conditions, "("+fieldName+" IS NULL)")
			}
		default:
			return
		}
	})

	gFilter.EachStringSlice(func(fn *graph.FieldNode, v []string) {
		switch fn.Name() {
		case fieldnames.In:
			c.Conditions = append(c.Conditions, "("+fieldName+" IN (?))")
		case fieldnames.NotIn:
			c.Conditions = append(c.Conditions, "(NOT ("+fieldName+" IN (?)))")
		default:
			return
		}

		c.Values = append(c.Values, v)
	})

	gFilter.EachGenericSlice(func(fn *graph.FieldNode, v generic.Slice) {
		switch fn.Name() {
		case fieldnames.Or:
			for _, g := range v.Get() {
				c.Or = append(c.Or, GetStringConditions(g, fieldName))
			}
		case fieldnames.And:
			for _, g := range v.Get() {
				c.And = append(c.And, GetStringConditions(g, fieldName))
			}
		default:
			return
		}
	})

	return
}
