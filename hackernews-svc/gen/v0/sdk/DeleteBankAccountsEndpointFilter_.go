// generated by metactl sdk gen 
package sdk

const (
	DeleteBankAccountsEndpointFilterName = "DeleteBankAccountsEndpointFilter"
)

type DeleteBankAccountsEndpointFilter struct {
    And []DeleteBankAccountsEndpointFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Not []DeleteBankAccountsEndpointFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []DeleteBankAccountsEndpointFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}