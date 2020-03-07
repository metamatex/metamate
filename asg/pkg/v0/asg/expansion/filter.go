package expansion

import (
	"errors"
	"fmt"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph/fieldflags"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/typenames"
)

func addBasicFilters(root *graph.RootNode) {
	flags := graph.Flags{
		fieldflags.Filter: false,
	}

	root.AddTypeNode(typenames.StringFilter, graph.FieldNodeSlice{
		graph.BoolField(fieldnames.CaseSensitive, flags),
		graph.BoolField(fieldnames.Set, flags),
		graph.StringField(fieldnames.Is, flags),
		graph.StringField(fieldnames.Not, flags),
		graph.StringField(fieldnames.Contains, flags),
		graph.StringField(fieldnames.NotContains, flags),
		graph.StringField(fieldnames.StartsWith, flags),
		graph.StringField(fieldnames.NotStartsWith, flags),
		graph.StringField(fieldnames.EndsWith, flags),
		graph.StringField(fieldnames.NotEndsWith, flags),
		graph.ListField(graph.StringField(fieldnames.In, flags)),
		graph.ListField(graph.StringField(fieldnames.NotIn, flags)),
		graph.ListField(graph.TypeField(fieldnames.And, typenames.StringFilter)),
		graph.ListField(graph.TypeField(fieldnames.Or, typenames.StringFilter)),
	}, graph.Flags{
		typeflags.GetFilter:     false,
		typeflags.IsFilter:      true,
		typeflags.IsBasicFilter: true,
	})

	root.AddTypeNode(typenames.StringListFilter, graph.FieldNodeSlice{
		graph.TypeField(fieldnames.And, typenames.StringFilter),
		graph.TypeField(fieldnames.Or, typenames.StringFilter),
		graph.TypeField(fieldnames.Not, typenames.StringFilter),
	}, graph.Flags{
		typeflags.GetFilter: false,
		typeflags.IsFilter:  true,
	})

	root.AddTypeNode(typenames.EnumFilter, graph.FieldNodeSlice{
		graph.BoolField(fieldnames.Set, flags),
		graph.StringField(fieldnames.Is, flags),
		graph.StringField(fieldnames.Not, flags),
		graph.ListField(graph.StringField(fieldnames.In, flags)),
		graph.ListField(graph.StringField(fieldnames.NotIn, flags)),
		graph.ListField(graph.TypeField(fieldnames.And, typenames.EnumFilter)),
		graph.ListField(graph.TypeField(fieldnames.Or, typenames.EnumFilter)),
	}, graph.Flags{
		typeflags.GetFilter:     false,
		typeflags.IsFilter:      true,
		typeflags.IsBasicFilter: true,
	})

	root.AddTypeNode(typenames.EnumListFilter, graph.FieldNodeSlice{
		graph.TypeField(fieldnames.And, typenames.EnumFilter),
		graph.TypeField(fieldnames.Or, typenames.EnumFilter),
		graph.TypeField(fieldnames.Not, typenames.EnumFilter),
	}, graph.Flags{
		typeflags.GetFilter: false,
		typeflags.IsFilter:  true,
	})

	root.AddTypeNode(typenames.Int32Filter, graph.FieldNodeSlice{
		graph.BoolField(fieldnames.Set, flags),
		graph.Int32Field(fieldnames.Is, flags),
		graph.Int32Field(fieldnames.Not, flags),
		graph.Int32Field(fieldnames.Lt, flags),
		graph.Int32Field(fieldnames.Lte, flags),
		graph.Int32Field(fieldnames.Gt, flags),
		graph.Int32Field(fieldnames.Gte, flags),
		graph.ListField(graph.Int32Field(fieldnames.In, flags)),
		graph.ListField(graph.Int32Field(fieldnames.NotIn, flags)),
		graph.ListField(graph.TypeField(fieldnames.And, typenames.Int32Filter)),
		graph.ListField(graph.TypeField(fieldnames.Or, typenames.Int32Filter)),
	}, graph.Flags{
		typeflags.GetFilter:     false,
		typeflags.IsFilter:      true,
		typeflags.IsBasicFilter: true,
	})

	root.AddTypeNode(typenames.Int32ListFilter, graph.FieldNodeSlice{
		graph.TypeField(fieldnames.And, typenames.Int32Filter),
		graph.TypeField(fieldnames.Or, typenames.Int32Filter),
		graph.TypeField(fieldnames.Not, typenames.Int32Filter),
	}, graph.Flags{
		typeflags.GetFilter: false,
		typeflags.IsFilter:  true,
	})

	root.AddTypeNode(typenames.Float64Filter, graph.FieldNodeSlice{
		graph.BoolField(fieldnames.Set, flags),
		graph.Float64Field(fieldnames.Is, flags),
		graph.Float64Field(fieldnames.Not, flags),
		graph.Float64Field(fieldnames.Lt, flags),
		graph.Float64Field(fieldnames.Lte, flags),
		graph.Float64Field(fieldnames.Gt, flags),
		graph.Float64Field(fieldnames.Gte, flags),
		graph.ListField(graph.Float64Field(fieldnames.In, flags)),
		graph.ListField(graph.Float64Field(fieldnames.NotIn, flags)),
		graph.ListField(graph.TypeField(fieldnames.And, typenames.Float64Filter)),
		graph.ListField(graph.TypeField(fieldnames.Or, typenames.Float64Filter)),
	}, graph.Flags{
		typeflags.GetFilter:     false,
		typeflags.IsFilter:      true,
		typeflags.IsBasicFilter: true,
	})

	root.AddTypeNode(typenames.Float64ListFilter, graph.FieldNodeSlice{
		graph.TypeField(fieldnames.And, typenames.Float64Filter),
		graph.TypeField(fieldnames.Or, typenames.Float64Filter),
		graph.TypeField(fieldnames.Not, typenames.Float64Filter),
	}, graph.Flags{
		typeflags.GetFilter: false,
		typeflags.IsFilter:  true,
	})

	root.AddTypeNode(typenames.BoolFilter, graph.FieldNodeSlice{
		graph.BoolField(fieldnames.Set),
		graph.BoolField(fieldnames.Is),
		graph.BoolField(fieldnames.Not),
		graph.ListField(graph.TypeField(fieldnames.And, typenames.BoolFilter)),
		graph.ListField(graph.TypeField(fieldnames.Or, typenames.BoolFilter)),
	}, graph.Flags{
		typeflags.GetFilter:     false,
		typeflags.IsFilter:      true,
		typeflags.IsBasicFilter: true,
	})

	root.AddTypeNode(typenames.BoolListFilter, graph.FieldNodeSlice{
		graph.TypeField(fieldnames.And, typenames.BoolFilter),
		graph.TypeField(fieldnames.Or, typenames.BoolFilter),
		graph.TypeField(fieldnames.Not, typenames.BoolFilter),
	}, graph.Flags{
		typeflags.GetFilter: false,
		typeflags.IsFilter:  true,
	})
}

func generateListFilter(root *graph.RootNode, tn *graph.TypeNode) {
	if tn.Flags().Is(typeflags.HasListFilter, true) {
		return
	}

	root.AddTypeNode(typenames.ListFilter(tn.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Some, typenames.Filter(tn.Name())),
		graph.TypeField(fieldnames.Every, typenames.Filter(tn.Name())),
		graph.TypeField(fieldnames.None, typenames.Filter(tn.Name())),
	}, graph.Flags{
		typeflags.GetFilter: false,
		typeflags.IsFilter:  true,
		typeflags.IsListFilter:  true,
	})

	tn.Flags().Set(typeflags.HasListFilter, true)
}

func generateTypeFilter(root *graph.RootNode, tn *graph.TypeNode) (err error) {
	fns := graph.FieldNodeSlice{}
	fns = append(fns, graph.BoolField(fieldnames.Set))

	fns = append(getFilterFields(tn.Edges.Fields.Holds()), fns...)

	name := typenames.Filter(tn.Name())

	if tn.Flags().Is(typeflags.IsFilter, false) {
		fns = append(fns, graph.ListField(graph.TypeField(fieldnames.And, name, graph.Flags{
			fieldflags.Filter: false,
		})))

		fns = append(fns, graph.ListField(graph.TypeField(fieldnames.Or, name, graph.Flags{
			fieldflags.Filter: false,
		})))

		fns = append(fns, graph.ListField(graph.TypeField(fieldnames.Not, name, graph.Flags{
			fieldflags.Filter: false,
		})))
	}

	if len(fns) == 0 {
		err = errors.New(fmt.Sprintf("filter %v can't have 0 fields", name))

		return
	}

	filterNode := root.AddTypeNode(name, fns, graph.Flags{
		typeflags.IsFilter:  true,
		typeflags.GetFilter: true,
	})
	filterNode.Edges.Type.Resolver.SetFor(tn.Id())

	tn.Flags().Set(typeflags.HasFilter, true)
	tn.Edges.Type.Resolver.SetFilteredBy(filterNode.Id())

	return
}

func generateFilterRecursive(root *graph.RootNode, tn *graph.TypeNode) (err error) {
	if tn.Edges.Type.FilteredBy() != nil || tn.Flags().Is(typeflags.GetFilter, false) {
		return
	}

	if !IsFieldFlagDeep(tn, fieldflags.Filter, true) {
		return
	}

	err = generateTypeFilter(root, tn)
	if err != nil {
		return
	}

	for _, fn := range tn.Edges.Fields.Holds() {
		if fn.Edges.Type.Holds() == nil {
			continue
		}

		if fn.Flags().Is(fieldflags.Filter, true) && IsFieldFlagDeep(fn.Edges.Type.Holds(), fieldflags.Filter, true) {
			err = generateFilterRecursive(root, fn.Edges.Type.Holds())
			if err != nil {
				return
			}
		}
	}

	return
}

//func getFilterFields(fnm graph.FieldNodeMap) (fns graph.FieldNodeSlice) {
//	for _, fn := range fnm {
//		if fn.Flags().Is(fieldflags.Filter, false) {
//			continue
//		}
//
//		if fn.IsBasicType() || fn.IsBasicTypeList() {
//			name := fn.Name()
//
//			basicFilterName := ""
//			switch fn.Edges.BasicType.Resolver.Holds() {
//			case graph.STRING:
//				basicFilterName = typenames.StringFilter
//			case graph.INT32:
//				basicFilterName = typenames.Int32Filter
//			case graph.FLOAT64:
//				basicFilterName = typenames.Float64Filter
//			case graph.BOOL:
//				basicFilterName = typenames.BoolFilter
//			default:
//				panic("Fields kind did not match")
//			}
//
//			filterName := ""
//			switch fn.Kind() {
//			//case graph.FieldKindType:
//
//			case graph.FieldKindEnum:
//				filterName = typenames.EnumFilter
//			case graph.FieldKindString:
//				filterName = typenames.StringFilter
//			case graph.FieldKindInt32:
//				filterName = typenames.Int32Filter
//			case graph.FieldKindFloat64:
//				filterName = typenames.Float64Filter
//			case graph.FieldKindBool:
//				filterName = typenames.BoolFilter
//			//case graph.FieldKindTypeList:
//				//filterName = typenames.StringFilter
//			case graph.FieldKindEnumList:
//				filterName = typenames.EnumListFilter
//			case graph.FieldKindStringList:
//				filterName = typenames.StringListFilter
//			case graph.FieldKindInt32List:
//				filterName = typenames.Int32ListFilter
//			case graph.FieldKindFloat64List:
//				filterName = typenames.Float64ListFilter
//			case graph.FieldKindBoolList:
//				filterName = typenames.BoolListFilter
//			}
//
//			if fn.Flags().Is(fieldflags.IsList, true) {
//				name_some := fn.Name() + fieldnames.Some
//				name_every := fn.Name() + fieldnames.Every
//				name_none := fn.Name() + fieldnames.None
//
//				for _, s := range []string{name_some, name_every, name_none} {
//					fn0 := graph.TypeField(s, basicFilterName, graph.flags{
//						fieldflags.Filter: false,
//					})
//					fn0.Edges.Field.Resolver.SetFor(fn.Id())
//					fns = append(fns, fn0)
//				}
//			} else {
//				fn0 := graph.TypeField(name, basicFilterName, graph.flags{
//					fieldflags.Filter: false,
//				})
//				fn0.Edges.Field.Resolver.SetFor(fn.Id())
//				fns = append(fns, fn0)
//			}
//
//			continue
//		}
//
//		if fn.IsType() || fn.IsTypeList() {
//			filterTypeName := typenames.Filter(fn.Edges.Type.Holds().Name())
//
//			name := fn.Name()
//
//			if !IsFieldFlagDeep(fn.Edges.Type.Holds(), fieldflags.Filter, true) {
//				fn.Flags().Set(fieldflags.Filter, false)
//
//				continue
//			}
//
//			if fn.Flags().Is(fieldflags.IsList, true) {
//				name_some := fn.Name() + fieldnames.Some
//				name_every := fn.Name() + fieldnames.Every
//				name_none := fn.Name() + fieldnames.None
//
//				for _, s := range []string{name_some, name_every, name_none} {
//					fn0 := graph.TypeField(s, filterTypeName, graph.flags{
//						fieldflags.Filter: false,
//					})
//					fn0.Edges.Field.Resolver.SetFor(fn.Id())
//					fns = append(fns, fn0)
//				}
//			} else {
//				fn0 := graph.TypeField(name, filterTypeName, graph.flags{
//					fieldflags.Filter: false,
//				})
//				fn0.Edges.Field.Resolver.SetFor(fn.Id())
//				fns = append(fns, fn0)
//			}
//
//			continue
//		}
//
//		if fn.IsEnum() || fn.IsEnumList() {
//			name := fn.Name()
//
//			if fn.Flags().Is(fieldflags.IsList, true) {
//				name_some := fn.Name() + fieldnames.Some
//				name_every := fn.Name() + fieldnames.Every
//				name_none := fn.Name() + fieldnames.None
//
//				for _, s := range []string{name_some, name_every, name_none} {
//					fn0 := graph.TypeField(s, typenames.EnumFilter, graph.flags{
//						fieldflags.Filter: false,
//					})
//					fn0.Edges.Field.Resolver.SetFor(fn.Id())
//					fns = append(fns, fn0)
//				}
//			} else {
//				fn0 := graph.TypeField(name, typenames.EnumFilter, graph.flags{
//					fieldflags.Filter: false,
//				})
//				fn0.Edges.Field.Resolver.SetFor(fn.Id())
//				fns = append(fns, fn0)
//			}
//
//			continue
//		}
//
//		fn.Print()
//
//		panic("hi123")
//	}
//
//	return
//}

func getFilterFields(fnm graph.FieldNodeMap) (fns graph.FieldNodeSlice) {
	for _, fn := range fnm {
		if fn.Flags().Is(fieldflags.Filter, false) {
			continue
		}

		name := fn.Name()
		filterName := ""
		switch fn.Kind() {
		case graph.FieldKindType:
			filterName = typenames.Filter(fn.Edges.Type.Holds().Name())

			if !IsFieldFlagDeep(fn.Edges.Type.Holds(), fieldflags.Filter, true) {
				fn.Flags().Set(fieldflags.Filter, false)

				continue
			}
		case graph.FieldKindEnum:
			filterName = typenames.EnumFilter
		case graph.FieldKindString:
			filterName = typenames.StringFilter
		case graph.FieldKindInt32:
			filterName = typenames.Int32Filter
		case graph.FieldKindFloat64:
			filterName = typenames.Float64Filter
		case graph.FieldKindBool:
			filterName = typenames.BoolFilter
		case graph.FieldKindTypeList:
			filterName = typenames.ListFilter(fn.Edges.Type.Holds().Name())

			if !IsFieldFlagDeep(fn.Edges.Type.Holds(), fieldflags.Filter, true) {
				fn.Flags().Set(fieldflags.Filter, false)

				continue
			}
		case graph.FieldKindEnumList:
			filterName = typenames.EnumListFilter
		case graph.FieldKindStringList:
			filterName = typenames.StringListFilter
		case graph.FieldKindInt32List:
			filterName = typenames.Int32ListFilter
		case graph.FieldKindFloat64List:
			filterName = typenames.Float64ListFilter
		case graph.FieldKindBoolList:
			filterName = typenames.BoolListFilter
		default:
			panic("hi123")
		}

		fn0 := graph.TypeField(name, filterName, graph.Flags{
			fieldflags.Filter: false,
		})
		fn0.Edges.Field.Resolver.SetFor(fn.Id())
		fns = append(fns, fn0)
	}

	return
}

func IsFieldFlagDeep(tn *graph.TypeNode, flag string, b bool) (bool) {
	return isFieldFlagDeep(tn, map[graph.NodeId]bool{}, flag, b)
}

func isFieldFlagDeep(tn *graph.TypeNode, seen map[graph.NodeId]bool, flag string, b bool) (bool) {
	_, ok := seen[tn.Id()]
	if ok {
		return false
	}
	seen[tn.Id()] = true

	for _, fn := range tn.Edges.Fields.Holds() {
		if fn.Flags().Is(flag, b) {
			return true
		}

		if fn.Edges.Type.Resolver.Holds() != tn.Id() && fn.Edges.Type.Resolver.Holds() != "" && isFieldFlagDeep(fn.Edges.Type.Holds(), seen, flag, b) == true {
			return true
		}
	}

	return false
}
