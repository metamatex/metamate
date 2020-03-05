package graph

import (
	"fmt"
	"github.com/metamatex/asg/pkg/v0/asg/typenames"
	"strings"
	"errors"

	"github.com/metamatex/asg/pkg/v0/asg/graph/fieldflags"
	"github.com/metamatex/asg/pkg/v0/asg/graph/typeflags"
)

func (rn *RootNode) AddScalar(name string, unit string) (*TypeNode) {
	return rn.AddTypeNode(name, []*FieldNode{
		EnumField("unit", unit),
		Float64Field("value", Flags{
			fieldflags.ValidateIsSet: true,
		}),
		BoolField("isEstimate"),
	}, Flags{
		typeflags.IsScalar: true,
	})
}

func validateScalar(tn *TypeNode) (errs []error) {
	if !strings.HasSuffix(tn.Name(), "Scalar") {
		errs = append(errs, errors.New(fmt.Sprintf("Name %v needs to end with \"Scalar\"", tn.Name())))
	}

	return
}

func (rn *RootNode) AddValue(name string, unit string) (*TypeNode) {
	return rn.AddTypeNode(name, []*FieldNode{
		EnumField("kind", typenames.ValueKind),
		EnumField("unit", unit),
		TypeField("range", typenames.FloatRange),
		Float64Field("value"),
		BoolField("isEstimate"),
	}, Flags{
		typeflags.IsValue: true,
	})
}

func validateValue(tn *TypeNode) (errs []error) {
	if !strings.HasSuffix(tn.Name(), "Value") {
		errs = append(errs, errors.New(fmt.Sprintf("Name %v needs to end with \"Value\"", tn.Name())))
	}

	return
}

func (rn *RootNode) AddRange(name string, unit string) (*TypeNode) {
	return rn.AddTypeNode(name, []*FieldNode{
		EnumField("unit", unit),
		TypeField("range", typenames.FloatRange),
		BoolField("isEstimate"),
	}, Flags{
		typeflags.IsRange: true,
	})
}

func validateRange(tn *TypeNode) (errs []error) {
	if !strings.HasSuffix(tn.Name(), "Range") {
		errs = append(errs, errors.New(fmt.Sprintf("Name %v needs to end with \"Range\"", tn.Name())))
	}

	return
}

func (rn *RootNode) AddRatio(ratioName, counterName, dividerName string) (*TypeNode) {
	return rn.AddTypeNode(ratioName, []*FieldNode{
		TypeField("counter", counterName),
		TypeField("divider", dividerName),
	}, Flags{
		typeflags.IsRatio: true,
	})
}

func validateRatio(tn *TypeNode) (errs []error) {
	if !strings.HasSuffix(tn.Name(), "Ratio") {
		errs = append(errs, errors.New(fmt.Sprintf("Name %v needs to end with \"Ratio\"", tn.Name())))
	}

	return
}