// generated by metactl sdk gen 
package sdk

const (
	DeleteStatusesEndpointFilterName = "DeleteStatusesEndpointFilter"
)

type DeleteStatusesEndpointFilter struct {
    And []DeleteStatusesEndpointFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    Not []DeleteStatusesEndpointFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []DeleteStatusesEndpointFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}