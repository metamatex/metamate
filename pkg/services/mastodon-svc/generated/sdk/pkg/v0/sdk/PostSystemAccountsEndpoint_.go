// generated by metactl sdk gen 
package sdk

const (
	PostClientAccountsEndpointName = "PostClientAccountsEndpoint"
)

type PostClientAccountsEndpoint struct {
    Filter *PostClientAccountsRequestFilter `json:"filter,omitempty",yaml:"filter,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
}