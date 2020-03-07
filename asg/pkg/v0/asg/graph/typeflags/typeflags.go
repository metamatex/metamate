package typeflags

const (
	// Sort
	IsSort  = "isSort"
	HasSort = "hasSort"
	GetSort = "getSort"

	// Filter
	IsFilter      = "isFilter"
	IsBasicFilter = "isBasicFilter"
	GetFilter     = "getFilter"
	HasFilter     = "hasFilter"
	IsListFilter  = "isListFilter"
	HasListFilter = "hasListFilter"

	// Filter
	IsSelect  = "isSelect"
	HasSelect = "hasSelect"

	// Characteristics
	HasListKind        = "hasListKind"
	HasTypeMeta        = "hasTypeMeta"
	IsInList           = "isInList"
	HasFieldValidation = "hasFieldValidation"

	// Request
	IsRequest       = "isRequest"
	IsResponse      = "isResponse"
	IsGetRequest    = "isGetRequest"
	IsActionRequest = "isActionRequest"
	IsGetCollection = "isGetCollection"
	IsGetRelations  = "isGetRelations"
	IsPostRequest   = "isPostRequest"
	IsPipeRequest   = "isPipeRequest"
	IsPutRequest    = "isPutRequest"
	IsDeleteRequest = "isDeleteRequest"

	// Response
	IsCollection = "isCollection"

	// Scope
	RequestScope  = "requestScope"
	ResponseScope = "responseScope"

	// AddEndpoint
	GetEndpoints    = "getEndpoints"
	GetPassEndpoint = "getPassEndpoint"
	HasPassEndpoint = "hasPassEndpoint"

	All         = "all"
	Base        = "base"
	IsEndpoint  = "isEndpoint"
	IsEntity    = "isEntity"
	IsExpansion = "isExpansion"

	// AddType
	Enum                 = "enum"
	Object               = "object"
	IsUnion              = "isUnion"
	IsOptionalValueUnion = "isOptionalValueUnion"
	IsScalar             = "isScalar"
	IsValue              = "isValue"
	IsRange              = "isRange"
	IsRatio              = "isRatio"
	IsRelationships      = "isRelationships"
	IsRelations          = "isRelations"
	IsRelation           = "isRelation"
	IsFromToManyRelation = "isFromToManyRelation"
	IsFromToOneRelation  = "isFromToOneRelation"
)

var Defaults = map[string]bool{
	// Sort
	GetSort: true,
	IsSort:  false,
	HasSort: false,

	// Filter
	IsFilter:      false,
	IsBasicFilter: false,
	GetFilter:     true,
	HasFilter:     false,
	IsListFilter:  false,
	HasListFilter: false,

	// Filter
	IsSelect:  false,
	HasSelect: false,

	// Characteristics
	HasListKind:        false,
	HasTypeMeta:        false,
	IsInList:           false,
	HasFieldValidation: false,

	// Request
	IsRequest:       false,
	IsGetRequest:    false,
	IsActionRequest: false,
	IsGetCollection: false,
	IsGetRelations:  false,
	IsPostRequest:   false,
	IsPutRequest:    false,
	IsDeleteRequest: false,
	IsResponse:      false,

	// Response
	IsCollection: false,

	// Scope
	RequestScope:  false,
	ResponseScope: false,

	// AddEndpoint
	GetEndpoints:    false,
	GetPassEndpoint: false,
	HasPassEndpoint: false,

	All:         true,
	Base:        false,
	IsEndpoint:  false,
	IsEntity:    false,
	IsExpansion: false,

	// AddType
	IsUnion:              false,
	IsOptionalValueUnion: false,
	IsScalar:             false,
	IsValue:              false,
	IsRange:              false,
	IsRatio:              false,
	IsRelationships:      false,
	IsRelations:          false,
	IsRelation:           false,
	IsFromToManyRelation: false,
	IsFromToOneRelation:  false,
}
