// generated by metactl sdk gen 
package mql

const (
	PipeGetBlueDummiesContextFilterName = "PipeGetBlueDummiesContextFilter"
)

type PipeGetBlueDummiesContextFilter struct {
    And []PipeGetBlueDummiesContextFilter `json:"and,omitempty" yaml:"and,omitempty"`
    ClientRequest *GetBlueDummiesRequestFilter `json:"clientRequest,omitempty" yaml:"clientRequest,omitempty"`
    ClientResponse *GetBlueDummiesResponseFilter `json:"clientResponse,omitempty" yaml:"clientResponse,omitempty"`
    Not []PipeGetBlueDummiesContextFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PipeGetBlueDummiesContextFilter `json:"or,omitempty" yaml:"or,omitempty"`
    ServiceRequest *GetBlueDummiesRequestFilter `json:"serviceRequest,omitempty" yaml:"serviceRequest,omitempty"`
    ServiceResponse *GetBlueDummiesResponseFilter `json:"serviceResponse,omitempty" yaml:"serviceResponse,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}