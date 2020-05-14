package pkg

import (
	"github.com/mattn/go-mastodon"
	"github.com/metamatex/metamate/gen/v0/mql"

)

func MapPostsFromStatuses(statuses []*mastodon.Status) (posts []mql.Post) {
	for _, s := range statuses {
		posts = append(posts, MapPostFromStatus(*s))
	}

	return
}

func MapPostFromStatus(status mastodon.Status) (post mql.Post) {
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

	post = mql.Post{
		Id: &mql.ServiceId{
			Value: mql.String(string(status.ID)),
		},
		AlternativeIds: []mql.Id{
			{
				Kind: &mql.IdKind.Url,
				Url: &mql.Url{
					Value: mql.String(string(status.URL)),
				},
			},
			{
				Kind: &mql.IdKind.Url,
				Url: &mql.Url{
					Value: mql.String(string(status.URI)),
				},
			},
		},
		SpoilerText: &mql.Text{
			Formatting: &mql.FormattingKind.Plain,
			Value:      mql.String(status.SpoilerText),
			Language:   mql.String(FromISO6391[status.Language]),
		},
		IsSensitive: mql.Bool(status.Sensitive),
		Content: &mql.Text{
			Formatting: &mql.FormattingKind.Plain,
			Value:      mql.String(status.Content),
			Language:   mql.String(FromISO6391[status.Language]),
		},
		Relations: &mql.PostRelations{
			MentionsSocialAccounts: &mql.SocialAccountsCollection{
				SocialAccounts: MapFromMastodonMentions(status.Mentions),
			},
			FavoredBySocialAccounts: &mql.SocialAccountsCollection{
				Count: mql.Int32(int32(status.FavouritesCount)),
			},
			RebloggedByPosts: &mql.PostsCollection{
				Count: mql.Int32(int32(status.ReblogsCount)),
			},
			WasRepliedToByPosts: &mql.PostsCollection{
				Count: mql.Int32(int32(status.RepliesCount)),
			},
			AuthoredBySocialAccount: &author,
		},
		Relationships: &mql.PostRelationships{
			FavoredByMe: favourited,
			MutedByMe: muted,
			RebloggedByMe: reblogged,
		},
	}

	if repliesToPostId != nil {
		post.Relations.RepliesToPost = &mql.Post{
			Id: &mql.ServiceId{
				Value: repliesToPostId,
			},
		}
	}

	if repliesToAccountId != nil {
		post.Relations.RepliesToSocialAccount = &mql.SocialAccount{
			Id: &mql.ServiceId{
				Value: repliesToAccountId,
			},
		}
	}

	return
}

func MapAttachmentFromMastodonAttachment(attachment *mastodon.Attachment) (attachment0 mql.Attachment) {
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

	attachment0 = mql.Attachment{
		Id: &mql.ServiceId{
			Value: mql.String(string(attachment.ID)),
		},
		AlternativeIds: []mql.Id{
			{
				Kind: &mql.IdKind.Url,
				Url: &mql.Url{
					Value: mql.String(string(attachment.URL)),
				},
			},
			{
				Kind: &mql.IdKind.Url,
				Url: &mql.Url{
					Value: mql.String(string(attachment.RemoteURL)),
				},
			},
			{
				Kind: &mql.IdKind.Url,
				Url: &mql.Url{
					Value: mql.String(string(attachment.PreviewURL)),
				},
			},
			{
				Kind: &mql.IdKind.Url,
				Url: &mql.Url{
					Value: mql.String(string(attachment.TextURL)),
				},
			},
		},
	}


	return
}

func MapSocialAccountsFromMastodonAccounts(accounts []*mastodon.Account) (people []mql.SocialAccount) {
	for _, account := range accounts {
		people = append(people, MapSocialAccountFromMastodonAccount(*account))
	}

	return
}

func MapSocialAccountFromMastodonAccount(account mastodon.Account) (person mql.SocialAccount) {
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

	person = mql.SocialAccount{
		Id: &mql.ServiceId{
			Value: mql.String(string(account.ID)),
		},
		AlternativeIds: []mql.Id{
			{
				Kind: &mql.IdKind.Url,
				Url: &mql.Url{
					Value: mql.String(string(account.URL)),
				},
			},
			{
				Kind:     &mql.IdKind.Username,
				Username: mql.String(account.Username),
			},
			{
				Kind: &mql.IdKind.Name,
				Name: mql.String(account.Acct),
			},
		},
		Note: &mql.Text{
			Formatting: &mql.FormattingKind.Html,
			Value:      mql.String(account.Note),
		},
		Avatar: &mql.Image{
			Url: &mql.Url{
				Value: mql.String(account.Avatar),
			},
		},
		Header: &mql.Image{
			Url: &mql.Url{
				Value: mql.String(account.Header),
			},
		},
		Relations: &mql.SocialAccountRelations{
			FollowedBySocialAccounts: &mql.SocialAccountsCollection{
				Count: mql.Int32(int32(account.FollowersCount)),
			},
			FollowsSocialAccounts: &mql.SocialAccountsCollection{
				Count: mql.Int32(int32(account.FollowingCount)),
			},
			AuthorsPosts: &mql.PostsCollection{
				Count: mql.Int32(int32(account.StatusesCount)),
			},
		},
	}

	return
}

func MapFromMastodonMentions(mentions []mastodon.Mention) (people []mql.SocialAccount) {
	for _, m := range mentions {
		people = append(people, MapFromMastodonMention(m))
	}

	return
}

func MapFromMastodonMention(mention mastodon.Mention) (person mql.SocialAccount) {
	// ✔ mention.URL      string
	// ✔ mention.Username string
	// ✔ mention.Acct     string
	// ✔ mention.ID       mastodon.ID

	person = mql.SocialAccount{
		Id: &mql.ServiceId{
			Value: mql.String(string(mention.ID)),
		},
		AlternativeIds: []mql.Id{
			{
				Kind: &mql.IdKind.Url,
				Url: &mql.Url{
					Value: mql.String(string(mention.URL)),
				},
			},
			{
				Kind:     &mql.IdKind.Username,
				Username: mql.String(mention.Username),
			},
			{
				Kind: &mql.IdKind.Name,
				Name: mql.String(mention.Acct),
			},
		},
	}

	return
}

func MapPostToMastodonToot(status mql.Post) (toot *mastodon.Toot) {
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
