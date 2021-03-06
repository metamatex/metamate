// generated by metactl sdk gen 
package mql

const (
	GetPostFeedsCollectionName = "GetPostFeedsCollection"
)

type GetPostFeedsCollection struct {
    Filter *PostFeedFilter `json:"filter,omitempty" yaml:"filter,omitempty"`
    Pages []ServicePage `json:"pages,omitempty" yaml:"pages,omitempty"`
    Relations *GetPostFeedsRelations `json:"relations,omitempty" yaml:"relations,omitempty"`
    Select *PostFeedsCollectionSelect `json:"select,omitempty" yaml:"select,omitempty"`
    ServiceFilter *ServiceFilter `json:"serviceFilter,omitempty" yaml:"serviceFilter,omitempty"`
    Sort *PostFeedSort `json:"sort,omitempty" yaml:"sort,omitempty"`
}