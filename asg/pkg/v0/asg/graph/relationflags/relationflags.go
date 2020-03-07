package relationflags

const (
	All  = "all"
	Base = "base"

	One2Many  = "one2Many"
	Many2Many = "many2Many"

	IsSelfReferencing = "isSelfReferencing"
)

var Defaults = map[string]bool{
	All:  true,
	Base: false,

	One2Many:  false,
	Many2Many:  false,

	IsSelfReferencing: false,
}
