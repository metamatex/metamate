// generated by metactl sdk gen 
package mql

const (
	PipeAttachmentsEndpointFilterName = "PipeAttachmentsEndpointFilter"
)

type PipeAttachmentsEndpointFilter struct {
    And []PipeAttachmentsEndpointFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Not []PipeAttachmentsEndpointFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PipeAttachmentsEndpointFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}