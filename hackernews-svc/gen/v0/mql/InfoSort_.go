// generated by metactl sdk gen 
package mql

const (
	InfoSortName = "InfoSort"
)

type InfoSort struct {
    Description *TextSort `json:"description,omitempty" yaml:"description,omitempty"`
    Name *TextSort `json:"name,omitempty" yaml:"name,omitempty"`
    Purpose *TextSort `json:"purpose,omitempty" yaml:"purpose,omitempty"`
}