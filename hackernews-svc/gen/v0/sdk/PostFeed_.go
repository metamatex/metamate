// generated by metactl sdk gen 
package sdk

const (
	PostFeedName = "PostFeed"
)

type PostFeed struct {
    AlternativeIds []Id `json:"alternativeIds,omitempty" yaml:"alternativeIds,omitempty"`
    Id *ServiceId `json:"id,omitempty" yaml:"id,omitempty"`
    Info *Info `json:"info,omitempty" yaml:"info,omitempty"`
    Kind *string `json:"kind,omitempty" yaml:"kind,omitempty"`
    Meta *TypeMeta `json:"meta,omitempty" yaml:"meta,omitempty"`
    Relations *PostFeedRelations `json:"relations,omitempty" yaml:"relations,omitempty"`
    Relationships *PostFeedRelationships `json:"relationships,omitempty" yaml:"relationships,omitempty"`
}