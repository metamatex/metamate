// generated by metactl sdk gen 
package sdk

const (
	PipePostTransactionsContextFilterName = "PipePostTransactionsContextFilter"
)

type PipePostTransactionsContextFilter struct {
    And []PipePostTransactionsContextFilter `json:"and,omitempty" yaml:"and,omitempty"`
    ClientRequest *PostTransactionsRequestFilter `json:"clientRequest,omitempty" yaml:"clientRequest,omitempty"`
    ClientResponse *PostTransactionsResponseFilter `json:"clientResponse,omitempty" yaml:"clientResponse,omitempty"`
    Not []PipePostTransactionsContextFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PipePostTransactionsContextFilter `json:"or,omitempty" yaml:"or,omitempty"`
    ServiceRequest *PostTransactionsRequestFilter `json:"serviceRequest,omitempty" yaml:"serviceRequest,omitempty"`
    ServiceResponse *PostTransactionsResponseFilter `json:"serviceResponse,omitempty" yaml:"serviceResponse,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}