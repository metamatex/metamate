// generated by metactl sdk gen 
package mql

const (
	PipeDummiesRequestName = "PipeDummiesRequest"
)

type PipeDummiesRequest struct {
    Context *PipeDummiesContext `json:"context,omitempty" yaml:"context,omitempty"`
    Meta *RequestMeta `json:"meta,omitempty" yaml:"meta,omitempty"`
    Mode *PipeMode `json:"mode,omitempty" yaml:"mode,omitempty"`
}