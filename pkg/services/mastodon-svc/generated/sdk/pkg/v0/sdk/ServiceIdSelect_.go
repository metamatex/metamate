// generated by metactl sdk gen 
package sdk

const (
	ServiceIdSelectName = "ServiceIdSelect"
)

type ServiceIdSelect struct {
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    All *bool `json:"selectAll,omitempty",yaml:"selectAll,omitempty"`
    ServiceName *bool `json:"serviceName,omitempty",yaml:"serviceName,omitempty"`
    Value *bool `json:"value,omitempty",yaml:"value,omitempty"`
}