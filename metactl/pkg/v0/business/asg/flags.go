package asg

import (
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph/basictypeflags"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph/endpointflags"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph/enumflags"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph/fieldflags"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph/relationflags"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/metamatemono/metactl/pkg/v0/types"
	"github.com/metamatex/metamatemono/metactl/pkg/v0/utils"
	"sort"
)

type FlagsReport struct {
	BasicType map[string]bool `yaml:",omitempty" json:"basictype,omitempty"`
	Endpoint  map[string]bool `yaml:",omitempty" json:"endpoint,omitempty"`
	Enum      map[string]bool `yaml:",omitempty" json:"enum,omitempty"`
	Field     map[string]bool `yaml:",omitempty" json:"field,omitempty"`
	Interface map[string]bool `yaml:",omitempty" json:"interface,omitempty"`
	Relation  map[string]bool `yaml:",omitempty" json:"relation,omitempty"`
	Type      map[string]bool `yaml:",omitempty" json:"type,omitempty"`
}

func Flags(returnData bool) (o types.Output) {
	if returnData {
		o.Data = GetFlagsReport()
	} else {
		o.Text = GetFlagsTableString()
	}

	return
}

func getKeys(m map[string]bool) (keys []string) {
	for k, v := range m {
		suffix := ""
		if v {
			suffix = " (default: true)"
		}

		keys = append(keys, string(k)+suffix)
	}

	sort.Strings(keys)

	return
}

func toRows(columns [][]string) (rows [][]string) {
	c := 0
	for {
		row := []string{}
		cSkipped := 0
		for i, _ := range columns {
			if len(columns[i]) <= c {
				row = append(row, "")

				cSkipped++

				continue
			}

			row = append(row, columns[i][c])
		}

		if cSkipped == len(columns) {
			break
		}

		c++

		rows = append(rows, row)
	}

	return
}

func GetFlagsReport() FlagsReport {
	return FlagsReport{
		BasicType: basictypeflags.Defaults,
		Endpoint:  endpointflags.Defaults,
		Enum:      enumflags.Defaults,
		Field:     fieldflags.Defaults,
		Relation:  relationflags.Defaults,
		Type:      typeflags.Defaults,
	}
}

func GetFlagsTableString() string {
	basictypeflagKeys := getKeys(basictypeflags.Defaults)
	endpointflagKeys := getKeys(endpointflags.Defaults)
	enumflagKeys := getKeys(enumflags.Defaults)
	fieldflagKeys := getKeys(fieldflags.Defaults)
	relationflagKeys := getKeys(relationflags.Defaults)
	typeflagKeys := getKeys(typeflags.Defaults)

	columns := [][]string{
		basictypeflagKeys,
		endpointflagKeys,
		enumflagKeys,
		fieldflagKeys,
		relationflagKeys,
		typeflagKeys,
	}

	rows := toRows(columns)

	return utils.GetTableString([]string{"basictype", "endpoint", "enum", "field", "interface", "relation", "type"}, rows, []int{})
}
