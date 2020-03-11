package types

import (
	"fmt"
	"reflect"
)

type MessageReport struct {
	Debug   []string `yaml:",omitempty" json:"debug,omitempty"`
	Info    []string `yaml:",omitempty" json:"info,omitempty"`
	Warning []string `yaml:",omitempty" json:"warning,omitempty"`
	Error   []string `yaml:",omitempty" json:"error,omitempty"`
	Hint   []string `yaml:",omitempty" json:"hint,omitempty"`
}

func (r *MessageReport) AddDebug(any interface{}) {
	switch sth := any.(type) {
	case string:
		r.Debug = append(r.Debug, sth)
	case error:
		r.Debug = append(r.Debug, sth.Error())
	case []error:
		for _, err := range sth {
			r.Debug = append(r.Debug, err.Error())
		}
	default:
		panic(fmt.Sprintf("must provide string, error or []error to MessageReport.AddDebug(), got %v", reflect.TypeOf(sth).Name()))
	}
}

func (r *MessageReport) AddInfo(any interface{}) {
	switch sth := any.(type) {
	case string:
		r.Info = append(r.Info, sth)
	case error:
		r.Info = append(r.Info, sth.Error())
	case []error:
		for _, err := range sth {
			r.Info = append(r.Info, err.Error())
		}
	default:
		panic(fmt.Sprintf("must provide string, error or []error to MessageReport.AddInfo(), got %v", reflect.TypeOf(sth).Name()))
	}
}

func (r *MessageReport) AddError(any interface{}) {
	switch sth := any.(type) {
	case string:
		r.Error = append(r.Error, sth)
	case error:
		r.Error = append(r.Error, sth.Error())
	case []error:
		for _, err := range sth {
			r.Error = append(r.Error, err.Error())
		}
	default:
		panic(fmt.Sprintf("must provide string, error or []error to MessageReport.AddError(), got %v", reflect.TypeOf(sth).Name()))
	}
}

func (r *MessageReport) AddWarning(any interface{}) {
	switch sth := any.(type) {
	case string:
		r.Warning = append(r.Warning, sth)
	case error:
		r.Warning = append(r.Warning, sth.Error())
	case []error:
		for _, err := range sth {
			r.Warning = append(r.Warning, err.Error())
		}
	default:
		panic(fmt.Sprintf("must provide string, error or []error to MessageReport.AddWarning(), got %v", reflect.TypeOf(sth).Name()))
	}
}

func (r *MessageReport) AddHint(any interface{}) {
	switch sth := any.(type) {
	case string:
		r.Hint = append(r.Hint, sth)
	case error:
		r.Hint = append(r.Hint, sth.Error())
	case []error:
		for _, err := range sth {
			r.Hint = append(r.Hint, err.Error())
		}
	default:
		panic(fmt.Sprintf("must provide string, error or []error to MessageReport.AddHint(), got %v", reflect.TypeOf(sth).Name()))
	}
}
