// generated by metactl sdk gen 
package mql

const (
	HyperLinkSelectName = "HyperLinkSelect"
)

type HyperLinkSelect struct {
    All *bool `json:"all,omitempty" yaml:"all,omitempty"`
    Label *bool `json:"label,omitempty" yaml:"label,omitempty"`
    Url *UrlSelect `json:"url,omitempty" yaml:"url,omitempty"`
}