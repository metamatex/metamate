// generated by metactl sdk gen 
package sdk

const (
	DeleteServicesResponseFilterName = "DeleteServicesResponseFilter"
)

type DeleteServicesResponseFilter struct {
    And []DeleteServicesResponseFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Meta *ResponseMetaFilter `json:"meta,omitempty",yaml:"meta,omitempty"`
    Not []DeleteServicesResponseFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []DeleteServicesResponseFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}