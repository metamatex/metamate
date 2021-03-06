package generic

import "github.com/metamatex/metamate/asg/pkg/v0/asg/graph"

type Generic interface {
	MustDelete(...string)
	Delete(...string) error
	WalkDelete(func(fn *graph.FieldNode) bool)
	String(...string) (string, bool)
	MustString(...string) string
	Int32(...string) (int32, bool)
	MustInt32(...string) int32
	Float64(...string) (float64, bool)
	MustFloat64(...string) float64
	Bool(...string) (bool, bool)
	MustBool(...string) bool
	Generic(...string) (Generic, bool)
	MustGeneric(...string) Generic
	Int32Slice(...string) ([]int32, bool)
	MustInt32Slice(...string) []int32
	Float64Slice(...string) ([]float64, bool)
	MustFloat64Slice(...string) []float64
	BoolSlice(...string) ([]bool, bool)
	MustBoolSlice(...string) []bool
	StringSlice(...string) ([]string, bool)
	MustStringSlice(...string) []string
	GenericSlice(...string) (Slice, bool)
	MustGenericSlice(...string) Slice
	FieldNames() []string
	EachInt32(func(fn *graph.FieldNode, i int32))
	EachFloat64(func(fn *graph.FieldNode, f float64))
	EachBool(func(fn *graph.FieldNode, b bool))
	EachString(func(fn *graph.FieldNode, s string))
	EachGeneric(func(fn *graph.FieldNode, g Generic))
	EachInt32Slice(func(fn *graph.FieldNode, is []int32))
	EachFloat64Slice(func(fn *graph.FieldNode, fs []float64))
	EachBoolSlice(func(fn *graph.FieldNode, bs []bool))
	EachStringSlice(func(fn *graph.FieldNode, ss []string))
	EachGenericSlice(func(fn *graph.FieldNode, gSlice Slice))
	SetString([]string, string) error
	MustSetString([]string, string) Generic
	SetInt32([]string, int32) error
	MustSetInt32([]string, int32) Generic
	SetFloat64([]string, float64) error
	MustSetFloat64([]string, float64) Generic
	SetBool([]string, bool) error
	MustSetBool([]string, bool) Generic
	SetEnum([]string, string) error
	MustSetEnum([]string, string) Generic
	SetGeneric([]string, Generic) error
	MustSetGeneric([]string, Generic) Generic
	SetInt32Slice([]string, []int32) error
	MustSetInt32Slice([]string, []int32) Generic
	SetFloat64Slice([]string, []float64) error
	MustSetFloat64Slice([]string, []float64) Generic
	SetBoolSlice([]string, []bool) error
	MustSetBoolSlice([]string, []bool) Generic
	SetStringSlice([]string, []string) error
	MustSetStringSlice([]string, []string) Generic
	SetGenericSlice([]string, Slice) error
	MustSetGenericSlice([]string, Slice) Generic
	Copy() Generic
	Select(gSelect Generic)
	Type() *graph.TypeNode
	Print()
	PrintDebug()
	Sprint() string
	Sanitize()
	Hash()
	GetHash() string
	ToStruct(interface{}) error
	MustToStruct(interface{})
	ToStringInterfaceMap() map[string]interface{}
	Flatten(delimiter string) (map[string]interface{}, error)
	MustFlatten(delimiter string) map[string]interface{}
}

type Factory interface {
	New(*graph.TypeNode) Generic
	NewSlice(*graph.TypeNode) Slice
	FromStruct(interface{}) (Generic, error)
	MustFromStruct(interface{}) Generic
	FromStructs(interface{}) (Slice, error)
	MustFromStructs(interface{}) Slice
	FromStringInterfaceMap(*graph.TypeNode, map[string]interface{}) (Generic, error)
	MustFromStringInterfaceMap(*graph.TypeNode, map[string]interface{}) Generic
	Unflatten(tn *graph.TypeNode, delimiter string, m map[string]interface{}) (Generic, error)
	MustUnflatten(tn *graph.TypeNode, delimiter string, m map[string]interface{}) Generic
	UnflattenSlice(tn *graph.TypeNode, delimiter string, ms []map[string]interface{}) (Slice, error)
	MustUnflattenSlice(tn *graph.TypeNode, delimiter string, ms []map[string]interface{}) Slice
}

type Slice interface {
	Type() *graph.TypeNode
	Print()
	Sprint() string
	Copy() Slice
	Get() []Generic
	Set([]Generic)
	Append(...Generic)
	Filter(soft bool, gFilter Generic) Slice
	Select(gSelect Generic)
	Hash()
	Flatten(delimiter string) (map[string][]string, error)
	ToStringInterfaceMaps() []map[string]interface{}
	ToStructs(interface{}) error
	MustToStructs(interface{})
}
