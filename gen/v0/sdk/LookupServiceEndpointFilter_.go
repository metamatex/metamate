// generated by metactl sdk gen 
package sdk

const (
	LookupServiceEndpointFilterName = "LookupServiceEndpointFilter"
)

type LookupServiceEndpointFilter struct {
    And []LookupServiceEndpointFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Not []LookupServiceEndpointFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []LookupServiceEndpointFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}