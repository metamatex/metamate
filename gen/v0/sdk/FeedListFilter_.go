// generated by metactl sdk gen 
package sdk

const (
	FeedListFilterName = "FeedListFilter"
)

type FeedListFilter struct {
    Every *FeedFilter `json:"every,omitempty",yaml:"every,omitempty"`
    None *FeedFilter `json:"none,omitempty",yaml:"none,omitempty"`
    Some *FeedFilter `json:"some,omitempty",yaml:"some,omitempty"`
}