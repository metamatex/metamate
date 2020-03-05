package fieldflags

const (
	All            = "all"
	Base           = "base"
	Sort           = "sort"
	Filter         = "filter"
	Hash           = "hash"
	IsHash         = "isHash"
	Select         = "select"
	IsList         = "isList"
	IsForList      = "isForList"
	IsUnionField   = "isUnionField"
	IsStringField  = "isStringField"
	IsInt32Field   = "isInt32Field"
	IsFloat64Field = "isFloat64Field"
	IsBoolField    = "isBoolField"
	IsEnumField    = "isEnumField"
	IsObjectField  = "isObjectField"

	// Validate
	ValidateIsSet = "validateIsSet"
	ValidateEmail = "validateEmail"
	ValidateUrl   = "validateUrl"
	ValidateHost  = "validateHost"
)

var Defaults = map[string]bool{
	All:            true,
	Base:           false,
	Sort:           true,
	Filter:         true,
	Hash:           true,
	IsHash:         false,
	Select:         false,
	IsList:         false,
	IsForList:      false,
	IsUnionField:   false,
	IsStringField:  false,
	IsInt32Field:   false,
	IsFloat64Field: false,
	IsBoolField:    false,
	IsEnumField:    false,
	IsObjectField:  false,

	// Validate
	ValidateIsSet: false,
	ValidateEmail: false,
	ValidateUrl:   false,
	ValidateHost:  false,
}
