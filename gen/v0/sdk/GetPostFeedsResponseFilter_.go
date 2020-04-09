// generated by metactl sdk gen 
package sdk

const (
	GetPostFeedsResponseFilterName = "GetPostFeedsResponseFilter"
)

type GetPostFeedsResponseFilter struct {
    And []GetPostFeedsResponseFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Meta *CollectionMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Not []GetPostFeedsResponseFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []GetPostFeedsResponseFilter `json:"or,omitempty" yaml:"or,omitempty"`
    PostFeeds *PostFeedListFilter `json:"postFeeds,omitempty" yaml:"postFeeds,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}