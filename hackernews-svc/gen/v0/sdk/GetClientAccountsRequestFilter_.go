// generated by metactl sdk gen 
package sdk

const (
	GetClientAccountsRequestFilterName = "GetClientAccountsRequestFilter"
)

type GetClientAccountsRequestFilter struct {
    And []GetClientAccountsRequestFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Meta *RequestMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Mode *GetModeFilter `json:"mode,omitempty" yaml:"mode,omitempty"`
    Not []GetClientAccountsRequestFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []GetClientAccountsRequestFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Pages *ServicePageListFilter `json:"pages,omitempty" yaml:"pages,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}