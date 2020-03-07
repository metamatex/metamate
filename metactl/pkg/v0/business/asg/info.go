package asg

import (
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph"
)

type InfoReport struct {
	Total     NodeReport       `yaml:",omitempty" json:"basictype,omitempty"`
	BasicType NodeReport       `yaml:",omitempty" json:"basictype,omitempty"`
	Endpoint  NodeReport       `yaml:",omitempty" json:"endpoint,omitempty"`
	Enum      NodeReport       `yaml:",omitempty" json:"enum,omitempty"`
	Field     NodeReport       `yaml:",omitempty" json:"field,omitempty"`
	Interface NodeReport       `yaml:",omitempty" json:"interface,omitempty"`
	Relation  NodeReport       `yaml:",omitempty" json:"relation,omitempty"`
	Type      NodeReport       `yaml:",omitempty" json:"type,omitempty"`
	Missing   graph.GraphStats `yaml:",omitempty" json:"missing,omitempty"`
}

type NodeReport struct {
	Count int            `yaml:",omitempty" json:"count,omitempty"`
	Flags map[string]int `yaml:",omitempty" json:"flags,omitempty"`
}

func Info(root *graph.RootNode) (r InfoReport, err error) {
	r.Total = NodeReport{
		Count: len(root.GetNodes()),
	}

	r.BasicType = NodeReport{
		Count: len(root.BasicTypes),
		Flags: countFlags(root.BasicTypes.ToNodeMap()),
	}

	r.Endpoint = NodeReport{
		Count: len(root.Endpoints),
		Flags: countFlags(root.Endpoints.ToNodeMap()),
	}

	r.Enum = NodeReport{
		Count: len(root.Enums),
		Flags: countFlags(root.Enums.ToNodeMap()),
	}

	r.Field = NodeReport{
		Count: len(root.Fields),
		Flags: countFlags(root.Fields.ToNodeMap()),
	}

	r.Relation = NodeReport{
		Count: len(root.Relations),
		Flags: countFlags(root.Relations.ToNodeMap()),
	}

	r.Type = NodeReport{
		Count: len(root.Types),
		Flags: countFlags(root.Types.ToNodeMap()),
	}

	r.Missing = graph.GetMissing(root)

	return
}

func countFlags(nm graph.NodeMap) (c map[string]int) {
	c = map[string]int{}

	nm.Each(func(n graph.Node) {
		for k, v := range n.Flags().Flags {
			if v {
				c[k]++
			}
		}
	})

	return
}
