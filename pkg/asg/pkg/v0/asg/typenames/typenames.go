package typenames

import (
	"github.com/iancoleman/strcase"
	"github.com/metamatex/asg/pkg/v0/asg/utils"
)

const (
	Endpoints         = "Endpoints"
	StringFilter      = "StringFilter"
	Int32Filter       = "Int32Filter"
	Float64Filter     = "Float64Filter"
	BoolFilter        = "BoolFilter"
	EnumFilter        = "EnumFilter"
	StringListFilter  = "StringListFilter"
	EnumListFilter    = "EnumListFilter"
	Int32ListFilter   = "Int32ListFilter"
	Float64ListFilter = "Float64ListFilter"
	BoolListFilter    = "BoolListFilter"
	ValueKind         = "ValueKind"
	FloatRange        = "FloatRange"
	ServiceTransport  = "ServiceTransport"
	Service           = "Service"
	Error             = "Error"
	Url               = "Url"
	Email             = "Email"
	ServiceId         = "ServiceId"
	Id                = "Id"
	MessageError      = "MessageError"
	TimestampKind     = "TimestampKind"
	Timestamp         = "Timestamp"
	RequestMeta       = "RequestMeta"
	ResponseKind      = "ResponseKind"
	ResponseMeta      = "ResponseMeta"
	CollectionMeta    = "CollectionMeta"
	TypeMeta          = "TypeMeta"

	Translation = "Translation"
	Text        = "Text"

	Info       = "Info"
	Person     = "Person"
	Status     = "Status"
	Feed       = "Feed"
	FeedKind   = "FeedKind"
	Attachment = "Attachment"

	DurationScalar = "DurationScalar"

	LocationQuery = "LocationQuery"
	Cache         = "Cache"

	LengthValue = "LengthValue"
	LengthScalar = "LengthScalar"

	Image = "Image"

	OffsetPage  = "OffsetPage"
	CursorPage  = "CursorPage"
	IndexPage   = "IndexPage"
	Page        = "Page"
	Pagination  = "Pagination"
	ServicePage = "ServicePage"

	RelationOperation = "RelationOperation"

	PostMode           = "PostMode"
	CollectionPostMode = "CollectionPostMode"

	DeleteMode           = "DeleteMode"
	CollectionDeleteMode = "CollectionDeleteMode"

	GetMode           = "GetMode"
	CollectionGetMode = "CollectionGetMode"
	IdGetMode         = "IdGetMode"
	RelationGetMode   = "RelationGetMode"
	SearchGetMode     = "SearchGetMode"

	PutMode         = "PutMode"
	RelationPutMode = "RelationPutMode"

	PipeMode        = "PipeMode"
	ContextPipeMode = "ContextPipeMode"

	Trace = "Trace"
	Span  = "Span"

	Whatever      = "Whatever"
	BlueWhatever  = "BlueWhatever"
	WhateverUnion = "WhateverUnion"

	ClientAccount  = "ClientAccount"
	ServiceAccount = "ServiceAccount"
	Password       = "Password"
	Token          = "Token"
	Auth           = "Auth"

	CurrencyScalar = "CurrencyScalar"
	Transaction = "Transaction"
	BankAccount     = "BankAccount"
)

func Sort(name string) (string) {
	return name + "Sort"
}

func Filter(name string) (string) {
	return name + "Filter"
}

func addSelectPluralSuffix(name string) (string) {
	return utils.Plural(name) + "Filter"
}

func ListKind(name string) (string) {
	return utils.Plural(name) + "ListKind"
}

func Select(name string) (string) {
	return name + "Select"
}

func Relationship(from, to string) (string) {
	return from + to + "Relationship"
}

func Relationships(name string) (string) {
	return name + "Relationships"
}

func Relation(name string) (string) {
	return utils.Plural(name) + "Relation"
}

func FromToManyRelation(from, to string) (string) {
	return from + "To" + utils.Plural(to) + "Relation"
}

func FromToOneRelation(from, to string) (string) {
	return from + "To" + to + "Relation"
}

func Relations(name string) (string) {
	return name + "Relations"
}

func RelationName(name string) (string) {
	return name + "RelationName"
}

func Request(name string) (string) {
	return name + "Request"
}

func Input(name string) (string) {
	return name + "Input"
}

func Output(name string) (string) {
	return name + "Output"
}

func ListFilter(name string) (string) {
	return name + "ListFilter"
}

func Response(name string) (string) {
	return name + "Response"
}

func Endpoint(name string) (string) {
	return name + "Endpoint"
}

func FieldName(name string) (string) {
	return strcase.ToLowerCamel(name)
}

func GetRequest(name string) (string) {
	return "Get" + utils.Plural(name) + "Request"
}

func GetResponse(name string) (string) {
	return "Get" + utils.Plural(name) + "Response"
}

func PostRequest(name string) (string) {
	return "Post" + utils.Plural(name) + "Request"
}

func PostResponse(name string) (string) {
	return "Post" + utils.Plural(name) + "Response"
}

func PutRequest(name string) (string) {
	return "Put" + utils.Plural(name) + "Request"
}

func PutResponse(name string) (string) {
	return "Put" + utils.Plural(name) + "Response"
}

func DeleteRequest(name string) (string) {
	return "Delete" + utils.Plural(name) + "Request"
}

func DeleteResponse(name string) (string) {
	return "Delete" + utils.Plural(name) + "Response"
}

func PipeRequest(name string) (string) {
	return "Pipe" + utils.Plural(name) + "Request"
}

func PipeResponse(name string) (string) {
	return "Pipe" + utils.Plural(name) + "Response"
}

func PipeContext(name string) (string) {
	return "Pipe" + utils.Plural(name) + "Context"
}

func PipePostContext(name string) (string) {
	return "PipePost" + utils.Plural(name) + "Context"
}

func PipeGetContext(name string) (string) {
	return "PipeGet" + utils.Plural(name) + "Context"
}

func PipePutContext(name string) (string) {
	return "PipePut" + utils.Plural(name) + "Context"
}

func PipeDeleteContext(name string) (string) {
	return "PipeDelete" + utils.Plural(name) + "Context"
}

func GetRelations(name string) (string) {
	return "Get" + utils.Plural(name) + "Relations"
}

func GetCollection(name string) (string) {
	return "Get" + utils.Plural(name) + "Collection"
}

func Collection(name string) (string) {
	return utils.Plural(name) + "Collection"
}

func GetNode(name string) (string) {
	return "Get" + utils.Plural(name) + "Node"
}
