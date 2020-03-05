// generated by metactl sdk gen 
package sdk

const (
	PostBlueWhateversRequestFilterName = "PostBlueWhateversRequestFilter"
)

type PostBlueWhateversRequestFilter struct {
    And []PostBlueWhateversRequestFilter `json:"and,omitempty",yaml:"and,omitempty"`
    BlueWhatevers *BlueWhateverListFilter `json:"blueWhatevers,omitempty",yaml:"blueWhatevers,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    Meta *RequestMetaFilter `json:"meta,omitempty",yaml:"meta,omitempty"`
    Not []PostBlueWhateversRequestFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []PostBlueWhateversRequestFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}