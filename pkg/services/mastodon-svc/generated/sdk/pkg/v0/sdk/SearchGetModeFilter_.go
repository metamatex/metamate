// generated by metactl sdk gen 
package sdk

const (
	SearchGetModeFilterName = "SearchGetModeFilter"
)

type SearchGetModeFilter struct {
    And []SearchGetModeFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    Location *LocationQueryFilter `json:"location,omitempty",yaml:"location,omitempty"`
    Not []SearchGetModeFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []SearchGetModeFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
    Term *StringFilter `json:"term,omitempty",yaml:"term,omitempty"`
}