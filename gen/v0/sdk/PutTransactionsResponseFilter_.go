// generated by metactl sdk gen 
package sdk

const (
	PutTransactionsResponseFilterName = "PutTransactionsResponseFilter"
)

type PutTransactionsResponseFilter struct {
    And []PutTransactionsResponseFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Meta *ResponseMetaFilter `json:"meta,omitempty",yaml:"meta,omitempty"`
    Not []PutTransactionsResponseFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []PutTransactionsResponseFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}