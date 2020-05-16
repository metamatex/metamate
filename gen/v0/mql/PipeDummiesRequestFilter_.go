// generated by metactl sdk gen 
package mql

const (
	PipeDummiesRequestFilterName = "PipeDummiesRequestFilter"
)

type PipeDummiesRequestFilter struct {
    And []PipeDummiesRequestFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Context *PipeDummiesContextFilter `json:"context,omitempty" yaml:"context,omitempty"`
    Meta *RequestMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Mode *PipeModeFilter `json:"mode,omitempty" yaml:"mode,omitempty"`
    Not []PipeDummiesRequestFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PipeDummiesRequestFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}