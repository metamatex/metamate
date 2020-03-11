package generic

import (
	"log"
	"reflect"
	"strings"

	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/fieldflags"
	"github.com/wolfeidau/unflatten"
)

type MultiMapFactory struct {
	root *graph.RootNode
}

func NewFactory(root *graph.RootNode) (Factory) {
	return MultiMapFactory{
		root: root,
	}
}

func (f MultiMapFactory) New(tn *graph.TypeNode) (Generic) {
	return &MultiMapGeneric{
		tn: tn,
	}
}

func (f MultiMapFactory) NewSlice(tn *graph.TypeNode) (Slice) {
	return &MultiMapSlice{
		tn: tn,
	}
}

func (f MultiMapFactory) FromStructs(is interface{}) (gs Slice, err error) {
	return f.fromStructs(is)
}

func (f MultiMapFactory) fromStructs(is interface{}) (gs *MultiMapSlice, err error) {
	gs = &MultiMapSlice{tn: f.root.Types.MustByName(reflect.TypeOf(is).Elem().Name())}

	switch reflect.TypeOf(is).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(is)

		for i := 0; i < s.Len(); i++ {
			var g *MultiMapGeneric
			g , err = f.fromStruct(s.Index(i).Interface())
			if err != nil {
			    return
			}

			gs.Gs = append(gs.Gs, g)
		}
	default:
		log.Fatal("not a slice")
	}

	return
}

func (f MultiMapFactory) FromStruct(i interface{}) (Generic, error) {
	return f.fromStruct(i)
}

func (f MultiMapFactory) fromStruct(i interface{}) (g *MultiMapGeneric, err error) {
	v := reflect.ValueOf(i)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		panic("empty point")
	}

	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("not struct")
	}

	t := v.Type()

	t0, err := f.root.Types.ByName(t.Name())
	if err != nil {
	    return
	}

	g = NewMultiMapGeneric(t0)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.PkgPath != "" {
			continue
		}

		v0 := v.FieldByIndex([]int{i})

		if v0.IsNil() {
			continue
		}

		fieldName := strings.Split(field.Tag.Get("json"), ",")[0]

		switch v0.Kind() {
		case reflect.Ptr:
			switch v0.Elem().Kind() {
			case reflect.String:
				if g.String0 == nil {
					g.String0 = map[string]string{}
				}

				g.String0[fieldName] = v0.Elem().Interface().(string)

				break
			case reflect.Int32:
				if g.Int320 == nil {
					g.Int320 = map[string]int32{}
				}

				g.Int320[fieldName] = v0.Elem().Interface().(int32)

				break
			case reflect.Uint32:
				if g.Uint320 == nil {
					g.Uint320 = map[string]uint32{}
				}

				g.Uint320[fieldName] = v0.Elem().Interface().(uint32)

				break
			case reflect.Float64:
				if g.Float640 == nil {
					g.Float640 = map[string]float64{}
				}

				g.Float640[fieldName] = v0.Elem().Interface().(float64)

				break
			case reflect.Bool:
				if g.Bool0 == nil {
					g.Bool0 = map[string]bool{}
				}

				g.Bool0[fieldName] = v0.Elem().Interface().(bool)

				break
			case reflect.Struct:
				if g.Generic0 == nil {
					g.Generic0 = map[string]*MultiMapGeneric{}
				}

				g.Generic0[fieldName], err = f.fromStruct(v0.Elem().Interface())
				if err != nil {
				    return
				}

				break
			}
		case reflect.Slice:
			switch v0.Type().Elem().Kind() {
			case reflect.String:
				if g.StringSlice0 == nil {
					g.StringSlice0 = map[string][]string{}
				}

				g.StringSlice0[fieldName] = v0.Interface().([]string)

				break
			case reflect.Int32:
				if g.Int32Slice0 == nil {
					g.Int32Slice0 = map[string][]int32{}
				}

				g.Int32Slice0[fieldName] = v0.Interface().([]int32)

				break
			case reflect.Uint32:
				if g.Uint32Slice0 == nil {
					g.Uint32Slice0 = map[string][]uint32{}
				}

				g.Uint32Slice0[fieldName] = v0.Interface().([]uint32)

				break
			case reflect.Float64:
				if g.Float64Slice0 == nil {
					g.Float64Slice0 = map[string][]float64{}
				}

				g.Float64Slice0[fieldName] = v0.Interface().([]float64)

				break
			case reflect.Bool:
				if g.BoolSlice0 == nil {
					g.BoolSlice0 = map[string][]bool{}
				}

				g.BoolSlice0[fieldName] = v0.Interface().([]bool)

				break
			case reflect.Struct:
				if g.GenericSlice0 == nil {
					g.GenericSlice0 = map[string]*MultiMapSlice{}
				}

				t1 := f.root.Types.MustByName(v0.Type().Elem().Name())

				gs := &MultiMapSlice{tn: t1}

				for i := 0; i < v0.Len(); i++ {
					var g0 *MultiMapGeneric
					g0, err = f.fromStruct(v0.Index(i).Interface())
					if err != nil {
					    return
					}

					gs.Gs = append(gs.Gs, g0)
				}

				g.GenericSlice0[fieldName] = gs

				break
			}
		}
	}

	return
}

func (f MultiMapFactory) FromStringInterfaceMap(tn *graph.TypeNode, m map[string]interface{}) (g Generic, err error) {
	return f.fromStringInterfaceMap(tn, m), nil
}

func (f MultiMapFactory) fromStringInterfaceMap(tn *graph.TypeNode, m map[string]interface{}) (g *MultiMapGeneric) {
	g = NewMultiMapGeneric(tn)

	tn.Edges.Fields.Holds().Each(func(fn *graph.FieldNode) {
		v, ok := m[fn.Name()]
		if !ok {
			return
		}

		if v == nil {
			return
		}

		if fn.Flags().Is(fieldflags.IsList, true) {
			if fn.Edges.Type.Resolver.Holds() != "" {
				if g.GenericSlice0 == nil {
					g.GenericSlice0 = map[string]*MultiMapSlice{}
				}

				var ms []map[string]interface{}
				switch v := v.(type) {
				case []map[string]interface{}:
					ms = v
					break
				case []interface{}:
					for _, m := range v {
						ms = append(ms, m.(map[string]interface{}))
					}
					break
				}

				gs := &MultiMapSlice{tn: fn.Edges.Type.Holds()}
				for _, m := range ms {
					gs.Gs = append(gs.Gs, f.fromStringInterfaceMap(fn.Edges.Type.Holds(), m))
				}

				g.GenericSlice0[fn.Name()] = gs
			} else if fn.Edges.BasicType.Resolver.Holds() != "" {
				switch fn.Edges.BasicType.Holds().Id() {
				case graph.STRING:
					if g.StringSlice0 == nil {
						g.StringSlice0 = map[string][]string{}
					}

					var ss []string
					switch v := v.(type) {
					case []string:
						ss = v
						break
					case []interface{}:
						for _, s := range v {
							ss = append(ss, s.(string))
						}
						break
					}

					g.StringSlice0[fn.Name()] = ss

					break
				case graph.INT32:
					if g.Int32Slice0 == nil {
						g.Int32Slice0 = map[string][]int32{}
					}

					var is []int32
					switch v := v.(type) {
					case []int32:
						is = v
						break
					case []interface{}:
						for _, i := range v {
							var i0 int32
							switch i := i.(type) {
							case float64:
								i0 = int32(i)
								break
							case int32:
								i0 = i
							}

							is = append(is, i0)
						}

						break
					}

					g.Int32Slice0[fn.Name()] = is

					break
				case graph.FLOAT64:
					if g.Float64Slice0 == nil {
						g.Float64Slice0 = map[string][]float64{}
					}

					var fs []float64
					switch v := v.(type) {
					case []float64:
						fs = v
						break
					case []interface{}:
						for _, f := range v {
							fs = append(fs, f.(float64))
						}
						break
					}

					g.Float64Slice0[fn.Name()] = fs

					break
				case graph.BOOL:
					if g.BoolSlice0 == nil {
						g.BoolSlice0 = map[string][]bool{}
					}

					var bs []bool
					switch v := v.(type) {
					case []bool:
						bs = v
						break
					case []interface{}:
						for _, b := range v {
							bs = append(bs, b.(bool))
						}
						break
					}

					g.BoolSlice0[fn.Name()] = bs

					break
				}
			} else if fn.Edges.Enum.Resolver.Holds() != "" {
				if g.StringSlice0 == nil {
					g.StringSlice0 = map[string][]string{}
				}

				var ss []string
				switch v := v.(type) {
				case []string:
					ss = v
					break
				case []interface{}:
					for _, s := range v {
						ss = append(ss, s.(string))
					}
					break
				}

				g.StringSlice0[fn.Name()] = ss
			}
		} else {
			if fn.Edges.Type.Resolver.Holds() != "" {
				if g.Generic0 == nil {
					g.Generic0 = map[string]*MultiMapGeneric{}
				}

				g.Generic0[fn.Name()] = f.fromStringInterfaceMap(fn.Edges.Type.Holds(), v.(map[string]interface{}))
			} else if fn.Edges.BasicType.Resolver.Holds() != "" {
				switch fn.Edges.BasicType.Holds().Id() {
				case graph.STRING:
					if g.String0 == nil {
						g.String0 = map[string]string{}
					}

					v0 := v.(string)

					g.String0[fn.Name()] = v0

					break
				case graph.INT32:
					if g.Int320 == nil {
						g.Int320 = map[string]int32{}
					}

					var v0 int32
					switch v := v.(type) {
					case float64:
						v0 = int32(v)
						break
					case int32:
						v0 = v
					}

					g.Int320[fn.Name()] = v0

					break
				case graph.FLOAT64:
					if g.Float640 == nil {
						g.Float640 = map[string]float64{}
					}

					v0 := v.(float64)

					g.Float640[fn.Name()] = v0

					break
				case graph.BOOL:
					if g.Bool0 == nil {
						g.Bool0 = map[string]bool{}
					}

					switch v := v.(type) {
					case int, int8, int16, int32, int64:
						g.Bool0[fn.Name()] = v == 1
					case bool:
						g.Bool0[fn.Name()] = v
					default:
						panic("expect bool, int, int8, int16, Int320 or int64")
					}

					break
				}
			} else if fn.Edges.Enum.Resolver.Holds() != "" {
				if g.String0 == nil {
					g.String0 = map[string]string{}
				}

				v0 := v.(string)

				g.String0[fn.Name()] = v0
			}
		}
	})

	return
}

func (f MultiMapFactory) Unflatten(tn *graph.TypeNode, delimiter string, m map[string]interface{}) (Generic, error) {
	m = unflatten.Unflatten(m, func (k string) []string { return strings.Split(k, delimiter) })

	return f.FromStringInterfaceMap(tn, m)
}

func (f MultiMapFactory) UnflattenSlice(tn *graph.TypeNode, delimiter string, ms []map[string]interface{}) (Slice, error) {
	gs := &MultiMapSlice{tn: tn}
	for _, m := range ms {
		m = unflatten.Unflatten(m, func (k string) []string { return strings.Split(k, delimiter) })

		g := f.fromStringInterfaceMap(tn, m)

		gs.Gs = append(gs.Gs, g)
	}

	return gs, nil
}


func (f MultiMapFactory) MustFromStruct(any interface{}) (Generic) {
	g, err := f.FromStruct(any)
	if err != nil {
	    panic(err)
	}

	return g
}

func (f MultiMapFactory) MustFromStructs(any interface{}) (Slice) {
	gSlice, err := f.FromStructs(any)
	if err != nil {
		panic(err)
	}

	return gSlice
}

func (f MultiMapFactory) MustFromStringInterfaceMap(tn *graph.TypeNode, m map[string]interface{}) (Generic) {
	g, err := f.FromStringInterfaceMap(tn, m)
	if err != nil {
		panic(err)
	}

	return g
}

func (f MultiMapFactory) MustUnflatten(tn *graph.TypeNode, delimiter string, m map[string]interface{}) (Generic) {
	g, err := f.Unflatten(tn, delimiter, m)
	if err != nil {
		panic(err)
	}

	return g
}

func (f MultiMapFactory) MustUnflattenSlice(tn *graph.TypeNode, delimiter string, ms []map[string]interface{}) (Slice) {
	gSlice, err := f.UnflattenSlice(tn, delimiter, ms)
	if err != nil {
		panic(err)
	}

	return gSlice
}
