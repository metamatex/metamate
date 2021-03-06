// generated by metactl sdk gen 
package mql

const (
	SocialAccountRelationshipsName = "SocialAccountRelationships"
)

type SocialAccountRelationships struct {
    BlockedByMe *bool `json:"blockedByMe,omitempty" yaml:"blockedByMe,omitempty"`
    BlocksMe *bool `json:"blocksMe,omitempty" yaml:"blocksMe,omitempty"`
    FollowedByMe *bool `json:"followedByMe,omitempty" yaml:"followedByMe,omitempty"`
    FollowsMe *bool `json:"followsMe,omitempty" yaml:"followsMe,omitempty"`
    MutedByMe *bool `json:"mutedByMe,omitempty" yaml:"mutedByMe,omitempty"`
    MutesMe *bool `json:"mutesMe,omitempty" yaml:"mutesMe,omitempty"`
    RequestedToBeFollowedByMe *bool `json:"requestedToBeFollowedByMe,omitempty" yaml:"requestedToBeFollowedByMe,omitempty"`
    RequestsToFollowMe *bool `json:"requestsToFollowMe,omitempty" yaml:"requestsToFollowMe,omitempty"`
}