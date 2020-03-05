// generated by metactl sdk gen 
package sdk

const (
	InfoSelectName = "InfoSelect"
)

type InfoSelect struct {
    Description *TextSelect `json:"description,omitempty",yaml:"description,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    Name *TextSelect `json:"name,omitempty",yaml:"name,omitempty"`
    Purpose *TextSelect `json:"purpose,omitempty",yaml:"purpose,omitempty"`
    All *bool `json:"selectAll,omitempty",yaml:"selectAll,omitempty"`
}