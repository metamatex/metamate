// generated by metactl sdk gen 
package sdk

const (
	ServiceAccountName = "ServiceAccount"
)

type ServiceAccount struct {
    AlternativeIds []Id `json:"alternativeIds,omitempty",yaml:"alternativeIds,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    Host *HostServiceAccount `json:"host,omitempty",yaml:"host,omitempty"`
    Id *ServiceId `json:"id,omitempty",yaml:"id,omitempty"`
    Kind *string `json:"kind,omitempty",yaml:"kind,omitempty"`
    Meta *TypeMeta `json:"meta,omitempty",yaml:"meta,omitempty"`
    Relations *ServiceAccountRelations `json:"relations,omitempty",yaml:"relations,omitempty"`
}