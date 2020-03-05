package endpointflags

const (
	All  = "all"
	Base = "base"

	IsCacheable              = "isCacheable"
	IsQuery                  = "isQuery"
	IsMutation               = "isMutation"
	IsStream                 = "isStream"
	IsActionEndpoint         = "isActionEndpoint"
	IsGetEndpoint            = "isGetEndpoint"
	IsGetByIdEndpoint        = "isGetByIdEndpoint"
	IsGetListByIdEndpoint    = "isGetListByIdEndpoint"
	IsPostEndpoint           = "isPostEndpoint"
	IsPipeEndpoint           = "isPipeEndpoint"
	IsPutEndpoint            = "isPutEndpoint"
	IsDeleteEndpoint         = "isDeleteEndpoint"
	IsPassEndpoint           = "isPassEndpoint"
	IsModifyListByIdEndpoint = "isModifyListByIdEndpoint"
	IsStreamEndpoint         = "isStreamEndpoint"
)

var Defaults = map[string]bool{
	All:                      true,
	Base:                     false,
	IsCacheable:              false,
	IsQuery:                  false,
	IsMutation:               false,
	IsStream:                 false,
	IsActionEndpoint:         false,
	IsGetEndpoint:            false,
	IsGetByIdEndpoint:        false,
	IsGetListByIdEndpoint:    false,
	IsPostEndpoint:           false,
	IsPipeEndpoint:           false,
	IsPutEndpoint:            false,
	IsDeleteEndpoint:         false,
	IsPassEndpoint:           false,
	IsModifyListByIdEndpoint: false,
	IsStreamEndpoint:         false,
}
