// generated by metactl sdk gen 
package sdk

const (
	PostRelationsSelectName = "PostRelationsSelect"
)

type PostRelationsSelect struct {
    All *bool `json:"all,omitempty" yaml:"all,omitempty"`
    AttachesAttachments *AttachmentsCollectionSelect `json:"attachesAttachments,omitempty" yaml:"attachesAttachments,omitempty"`
    AuthoredBySocialAccount *SocialAccountSelect `json:"authoredBySocialAccount,omitempty" yaml:"authoredBySocialAccount,omitempty"`
    ContainedByPostFeeds *PostFeedsCollectionSelect `json:"containedByPostFeeds,omitempty" yaml:"containedByPostFeeds,omitempty"`
    FavoredBySocialAccounts *SocialAccountsCollectionSelect `json:"favoredBySocialAccounts,omitempty" yaml:"favoredBySocialAccounts,omitempty"`
    MentionsSocialAccounts *SocialAccountsCollectionSelect `json:"mentionsSocialAccounts,omitempty" yaml:"mentionsSocialAccounts,omitempty"`
    MutedBySocialAccounts *SocialAccountsCollectionSelect `json:"mutedBySocialAccounts,omitempty" yaml:"mutedBySocialAccounts,omitempty"`
    NotReadBySocialAccounts *SocialAccountsCollectionSelect `json:"notReadBySocialAccounts,omitempty" yaml:"notReadBySocialAccounts,omitempty"`
    ReadBySocialAccounts *SocialAccountsCollectionSelect `json:"readBySocialAccounts,omitempty" yaml:"readBySocialAccounts,omitempty"`
    RebloggedByPosts *PostsCollectionSelect `json:"rebloggedByPosts,omitempty" yaml:"rebloggedByPosts,omitempty"`
    RebloggedBySocialAccounts *SocialAccountsCollectionSelect `json:"rebloggedBySocialAccounts,omitempty" yaml:"rebloggedBySocialAccounts,omitempty"`
    ReblogsPost *PostSelect `json:"reblogsPost,omitempty" yaml:"reblogsPost,omitempty"`
    RepliesToPost *PostSelect `json:"repliesToPost,omitempty" yaml:"repliesToPost,omitempty"`
    RepliesToSocialAccount *SocialAccountSelect `json:"repliesToSocialAccount,omitempty" yaml:"repliesToSocialAccount,omitempty"`
    WasRepliedToByPosts *PostsCollectionSelect `json:"wasRepliedToByPosts,omitempty" yaml:"wasRepliedToByPosts,omitempty"`
}