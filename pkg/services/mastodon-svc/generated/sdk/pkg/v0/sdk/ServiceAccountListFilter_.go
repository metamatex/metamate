// generated by metactl sdk gen 
package sdk

const (
	ServiceAccountListFilterName = "ServiceAccountListFilter"
)

type ServiceAccountListFilter struct {
    Every *ServiceAccountFilter `json:"every,omitempty",yaml:"every,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    None *ServiceAccountFilter `json:"none,omitempty",yaml:"none,omitempty"`
    Some *ServiceAccountFilter `json:"some,omitempty",yaml:"some,omitempty"`
}