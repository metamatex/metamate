// generated by metactl sdk gen 
package sdk

const (
	GetSocialAccountsRelationsName = "GetSocialAccountsRelations"
)

type GetSocialAccountsRelations struct {
    AuthorsStatuses *GetStatusesCollection `json:"authorsStatuses,omitempty" yaml:"authorsStatuses,omitempty"`
    BlockedBySocialAccounts *GetSocialAccountsCollection `json:"blockedBySocialAccounts,omitempty" yaml:"blockedBySocialAccounts,omitempty"`
    BlocksSocialAccounts *GetSocialAccountsCollection `json:"blocksSocialAccounts,omitempty" yaml:"blocksSocialAccounts,omitempty"`
    FavorsStatuses *GetStatusesCollection `json:"favorsStatuses,omitempty" yaml:"favorsStatuses,omitempty"`
    FollowedBySocialAccounts *GetSocialAccountsCollection `json:"followedBySocialAccounts,omitempty" yaml:"followedBySocialAccounts,omitempty"`
    FollowsSocialAccounts *GetSocialAccountsCollection `json:"followsSocialAccounts,omitempty" yaml:"followsSocialAccounts,omitempty"`
    MentionedByStatuses *GetStatusesCollection `json:"mentionedByStatuses,omitempty" yaml:"mentionedByStatuses,omitempty"`
    MutedBySocialAccounts *GetSocialAccountsCollection `json:"mutedBySocialAccounts,omitempty" yaml:"mutedBySocialAccounts,omitempty"`
    MutesSocialAccounts *GetSocialAccountsCollection `json:"mutesSocialAccounts,omitempty" yaml:"mutesSocialAccounts,omitempty"`
    MutesStatuses *GetStatusesCollection `json:"mutesStatuses,omitempty" yaml:"mutesStatuses,omitempty"`
    NotReadStatuses *GetStatusesCollection `json:"notReadStatuses,omitempty" yaml:"notReadStatuses,omitempty"`
    ParticipatesFeeds *GetFeedsCollection `json:"participatesFeeds,omitempty" yaml:"participatesFeeds,omitempty"`
    ReadStatuses *GetStatusesCollection `json:"readStatuses,omitempty" yaml:"readStatuses,omitempty"`
    ReblogsStatuses *GetStatusesCollection `json:"reblogsStatuses,omitempty" yaml:"reblogsStatuses,omitempty"`
    RequestedToBeFollowedBySocialAccounts *GetSocialAccountsCollection `json:"requestedToBeFollowedBySocialAccounts,omitempty" yaml:"requestedToBeFollowedBySocialAccounts,omitempty"`
    RequestsToFollowSocialAccounts *GetSocialAccountsCollection `json:"requestsToFollowSocialAccounts,omitempty" yaml:"requestsToFollowSocialAccounts,omitempty"`
    WasRepliedToByStatuses *GetStatusesCollection `json:"wasRepliedToByStatuses,omitempty" yaml:"wasRepliedToByStatuses,omitempty"`
}