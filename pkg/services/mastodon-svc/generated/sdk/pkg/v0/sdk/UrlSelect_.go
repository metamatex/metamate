// generated by metactl sdk gen 
package sdk

const (
	UrlSelectName = "UrlSelect"
)

type UrlSelect struct {
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    All *bool `json:"selectAll,omitempty",yaml:"selectAll,omitempty"`
    Value *bool `json:"value,omitempty",yaml:"value,omitempty"`
}