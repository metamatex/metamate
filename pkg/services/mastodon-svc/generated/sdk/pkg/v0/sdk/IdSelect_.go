// generated by metactl sdk gen 
package sdk

const (
	IdSelectName = "IdSelect"
)

type IdSelect struct {
    Ean *bool `json:"ean,omitempty",yaml:"ean,omitempty"`
    Email *EmailSelect `json:"email,omitempty",yaml:"email,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    Kind *bool `json:"kind,omitempty",yaml:"kind,omitempty"`
    Local *bool `json:"local,omitempty",yaml:"local,omitempty"`
    Me *bool `json:"me,omitempty",yaml:"me,omitempty"`
    Name *bool `json:"name,omitempty",yaml:"name,omitempty"`
    All *bool `json:"selectAll,omitempty",yaml:"selectAll,omitempty"`
    ServiceId *ServiceIdSelect `json:"serviceId,omitempty",yaml:"serviceId,omitempty"`
    Token *TokenSelect `json:"token,omitempty",yaml:"token,omitempty"`
    Url *UrlSelect `json:"url,omitempty",yaml:"url,omitempty"`
    Username *bool `json:"username,omitempty",yaml:"username,omitempty"`
}