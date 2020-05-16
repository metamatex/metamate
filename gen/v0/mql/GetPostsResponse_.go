// generated by metactl sdk gen 
package mql

const (
	GetPostsResponseName = "GetPostsResponse"
)

type GetPostsResponse struct {
    Count *int32 `json:"count,omitempty" yaml:"count,omitempty"`
    Errors []Error `json:"errors,omitempty" yaml:"errors,omitempty"`
    Pagination *Pagination `json:"pagination,omitempty" yaml:"pagination,omitempty"`
    Posts []Post `json:"posts,omitempty" yaml:"posts,omitempty"`
    Warnings []Warning `json:"warnings,omitempty" yaml:"warnings,omitempty"`
}