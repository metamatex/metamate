// generated by metactl sdk gen 
package mql

const (
	DummiesCollectionSelectName = "DummiesCollectionSelect"
)

type DummiesCollectionSelect struct {
    All *bool `json:"all,omitempty" yaml:"all,omitempty"`
    Count *bool `json:"count,omitempty" yaml:"count,omitempty"`
    Dummies *DummySelect `json:"dummies,omitempty" yaml:"dummies,omitempty"`
    Errors *ErrorSelect `json:"errors,omitempty" yaml:"errors,omitempty"`
    Pagination *PaginationSelect `json:"pagination,omitempty" yaml:"pagination,omitempty"`
    Warnings *WarningSelect `json:"warnings,omitempty" yaml:"warnings,omitempty"`
}