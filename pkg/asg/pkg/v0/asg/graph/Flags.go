package graph

import (
	"fmt"
)

type Flags map[string]bool

type FlagsContainer struct {
	Flags Flags `yaml:",omitempty" json:"flags,inline,omitempty"`
	defaults map[string]bool
}

func (c FlagsContainer) Copy() (c0 FlagsContainer) {
	c0.defaults = c.defaults
	c0.Flags = Flags{}
	for k, v := range c.Flags {
		c0.Flags[k] = v
	}

	return
}

func (c FlagsContainer) Set(f string, b bool) (FlagsContainer) {
	c.Flags[f] = b

	return c
}

func (c FlagsContainer) Is(f string, b bool) (bool) {
	b0, ok := c.Flags[f]
	if ok {
		return b0 == b
	}

	b0, ok = c.defaults[f]
	if !ok {
		panic(fmt.Sprintf("no default set for flag %v", f))
	}

	return b0 == b
}

func (c FlagsContainer) Or(flags ...string) (bool) {
	for _, f := range flags {
		if c.Is(f, true) {
			return true
		}
	}

	return false
}
