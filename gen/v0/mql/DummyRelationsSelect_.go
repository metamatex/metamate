// generated by metactl sdk gen 
package mql

const (
	DummyRelationsSelectName = "DummyRelationsSelect"
)

type DummyRelationsSelect struct {
    All *bool `json:"all,omitempty" yaml:"all,omitempty"`
    KnewByDummies *DummiesCollectionSelect `json:"knewByDummies,omitempty" yaml:"knewByDummies,omitempty"`
    KnowsBlueDummies *BlueDummiesCollectionSelect `json:"knowsBlueDummies,omitempty" yaml:"knowsBlueDummies,omitempty"`
    KnowsDummies *DummiesCollectionSelect `json:"knowsDummies,omitempty" yaml:"knowsDummies,omitempty"`
}