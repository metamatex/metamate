// generated by metactl sdk gen 
package sdk

const (
	PipePutBlueWhateversContextFilterName = "PipePutBlueWhateversContextFilter"
)

type PipePutBlueWhateversContextFilter struct {
    And []PipePutBlueWhateversContextFilter `json:"and,omitempty" yaml:"and,omitempty"`
    ClientRequest *PutBlueWhateversRequestFilter `json:"clientRequest,omitempty" yaml:"clientRequest,omitempty"`
    ClientResponse *PutBlueWhateversResponseFilter `json:"clientResponse,omitempty" yaml:"clientResponse,omitempty"`
    Not []PipePutBlueWhateversContextFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PipePutBlueWhateversContextFilter `json:"or,omitempty" yaml:"or,omitempty"`
    ServiceRequest *PutBlueWhateversRequestFilter `json:"serviceRequest,omitempty" yaml:"serviceRequest,omitempty"`
    ServiceResponse *PutBlueWhateversResponseFilter `json:"serviceResponse,omitempty" yaml:"serviceResponse,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}