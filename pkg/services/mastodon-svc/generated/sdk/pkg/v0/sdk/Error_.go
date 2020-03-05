// generated by metactl sdk gen 
package sdk

const (
	ErrorName = "Error"
)

type Error struct {
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    Id *Id `json:"id,omitempty",yaml:"id,omitempty"`
    Kind *string `json:"kind,omitempty",yaml:"kind,omitempty"`
    Message *Text `json:"message,omitempty",yaml:"message,omitempty"`
    Service *Service `json:"service,omitempty",yaml:"service,omitempty"`
    Upstream *Error `json:"upstream,omitempty",yaml:"upstream,omitempty"`
}