// generated by metactl sdk gen 
package sdk

const (
	PutServicesRequestFilterName = "PutServicesRequestFilter"
)

type PutServicesRequestFilter struct {
    And []PutServicesRequestFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Meta *RequestMetaFilter `json:"meta,omitempty",yaml:"meta,omitempty"`
    Mode *PutModeFilter `json:"mode,omitempty",yaml:"mode,omitempty"`
    Not []PutServicesRequestFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []PutServicesRequestFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Services *ServiceListFilter `json:"services,omitempty",yaml:"services,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}