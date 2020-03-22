// generated by metactl sdk gen 
package sdk

const (
	PostStatusesResponseFilterName = "PostStatusesResponseFilter"
)

type PostStatusesResponseFilter struct {
    And []PostStatusesResponseFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Meta *ResponseMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Not []PostStatusesResponseFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PostStatusesResponseFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
    Statuses *StatusListFilter `json:"statuses,omitempty" yaml:"statuses,omitempty"`
}