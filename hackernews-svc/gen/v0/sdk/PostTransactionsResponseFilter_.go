// generated by metactl sdk gen 
package sdk

const (
	PostTransactionsResponseFilterName = "PostTransactionsResponseFilter"
)

type PostTransactionsResponseFilter struct {
    And []PostTransactionsResponseFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Meta *ResponseMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Not []PostTransactionsResponseFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PostTransactionsResponseFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
    Transactions *TransactionListFilter `json:"transactions,omitempty" yaml:"transactions,omitempty"`
}