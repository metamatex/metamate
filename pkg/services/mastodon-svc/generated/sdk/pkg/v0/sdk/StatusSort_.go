// generated by metactl sdk gen 
package sdk

const (
	StatusSortName = "StatusSort"
)

type StatusSort struct {
    Content *TextSort `json:"content,omitempty",yaml:"content,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    Id *ServiceIdSort `json:"id,omitempty",yaml:"id,omitempty"`
    Meta *TypeMetaSort `json:"meta,omitempty",yaml:"meta,omitempty"`
    Pinned *string `json:"pinned,omitempty",yaml:"pinned,omitempty"`
    Sensitive *string `json:"sensitive,omitempty",yaml:"sensitive,omitempty"`
    SpoilerText *TextSort `json:"spoilerText,omitempty",yaml:"spoilerText,omitempty"`
}