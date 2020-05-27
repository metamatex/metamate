package types

type GetSubredditSubmissionsResponse struct {
	Error
	Kind string          `json:"kind"`
	Data PostListingData `json:"data"`
}

type PostListingData struct {
	Modhash  string             `json:"modhash"`
	Dist     float64            `json:"dist"`
	Children []PostListingChild `json:"children"`
	After    *string            `json:"after"`
	Before   *string            `json:"before"`
}

type PostListingChild struct {
	Kind   string               `json:"kind"`
	Data   PostListingChildData `json:"data"`
	After  string               `json:"after"`
	Before string               `json:"before"`
}

type PostListingChildData struct {
	ApprovedAtUtc              *float64           `json:"approved_at_utc"`
	Subreddit                  *string            `json:"subreddit"`
	Selftext                   *string            `json:"selftext"`
	AuthorFullname             *string            `json:"author_fullname"`
	Saved                      *bool              `json:"saved"`
	ModReasonTitle             *string            `json:"mod_reason_title"`
	Gilded                     *float64           `json:"gilded"`
	Clicked                    *bool              `json:"clicked"`
	Title                      *string            `json:"title"`
	LinkFlairRichtext          interface{}        `json:"link_flair_richtext"`
	SubredditNamePrefixed      *string            `json:"subreddit_name_prefixed"`
	Hidden                     *bool              `json:"hidden"`
	Pwls                       *float64           `json:"pwls"`
	LinkFlairCssClass          *string            `json:"link_flair_css_class"`
	Downs                      *float64           `json:"downs"`
	ThumbnailHeight            *float64           `json:"thumbnail_height"`
	HideScore                  *bool              `json:"hide_score"`
	Name                       *string            `json:"name"`
	Quarantine                 *bool              `json:"quarantine"`
	LinkFlairTextColor         *string            `json:"link_flair_text_color"`
	AuthorFlairBackgroundColor *string            `json:"author_flair_background_color"`
	SubredditType              *string            `json:"subreddit_type"`
	Ups                        *int32             `json:"ups"`
	TotalAwardsReceived        *float64           `json:"total_awards_received"`
	MediaEmbed                 interface{}        `json:"media_embed"`
	ThumbnailWidth             *float64           `json:"thumbnail_width"`
	AuthorFlairTemplateId      *string            `json:"author_flair_template_id"`
	IsOriginalContent          *bool              `json:"is_original_content"`
	UserReports                []string           `json:"user_reports"`
	SecureMedia                *string            `json:"secure_media"`
	IsRedditMediaDomain        *bool              `json:"is_reddit_media_domain"`
	IsMeta                     *bool              `json:"is_meta"`
	Category                   *string            `json:"category"`
	SecureMediaEmbed           interface{}        `json:"secure_media_embed"`
	LinkFlairText              *string            `json:"link_flair_text"`
	CanModPost                 *bool              `json:"can_mod_post"`
	Score                      *float64           `json:"score"`
	ApprovedBy                 *string            `json:"approved_by"`
	Thumbnail                  *string            `json:"thumbnail"`
	Edited                     *bool              `json:"edited"`
	AuthorFlairCssClass        *string            `json:"author_flair_css_class"`
	AuthorFlairRichtext        []string           `json:"author_flair_richtext"`
	Gildings                   map[string]float64 `json:"gildings"`
	PostHint                   *string            `json:"post_hint"`
	ContentCategories          []string           `json:"content_categories"`
	IsSelf                     *bool              `json:"is_self"`
	ModNote                    *string            `json:"mod_note"`
	Created                    *float64           `json:"created"`
	LinkFlairType              *string            `json:"link_flair_type"`
	Wls                        *float64           `json:"wls"`
	BannedBy                   *string            `json:"banned_by"`
	AuthorFlairType            *string            `json:"author_flair_type"`
	Domain                     *string            `json:"domain"`
	SelftextHtml               *string            `json:"selftext_html"`
	Likes                      *float64           `json:"likes"`
	SuggestedSort              *string            `json:"suggested_sort"`
	BannedAtUtc                *float64           `json:"banned_at_utc"`
	ViewCount                  *float64           `json:"view_count"`
	Archived                   *bool              `json:"archived"`
	NoFollow                   *bool              `json:"no_follow"`
	IsCrosspostable            *bool              `json:"is_crosspostable"`
	Pinned                     *bool              `json:"pinned"`
	Over18                     *bool              `json:"over_18"`
	Preview                    interface{}        `json:"preview"`
	Awardings                  []PostAward        `json:"all_awardings"`
	MediaOnly                  *bool              `json:"media_only"`
	CanGild                    *bool              `json:"can_gild"`
	Spoiler                    *bool              `json:"spoiler"`
	Locked                     *bool              `json:"locked"`
	AuthorFlairText            *string            `json:"author_flair_text"`
	Visited                    *bool              `json:"visited"`
	NumReports                 *float64           `json:"num_reports"`
	Distinguished              *bool              `json:"distinguished"`
	SubredditId                *string            `json:"subreddit_id"`
	ModReasonBy                *string            `json:"mod_reason_by"`
	RemovalReason              *string            `json:"removal_reason"`
	LinkFlairBackgroundColor   *string            `json:"link_flair_background_color"`
	Id                         *string            `json:"id"`
	IsRobotIndexable           *bool              `json:"is_robot_indexable"`
	ReportReasons              *string            `json:"report_reasons"`
	Author                     *string            `json:"author"`
	NumCrossposts              *float64           `json:"num_crossposts"`
	NumComments                *int32             `json:"num_comments"`
	SendReplies                *bool              `json:"send_replies"`
	WhitelistStatus            *string            `json:"whitelist_status"`
	ContestMode                *bool              `json:"contest_mode"`
	ModReports                 []string           `json:"mod_reports"`
	AuthorPatreonFlair         *bool              `json:"author_patreon_flair"`
	AuthorFlairTextColor       *string            `json:"author_flair_text_color"`
	Permalink                  *string            `json:"permalink"`
	ParentWhitelistStatus      *string            `json:"parent_whitelist_status"`
	Stickied                   *bool              `json:"stickied"`
	Url                        *string            `json:"url"`
	SubredditSubscribers       *float64           `json:"subreddit_subscribers"`
	CreatedUtc                 *float64           `json:"created_utc"`
	Media                      []string           `json:"media"`
	IsVideo                    *bool              `json:"is_video"`
}

type PostAward struct {
	IsEnabled           bool            `json:"is_enabled"`
	Count               float64         `json:"count"`
	SubredditId         string          `json:"subreddit_id"`
	Description         string          `json:"description"`
	CoinReward          float64         `json:"coin_reward"`
	IconWidth           float64         `json:"icon_width"`
	IconUrl             string          `json:"icon_url"`
	DaysOfPremium       float64         `json:"days_of_premium"`
	IconHeight          float64         `json:"icon_height"`
	ResizedIcons        []PostAwardIcon `json:"resized_icons"`
	DaysOfDripExtension float64         `json:"days_of_drip_extension"`
	AwardType           string          `json:"award_type"`
	CoinPrice           float64         `json:"coin_price"`
	Id                  string          `json:"id"`
	Name                string          `json:"name"`
}

type PostAwardIcon struct {
	Url    string  `json:"url"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type PostPreview struct {
	Images  []PostPreviewImage `json:"images"`
	Enabled bool               `json:"enabled"`
}

type PostPreviewImage struct {
	Source      PostPreviewImageSource      `json:"source"`
	Resolutions PostPreviewImageResolutions `json:"resolutions"`
	Variants    []string                    `json:"variants"`
	Id          string                      `json:"id"`
}

type PostPreviewImageResolutions struct {
	Url    string  `json:"url"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type PostPreviewImageSource struct {
	Url    string  `json:"url"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}
