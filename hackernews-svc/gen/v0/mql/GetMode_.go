// generated by metactl sdk gen 
package mql

const (
	GetModeName = "GetMode"
)

type GetMode struct {
    Collection *CollectionGetMode `json:"collection,omitempty" yaml:"collection,omitempty"`
    Id *Id `json:"id,omitempty" yaml:"id,omitempty"`
    Kind *string `json:"kind,omitempty" yaml:"kind,omitempty"`
    Relation *RelationGetMode `json:"relation,omitempty" yaml:"relation,omitempty"`
    Search *SearchGetMode `json:"search,omitempty" yaml:"search,omitempty"`
}