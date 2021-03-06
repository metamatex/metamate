// generated by metactl sdk gen 
package mql

const (
	PostSelectName = "PostSelect"
)

type PostSelect struct {
    All *bool `json:"all,omitempty" yaml:"all,omitempty"`
    AlternativeIds *IdSelect `json:"alternativeIds,omitempty" yaml:"alternativeIds,omitempty"`
    Content *TextSelect `json:"content,omitempty" yaml:"content,omitempty"`
    CreatedAt *TimestampSelect `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
    Id *ServiceIdSelect `json:"id,omitempty" yaml:"id,omitempty"`
    IsPinned *bool `json:"isPinned,omitempty" yaml:"isPinned,omitempty"`
    IsSensitive *bool `json:"isSensitive,omitempty" yaml:"isSensitive,omitempty"`
    Kind *bool `json:"kind,omitempty" yaml:"kind,omitempty"`
    Links *HyperLinkSelect `json:"links,omitempty" yaml:"links,omitempty"`
    Relations *PostRelationsSelect `json:"relations,omitempty" yaml:"relations,omitempty"`
    Relationships *PostRelationshipsSelect `json:"relationships,omitempty" yaml:"relationships,omitempty"`
    SpoilerText *TextSelect `json:"spoilerText,omitempty" yaml:"spoilerText,omitempty"`
    Title *TextSelect `json:"title,omitempty" yaml:"title,omitempty"`
    TotalWasRepliedToByPostsCount *bool `json:"totalWasRepliedToByPostsCount,omitempty" yaml:"totalWasRepliedToByPostsCount,omitempty"`
}