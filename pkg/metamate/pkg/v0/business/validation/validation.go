package validation

import (
	"errors"
	"fmt"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/metamatex/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/asg/pkg/v0/asg/graph"
	"github.com/metamatex/asg/pkg/v0/asg/graph/fieldflags"
	"github.com/metamatex/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/metamatemono/pkg/generic/pkg/v0/generic"
	"github.com/metamatex/metamatemono/gen/v0/sdk"
	"github.com/metamatex/metamatemono/gen/v0/sdk/utils/ptr"
	"github.com/metamatex/metamatemono/pkg/metamate/pkg/v0/types"
)

var validators = map[string][]func(stage string, svc *sdk.Service, g generic.Generic) (errs []sdk.Error){}
var sliceValidators = map[string][]func(stage string, svc *sdk.Service, g generic.Slice) (errs []sdk.Error){}

func init() {
	sliceValidators[sdk.IdName] = append(sliceValidators[sdk.IdName], func(stage string, svc *sdk.Service, ids generic.Slice) (errs []sdk.Error) {
		serviceNameCounter := map[string]int{}

		for _, gId := range ids.Get() {
			kind, _ := gId.String(fieldnames.Kind)

			if kind != sdk.IdKind.ServiceId {
				continue
			}

			_, ok := gId.Generic(fieldnames.ServiceId)
			if !ok {
				continue
			}

			serviceName, ok := gId.String(fieldnames.ServiceId, fieldnames.ServiceName)
			if !ok {
				continue
			}

			serviceNameCounter[serviceName]++
		}

		for serviceName, c := range serviceNameCounter {
			if c == 1 {
				continue
			}

			errs = append(errs, getError(stage, svc, fmt.Sprintf("ids can't contain two service ids of the same service, %v service ids of service %v occurred", c, serviceName)))
		}

		return
	})
}

func Validate(stage string, svc *sdk.Service, g generic.Generic) (errs []sdk.Error) {
	g.EachGeneric(func(fn *graph.FieldNode, g0 generic.Generic) {
		errs = append(errs, Validate(stage, svc, g0)...)
	})

	g.EachGenericSlice(func(fn *graph.FieldNode, gs generic.Slice) {
		errs = append(errs, ValidateSlice(stage, svc, gs)...)
	})

	isSort := g.Type().Flags().Is(typeflags.IsSort, true)
	isUnion := g.Type().Flags().Is(typeflags.IsUnion, true)

	if isSort {
		errs = append(errs, ValidateSort(stage, svc, g)...)
	}

	if isUnion {
		errs = append(errs, ValidateUnion(stage, svc, g)...)
	}

	errs = append(errs, ValidateFields(stage, svc, g)...)

	for _, p := range validators[g.Type().Name()] {
		errs = append(errs, p(stage, svc, g)...)
	}

	return
}

func ValidateSlice(stage string, svc *sdk.Service, gs generic.Slice) (errs []sdk.Error) {
	for _, g := range gs.Get() {
		errs = append(errs, Validate(stage, svc, g)...)
	}

	for _, p := range sliceValidators[gs.Type().Name()] {
		errs = append(errs, p(stage, svc, gs)...)
	}

	return
}

func ValidateSort(stage string, svc *sdk.Service, g generic.Generic) (errs []sdk.Error) {
	c := 0

	g.EachString(func(_ *graph.FieldNode, _ string) {
		c++
	})

	if c != 1 {
		errs = append(errs, getError(stage, svc, "exactly one sort value needs to be set"))
	}

	return
}

func ValidateFields(stage string, svc *sdk.Service, g generic.Generic) (errs []sdk.Error) {
	g.Type().Edges.Fields.Holds().Each(func(fn *graph.FieldNode) {
		if fn.Flags().Is(fieldflags.ValidateIsSet, true) {
			err := validateIsSet(g, fn)
			if err != nil {
				errs = append(errs, getError(stage, svc, err.Error()))
			}
		}

		if fn.Flags().Is(fieldflags.ValidateEmail, true) {
			err := validateEmail(g, fn.Name())
			if err != nil {
				errs = append(errs, getError(stage, svc, err.Error()))
			}
		}
	})

	return
}

func validateIsSet(g generic.Generic, fn *graph.FieldNode) (err error) {
	switch fn.Kind() {
	case graph.FieldKindType:
		_, ok := g.Generic(fn.Name())
		if ok {
			return
		}
	case graph.FieldKindEnum:
		_, ok := g.String(fn.Name())
		if ok {
			return
		}
	case graph.FieldKindString:
		_, ok := g.String(fn.Name())
		if ok {
			return
		}
	case graph.FieldKindInt32:
		_, ok := g.Int32(fn.Name())
		if ok {
			return
		}
	case graph.FieldKindFloat64:
		_, ok := g.Float64(fn.Name())
		if ok {
			return
		}
	case graph.FieldKindBool:
		_, ok := g.Bool(fn.Name())
		if ok {
			return
		}
	case graph.FieldKindTypeList:
		_, ok := g.GenericSlice(fn.Name())
		if ok {
			return
		}
	case graph.FieldKindEnumList:
		_, ok := g.StringSlice(fn.Name())
		if ok {
			return
		}
	case graph.FieldKindStringList:
		_, ok := g.StringSlice(fn.Name())
		if ok {
			return
		}
	case graph.FieldKindInt32List:
		_, ok := g.Int32Slice(fn.Name())
		if ok {
			return
		}
	case graph.FieldKindFloat64List:
		_, ok := g.Float64Slice(fn.Name())
		if ok {
			return
		}
	case graph.FieldKindBoolList:
		_, ok := g.BoolSlice(fn.Name())
		if ok {
			return
		}
	default:
		err = errors.New(fmt.Sprintf("unknown field kind %v", fn.Kind()))
	}

	err = errors.New(fmt.Sprintf("%s.%s is not set", g.Type().Name(), fn.Name()))

	return
}

func validateEmail(g generic.Generic, name string) (err error) {
	email, ok := g.String(name)
	if !ok {
		return
	}

	err = validation.Validate(email,
		is.Email,
	)

	return
}

func ValidateUnion(stage string, svc *sdk.Service, g generic.Generic) (errs []sdk.Error) {
	var set []string

	kind, ok := g.String(fieldnames.Kind)

	for _, fieldName := range g.FieldNames() {
		if fieldName == fieldnames.Hash || fieldName == fieldnames.Kind {
			continue
		}

		set = append(set, fieldName)

		if ok && kind != fieldName {
			errs = append(errs, getError(stage, svc, fmt.Sprintf("union %s.%s is set, but %s.kind is %s", g.Type().Name(), fieldName, g.Type().Name(), kind)))
		}
	}

	isOptional := g.Type().Flags().Is(typeflags.IsOptionalValueUnion, true)

	if !isOptional && len(set) == 0 {
		errs = append(errs, getError(stage, svc, fmt.Sprintf("union %s needs a value to be set \n", g.Type().Name())))
	}

	if len(set) > 1 {
		errs = append(errs, getError(stage, svc, fmt.Sprintf("union %s has more than one value set: %s", g.Type().Name(), set)))
	}

	return
}

func getError(stage string, svc *sdk.Service, message string) (sdk.Error) {
	var kind string

	switch stage {
	case types.SvcRsp:
		kind = sdk.ErrorKind.ResponseValidation
	case types.CliReq:
		kind = sdk.ErrorKind.RequestValidation
	}

	return sdk.Error{
		Kind: ptr.String(kind),
		Message: &sdk.Text{
			Formatting: &sdk.FormattingKind.Plain,
			Value: &message,
		},
		Service: svc,
	}
}
