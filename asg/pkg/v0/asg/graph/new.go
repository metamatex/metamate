package graph

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/endpointnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/enumnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/fieldflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/words/cardinality"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/words/preposition"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/words/verbs/infinitive"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/words/verbs/past"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/words/verbs/present"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/typenames"
)

const (
	String  = "string"
	Int32   = "int32"
	Float64 = "float64"
	Bool    = "bool"
)

func NewRoot() (root *RootNode) {
	root = NewRootNode()

	addEntities(root)

	root.Types.Each(func(tn *TypeNode) {
		tn.Flags().Set(typeflags.IsEntity, true)
	})

	addExpansion(root)

	root.Types.Flagged(typeflags.IsEntity, false).Each(func(tn *TypeNode) {
		tn.Flags().Set(typeflags.IsExpansion, true)
	})

	addActionEndpoints(root)

	return
}

func addActionEndpoints(root *RootNode) () {
	root.AddActionEndpoint(endpointnames.AuthenticateClientAccount,
		FieldNodeSlice{
			TypeField(fieldnames.Id, typenames.Id),
			StringField(fieldnames.Password),
		},
		FieldNodeSlice{
			TypeField(fieldnames.Token, typenames.Token),
		},
	)

	root.AddActionEndpoint(endpointnames.VerifyToken,
		FieldNodeSlice{
			TypeField(fieldnames.Token, typenames.Token),
		},
		FieldNodeSlice{
			BoolField("isValid"),
			TypeField("clientAccountId", typenames.ServiceId),
		},
	)

	root.AddActionEndpoint(endpointnames.LookupService,
		FieldNodeSlice{},
		FieldNodeSlice{
			TypeField("service", typenames.Service),
		},
	)
}

func addExpansion(root *RootNode) () {
	root.AddTypeNode(typenames.Auth, FieldNodeSlice{
		TypeField(fieldnames.Token, typenames.Token),
		TypeField(fieldnames.ClientAccount, typenames.ClientAccount),
	})

	root.AddTypeNode(typenames.IdGetMode, FieldNodeSlice{
		TypeField(fieldnames.Id, typenames.Id),
	})

	root.AddTypeNode(typenames.RelationGetMode, FieldNodeSlice{
		TypeField(fieldnames.Id, typenames.ServiceId),
		StringField(fieldnames.Relation),
	})

	root.AddTypeNode(typenames.SearchGetMode, FieldNodeSlice{
		StringField("term"),
		TypeField("location", typenames.LocationQuery),
	})

	root.AddTypeNode(typenames.CollectionGetMode, FieldNodeSlice{
		ListField(TypeField("pages", typenames.ServicePage)),
	})

	root.AddUnion(typenames.PostMode, []interface{}{
		TypeField(fieldnames.Collection, typenames.CollectionPostMode),
	})

	root.AddTypeNode(typenames.CollectionPostMode, FieldNodeSlice{})

	root.AddUnion(typenames.GetMode, []interface{}{
		TypeField(fieldnames.Collection, typenames.CollectionGetMode),
		TypeField(fieldnames.Id, typenames.Id),
		TypeField(fieldnames.Relation, typenames.RelationGetMode),
		TypeField(fieldnames.Search, typenames.SearchGetMode),
	})

	root.AddTypeNode(typenames.RelationPutMode, FieldNodeSlice{
		TypeField(fieldnames.Id, typenames.ServiceId),
		StringField(fieldnames.Relation),
		EnumField(fieldnames.Operation, typenames.RelationOperation),
		ListField(TypeField(fieldnames.Ids, typenames.ServiceId)),
	})

	root.AddUnion(typenames.PutMode, []interface{}{
		TypeField(fieldnames.Relation, typenames.RelationPutMode),
	})

	root.AddUnion(typenames.DeleteMode, []interface{}{
		BoolField(fieldnames.Archive),
		BoolField(fieldnames.Permanent),
	})

	root.AddUnion(typenames.PipeMode, []interface{}{
		TypeField(fieldnames.Context, typenames.ContextPipeMode),
	})

	root.AddTypeNode(typenames.ContextPipeMode, FieldNodeSlice{
		EnumField(fieldnames.Stage, enumnames.RequestStage),
		EnumField(fieldnames.Method, enumnames.Methods),
		EnumField(fieldnames.Requester, enumnames.BusActor),
	})

	//root.AddTypeNode(typenames.Trace, FieldNodeSlice{
	//}, Flags{
	//	typeflags.GetEndpoints: true,
	//})
	//
	//root.AddTypeNode(typenames.Span, FieldNodeSlice{
	//	TypeField(fieldnames.Data, typenames.Text),
	//}, Flags{
	//	typeflags.GetEndpoints: true,
	//})
	//
	//root.AddRelationNode(
	//	RelationPath{typenames.Trace, []string{present.Spans}, cardinality.One, typenames.Span},
	//	RelationPath{typenames.Span, []string{past.Spanned, preposition.By}, cardinality.One, typenames.Trace},
	//)
	//
	//root.AddRelationNode(
	//	RelationPath{typenames.Span, []string{present.Spans}, cardinality.Many, typenames.Span},
	//	RelationPath{typenames.Span, []string{past.Spanned, preposition.By}, cardinality.One, typenames.Trace},
	//)

	root.AddEnumNode(enumnames.RequestStage, []string{
		"request",
		"response",
	})

	root.AddEnumNode(enumnames.BusActor, []string{
		"client",
		"bus",
		"service",
	})

	root.AddEnumNode(enumnames.Methods, []string{
		"post",
		"get",
		"put",
		"delete",
		"pipe",
		"action",
	})

	//root.AddTypeNode(Translation, FieldNodeSlice{
	//	EnumField("language", Language),
	//	EnumField("formatting", FormattingKind),
	//	StringField("value"),
	//})

	root.AddEnumNode(enumnames.ErrorKind, []string{
		"service",
		"upstream",
		"pipe",
		"internal",
		"requestValidation",
		"responseValidation",
		"serviceIdNotPresent",
		"serviceIdAlreadyPresent",
		"noServiceMatch",
	})

	root.AddTypeNode(typenames.Error, FieldNodeSlice{
		EnumField(fieldnames.Kind, enumnames.ErrorKind),
		TypeField(fieldnames.Message, typenames.Text),
		TypeField(fieldnames.Id, typenames.Id),
		TypeField(fieldnames.Service, typenames.Service),
		TypeField(fieldnames.Wraps, typenames.Error),
	})

	root.AddTypeNode(typenames.RequestMeta, FieldNodeSlice{
		TypeField("createdAt", typenames.Timestamp),
	})

	root.AddEnumNode(enumnames.ResponseKind, []string{
		"service",
		"bus",
	})

	root.AddTypeNode(typenames.ResponseMeta, FieldNodeSlice{
		EnumField("kind", enumnames.ResponseKind),
		ListField(TypeField("errors", typenames.Error)),
		ListField(TypeField("services", typenames.Service)),
	})

	root.AddTypeNode(typenames.CollectionMeta, FieldNodeSlice{
		ListField(TypeField(fieldnames.Errors, typenames.Error)),
		TypeField(fieldnames.Pagination, typenames.Pagination),
		Int32Field(fieldnames.Count),
	})

	root.AddTypeNode(typenames.IndexPage, FieldNodeSlice{
		Int32Field("page"),
	})

	root.AddTypeNode(typenames.CursorPage, FieldNodeSlice{
		StringField("value"),
	})

	root.AddTypeNode(typenames.OffsetPage, FieldNodeSlice{
		Int32Field("offset"),
		Int32Field("limit"),
	})

	root.AddUnion(typenames.Page, []interface{}{
		typenames.IndexPage,
		typenames.OffsetPage,
		typenames.CursorPage,
	})

	root.AddTypeNode(typenames.ServicePage, FieldNodeSlice{
		TypeField(fieldnames.Service, typenames.Service),
		TypeField("page", typenames.Page),
	})

	root.AddTypeNode(typenames.Pagination, FieldNodeSlice{
		TypeField("previous", typenames.Page),
		TypeField("current", typenames.Page),
		TypeField("next", typenames.Page),
	})

	root.AddEnumNode(enumnames.RelationOperation, []string{
		"add",
		"remove",
	})

	root.AddEnumNode(enumnames.SortKind, []string{
		"asc",
		"desc",
	})

	return
}

func addEntities(root *RootNode) () {
	root.AddEnumNode(enumnames.HashFunction, []string{
		"md5",
		"sha1",
		"sha224",
		"sha256",
		"sha384",
		"sha512",
		"sha3",
	})

	root.AddTypeNode(typenames.ClientAccount, FieldNodeSlice{
		TypeField(fieldnames.Password, typenames.Password),
	}, Flags{
		typeflags.GetEndpoints: true,
	})

	root.AddTypeNode(typenames.ServiceAccount, FieldNodeSlice{
		TypeField(fieldnames.Url, typenames.Url),
		StringField(fieldnames.Handle),
		TypeField(fieldnames.Password, typenames.Password),
	}, Flags{
		typeflags.GetEndpoints: true,
	})

	root.AddTypeNode(typenames.Password, FieldNodeSlice{
		BoolField("isHashed"),
		EnumField("hashFunction", enumnames.HashFunction),
		StringField("value"),
	})

	root.AddTypeNode(typenames.Token, FieldNodeSlice{
		StringField("value"),
	})

	root.AddTypeNode(typenames.TypeMeta, FieldNodeSlice{
		TypeField(fieldnames.Service, typenames.Service),
		BoolField("archived"),
		BoolField("sensitive"),
		TypeField("createdAt", typenames.Timestamp),
		TypeField("updateAt", typenames.Timestamp),
		TypeField("deletedAt", typenames.Timestamp),
	})

	root.AddEnumNode(enumnames.LengthUnit, []string{
		"nm",
		"mcm",
		"mm",
		"cm",
		"dm",
		"m",
		"km",
		"th",
		"in",
		"ft",
		"yd",
		"mi",
		"lea",
	})

	root.AddValue(typenames.LengthValue, enumnames.LengthUnit)

	root.AddScalar(typenames.LengthScalar, enumnames.LengthUnit)

	root.AddTypeNode(typenames.LocationQuery, FieldNodeSlice{
		TypeField("radiusLt", typenames.LengthScalar),
		StringField("street"),
		StringField("zipCode"),
		StringField("city"),
		StringField("cityDistrict"),
		StringField("country"),
		StringField("countryState"),
		StringField("countryStateDistrict"),
	})

	root.AddEnumNode(enumnames.DurationUnit, []string{
		"ns",
		"ms",
		"s",
		"m",
		"h",
		"d",
		"w",
		"y",
	})

	root.AddScalar(typenames.DurationScalar, enumnames.DurationUnit)

	root.AddEnumNode(enumnames.Language, []string{
		"en",
	})

	root.AddEnumNode(enumnames.FormattingKind, []string{
		"html",
		"plain",
		"markdown",
	})

	root.AddTypeNode(typenames.Text, FieldNodeSlice{
		EnumField("language", enumnames.Language),
		EnumField("formatting", enumnames.FormattingKind),
		StringField("value"),
		//ListField(TypeField("translations", Translation)),
	})

	root.AddEnumNode(enumnames.ValueKind, []string{
		"value",
		"range",
	})

	root.AddTypeNode(typenames.FloatRange, FieldNodeSlice{
		Float64Field("from"),
		Float64Field("to"),
	})

	root.AddEnumNode(typenames.ServiceTransport, []string{
		"httpJson",
	})

	root.AddTypeNode(typenames.Service, FieldNodeSlice{
		BoolField(fieldnames.IsVirtual),
		StringField(fieldnames.Name),
		TypeField(fieldnames.Url, typenames.Url),
		EnumField(fieldnames.Transport, typenames.ServiceTransport),
		Int32Field(fieldnames.Port),
	}, Flags{
		typeflags.GetEndpoints: true,
	})

	root.AddTypeNode(typenames.Url, FieldNodeSlice{
		StringField("value", Flags{fieldflags.ValidateUrl: true}),
	}, Flags{
		typeflags.GetPassEndpoint: true,
	})

	root.AddTypeNode(typenames.Email, FieldNodeSlice{
		StringField("value", Flags{fieldflags.ValidateEmail: true, fieldflags.ValidateIsSet: true}),
	})

	root.AddTypeNode(typenames.ServiceId, FieldNodeSlice{
		StringField("value", Flags{fieldflags.ValidateIsSet: true}),
		StringField(fieldnames.ServiceName),
	})

	root.AddUnion(typenames.Id, []interface{}{
		typenames.Url,
		StringField(fieldnames.Name),
		StringField("username"),
		StringField("ean"),
		TypeField(fieldnames.ServiceId, typenames.ServiceId),
		StringField(fieldnames.Local),
		BoolField(fieldnames.Me),
		TypeField(fieldnames.Token, typenames.Token),
		typenames.Email,
	})

	root.AddEnumNode(typenames.TimestampKind, []string{
		"unix",
		"unixNs",
	})

	root.AddTypeNode(typenames.Timestamp, FieldNodeSlice{
		EnumField("kind", typenames.TimestampKind),
		StringField("value"),
	})

	root.AddTypeNode(typenames.Image, FieldNodeSlice{
		TypeField("url", typenames.Url),
		BoolField("isPreview"),
		Int32Field("width"),
		Int32Field("height"),
		TypeField("description", typenames.Text),
	})

	root.AddTypeNode(typenames.Person, FieldNodeSlice{
		TypeField("username", typenames.Text),
		TypeField("displayName", typenames.Text),
		TypeField("note", typenames.Text),
		TypeField("avatar", typenames.Image),
		TypeField("header", typenames.Image),
	}, Flags{
		typeflags.GetEndpoints: true,
	})

	root.AddTypeNode(typenames.Attachment, FieldNodeSlice{
		StringField("description"),
	}, Flags{
		typeflags.GetEndpoints: true,
	})

	root.AddTypeNode(typenames.Info, FieldNodeSlice{
		TypeField("name", typenames.Text),
		TypeField("description", typenames.Text),
		TypeField("purpose", typenames.Text),
	})

	root.AddEnumNode(typenames.FeedKind, []string{
		"channel",
		"privateChannel",
		"conversation",
	})

	root.AddTypeNode(typenames.Feed, FieldNodeSlice{
		TypeField("info", typenames.Info),
		EnumField("kind", typenames.FeedKind),
	}, Flags{
		typeflags.GetEndpoints: true,
	})

	root.AddTypeNode(typenames.Status, FieldNodeSlice{
		TypeField("content", typenames.Text),
		TypeField("spoilerText", typenames.Text),
		BoolField("sensitive"),
		BoolField("pinned"),
	}, Flags{
		typeflags.GetEndpoints: true,
	})

	root.AddRelationNode(
		RelationPath{typenames.Person, []string{present.Authors}, cardinality.Many, typenames.Status},
		RelationPath{typenames.Status, []string{past.Authored, preposition.By}, cardinality.One, typenames.Person},
	)

	root.AddRelationNode(
		RelationPath{typenames.Person, []string{present.Follows}, cardinality.Many, typenames.Person},
		RelationPath{typenames.Person, []string{past.Followed, preposition.By}, cardinality.Many, typenames.Person},
	)

	root.AddRelationNode(
		RelationPath{typenames.Person, []string{present.Mutes}, cardinality.Many, typenames.Person},
		RelationPath{typenames.Person, []string{past.Muted, preposition.By}, cardinality.Many, typenames.Person},
	)

	root.AddRelationNode(
		RelationPath{typenames.Person, []string{present.Requests, preposition.To, infinitive.Follow}, cardinality.Many, typenames.Person},
		RelationPath{typenames.Person, []string{past.Requested, preposition.To, infinitive.Be, past.Followed, preposition.By}, cardinality.Many, typenames.Person},
	)

	root.AddRelationNode(
		RelationPath{typenames.Person, []string{present.Blocks}, cardinality.Many, typenames.Person},
		RelationPath{typenames.Person, []string{past.Blocked, preposition.By}, cardinality.Many, typenames.Person},
	)

	root.AddRelationNode(
		RelationPath{typenames.Status, []string{present.Replies, preposition.To}, cardinality.One, typenames.Person},
		RelationPath{typenames.Person, []string{past.Was, past.Replied, preposition.To, preposition.By}, cardinality.Many, typenames.Status},
	)

	root.AddRelationNode(
		RelationPath{typenames.Status, []string{present.Replies, preposition.To}, cardinality.One, typenames.Status},
		RelationPath{typenames.Status, []string{past.Was, past.Replied, preposition.To, preposition.By}, cardinality.Many, typenames.Status},
	)

	root.AddRelationNode(
		RelationPath{typenames.Status, []string{present.Reblogs}, cardinality.One, typenames.Status},
		RelationPath{typenames.Status, []string{past.Reblogged, preposition.By}, cardinality.Many, typenames.Status},
	)

	root.AddRelationNode(
		RelationPath{typenames.Status, []string{present.Mentions}, cardinality.Many, typenames.Person},
		RelationPath{typenames.Person, []string{past.Mentioned, preposition.By}, cardinality.Many, typenames.Status},
	)

	root.AddRelationNode(
		RelationPath{typenames.Person, []string{present.Favors}, cardinality.Many, typenames.Status},
		RelationPath{typenames.Status, []string{past.Favored, preposition.By}, cardinality.Many, typenames.Person},
	)

	root.AddRelationNode(
		RelationPath{typenames.Status, []string{present.Attaches}, cardinality.Many, typenames.Attachment},
		RelationPath{typenames.Attachment, []string{past.Attached, preposition.To}, cardinality.One, typenames.Status},
	)

	root.AddRelationNode(
		RelationPath{typenames.Feed, []string{present.Contains}, cardinality.Many, typenames.Status},
		RelationPath{typenames.Status, []string{past.Contained, preposition.By}, cardinality.Many, typenames.Feed},
	)

	root.AddRelationNode(
		RelationPath{typenames.Person, []string{present.Participates}, cardinality.Many, typenames.Feed},
		RelationPath{typenames.Feed, []string{past.Participated, preposition.By}, cardinality.Many, typenames.Person},
	)

	root.AddRelationNode(
		RelationPath{typenames.Person, []string{past.Read}, cardinality.Many, typenames.Status},
		RelationPath{typenames.Status, []string{past.Read, preposition.By}, cardinality.Many, typenames.Person},
	)

	root.AddRelationNode(
		RelationPath{typenames.Person, []string{preposition.Not, past.Read}, cardinality.Many, typenames.Status},
		RelationPath{typenames.Status, []string{preposition.Not, past.Read, preposition.By}, cardinality.Many, typenames.Person},
	)

	root.AddRelationNode(
		RelationPath{typenames.Person, []string{present.Reblogs}, cardinality.Many, typenames.Status},
		RelationPath{typenames.Status, []string{past.Reblogged, preposition.By}, cardinality.Many, typenames.Person},
	)

	root.AddRelationNode(
		RelationPath{typenames.Person, []string{present.Mutes}, cardinality.Many, typenames.Status},
		RelationPath{typenames.Status, []string{past.Muted, preposition.By}, cardinality.Many, typenames.Person},
	)

	root.AddRelationNode(
		RelationPath{typenames.Whatever, []string{present.Knows}, cardinality.Many, typenames.Whatever},
		RelationPath{typenames.Whatever, []string{past.Knew, preposition.By}, cardinality.Many, typenames.Whatever},
	)

	root.AddRelationNode(
		RelationPath{typenames.Whatever, []string{present.Knows}, cardinality.Many, typenames.BlueWhatever},
		RelationPath{typenames.BlueWhatever, []string{past.Knew, preposition.By}, cardinality.Many, typenames.Whatever},
	)

	root.AddRelationNode(
		RelationPath{typenames.ClientAccount, []string{present.Owns}, cardinality.Many, typenames.ServiceAccount},
		RelationPath{typenames.ServiceAccount, []string{past.Owned, preposition.By}, cardinality.One, typenames.ClientAccount},
	)

	root.AddRelationNode(
		RelationPath{typenames.Service, []string{present.Uses}, cardinality.Many, typenames.ServiceAccount},
		RelationPath{typenames.ServiceAccount, []string{past.Used, preposition.By}, cardinality.Many, typenames.Service},
	)

	root.AddEnumNode(enumnames.WhateverKind, []string{
		"red",
		"blue",
		"green",
	})

	root.AddTypeNode(typenames.Whatever, FieldNodeSlice{
		EnumField("enumField", enumnames.WhateverKind),
		StringField("stringField"),
		Int32Field("int32Field"),
		Float64Field("float64Field"),
		BoolField("boolField"),
		TypeField("unionField", typenames.WhateverUnion),
		ListField(EnumField("enumList", enumnames.WhateverKind)),
		ListField(StringField("stringList")),
		ListField(Int32Field("int32List")),
		ListField(Float64Field("float64List")),
		ListField(BoolField("boolList")),
		//ListField(TypeField("typeList", Whatever)),
		ListField(TypeField("unionList", typenames.WhateverUnion)),
	}, Flags{
		typeflags.GetEndpoints: true,
	})

	root.AddTypeNode(typenames.BlueWhatever, FieldNodeSlice{
		EnumField("enumField", enumnames.WhateverKind),
		StringField("stringField"),
		Int32Field("int32Field"),
		Float64Field("float64Field"),
		BoolField("boolField"),
		TypeField("unionField", typenames.WhateverUnion),
		ListField(EnumField("enumList", enumnames.WhateverKind)),
		ListField(StringField("stringList")),
		ListField(Int32Field("int32List")),
		ListField(Float64Field("float64List")),
		ListField(BoolField("boolList")),
		//ListField(TypeField("typeList", Whatever)),
		ListField(TypeField("unionList", typenames.WhateverUnion)),
	}, Flags{
		typeflags.GetEndpoints: true,
	})

	root.AddUnion(typenames.WhateverUnion, []interface{}{
		StringField("stringField"),
		Int32Field("int32Field"),
		Float64Field("float64Field"),
		BoolField("boolField"),
		EnumField("enumField", enumnames.WhateverKind),
	})

	root.AddEnumNode(enumnames.CurrencyUnit, []string{
		"eur",
	})

	root.AddScalar(typenames.CurrencyScalar, enumnames.CurrencyUnit)

	root.AddTypeNode(typenames.BankAccount, FieldNodeSlice{
		TypeField("balance", typenames.CurrencyScalar),
		StringField("iban"),
		StringField("bankName"),
	}, Flags{
		typeflags.GetEndpoints: true,
	})

	root.AddTypeNode(typenames.Transaction, FieldNodeSlice{
		TypeField("amount", typenames.CurrencyScalar),
		StringField("payee"),
		StringField("category"),
		StringField("reference"),
		TypeField("createdAt", typenames.Timestamp),
	}, Flags{
		typeflags.GetEndpoints: true,
	})
}
