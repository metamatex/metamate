package types

type GetUserAboutResponse struct {
	Error
	Kind string    `json:"kind"`
	Data *UserData `json:"data"`
}

type UserData struct {
	IsEmployee        *bool        `json:"is_employee"`
	IconImg           *string      `json:"icon_img"`
	PrefShowSnoovatar *bool        `json:"pref_show_snoovatar"`
	Name              *string      `json:"name"`
	IsFriend          *bool        `json:"is_friend"`
	Created           *float64     `json:"created"`
	HasSubscribed     *bool        `json:"has_subscribed"`
	HideFromRobots    *bool        `json:"hide_from_robots"`
	CreatedUtc        *float64     `json:"created_utc"`
	LinkKarma         *int32       `json:"link_karma"`
	CommentKarma      *int32       `json:"comment_karma"`
	IsGold            *bool        `json:"is_gold"`
	IsMod             *bool        `json:"is_mod"`
	Verified          *bool        `json:"verified"`
	Subreddit         *Subreddit_s `json:"subreddit"`
	HasVerifiedEmail  *bool        `json:"has_verified_email"`
	Id                *string      `json:"id"`
}

type Subreddit_s struct {
	DefaultSet                 *bool     `json:"default_set"`
	UserIsContributor          *bool     `json:"user_is_contributor"`
	BannerImg                  *string   `json:"banner_img"`
	DisableContributorRequests *bool     `json:"disable_contributor_requests"`
	UserIsBanned               *bool     `json:"user_is_banned"`
	FreeFormReports            *bool     `json:"free_form_reports"`
	CommunityIcon              *string   `json:"community_icon"`
	ShowMedia                  *bool     `json:"show_media"`
	IconColor                  *string   `json:"icon_color"`
	UserIsMuted                *bool     `json:"user_is_muted"`
	DisplayName                *string   `json:"display_name"`
	HeaderImg                  *string   `json:"header_img"` // *
	Title                      *string   `json:"title"`
	Over18                     *bool     `json:"over_18"`
	IconSize                   []float64 `json:"icon_size"`
	PrimaryColor               *string   `json:"primary_color"`
	IconImg                    *string   `json:"icon_img"`
	Description                *string   `json:"description"`
	HeaderSize                 *string   `json:"header_size"` // *
	RestrictPosting            *bool     `json:"restrict_posting"`
	RestrictCommenting         *bool     `json:"restrict_commenting"`
	Subscribers                *float64  `json:"subscribers"`
	IsDefaultIcon              *bool     `json:"is_default_icon"`
	LinkFlairPosition          *string   `json:"link_flair_position"`
	DisplayNamePrefixed        *string   `json:"display_name_prefixed"`
	KeyColor                   *string   `json:"key_color"`
	Name                       *string   `json:"name"`
	IsDefaultBanner            *bool     `json:"is_default_banner"`
	Url                        *string   `json:"url"`
	BannerSize                 []float64 `json:"banner_size"`
	UserIsModerator            *bool     `json:"user_is_moderator"`
	PublicDescription          *string   `json:"public_description"`
	LinkFlairEnabled           *bool     `json:"link_flair_enabled"`
	SubredditType              *string   `json:"subreddit_type"`
	UserIsSubscriber           *bool     `json:"user_is_subscriber"`
}
