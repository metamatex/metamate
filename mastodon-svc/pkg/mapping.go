package pkg

import (
	"github.com/mattn/go-mastodon"
	"github.com/metamatex/metamate/gen/v0/sdk"

)

func MapPostsFromStatuses(statuses []*mastodon.Status) (posts []sdk.Post) {
	for _, s := range statuses {
		posts = append(posts, MapPostFromStatus(*s))
	}

	return
}

func MapPostFromStatus(status mastodon.Status) (post sdk.Post) {
	// ✔ status.ID
	// ✔ status.URI
	// ✔ status.URL
	// ✔ status.Account
	// ✔ status.InReplyToID
	// ✔ status.InReplyToAccountID
	// status.Reblog
	// ✔ status.Content
	// M status.CreatedAt
	// status.Emojis
	// ✔ status.RepliesCount
	// ✔ status.ReblogsCount
	// ✔ status.FavouritesCount
	// ✔ status.Reblogged
	// ✔ status.Favourited
	// ✔ status.Muted
	// ✔ status.Sensitive
	// ✔ status.SpoilerText
	// status.Visibility
	// status.MediaAttachments
	// ✔ status.Mentions
	// status.Tags
	// status.Card
	// status.Application
	// ✔ status.Language
	// status.Pinned

	author := MapSocialAccountFromMastodonAccount(status.Account)

	var muted *bool
	muted0, ok := status.Muted.(bool)
	if ok {
		muted = &muted0
	}

	var reblogged *bool
	reblogged0, ok := status.Reblogged.(bool)
	if ok {
		reblogged = &reblogged0
	}

	var favourited *bool
	favourited0, ok := status.Favourited.(bool)
	if ok {
		favourited = &favourited0
	}

	var repliesToPostId *string
	repliesToPostId0, ok := status.InReplyToID.(string)
	if ok {
		repliesToPostId = &repliesToPostId0
	}

	var repliesToAccountId *string
	repliesToAccountId0, ok := status.InReplyToAccountID.(string)
	if ok {
		repliesToAccountId = &repliesToAccountId0
	}

	post = sdk.Post{
		Id: &sdk.ServiceId{
			Value: sdk.String(string(status.ID)),
		},
		AlternativeIds: []sdk.Id{
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: sdk.String(string(status.URL)),
				},
			},
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: sdk.String(string(status.URI)),
				},
			},
		},
		SpoilerText: &sdk.Text{
			Formatting: &sdk.FormattingKind.Plain,
			Value:      sdk.String(status.SpoilerText),
			Language:   sdk.String(FromISO6391[status.Language]),
		},
		IsSensitive: sdk.Bool(status.Sensitive),
		Content: &sdk.Text{
			Formatting: &sdk.FormattingKind.Plain,
			Value:      sdk.String(status.Content),
			Language:   sdk.String(FromISO6391[status.Language]),
		},
		Relations: &sdk.PostRelations{
			MentionsSocialAccounts: &sdk.SocialAccountsCollection{
				SocialAccounts: MapFromMastodonMentions(status.Mentions),
			},
			FavoredBySocialAccounts: &sdk.SocialAccountsCollection{
				Count: sdk.Int32(int32(status.FavouritesCount)),
			},
			RebloggedByPosts: &sdk.PostsCollection{
				Count: sdk.Int32(int32(status.ReblogsCount)),
			},
			WasRepliedToByPosts: &sdk.PostsCollection{
				Count: sdk.Int32(int32(status.RepliesCount)),
			},
			AuthoredBySocialAccount: &author,
		},
		Relationships: &sdk.PostRelationships{
			FavoredByMe: favourited,
			MutedByMe: muted,
			RebloggedByMe: reblogged,
		},
	}

	if repliesToPostId != nil {
		post.Relations.RepliesToPost = &sdk.Post{
			Id: &sdk.ServiceId{
				Value: repliesToPostId,
			},
		}
	}

	if repliesToAccountId != nil {
		post.Relations.RepliesToSocialAccount = &sdk.SocialAccount{
			Id: &sdk.ServiceId{
				Value: repliesToAccountId,
			},
		}
	}

	return
}

func MapAttachmentFromMastodonAttachment(attachment *mastodon.Attachment) (attachment0 sdk.Attachment) {
	// https://docs.joinmastodon.org/api/entities/#attachment

	// id: "7380901"
	// type: image
	// url: https://files.mastodon.social/media_attachments/files/007/380/901/original/43e66b610e759fbd.jpeg
	// remoteurl: https://bsd.network/system/media_attachments/files/000/561/551/original/65075e8e58cc33c7.jpg
	// previewurl: https://files.mastodon.social/media_attachments/files/007/380/901/small/43e66b610e759fbd.jpeg
	// texturl: ""
	// description: ""

	// id: "18386357"
	// type: image
	// url: https://files.mastodon.social/media_attachments/files/018/386/357/original/6b899dc2eba32eff.jpg
	// remoteurl: ""
	// previewurl: https://files.mastodon.social/media_attachments/files/018/386/357/small/6b899dc2eba32eff.jpg
	// texturl: https://mastodon.social/media/Jc3CDktc_NlcnskS2_E
	// description: ' lilttle prince standing on a planet'

	// mastodon.Attachment
	// ✔ ID          ID
	// Type          string [unknown,image,gifv,video]
	// ✔ URL         string
	// ✔ RemoteURL   string
	// ✔ PreviewURL  string
	// ✔ TextURL     string
	// Description string

	attachment0 = sdk.Attachment{
		Id: &sdk.ServiceId{
			Value: sdk.String(string(attachment.ID)),
		},
		AlternativeIds: []sdk.Id{
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: sdk.String(string(attachment.URL)),
				},
			},
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: sdk.String(string(attachment.RemoteURL)),
				},
			},
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: sdk.String(string(attachment.PreviewURL)),
				},
			},
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: sdk.String(string(attachment.TextURL)),
				},
			},
		},
	}


	return
}

func MapSocialAccountsFromMastodonAccounts(accounts []*mastodon.Account) (people []sdk.SocialAccount) {
	for _, account := range accounts {
		people = append(people, MapSocialAccountFromMastodonAccount(*account))
	}

	return
}

func MapSocialAccountFromMastodonAccount(account mastodon.Account) (person sdk.SocialAccount) {
	// ✔ account.ID             ID
	// ✔ account.Username       string
	// ✔ account.Acct           string
	// account.DisplayName    string
	// account.Locked         bool
	// m account.CreatedAt      time.Time
	// ✔ account.FollowersCount int64
	// ✔ account.FollowingCount int64
	// ✔ account.PostsCount  int64
	// ✔ account.Note           string
	// ✔ account.URL            string
	// ✔ account.Avatar         string
	// account.AvatarStatic   string
	// ✔ account.Header         string
	// account.HeaderStatic   string
	// account.Emojis         []Emoji
	// account.Moved          *Account
	// account.Fields         []Field
	// account.Bot            bool

	person = sdk.SocialAccount{
		Id: &sdk.ServiceId{
			Value: sdk.String(string(account.ID)),
		},
		AlternativeIds: []sdk.Id{
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: sdk.String(string(account.URL)),
				},
			},
			{
				Kind:     &sdk.IdKind.Username,
				Username: sdk.String(account.Username),
			},
			{
				Kind: &sdk.IdKind.Name,
				Name: sdk.String(account.Acct),
			},
		},
		Note: &sdk.Text{
			Formatting: &sdk.FormattingKind.Html,
			Value:      sdk.String(account.Note),
		},
		Avatar: &sdk.Image{
			Url: &sdk.Url{
				Value: sdk.String(account.Avatar),
			},
		},
		Header: &sdk.Image{
			Url: &sdk.Url{
				Value: sdk.String(account.Header),
			},
		},
		Relations: &sdk.SocialAccountRelations{
			FollowedBySocialAccounts: &sdk.SocialAccountsCollection{
				Count: sdk.Int32(int32(account.FollowersCount)),
			},
			FollowsSocialAccounts: &sdk.SocialAccountsCollection{
				Count: sdk.Int32(int32(account.FollowingCount)),
			},
			AuthorsPosts: &sdk.PostsCollection{
				Count: sdk.Int32(int32(account.StatusesCount)),
			},
		},
	}

	return
}

func MapFromMastodonMentions(mentions []mastodon.Mention) (people []sdk.SocialAccount) {
	for _, m := range mentions {
		people = append(people, MapFromMastodonMention(m))
	}

	return
}

func MapFromMastodonMention(mention mastodon.Mention) (person sdk.SocialAccount) {
	// ✔ mention.URL      string
	// ✔ mention.Username string
	// ✔ mention.Acct     string
	// ✔ mention.ID       mastodon.ID

	person = sdk.SocialAccount{
		Id: &sdk.ServiceId{
			Value: sdk.String(string(mention.ID)),
		},
		AlternativeIds: []sdk.Id{
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: sdk.String(string(mention.URL)),
				},
			},
			{
				Kind:     &sdk.IdKind.Username,
				Username: sdk.String(mention.Username),
			},
			{
				Kind: &sdk.IdKind.Name,
				Name: sdk.String(mention.Acct),
			},
		},
	}

	return
}

func MapPostToMastodonToot(status sdk.Post) (toot *mastodon.Toot) {
	toot = &mastodon.Toot{}

	if status.Content != nil && status.Content.Value != nil {
		toot.Status = *status.Content.Value
	}

	if status.IsSensitive != nil {
		toot.Sensitive = *status.IsSensitive
	}

	if status.SpoilerText != nil && status.SpoilerText.Value != nil {
		toot.SpoilerText = *status.SpoilerText.Value
	}

	if status.Relations != nil &&
		status.Relations.RepliesToPost != nil &&
		status.Relations.RepliesToPost.Id != nil &&
		status.Relations.RepliesToPost.Id.Value != nil {
		toot.InReplyToID = mastodon.ID(*status.Relations.RepliesToPost.Id.Value)
	}

	return
}
