// generated by metactl sdk gen 
package sdk

const (
	GetAttachmentsRelationsName = "GetAttachmentsRelations"
)

type GetAttachmentsRelations struct {
    AttachedToStatuses *GetStatusesCollection `json:"attachedToStatuses,omitempty",yaml:"attachedToStatuses,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
}