// generated by metactl sdk gen 
package sdk

const (
	PutPeopleResponseSelectName = "PutPeopleResponseSelect"
)

type PutPeopleResponseSelect struct {
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    Meta *ResponseMetaSelect `json:"meta,omitempty",yaml:"meta,omitempty"`
    All *bool `json:"selectAll,omitempty",yaml:"selectAll,omitempty"`
}