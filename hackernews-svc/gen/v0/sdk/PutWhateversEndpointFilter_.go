// generated by metactl sdk gen 
package sdk

const (
	PutWhateversEndpointFilterName = "PutWhateversEndpointFilter"
)

type PutWhateversEndpointFilter struct {
    And []PutWhateversEndpointFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Not []PutWhateversEndpointFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PutWhateversEndpointFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}