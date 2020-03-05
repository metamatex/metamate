// generated by metactl sdk gen 
package sdk

const (
	Float64FilterName = "Float64Filter"
)

type Float64Filter struct {
    And []Float64Filter `json:"and,omitempty",yaml:"and,omitempty"`
    Gt *float64 `json:"gt,omitempty",yaml:"gt,omitempty"`
    Gte *float64 `json:"gte,omitempty",yaml:"gte,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    In []float64 `json:"in,omitempty",yaml:"in,omitempty"`
    Is *float64 `json:"is,omitempty",yaml:"is,omitempty"`
    Lt *float64 `json:"lt,omitempty",yaml:"lt,omitempty"`
    Lte *float64 `json:"lte,omitempty",yaml:"lte,omitempty"`
    Not *float64 `json:"not,omitempty",yaml:"not,omitempty"`
    NotIn []float64 `json:"notIn,omitempty",yaml:"notIn,omitempty"`
    Or []Float64Filter `json:"or,omitempty",yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}