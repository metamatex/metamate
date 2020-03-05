package pkg

import (
	"github.com/mattn/go-mastodon"
	"github.com/metamatex/metamatemono/gen/v0/sdk"
	"github.com/metamatex/metamatemono/gen/v0/sdk/utils/ptr"
)

func MapStatusesFromMastodonStatuses(statuses []*mastodon.Status) (statuses0 []sdk.Status) {
	for _, s := range statuses {
		statuses0 = append(statuses0, MapStatusFromMastodonStatus(*s))
	}

	return
}

func MapStatusFromMastodonStatus(status mastodon.Status) (status0 sdk.Status) {
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

	author := MapPersonFromMastodonAccount(status.Account)

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

	var repliesToStatusId *string
	repliesToStatusId0, ok := status.InReplyToID.(string)
	if ok {
		repliesToStatusId = &repliesToStatusId0
	}

	var repliesToAccountId *string
	repliesToAccountId0, ok := status.InReplyToAccountID.(string)
	if ok {
		repliesToAccountId = &repliesToAccountId0
	}

	status0 = sdk.Status{
		Id: &sdk.ServiceId{
			Value: ptr.String(string(status.ID)),
		},
		AlternativeIds: []sdk.Id{
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: ptr.String(string(status.URL)),
				},
			},
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: ptr.String(string(status.URI)),
				},
			},
		},
		SpoilerText: &sdk.Text{
			Formatting: &sdk.FormattingKind.Plain,
			Value:      ptr.String(status.SpoilerText),
			Language:   ptr.String(FromISO6391[status.Language]),
		},
		Sensitive: ptr.Bool(status.Sensitive),
		Content: &sdk.Text{
			Formatting: &sdk.FormattingKind.Plain,
			Value:      ptr.String(status.Content),
			Language:   ptr.String(FromISO6391[status.Language]),
		},
		Relations: &sdk.StatusRelations{
			MentionsPeople: &sdk.PeopleCollection{
				People: MapFromMastodonMentions(status.Mentions),
			},
			FavoredByPeople: &sdk.PeopleCollection{
				Meta: &sdk.CollectionMeta{
					Count: ptr.Int32(int32(status.FavouritesCount)),
				},
			},
			RebloggedByStatuses: &sdk.StatusesCollection{
				Meta: &sdk.CollectionMeta{
					Count: ptr.Int32(int32(status.ReblogsCount)),
				},
			},
			WasRepliedToByStatuses: &sdk.StatusesCollection{
				Meta: &sdk.CollectionMeta{
					Count: ptr.Int32(int32(status.RepliesCount)),
				},
			},
			AuthoredByPerson: &author,
		},
		Relationships: &sdk.StatusRelationships{
			FavoredByMe: favourited,
			MutedByMe: muted,
			RebloggedByMe: reblogged,
		},
	}

	if repliesToStatusId != nil {
		status0.Relations.RepliesToStatus = &sdk.Status{
			Id: &sdk.ServiceId{
				Value: repliesToStatusId,
			},
		}
	}

	if repliesToAccountId != nil {
		status0.Relations.RepliesToPerson = &sdk.Person{
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
			Value: ptr.String(string(attachment.ID)),
		},
		AlternativeIds: []sdk.Id{
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: ptr.String(string(attachment.URL)),
				},
			},
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: ptr.String(string(attachment.RemoteURL)),
				},
			},
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: ptr.String(string(attachment.PreviewURL)),
				},
			},
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: ptr.String(string(attachment.TextURL)),
				},
			},
		},
	}


	return
}

func MapPeopleFromMastodonAccounts(accounts []*mastodon.Account) (people []sdk.Person) {
	for _, account := range accounts {
		people = append(people, MapPersonFromMastodonAccount(*account))
	}

	return
}

func MapPersonFromMastodonAccount(account mastodon.Account) (person sdk.Person) {
	// ✔ account.ID             ID
	// ✔ account.Username       string
	// ✔ account.Acct           string
	// account.DisplayName    string
	// account.Locked         bool
	// m account.CreatedAt      time.Time
	// ✔ account.FollowersCount int64
	// ✔ account.FollowingCount int64
	// ✔ account.StatusesCount  int64
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

	person = sdk.Person{
		Id: &sdk.ServiceId{
			Value: ptr.String(string(account.ID)),
		},
		AlternativeIds: []sdk.Id{
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: ptr.String(string(account.URL)),
				},
			},
			{
				Kind:     &sdk.IdKind.Username,
				Username: ptr.String(account.Username),
			},
			{
				Kind: &sdk.IdKind.Name,
				Name: ptr.String(account.Acct),
			},
		},
		Note: &sdk.Text{
			Formatting: &sdk.FormattingKind.Html,
			Value:      ptr.String(account.Note),
		},
		Avatar: &sdk.Image{
			Url: &sdk.Url{
				Value: ptr.String(account.Avatar),
			},
		},
		Header: &sdk.Image{
			Url: &sdk.Url{
				Value: ptr.String(account.Header),
			},
		},
		Relations: &sdk.PersonRelations{
			FollowedByPeople: &sdk.PeopleCollection{
				Meta: &sdk.CollectionMeta{
					Count: ptr.Int32(int32(account.FollowersCount)),
				},
			},
			FollowsPeople: &sdk.PeopleCollection{
				Meta: &sdk.CollectionMeta{
					Count: ptr.Int32(int32(account.FollowingCount)),
				},
			},
			AuthorsStatuses: &sdk.StatusesCollection{
				Meta: &sdk.CollectionMeta{
					Count: ptr.Int32(int32(account.StatusesCount)),
				},
			},
		},
	}

	return
}

func MapFromMastodonMentions(mentions []mastodon.Mention) (people []sdk.Person) {
	for _, m := range mentions {
		people = append(people, MapFromMastodonMention(m))
	}

	return
}

func MapFromMastodonMention(mention mastodon.Mention) (person sdk.Person) {
	// ✔ mention.URL      string
	// ✔ mention.Username string
	// ✔ mention.Acct     string
	// ✔ mention.ID       mastodon.ID

	person = sdk.Person{
		Id: &sdk.ServiceId{
			Value: ptr.String(string(mention.ID)),
		},
		AlternativeIds: []sdk.Id{
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: ptr.String(string(mention.URL)),
				},
			},
			{
				Kind:     &sdk.IdKind.Username,
				Username: ptr.String(mention.Username),
			},
			{
				Kind: &sdk.IdKind.Name,
				Name: ptr.String(mention.Acct),
			},
		},
	}

	return
}

func MapStatusToMastodonToot(status sdk.Status) (toot *mastodon.Toot) {
	toot = &mastodon.Toot{}

	if status.Content != nil && status.Content.Value != nil {
		toot.Status = *status.Content.Value
	}

	if status.Sensitive != nil {
		toot.Sensitive = *status.Sensitive
	}

	if status.SpoilerText != nil && status.SpoilerText.Value != nil {
		toot.SpoilerText = *status.SpoilerText.Value
	}

	if status.Relations != nil &&
		status.Relations.RepliesToStatus != nil &&
		status.Relations.RepliesToStatus.Id != nil &&
		status.Relations.RepliesToStatus.Id.Value != nil {
		toot.InReplyToID = mastodon.ID(*status.Relations.RepliesToStatus.Id.Value)
	}

	return
}
