// generated by metactl sdk gen 
package mql

const (
	HyperLinkListFilterName = "HyperLinkListFilter"
)

type HyperLinkListFilter struct {
    Every *HyperLinkFilter `json:"every,omitempty" yaml:"every,omitempty"`
    None *HyperLinkFilter `json:"none,omitempty" yaml:"none,omitempty"`
    Some *HyperLinkFilter `json:"some,omitempty" yaml:"some,omitempty"`
}