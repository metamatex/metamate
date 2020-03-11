# schema

MetaMate comes with an abstract data schema that is progressively modeled around real world data. The schema is defined as code to allow applications to interact programmatically with it. 

It comes with two layers:

The [entities](pkg/entities) layer adds [base](pkg/entities/base) and [domain-specific](pkg/entities/domain) entities to the schema. 

The [expansion](pkg/expansion) layer processes the schema and adds types like [Select](pkg/expansion/select.go), [Filter](pkg/expansion/filter.go), [Sort](pkg/expansion/sort.go) and [Requests, Responses and Endpoints](pkg/expansion/communication.go) for the entities.

### Contents

- [Usage](#usage)
- [Type system](#type-system)
  - [Endpoint](#endpoint)
  - [Type](#type)
  - [Enum](#enum)
- [Type Construtors](#type-construtors)
  - [Union](#union)
  - [Value](#value)
  - [Scalar](#scalar)
  - [Range](#range)
  - [Ratio](#ratio)
- [Fields](#fields)

## Usage

`metamatex/schema` simply provides programmatic access of the schema to other applications.

[metamatex/metactl](https://github.com/metamatex/metamate/metactl) uses it to populate user supplied templates, allowing to e.g. generate `.proto2` definitions, `graphql` schemas, docs, sdks and much more.

[metamatex/metamate](https://github.com/metamatex/metamate) uses it to power it's generic type system.


## Type system

The type system is derived as the lowest common denominator of the capabilities of the most widely adopted communication protocols, like grpc, graphql, thrift and openapi. Therefore the type system can be seamlessly mapped onto these protocols.

The 3 kinds of top-level types are `Endpoint`, `Type`, and `Enum`

### Endpoint 

[schema/pkg/expand/communication.go](pkg/expand/communication.go)
```golang
schema.AddEndpoint(GET, "Get", t.Name, GET_REQUEST, GET_RESPONSE, endpointflags.EndpointFlags{
  endpointflags.IsGetEndpoint: true,
  endpointflags.IsQuery:       true,
  endpointflags.IsCacheable:   true,
})
```
- defines the communication layer
- solely relies on `Type`
- holds binary flags - [available flags and their defaults](pkg/def/endpointflags/endpointFlags.go)

### Type

[schema/schemas/base/text.go](pkg/schemas/base/text.go)
```golang
schema.AddType(Text, []def.Field{
  Enum("language", Language),
  Enum("formatting", FormattingKind),
  String("value"),
  List(Object("translations", Translation)),
})
```
- defines data structures
- holds `Fields`
- holds binary flags - [available flags and their defaults](pkg/def/typeflags/typeFlags.go)
- holds relations to other `Type` - [available relations](pkg/def/relation/relation.go)

[schema/pkg/expand/filter.go](pkg/expand/filter.go)
```golang
schema.AddType(filterName, fields, typeflags.TypeFlags{
  typeflags.IsFilter: true,
  typeflags.GetFilter: t.HasAnyFieldFlagDeep(fieldflags.Filter, true),
}, relation.RelationNames{
  relation.FOR: t.Name,
})
```

### Enum

[schema/schemas/base/values.go](pkg/schemas/base/values.go)
```golang
schema.AddEnum(DurationUnit, []string{
  "NANOSECOND",
  "MILLISECOND",
  "SECOND",
  "MINUTE",
  "HOUR",
  "DAY",
  "WEEK",
  "MONTH",
  "YEAR",
})
```
- defines a named set of values

## Type constructors

Type constructors are utilities that create commonly used type constructs. Under the hood they may create multiple `Type` and `Enum` and set `TypeFlags`

### Union

[schema/schemas/base/id.go](pkg/schemas/base/id.go)
```golang
schema.AddUnion(Id, []interface{}{
  Url,
  String("name"),
  String("username"),
  String("ean"),
  ServiceId,
  String("local"),
  Bool("me"),
})
```

- holds one of a defined set of `Field`
- holds a kind `Enum` indicating what value the union holds
- corresponds to

```golang
schema.AddEnum(IdKind, []interface{}{
  "ID_URL",
  "ID_NAME",
  "ID_USERNAME",
  "ID_EAN",
  "ID_SERVICE_ID",
  "ID_LOCAL",
  "ID_ME",
})

schema.AddType(Id, []interface{}{
  Enum("kind", IdKind, FieldValidation{ValidateIsSet: true},
  Object("url", Url, fieldflags.FieldFlags{IsUnionField: true}),
  String("name", fieldflags.FieldFlags{IsUnionField: true}),
  String("username", fieldflags.FieldFlags{IsUnionField: true}),
  String("ean", fieldflags.FieldFlags{IsUnionField: true}),
  Object("serviceId", ServiceId, fieldflags.FieldFlags{IsUnionField: true}),
  String("local", fieldflags.FieldFlags{IsUnionField: true}),
  Bool("me", fieldflags.FieldFlags{IsUnionField: true}),
}, typeflags.TypeFlags{
  typeflags.IsUnion: true,
})
```

### Value

[schema/schemas/base/values.go](pkg/schemas/base/values.go)
```golang
schema.AddValue(DurationValue, DurationUnit)
```

- holds a unit `Enum`
- either holds a scalar `Float64` value or a `FloatRange`
- corresponds to

```golang
schema.AddType(DurationValue, []Field{
  EnumField("kind", ValueKind),
  EnumField("unit", DurationUnit),
  ObjectField("range", FloatRange),
  Float64Field("value"),
  BoolField("isEstimate"),
}, typeflags.TypeFlags{
  typeflags.IsValue: true,
})
```

### Scalar

[schema/schemas/base/values.go](pkg/schemas/base/values.go)
```golang
schema.AddScalar(DurationScalar, DurationUnit)
```

- holds a unit `Enum`
- holds a scalar `Float64`
- corresponds to

```golang
schema.AddType(DurationScalar, []Field{
  EnumField("unit", DurationUnit),
  Float64Field("value", FieldValidation{ValidateIsSet: true},
  BoolField("isEstimate"),
}, typeflags.TypeFlags{
  typeflags.IsScalar: true,
})
```

### Range

[schema/schemas/base/values.go](pkg/schemas/base/values.go)
```golang
schema.AddRange(DurationRange, DurationUnit)
```

- holds a unit `Enum`
- holds a `FloatRange`
- corresponds to

```golang
schema.AddType(DurationRange, []Field{
  EnumField("unit", DurationUnit),
  ObjectField("range", FloatRange),
  BoolField("isEstimate"),
}, typeflags.TypeFlags{
  typeflags.IsRange: true,
})
```

### Ratio

[schema/schemas/base/values.go](pkg/schemas/base/values.go)
```golang
schema.AddRatio(CurrencyPerDurationRatio, CurrencyValue, DurationValue)
```

- holds a counter `Object`
- holds a divider `Object`

```golang
schema.AddType(CurrencyPerDurationRatio, []Field{
  ObjectField("counter", CurrencyValue),
  ObjectField("divider", DurationValue),
}, typeflags.TypeFlags{
  typeflags.IsRatio: true,
})
```

## Fields

[schema/pkg/def/specification.go](pkg/schemas/base/text.go)
```golang
schema.AddType(Whatever, []def.Field{
  Enum("enumField", WhateverKind),
  String("stringField"),
  Int32("intField"),
  Float64("floatField"),
  Bool("boolField"),
  Object("objectField", Whatever),
  Object("unionField", WhateverUnion),
  List(Enum("enumList", WhateverKind)),
  List(String("stringList")),
  List(Int32("intList")),
  List(Float64("floatList")),
  List(Bool("boolList")),
  List(Object("objectList", Whatever)),
  List(Object("unionList", WhateverUnion)),
}, typeflags.TypeFlags{
  typeflags.GetDefaultEndpoints: true,
})
```

- field types are `Object`, `Enum`, `List`, `Float64`, `Int32`, `String` and `Bool`

#### show tag usage
