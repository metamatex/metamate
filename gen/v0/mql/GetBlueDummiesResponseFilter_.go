// generated by metactl sdk gen 
package mql

const (
	GetBlueDummiesResponseFilterName = "GetBlueDummiesResponseFilter"
)

type GetBlueDummiesResponseFilter struct {
    And []GetBlueDummiesResponseFilter `json:"and,omitempty" yaml:"and,omitempty"`
    BlueDummies *BlueDummyListFilter `json:"blueDummies,omitempty" yaml:"blueDummies,omitempty"`
    Count *Int32Filter `json:"count,omitempty" yaml:"count,omitempty"`
    Errors *ErrorListFilter `json:"errors,omitempty" yaml:"errors,omitempty"`
    Not []GetBlueDummiesResponseFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []GetBlueDummiesResponseFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Pagination *PaginationFilter `json:"pagination,omitempty" yaml:"pagination,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
    Warnings *WarningListFilter `json:"warnings,omitempty" yaml:"warnings,omitempty"`
}