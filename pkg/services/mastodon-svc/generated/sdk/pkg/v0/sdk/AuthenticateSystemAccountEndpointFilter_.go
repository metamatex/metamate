// generated by metactl sdk gen 
package sdk

const (
	AuthenticateClientAccountEndpointFilterName = "AuthenticateClientAccountEndpointFilter"
)

type AuthenticateClientAccountEndpointFilter struct {
    And []AuthenticateClientAccountEndpointFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    Not []AuthenticateClientAccountEndpointFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []AuthenticateClientAccountEndpointFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}