package enumflags

const (
	All = "all"
	Base = "base"

	IsRelationNames = "isRelationNames"
)

var Defaults = map[string]bool{
	All:  true,
	Base: false,
	IsRelationNames: false,
}