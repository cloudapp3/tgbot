// Telegram Bot API type definitions aligned with the official docs.
// Source: https://core.telegram.org/bots/api (Bot API 9.5, March 1, 2026)

package tgbot

// AcceptedGiftTypes maps to Telegram Bot API type "AcceptedGiftTypes".
type AcceptedGiftTypes struct {
	UnlimitedGifts      bool `json:"unlimited_gifts"`
	LimitedGifts        bool `json:"limited_gifts"`
	UniqueGifts         bool `json:"unique_gifts"`
	PremiumSubscription bool `json:"premium_subscription"`
	GiftsFromChannels   bool `json:"gifts_from_channels"`
}

// AffiliateInfo maps to Telegram Bot API type "AffiliateInfo".
type AffiliateInfo struct {
	AffiliateUser      *User `json:"affiliate_user"`
	AffiliateChat      *Chat `json:"affiliate_chat"`
	CommissionPerMille int64 `json:"commission_per_mille"`
	Amount             int64 `json:"amount"`
	NanostarAmount     int64 `json:"nanostar_amount"`
}

// Animation maps to Telegram Bot API type "Animation".
type Animation struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int64      `json:"width"`
	Height       int64      `json:"height"`
	Duration     int64      `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail"`
	FileName     string     `json:"file_name"`
	MimeType     string     `json:"mime_type"`
	FileSize     int64      `json:"file_size"`
}

// Audio maps to Telegram Bot API type "Audio".
type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int64      `json:"duration"`
	Performer    string     `json:"performer"`
	Title        string     `json:"title"`
	FileName     string     `json:"file_name"`
	MimeType     string     `json:"mime_type"`
	FileSize     int64      `json:"file_size"`
	Thumbnail    *PhotoSize `json:"thumbnail"`
}

// BackgroundFill is a union type in Telegram Bot API.
type BackgroundFill interface {
	isBackgroundFill()
}

// BackgroundFillFreeformGradient maps to Telegram Bot API type "BackgroundFillFreeformGradient".
type BackgroundFillFreeformGradient struct {
	Type   string  `json:"type"`
	Colors []int64 `json:"colors"`
}

func (*BackgroundFillFreeformGradient) isBackgroundFill() {}

// BackgroundFillGradient maps to Telegram Bot API type "BackgroundFillGradient".
type BackgroundFillGradient struct {
	Type          string `json:"type"`
	TopColor      int64  `json:"top_color"`
	BottomColor   int64  `json:"bottom_color"`
	RotationAngle int64  `json:"rotation_angle"`
}

func (*BackgroundFillGradient) isBackgroundFill() {}

// BackgroundFillSolid maps to Telegram Bot API type "BackgroundFillSolid".
type BackgroundFillSolid struct {
	Type  string `json:"type"`
	Color int64  `json:"color"`
}

func (*BackgroundFillSolid) isBackgroundFill() {}

// BackgroundType is a union type in Telegram Bot API.
type BackgroundType interface {
	isBackgroundType()
}

// BackgroundTypeChatTheme maps to Telegram Bot API type "BackgroundTypeChatTheme".
type BackgroundTypeChatTheme struct {
	Type      string `json:"type"`
	ThemeName string `json:"theme_name"`
}

func (*BackgroundTypeChatTheme) isBackgroundType() {}

// BackgroundTypeFill maps to Telegram Bot API type "BackgroundTypeFill".
type BackgroundTypeFill struct {
	Type             string         `json:"type"`
	Fill             BackgroundFill `json:"fill"`
	DarkThemeDimming int64          `json:"dark_theme_dimming"`
}

func (*BackgroundTypeFill) isBackgroundType() {}

// BackgroundTypePattern maps to Telegram Bot API type "BackgroundTypePattern".
type BackgroundTypePattern struct {
	Type       string         `json:"type"`
	Document   *Document      `json:"document"`
	Fill       BackgroundFill `json:"fill"`
	Intensity  int64          `json:"intensity"`
	IsInverted bool           `json:"is_inverted"`
	IsMoving   bool           `json:"is_moving"`
}

func (*BackgroundTypePattern) isBackgroundType() {}

// BackgroundTypeWallpaper maps to Telegram Bot API type "BackgroundTypeWallpaper".
type BackgroundTypeWallpaper struct {
	Type             string    `json:"type"`
	Document         *Document `json:"document"`
	DarkThemeDimming int64     `json:"dark_theme_dimming"`
	IsBlurred        bool      `json:"is_blurred"`
	IsMoving         bool      `json:"is_moving"`
}

func (*BackgroundTypeWallpaper) isBackgroundType() {}

// Birthdate maps to Telegram Bot API type "Birthdate".
type Birthdate struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

// BotCommand maps to Telegram Bot API type "BotCommand".
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// BotCommandScope is a union type in Telegram Bot API.
type BotCommandScope interface {
	isBotCommandScope()
}

// BotCommandScopeAllChatAdministrators maps to Telegram Bot API type "BotCommandScopeAllChatAdministrators".
type BotCommandScopeAllChatAdministrators struct {
	Type string `json:"type"`
}

func (*BotCommandScopeAllChatAdministrators) isBotCommandScope() {}

// BotCommandScopeAllGroupChats maps to Telegram Bot API type "BotCommandScopeAllGroupChats".
type BotCommandScopeAllGroupChats struct {
	Type string `json:"type"`
}

func (*BotCommandScopeAllGroupChats) isBotCommandScope() {}

// BotCommandScopeAllPrivateChats maps to Telegram Bot API type "BotCommandScopeAllPrivateChats".
type BotCommandScopeAllPrivateChats struct {
	Type string `json:"type"`
}

func (*BotCommandScopeAllPrivateChats) isBotCommandScope() {}

// BotCommandScopeChat maps to Telegram Bot API type "BotCommandScopeChat".
type BotCommandScopeChat struct {
	Type   string `json:"type"`
	ChatID any    `json:"chat_id"`
}

func (*BotCommandScopeChat) isBotCommandScope() {}

// BotCommandScopeChatAdministrators maps to Telegram Bot API type "BotCommandScopeChatAdministrators".
type BotCommandScopeChatAdministrators struct {
	Type   string `json:"type"`
	ChatID any    `json:"chat_id"`
}

func (*BotCommandScopeChatAdministrators) isBotCommandScope() {}

// BotCommandScopeChatMember maps to Telegram Bot API type "BotCommandScopeChatMember".
type BotCommandScopeChatMember struct {
	Type   string `json:"type"`
	ChatID any    `json:"chat_id"`
	UserID int64  `json:"user_id"`
}

func (*BotCommandScopeChatMember) isBotCommandScope() {}

// BotCommandScopeDefault maps to Telegram Bot API type "BotCommandScopeDefault".
type BotCommandScopeDefault struct {
	Type string `json:"type"`
}

func (*BotCommandScopeDefault) isBotCommandScope() {}

// BotDescription maps to Telegram Bot API type "BotDescription".
type BotDescription struct {
	Description string `json:"description"`
}

// BotName maps to Telegram Bot API type "BotName".
type BotName struct {
	Name string `json:"name"`
}

// BotShortDescription maps to Telegram Bot API type "BotShortDescription".
type BotShortDescription struct {
	ShortDescription string `json:"short_description"`
}

// BusinessBotRights maps to Telegram Bot API type "BusinessBotRights".
type BusinessBotRights struct {
	CanReply                   bool `json:"can_reply"`
	CanReadMessages            bool `json:"can_read_messages"`
	CanDeleteSentMessages      bool `json:"can_delete_sent_messages"`
	CanDeleteAllMessages       bool `json:"can_delete_all_messages"`
	CanEditName                bool `json:"can_edit_name"`
	CanEditBio                 bool `json:"can_edit_bio"`
	CanEditProfilePhoto        bool `json:"can_edit_profile_photo"`
	CanEditUsername            bool `json:"can_edit_username"`
	CanChangeGiftSettings      bool `json:"can_change_gift_settings"`
	CanViewGiftsAndStars       bool `json:"can_view_gifts_and_stars"`
	CanConvertGiftsToStars     bool `json:"can_convert_gifts_to_stars"`
	CanTransferAndUpgradeGifts bool `json:"can_transfer_and_upgrade_gifts"`
	CanTransferStars           bool `json:"can_transfer_stars"`
	CanManageStories           bool `json:"can_manage_stories"`
}

// BusinessConnection maps to Telegram Bot API type "BusinessConnection".
type BusinessConnection struct {
	ID         string             `json:"id"`
	User       *User              `json:"user"`
	UserChatID int64              `json:"user_chat_id"`
	Date       int64              `json:"date"`
	Rights     *BusinessBotRights `json:"rights"`
	IsEnabled  bool               `json:"is_enabled"`
}

// BusinessIntro maps to Telegram Bot API type "BusinessIntro".
type BusinessIntro struct {
	Title   string   `json:"title"`
	Message string   `json:"message"`
	Sticker *Sticker `json:"sticker"`
}

// BusinessLocation maps to Telegram Bot API type "BusinessLocation".
type BusinessLocation struct {
	Address  string    `json:"address"`
	Location *Location `json:"location"`
}

// BusinessMessagesDeleted maps to Telegram Bot API type "BusinessMessagesDeleted".
type BusinessMessagesDeleted struct {
	BusinessConnectionID string  `json:"business_connection_id"`
	Chat                 *Chat   `json:"chat"`
	MessageIds           []int64 `json:"message_ids"`
}

// BusinessOpeningHours maps to Telegram Bot API type "BusinessOpeningHours".
type BusinessOpeningHours struct {
	TimeZoneName string                         `json:"time_zone_name"`
	OpeningHours []BusinessOpeningHoursInterval `json:"opening_hours"`
}

// BusinessOpeningHoursInterval maps to Telegram Bot API type "BusinessOpeningHoursInterval".
type BusinessOpeningHoursInterval struct {
	OpeningMinute int64 `json:"opening_minute"`
	ClosingMinute int64 `json:"closing_minute"`
}

// CallbackGame maps to Telegram Bot API type "CallbackGame".
type CallbackGame struct {
}

// CallbackQuery maps to Telegram Bot API type "CallbackQuery".
type CallbackQuery struct {
	ID              string                   `json:"id"`
	From            *User                    `json:"from"`
	Message         MaybeInaccessibleMessage `json:"message"`
	InlineMessageID string                   `json:"inline_message_id"`
	ChatInstance    string                   `json:"chat_instance"`
	Data            string                   `json:"data"`
	GameShortName   string                   `json:"game_short_name"`
}

// Chat maps to Telegram Bot API type "Chat".
type Chat struct {
	ID               int64  `json:"id"`
	Type             string `json:"type"`
	Title            string `json:"title"`
	Username         string `json:"username"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	IsForum          bool   `json:"is_forum"`
	IsDirectMessages bool   `json:"is_direct_messages"`
}

// ChatAdministratorRights maps to Telegram Bot API type "ChatAdministratorRights".
type ChatAdministratorRights struct {
	IsAnonymous             bool `json:"is_anonymous"`
	CanManageChat           bool `json:"can_manage_chat"`
	CanDeleteMessages       bool `json:"can_delete_messages"`
	CanManageVideoChats     bool `json:"can_manage_video_chats"`
	CanRestrictMembers      bool `json:"can_restrict_members"`
	CanPromoteMembers       bool `json:"can_promote_members"`
	CanChangeInfo           bool `json:"can_change_info"`
	CanInviteUsers          bool `json:"can_invite_users"`
	CanPostStories          bool `json:"can_post_stories"`
	CanEditStories          bool `json:"can_edit_stories"`
	CanDeleteStories        bool `json:"can_delete_stories"`
	CanPostMessages         bool `json:"can_post_messages"`
	CanEditMessages         bool `json:"can_edit_messages"`
	CanPinMessages          bool `json:"can_pin_messages"`
	CanManageTopics         bool `json:"can_manage_topics"`
	CanManageDirectMessages bool `json:"can_manage_direct_messages"`
	CanManageTags           bool `json:"can_manage_tags"`
}

// ChatBackground maps to Telegram Bot API type "ChatBackground".
type ChatBackground struct {
	Type BackgroundType `json:"type"`
}

// ChatBoost maps to Telegram Bot API type "ChatBoost".
type ChatBoost struct {
	BoostID        string          `json:"boost_id"`
	AddDate        int64           `json:"add_date"`
	ExpirationDate int64           `json:"expiration_date"`
	Source         ChatBoostSource `json:"source"`
}

// ChatBoostAdded maps to Telegram Bot API type "ChatBoostAdded".
type ChatBoostAdded struct {
	BoostCount int64 `json:"boost_count"`
}

// ChatBoostRemoved maps to Telegram Bot API type "ChatBoostRemoved".
type ChatBoostRemoved struct {
	Chat       *Chat           `json:"chat"`
	BoostID    string          `json:"boost_id"`
	RemoveDate int64           `json:"remove_date"`
	Source     ChatBoostSource `json:"source"`
}

// ChatBoostSource is a union type in Telegram Bot API.
type ChatBoostSource interface {
	isChatBoostSource()
}

// ChatBoostSourceGiftCode maps to Telegram Bot API type "ChatBoostSourceGiftCode".
type ChatBoostSourceGiftCode struct {
	Source string `json:"source"`
	User   *User  `json:"user"`
}

func (*ChatBoostSourceGiftCode) isChatBoostSource() {}

// ChatBoostSourceGiveaway maps to Telegram Bot API type "ChatBoostSourceGiveaway".
type ChatBoostSourceGiveaway struct {
	Source            string `json:"source"`
	GiveawayMessageID int64  `json:"giveaway_message_id"`
	User              *User  `json:"user"`
	PrizeStarCount    int64  `json:"prize_star_count"`
	IsUnclaimed       bool   `json:"is_unclaimed"`
}

func (*ChatBoostSourceGiveaway) isChatBoostSource() {}

// ChatBoostSourcePremium maps to Telegram Bot API type "ChatBoostSourcePremium".
type ChatBoostSourcePremium struct {
	Source string `json:"source"`
	User   *User  `json:"user"`
}

func (*ChatBoostSourcePremium) isChatBoostSource() {}

// ChatBoostUpdated maps to Telegram Bot API type "ChatBoostUpdated".
type ChatBoostUpdated struct {
	Chat  *Chat      `json:"chat"`
	Boost *ChatBoost `json:"boost"`
}

// ChatFullInfo maps to Telegram Bot API type "ChatFullInfo".
type ChatFullInfo struct {
	ID                                 int64                 `json:"id"`
	Type                               string                `json:"type"`
	Title                              string                `json:"title"`
	Username                           string                `json:"username"`
	FirstName                          string                `json:"first_name"`
	LastName                           string                `json:"last_name"`
	IsForum                            bool                  `json:"is_forum"`
	IsDirectMessages                   bool                  `json:"is_direct_messages"`
	AccentColorID                      int64                 `json:"accent_color_id"`
	MaxReactionCount                   int64                 `json:"max_reaction_count"`
	Photo                              *ChatPhoto            `json:"photo"`
	ActiveUsernames                    []string              `json:"active_usernames"`
	Birthdate                          *Birthdate            `json:"birthdate"`
	BusinessIntro                      *BusinessIntro        `json:"business_intro"`
	BusinessLocation                   *BusinessLocation     `json:"business_location"`
	BusinessOpeningHours               *BusinessOpeningHours `json:"business_opening_hours"`
	PersonalChat                       *Chat                 `json:"personal_chat"`
	ParentChat                         *Chat                 `json:"parent_chat"`
	AvailableReactions                 []ReactionType        `json:"available_reactions"`
	BackgroundCustomEmojiID            string                `json:"background_custom_emoji_id"`
	ProfileAccentColorID               int64                 `json:"profile_accent_color_id"`
	ProfileBackgroundCustomEmojiID     string                `json:"profile_background_custom_emoji_id"`
	EmojiStatusCustomEmojiID           string                `json:"emoji_status_custom_emoji_id"`
	EmojiStatusExpirationDate          int64                 `json:"emoji_status_expiration_date"`
	Bio                                string                `json:"bio"`
	HasPrivateForwards                 bool                  `json:"has_private_forwards"`
	HasRestrictedVoiceAndVideoMessages bool                  `json:"has_restricted_voice_and_video_messages"`
	JoinToSendMessages                 bool                  `json:"join_to_send_messages"`
	JoinByRequest                      bool                  `json:"join_by_request"`
	Description                        string                `json:"description"`
	InviteLink                         string                `json:"invite_link"`
	PinnedMessage                      *Message              `json:"pinned_message"`
	Permissions                        *ChatPermissions      `json:"permissions"`
	AcceptedGiftTypes                  *AcceptedGiftTypes    `json:"accepted_gift_types"`
	CanSendPaidMedia                   bool                  `json:"can_send_paid_media"`
	SlowModeDelay                      int64                 `json:"slow_mode_delay"`
	UnrestrictBoostCount               int64                 `json:"unrestrict_boost_count"`
	MessageAutoDeleteTime              int64                 `json:"message_auto_delete_time"`
	HasAggressiveAntiSpamEnabled       bool                  `json:"has_aggressive_anti_spam_enabled"`
	HasHiddenMembers                   bool                  `json:"has_hidden_members"`
	HasProtectedContent                bool                  `json:"has_protected_content"`
	HasVisibleHistory                  bool                  `json:"has_visible_history"`
	StickerSetName                     string                `json:"sticker_set_name"`
	CanSetStickerSet                   bool                  `json:"can_set_sticker_set"`
	CustomEmojiStickerSetName          string                `json:"custom_emoji_sticker_set_name"`
	LinkedChatID                       int64                 `json:"linked_chat_id"`
	Location                           *ChatLocation         `json:"location"`
	Rating                             *UserRating           `json:"rating"`
	FirstProfileAudio                  *Audio                `json:"first_profile_audio"`
	UniqueGiftColors                   *UniqueGiftColors     `json:"unique_gift_colors"`
	PaidMessageStarCount               int64                 `json:"paid_message_star_count"`
}

// ChatInviteLink maps to Telegram Bot API type "ChatInviteLink".
type ChatInviteLink struct {
	InviteLink              string `json:"invite_link"`
	Creator                 *User  `json:"creator"`
	CreatesJoinRequest      bool   `json:"creates_join_request"`
	IsPrimary               bool   `json:"is_primary"`
	IsRevoked               bool   `json:"is_revoked"`
	Name                    string `json:"name"`
	ExpireDate              int64  `json:"expire_date"`
	MemberLimit             int64  `json:"member_limit"`
	PendingJoinRequestCount int64  `json:"pending_join_request_count"`
	SubscriptionPeriod      int64  `json:"subscription_period"`
	SubscriptionPrice       int64  `json:"subscription_price"`
}

// ChatJoinRequest maps to Telegram Bot API type "ChatJoinRequest".
type ChatJoinRequest struct {
	Chat       *Chat           `json:"chat"`
	From       *User           `json:"from"`
	UserChatID int64           `json:"user_chat_id"`
	Date       int64           `json:"date"`
	Bio        string          `json:"bio"`
	InviteLink *ChatInviteLink `json:"invite_link"`
}

// ChatLocation maps to Telegram Bot API type "ChatLocation".
type ChatLocation struct {
	Location *Location `json:"location"`
	Address  string    `json:"address"`
}

// ChatMember is a union type in Telegram Bot API.
type ChatMember interface {
	isChatMember()
}

// ChatMemberAdministrator maps to Telegram Bot API type "ChatMemberAdministrator".
type ChatMemberAdministrator struct {
	Status                  string `json:"status"`
	User                    *User  `json:"user"`
	CanBeEdited             bool   `json:"can_be_edited"`
	IsAnonymous             bool   `json:"is_anonymous"`
	CanManageChat           bool   `json:"can_manage_chat"`
	CanDeleteMessages       bool   `json:"can_delete_messages"`
	CanManageVideoChats     bool   `json:"can_manage_video_chats"`
	CanRestrictMembers      bool   `json:"can_restrict_members"`
	CanPromoteMembers       bool   `json:"can_promote_members"`
	CanChangeInfo           bool   `json:"can_change_info"`
	CanInviteUsers          bool   `json:"can_invite_users"`
	CanPostStories          bool   `json:"can_post_stories"`
	CanEditStories          bool   `json:"can_edit_stories"`
	CanDeleteStories        bool   `json:"can_delete_stories"`
	CanPostMessages         bool   `json:"can_post_messages"`
	CanEditMessages         bool   `json:"can_edit_messages"`
	CanPinMessages          bool   `json:"can_pin_messages"`
	CanManageTopics         bool   `json:"can_manage_topics"`
	CanManageDirectMessages bool   `json:"can_manage_direct_messages"`
	CanManageTags           bool   `json:"can_manage_tags"`
	CustomTitle             string `json:"custom_title"`
}

func (*ChatMemberAdministrator) isChatMember() {}

// ChatMemberBanned maps to Telegram Bot API type "ChatMemberBanned".
type ChatMemberBanned struct {
	Status    string `json:"status"`
	User      *User  `json:"user"`
	UntilDate int64  `json:"until_date"`
}

func (*ChatMemberBanned) isChatMember() {}

// ChatMemberLeft maps to Telegram Bot API type "ChatMemberLeft".
type ChatMemberLeft struct {
	Status string `json:"status"`
	User   *User  `json:"user"`
}

func (*ChatMemberLeft) isChatMember() {}

// ChatMemberMember maps to Telegram Bot API type "ChatMemberMember".
type ChatMemberMember struct {
	Status    string `json:"status"`
	Tag       string `json:"tag"`
	User      *User  `json:"user"`
	UntilDate int64  `json:"until_date"`
}

func (*ChatMemberMember) isChatMember() {}

// ChatMemberOwner maps to Telegram Bot API type "ChatMemberOwner".
type ChatMemberOwner struct {
	Status      string `json:"status"`
	User        *User  `json:"user"`
	IsAnonymous bool   `json:"is_anonymous"`
	CustomTitle string `json:"custom_title"`
}

func (*ChatMemberOwner) isChatMember() {}

// ChatMemberRestricted maps to Telegram Bot API type "ChatMemberRestricted".
type ChatMemberRestricted struct {
	Status                string `json:"status"`
	Tag                   string `json:"tag"`
	User                  *User  `json:"user"`
	IsMember              bool   `json:"is_member"`
	CanSendMessages       bool   `json:"can_send_messages"`
	CanSendAudios         bool   `json:"can_send_audios"`
	CanSendDocuments      bool   `json:"can_send_documents"`
	CanSendPhotos         bool   `json:"can_send_photos"`
	CanSendVideos         bool   `json:"can_send_videos"`
	CanSendVideoNotes     bool   `json:"can_send_video_notes"`
	CanSendVoiceNotes     bool   `json:"can_send_voice_notes"`
	CanSendPolls          bool   `json:"can_send_polls"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
	CanEditTag            bool   `json:"can_edit_tag"`
	CanChangeInfo         bool   `json:"can_change_info"`
	CanInviteUsers        bool   `json:"can_invite_users"`
	CanPinMessages        bool   `json:"can_pin_messages"`
	CanManageTopics       bool   `json:"can_manage_topics"`
	UntilDate             int64  `json:"until_date"`
}

func (*ChatMemberRestricted) isChatMember() {}

// ChatMemberUpdated maps to Telegram Bot API type "ChatMemberUpdated".
type ChatMemberUpdated struct {
	Chat                    *Chat           `json:"chat"`
	From                    *User           `json:"from"`
	Date                    int64           `json:"date"`
	OldChatMember           ChatMember      `json:"old_chat_member"`
	NewChatMember           ChatMember      `json:"new_chat_member"`
	InviteLink              *ChatInviteLink `json:"invite_link"`
	ViaJoinRequest          bool            `json:"via_join_request"`
	ViaChatFolderInviteLink bool            `json:"via_chat_folder_invite_link"`
}

// ChatOwnerChanged maps to Telegram Bot API type "ChatOwnerChanged".
type ChatOwnerChanged struct {
	NewOwner *User `json:"new_owner"`
}

// ChatOwnerLeft maps to Telegram Bot API type "ChatOwnerLeft".
type ChatOwnerLeft struct {
	NewOwner *User `json:"new_owner"`
}

// ChatPermissions maps to Telegram Bot API type "ChatPermissions".
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendAudios         bool `json:"can_send_audios"`
	CanSendDocuments      bool `json:"can_send_documents"`
	CanSendPhotos         bool `json:"can_send_photos"`
	CanSendVideos         bool `json:"can_send_videos"`
	CanSendVideoNotes     bool `json:"can_send_video_notes"`
	CanSendVoiceNotes     bool `json:"can_send_voice_notes"`
	CanSendPolls          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanEditTag            bool `json:"can_edit_tag"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`
	CanManageTopics       bool `json:"can_manage_topics"`
}

// ChatPhoto maps to Telegram Bot API type "ChatPhoto".
type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

// ChatShared maps to Telegram Bot API type "ChatShared".
type ChatShared struct {
	RequestID int64       `json:"request_id"`
	ChatID    int64       `json:"chat_id"`
	Title     string      `json:"title"`
	Username  string      `json:"username"`
	Photo     []PhotoSize `json:"photo"`
}

// Checklist maps to Telegram Bot API type "Checklist".
type Checklist struct {
	Title                    string          `json:"title"`
	TitleEntities            []MessageEntity `json:"title_entities"`
	Tasks                    []ChecklistTask `json:"tasks"`
	OthersCanAddTasks        bool            `json:"others_can_add_tasks"`
	OthersCanMarkTasksAsDone bool            `json:"others_can_mark_tasks_as_done"`
}

// ChecklistTask maps to Telegram Bot API type "ChecklistTask".
type ChecklistTask struct {
	ID              int64           `json:"id"`
	Text            string          `json:"text"`
	TextEntities    []MessageEntity `json:"text_entities"`
	CompletedByUser *User           `json:"completed_by_user"`
	CompletedByChat *Chat           `json:"completed_by_chat"`
	CompletionDate  int64           `json:"completion_date"`
}

// ChecklistTasksAdded maps to Telegram Bot API type "ChecklistTasksAdded".
type ChecklistTasksAdded struct {
	ChecklistMessage *Message        `json:"checklist_message"`
	Tasks            []ChecklistTask `json:"tasks"`
}

// ChecklistTasksDone maps to Telegram Bot API type "ChecklistTasksDone".
type ChecklistTasksDone struct {
	ChecklistMessage       *Message `json:"checklist_message"`
	MarkedAsDoneTaskIds    []int64  `json:"marked_as_done_task_ids"`
	MarkedAsNotDoneTaskIds []int64  `json:"marked_as_not_done_task_ids"`
}

// ChosenInlineResult maps to Telegram Bot API type "ChosenInlineResult".
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location"`
	InlineMessageID string    `json:"inline_message_id"`
	Query           string    `json:"query"`
}

// Contact maps to Telegram Bot API type "Contact".
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserID      int64  `json:"user_id"`
	Vcard       string `json:"vcard"`
}

// CopyTextButton maps to Telegram Bot API type "CopyTextButton".
type CopyTextButton struct {
	Text string `json:"text"`
}

// Dice maps to Telegram Bot API type "Dice".
type Dice struct {
	Emoji string `json:"emoji"`
	Value int64  `json:"value"`
}

// DirectMessagePriceChanged maps to Telegram Bot API type "DirectMessagePriceChanged".
type DirectMessagePriceChanged struct {
	AreDirectMessagesEnabled bool  `json:"are_direct_messages_enabled"`
	DirectMessageStarCount   int64 `json:"direct_message_star_count"`
}

// DirectMessagesTopic maps to Telegram Bot API type "DirectMessagesTopic".
type DirectMessagesTopic struct {
	TopicID int64 `json:"topic_id"`
	User    *User `json:"user"`
}

// Document maps to Telegram Bot API type "Document".
type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumbnail    *PhotoSize `json:"thumbnail"`
	FileName     string     `json:"file_name"`
	MimeType     string     `json:"mime_type"`
	FileSize     int64      `json:"file_size"`
}

// EncryptedCredentials maps to Telegram Bot API type "EncryptedCredentials".
type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

// EncryptedPassportElement maps to Telegram Bot API type "EncryptedPassportElement".
type EncryptedPassportElement struct {
	Type        string         `json:"type"`
	Data        string         `json:"data"`
	PhoneNumber string         `json:"phone_number"`
	Email       string         `json:"email"`
	Files       []PassportFile `json:"files"`
	FrontSide   *PassportFile  `json:"front_side"`
	ReverseSide *PassportFile  `json:"reverse_side"`
	Selfie      *PassportFile  `json:"selfie"`
	Translation []PassportFile `json:"translation"`
	Hash        string         `json:"hash"`
}

// ExternalReplyInfo maps to Telegram Bot API type "ExternalReplyInfo".
type ExternalReplyInfo struct {
	Origin             MessageOrigin       `json:"origin"`
	Chat               *Chat               `json:"chat"`
	MessageID          int64               `json:"message_id"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options"`
	Animation          *Animation          `json:"animation"`
	Audio              *Audio              `json:"audio"`
	Document           *Document           `json:"document"`
	PaidMedia          *PaidMediaInfo      `json:"paid_media"`
	Photo              []PhotoSize         `json:"photo"`
	Sticker            *Sticker            `json:"sticker"`
	Story              *Story              `json:"story"`
	Video              *Video              `json:"video"`
	VideoNote          *VideoNote          `json:"video_note"`
	Voice              *Voice              `json:"voice"`
	HasMediaSpoiler    bool                `json:"has_media_spoiler"`
	Checklist          *Checklist          `json:"checklist"`
	Contact            *Contact            `json:"contact"`
	Dice               *Dice               `json:"dice"`
	Game               *Game               `json:"game"`
	Giveaway           *Giveaway           `json:"giveaway"`
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners"`
	Invoice            *Invoice            `json:"invoice"`
	Location           *Location           `json:"location"`
	Poll               *Poll               `json:"poll"`
	Venue              *Venue              `json:"venue"`
}

// File maps to Telegram Bot API type "File".
type File struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	FilePath     string `json:"file_path"`
}

// ForceReply maps to Telegram Bot API type "ForceReply".
type ForceReply struct {
	ForceReply            bool   `json:"force_reply"`
	InputFieldPlaceholder string `json:"input_field_placeholder"`
	Selective             bool   `json:"selective"`
}

// ForumTopic maps to Telegram Bot API type "ForumTopic".
type ForumTopic struct {
	MessageThreadID   int64  `json:"message_thread_id"`
	Name              string `json:"name"`
	IconColor         int64  `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id"`
	IsNameImplicit    bool   `json:"is_name_implicit"`
}

// ForumTopicClosed maps to Telegram Bot API type "ForumTopicClosed".
type ForumTopicClosed struct {
}

// ForumTopicCreated maps to Telegram Bot API type "ForumTopicCreated".
type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int64  `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id"`
	IsNameImplicit    bool   `json:"is_name_implicit"`
}

// ForumTopicEdited maps to Telegram Bot API type "ForumTopicEdited".
type ForumTopicEdited struct {
	Name              string `json:"name"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id"`
}

// ForumTopicReopened maps to Telegram Bot API type "ForumTopicReopened".
type ForumTopicReopened struct {
}

// Game maps to Telegram Bot API type "Game".
type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities"`
	Animation    *Animation      `json:"animation"`
}

// GameHighScore maps to Telegram Bot API type "GameHighScore".
type GameHighScore struct {
	Position int64 `json:"position"`
	User     *User `json:"user"`
	Score    int64 `json:"score"`
}

// GeneralForumTopicHidden maps to Telegram Bot API type "GeneralForumTopicHidden".
type GeneralForumTopicHidden struct {
}

// GeneralForumTopicUnhidden maps to Telegram Bot API type "GeneralForumTopicUnhidden".
type GeneralForumTopicUnhidden struct {
}

// Gift maps to Telegram Bot API type "Gift".
type Gift struct {
	ID                     string          `json:"id"`
	Sticker                *Sticker        `json:"sticker"`
	StarCount              int64           `json:"star_count"`
	UpgradeStarCount       int64           `json:"upgrade_star_count"`
	IsPremium              bool            `json:"is_premium"`
	HasColors              bool            `json:"has_colors"`
	TotalCount             int64           `json:"total_count"`
	RemainingCount         int64           `json:"remaining_count"`
	PersonalTotalCount     int64           `json:"personal_total_count"`
	PersonalRemainingCount int64           `json:"personal_remaining_count"`
	Background             *GiftBackground `json:"background"`
	UniqueGiftVariantCount int64           `json:"unique_gift_variant_count"`
	PublisherChat          *Chat           `json:"publisher_chat"`
}

// GiftBackground maps to Telegram Bot API type "GiftBackground".
type GiftBackground struct {
	CenterColor int64 `json:"center_color"`
	EdgeColor   int64 `json:"edge_color"`
	TextColor   int64 `json:"text_color"`
}

// GiftInfo maps to Telegram Bot API type "GiftInfo".
type GiftInfo struct {
	Gift                    *Gift           `json:"gift"`
	OwnedGiftID             string          `json:"owned_gift_id"`
	ConvertStarCount        int64           `json:"convert_star_count"`
	PrepaidUpgradeStarCount int64           `json:"prepaid_upgrade_star_count"`
	IsUpgradeSeparate       bool            `json:"is_upgrade_separate"`
	CanBeUpgraded           bool            `json:"can_be_upgraded"`
	Text                    string          `json:"text"`
	Entities                []MessageEntity `json:"entities"`
	IsPrivate               bool            `json:"is_private"`
	UniqueGiftNumber        int64           `json:"unique_gift_number"`
}

// Gifts maps to Telegram Bot API type "Gifts".
type Gifts struct {
	Gifts []Gift `json:"gifts"`
}

// Giveaway maps to Telegram Bot API type "Giveaway".
type Giveaway struct {
	Chats                         []Chat   `json:"chats"`
	WinnersSelectionDate          int64    `json:"winners_selection_date"`
	WinnerCount                   int64    `json:"winner_count"`
	OnlyNewMembers                bool     `json:"only_new_members"`
	HasPublicWinners              bool     `json:"has_public_winners"`
	PrizeDescription              string   `json:"prize_description"`
	CountryCodes                  []string `json:"country_codes"`
	PrizeStarCount                int64    `json:"prize_star_count"`
	PremiumSubscriptionMonthCount int64    `json:"premium_subscription_month_count"`
}

// GiveawayCompleted maps to Telegram Bot API type "GiveawayCompleted".
type GiveawayCompleted struct {
	WinnerCount         int64    `json:"winner_count"`
	UnclaimedPrizeCount int64    `json:"unclaimed_prize_count"`
	GiveawayMessage     *Message `json:"giveaway_message"`
	IsStarGiveaway      bool     `json:"is_star_giveaway"`
}

// GiveawayCreated maps to Telegram Bot API type "GiveawayCreated".
type GiveawayCreated struct {
	PrizeStarCount int64 `json:"prize_star_count"`
}

// GiveawayWinners maps to Telegram Bot API type "GiveawayWinners".
type GiveawayWinners struct {
	Chat                          *Chat  `json:"chat"`
	GiveawayMessageID             int64  `json:"giveaway_message_id"`
	WinnersSelectionDate          int64  `json:"winners_selection_date"`
	WinnerCount                   int64  `json:"winner_count"`
	Winners                       []User `json:"winners"`
	AdditionalChatCount           int64  `json:"additional_chat_count"`
	PrizeStarCount                int64  `json:"prize_star_count"`
	PremiumSubscriptionMonthCount int64  `json:"premium_subscription_month_count"`
	UnclaimedPrizeCount           int64  `json:"unclaimed_prize_count"`
	OnlyNewMembers                bool   `json:"only_new_members"`
	WasRefunded                   bool   `json:"was_refunded"`
	PrizeDescription              string `json:"prize_description"`
}

// InaccessibleMessage maps to Telegram Bot API type "InaccessibleMessage".
type InaccessibleMessage struct {
	Chat      *Chat `json:"chat"`
	MessageID int64 `json:"message_id"`
	Date      int64 `json:"date"`
}

func (*InaccessibleMessage) isMaybeInaccessibleMessage() {}

// InlineKeyboardButton maps to Telegram Bot API type "InlineKeyboardButton".
type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`
	IconCustomEmojiID            string                       `json:"icon_custom_emoji_id"`
	Style                        string                       `json:"style"`
	URL                          string                       `json:"url"`
	CallbackData                 string                       `json:"callback_data"`
	WebApp                       *WebAppInfo                  `json:"web_app"`
	LoginURL                     *LoginUrl                    `json:"login_url"`
	SwitchInlineQuery            string                       `json:"switch_inline_query"`
	SwitchInlineQueryCurrentChat string                       `json:"switch_inline_query_current_chat"`
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat"`
	CopyText                     *CopyTextButton              `json:"copy_text"`
	CallbackGame                 *CallbackGame                `json:"callback_game"`
	Pay                          bool                         `json:"pay"`
}

// InlineKeyboardMarkup maps to Telegram Bot API type "InlineKeyboardMarkup".
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineQuery maps to Telegram Bot API type "InlineQuery".
type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	ChatType string    `json:"chat_type"`
	Location *Location `json:"location"`
}

// InlineQueryResult is a union type in Telegram Bot API.
type InlineQueryResult interface {
	isInlineQueryResult()
}

// InlineQueryResultArticle maps to Telegram Bot API type "InlineQueryResultArticle".
type InlineQueryResultArticle struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	URL                 string                `json:"url"`
	Description         string                `json:"description"`
	ThumbnailURL        string                `json:"thumbnail_url"`
	ThumbnailWidth      int64                 `json:"thumbnail_width"`
	ThumbnailHeight     int64                 `json:"thumbnail_height"`
}

func (*InlineQueryResultArticle) isInlineQueryResult() {}

// InlineQueryResultAudio maps to Telegram Bot API type "InlineQueryResultAudio".
type InlineQueryResultAudio struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	AudioURL            string                `json:"audio_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption"`
	ParseMode           string                `json:"parse_mode"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	Performer           string                `json:"performer"`
	AudioDuration       int64                 `json:"audio_duration"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
}

func (*InlineQueryResultAudio) isInlineQueryResult() {}

// InlineQueryResultCachedAudio maps to Telegram Bot API type "InlineQueryResultCachedAudio".
type InlineQueryResultCachedAudio struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	AudioFileID         string                `json:"audio_file_id"`
	Caption             string                `json:"caption"`
	ParseMode           string                `json:"parse_mode"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
}

func (*InlineQueryResultCachedAudio) isInlineQueryResult() {}

// InlineQueryResultCachedDocument maps to Telegram Bot API type "InlineQueryResultCachedDocument".
type InlineQueryResultCachedDocument struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	DocumentFileID      string                `json:"document_file_id"`
	Description         string                `json:"description"`
	Caption             string                `json:"caption"`
	ParseMode           string                `json:"parse_mode"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
}

func (*InlineQueryResultCachedDocument) isInlineQueryResult() {}

// InlineQueryResultCachedGif maps to Telegram Bot API type "InlineQueryResultCachedGif".
type InlineQueryResultCachedGif struct {
	Type                  string                `json:"type"`
	ID                    string                `json:"id"`
	GifFileID             string                `json:"gif_file_id"`
	Title                 string                `json:"title"`
	Caption               string                `json:"caption"`
	ParseMode             string                `json:"parse_mode"`
	CaptionEntities       []MessageEntity       `json:"caption_entities"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent   InputMessageContent   `json:"input_message_content"`
}

func (*InlineQueryResultCachedGif) isInlineQueryResult() {}

// InlineQueryResultCachedMpeg4Gif maps to Telegram Bot API type "InlineQueryResultCachedMpeg4Gif".
type InlineQueryResultCachedMpeg4Gif struct {
	Type                  string                `json:"type"`
	ID                    string                `json:"id"`
	Mpeg4FileID           string                `json:"mpeg4_file_id"`
	Title                 string                `json:"title"`
	Caption               string                `json:"caption"`
	ParseMode             string                `json:"parse_mode"`
	CaptionEntities       []MessageEntity       `json:"caption_entities"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent   InputMessageContent   `json:"input_message_content"`
}

func (*InlineQueryResultCachedMpeg4Gif) isInlineQueryResult() {}

// InlineQueryResultCachedPhoto maps to Telegram Bot API type "InlineQueryResultCachedPhoto".
type InlineQueryResultCachedPhoto struct {
	Type                  string                `json:"type"`
	ID                    string                `json:"id"`
	PhotoFileID           string                `json:"photo_file_id"`
	Title                 string                `json:"title"`
	Description           string                `json:"description"`
	Caption               string                `json:"caption"`
	ParseMode             string                `json:"parse_mode"`
	CaptionEntities       []MessageEntity       `json:"caption_entities"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent   InputMessageContent   `json:"input_message_content"`
}

func (*InlineQueryResultCachedPhoto) isInlineQueryResult() {}

// InlineQueryResultCachedSticker maps to Telegram Bot API type "InlineQueryResultCachedSticker".
type InlineQueryResultCachedSticker struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	StickerFileID       string                `json:"sticker_file_id"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
}

func (*InlineQueryResultCachedSticker) isInlineQueryResult() {}

// InlineQueryResultCachedVideo maps to Telegram Bot API type "InlineQueryResultCachedVideo".
type InlineQueryResultCachedVideo struct {
	Type                  string                `json:"type"`
	ID                    string                `json:"id"`
	VideoFileID           string                `json:"video_file_id"`
	Title                 string                `json:"title"`
	Description           string                `json:"description"`
	Caption               string                `json:"caption"`
	ParseMode             string                `json:"parse_mode"`
	CaptionEntities       []MessageEntity       `json:"caption_entities"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent   InputMessageContent   `json:"input_message_content"`
}

func (*InlineQueryResultCachedVideo) isInlineQueryResult() {}

// InlineQueryResultCachedVoice maps to Telegram Bot API type "InlineQueryResultCachedVoice".
type InlineQueryResultCachedVoice struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VoiceFileID         string                `json:"voice_file_id"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption"`
	ParseMode           string                `json:"parse_mode"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
}

func (*InlineQueryResultCachedVoice) isInlineQueryResult() {}

// InlineQueryResultContact maps to Telegram Bot API type "InlineQueryResultContact".
type InlineQueryResultContact struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	PhoneNumber         string                `json:"phone_number"`
	FirstName           string                `json:"first_name"`
	LastName            string                `json:"last_name"`
	Vcard               string                `json:"vcard"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
	ThumbnailURL        string                `json:"thumbnail_url"`
	ThumbnailWidth      int64                 `json:"thumbnail_width"`
	ThumbnailHeight     int64                 `json:"thumbnail_height"`
}

func (*InlineQueryResultContact) isInlineQueryResult() {}

// InlineQueryResultDocument maps to Telegram Bot API type "InlineQueryResultDocument".
type InlineQueryResultDocument struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption"`
	ParseMode           string                `json:"parse_mode"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	DocumentURL         string                `json:"document_url"`
	MimeType            string                `json:"mime_type"`
	Description         string                `json:"description"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
	ThumbnailURL        string                `json:"thumbnail_url"`
	ThumbnailWidth      int64                 `json:"thumbnail_width"`
	ThumbnailHeight     int64                 `json:"thumbnail_height"`
}

func (*InlineQueryResultDocument) isInlineQueryResult() {}

// InlineQueryResultGame maps to Telegram Bot API type "InlineQueryResultGame".
type InlineQueryResultGame struct {
	Type          string                `json:"type"`
	ID            string                `json:"id"`
	GameShortName string                `json:"game_short_name"`
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup"`
}

func (*InlineQueryResultGame) isInlineQueryResult() {}

// InlineQueryResultGif maps to Telegram Bot API type "InlineQueryResultGif".
type InlineQueryResultGif struct {
	Type                  string                `json:"type"`
	ID                    string                `json:"id"`
	GifURL                string                `json:"gif_url"`
	GifWidth              int64                 `json:"gif_width"`
	GifHeight             int64                 `json:"gif_height"`
	GifDuration           int64                 `json:"gif_duration"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	ThumbnailMimeType     string                `json:"thumbnail_mime_type"`
	Title                 string                `json:"title"`
	Caption               string                `json:"caption"`
	ParseMode             string                `json:"parse_mode"`
	CaptionEntities       []MessageEntity       `json:"caption_entities"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent   InputMessageContent   `json:"input_message_content"`
}

func (*InlineQueryResultGif) isInlineQueryResult() {}

// InlineQueryResultLocation maps to Telegram Bot API type "InlineQueryResultLocation".
type InlineQueryResultLocation struct {
	Type                 string                `json:"type"`
	ID                   string                `json:"id"`
	Latitude             float64               `json:"latitude"`
	Longitude            float64               `json:"longitude"`
	Title                string                `json:"title"`
	HorizontalAccuracy   float64               `json:"horizontal_accuracy"`
	LivePeriod           int64                 `json:"live_period"`
	Heading              int64                 `json:"heading"`
	ProximityAlertRadius int64                 `json:"proximity_alert_radius"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent  InputMessageContent   `json:"input_message_content"`
	ThumbnailURL         string                `json:"thumbnail_url"`
	ThumbnailWidth       int64                 `json:"thumbnail_width"`
	ThumbnailHeight      int64                 `json:"thumbnail_height"`
}

func (*InlineQueryResultLocation) isInlineQueryResult() {}

// InlineQueryResultMpeg4Gif maps to Telegram Bot API type "InlineQueryResultMpeg4Gif".
type InlineQueryResultMpeg4Gif struct {
	Type                  string                `json:"type"`
	ID                    string                `json:"id"`
	Mpeg4URL              string                `json:"mpeg4_url"`
	Mpeg4Width            int64                 `json:"mpeg4_width"`
	Mpeg4Height           int64                 `json:"mpeg4_height"`
	Mpeg4Duration         int64                 `json:"mpeg4_duration"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	ThumbnailMimeType     string                `json:"thumbnail_mime_type"`
	Title                 string                `json:"title"`
	Caption               string                `json:"caption"`
	ParseMode             string                `json:"parse_mode"`
	CaptionEntities       []MessageEntity       `json:"caption_entities"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent   InputMessageContent   `json:"input_message_content"`
}

func (*InlineQueryResultMpeg4Gif) isInlineQueryResult() {}

// InlineQueryResultPhoto maps to Telegram Bot API type "InlineQueryResultPhoto".
type InlineQueryResultPhoto struct {
	Type                  string                `json:"type"`
	ID                    string                `json:"id"`
	PhotoURL              string                `json:"photo_url"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	PhotoWidth            int64                 `json:"photo_width"`
	PhotoHeight           int64                 `json:"photo_height"`
	Title                 string                `json:"title"`
	Description           string                `json:"description"`
	Caption               string                `json:"caption"`
	ParseMode             string                `json:"parse_mode"`
	CaptionEntities       []MessageEntity       `json:"caption_entities"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent   InputMessageContent   `json:"input_message_content"`
}

func (*InlineQueryResultPhoto) isInlineQueryResult() {}

// InlineQueryResultVenue maps to Telegram Bot API type "InlineQueryResultVenue".
type InlineQueryResultVenue struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	Title               string                `json:"title"`
	Address             string                `json:"address"`
	FoursquareID        string                `json:"foursquare_id"`
	FoursquareType      string                `json:"foursquare_type"`
	GooglePlaceID       string                `json:"google_place_id"`
	GooglePlaceType     string                `json:"google_place_type"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
	ThumbnailURL        string                `json:"thumbnail_url"`
	ThumbnailWidth      int64                 `json:"thumbnail_width"`
	ThumbnailHeight     int64                 `json:"thumbnail_height"`
}

func (*InlineQueryResultVenue) isInlineQueryResult() {}

// InlineQueryResultVideo maps to Telegram Bot API type "InlineQueryResultVideo".
type InlineQueryResultVideo struct {
	Type                  string                `json:"type"`
	ID                    string                `json:"id"`
	VideoURL              string                `json:"video_url"`
	MimeType              string                `json:"mime_type"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	Title                 string                `json:"title"`
	Caption               string                `json:"caption"`
	ParseMode             string                `json:"parse_mode"`
	CaptionEntities       []MessageEntity       `json:"caption_entities"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media"`
	VideoWidth            int64                 `json:"video_width"`
	VideoHeight           int64                 `json:"video_height"`
	VideoDuration         int64                 `json:"video_duration"`
	Description           string                `json:"description"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent   InputMessageContent   `json:"input_message_content"`
}

func (*InlineQueryResultVideo) isInlineQueryResult() {}

// InlineQueryResultVoice maps to Telegram Bot API type "InlineQueryResultVoice".
type InlineQueryResultVoice struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VoiceURL            string                `json:"voice_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption"`
	ParseMode           string                `json:"parse_mode"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	VoiceDuration       int64                 `json:"voice_duration"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
}

func (*InlineQueryResultVoice) isInlineQueryResult() {}

// InlineQueryResultsButton maps to Telegram Bot API type "InlineQueryResultsButton".
type InlineQueryResultsButton struct {
	Text           string      `json:"text"`
	WebApp         *WebAppInfo `json:"web_app"`
	StartParameter string      `json:"start_parameter"`
}

// InputChecklist maps to Telegram Bot API type "InputChecklist".
type InputChecklist struct {
	Title                    string               `json:"title"`
	ParseMode                string               `json:"parse_mode"`
	TitleEntities            []MessageEntity      `json:"title_entities"`
	Tasks                    []InputChecklistTask `json:"tasks"`
	OthersCanAddTasks        bool                 `json:"others_can_add_tasks"`
	OthersCanMarkTasksAsDone bool                 `json:"others_can_mark_tasks_as_done"`
}

// InputChecklistTask maps to Telegram Bot API type "InputChecklistTask".
type InputChecklistTask struct {
	ID           int64           `json:"id"`
	Text         string          `json:"text"`
	ParseMode    string          `json:"parse_mode"`
	TextEntities []MessageEntity `json:"text_entities"`
}

// InputContactMessageContent maps to Telegram Bot API type "InputContactMessageContent".
type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Vcard       string `json:"vcard"`
}

func (*InputContactMessageContent) isInputMessageContent() {}

// InputInvoiceMessageContent maps to Telegram Bot API type "InputInvoiceMessageContent".
type InputInvoiceMessageContent struct {
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             string         `json:"provider_token"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	MaxTipAmount              int64          `json:"max_tip_amount"`
	SuggestedTipAmounts       []int64        `json:"suggested_tip_amounts"`
	ProviderData              string         `json:"provider_data"`
	PhotoURL                  string         `json:"photo_url"`
	PhotoSize                 int64          `json:"photo_size"`
	PhotoWidth                int64          `json:"photo_width"`
	PhotoHeight               int64          `json:"photo_height"`
	NeedName                  bool           `json:"need_name"`
	NeedPhoneNumber           bool           `json:"need_phone_number"`
	NeedEmail                 bool           `json:"need_email"`
	NeedShippingAddress       bool           `json:"need_shipping_address"`
	SendPhoneNumberToProvider bool           `json:"send_phone_number_to_provider"`
	SendEmailToProvider       bool           `json:"send_email_to_provider"`
	IsFlexible                bool           `json:"is_flexible"`
}

func (*InputInvoiceMessageContent) isInputMessageContent() {}

// InputLocationMessageContent maps to Telegram Bot API type "InputLocationMessageContent".
type InputLocationMessageContent struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy"`
	LivePeriod           int64   `json:"live_period"`
	Heading              int64   `json:"heading"`
	ProximityAlertRadius int64   `json:"proximity_alert_radius"`
}

func (*InputLocationMessageContent) isInputMessageContent() {}

// InputMedia is a union type in Telegram Bot API.
type InputMedia interface {
	isInputMedia()
}

// InputMediaAnimation maps to Telegram Bot API type "InputMediaAnimation".
type InputMediaAnimation struct {
	Type                  string          `json:"type"`
	Media                 string          `json:"media"`
	Thumbnail             string          `json:"thumbnail"`
	Caption               string          `json:"caption"`
	ParseMode             string          `json:"parse_mode"`
	CaptionEntities       []MessageEntity `json:"caption_entities"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media"`
	Width                 int64           `json:"width"`
	Height                int64           `json:"height"`
	Duration              int64           `json:"duration"`
	HasSpoiler            bool            `json:"has_spoiler"`
}

func (*InputMediaAnimation) isInputMedia() {}

// InputMediaAudio maps to Telegram Bot API type "InputMediaAudio".
type InputMediaAudio struct {
	Type            string          `json:"type"`
	Media           string          `json:"media"`
	Thumbnail       string          `json:"thumbnail"`
	Caption         string          `json:"caption"`
	ParseMode       string          `json:"parse_mode"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
	Duration        int64           `json:"duration"`
	Performer       string          `json:"performer"`
	Title           string          `json:"title"`
}

func (*InputMediaAudio) isInputMedia() {}

// InputMediaDocument maps to Telegram Bot API type "InputMediaDocument".
type InputMediaDocument struct {
	Type                        string          `json:"type"`
	Media                       string          `json:"media"`
	Thumbnail                   string          `json:"thumbnail"`
	Caption                     string          `json:"caption"`
	ParseMode                   string          `json:"parse_mode"`
	CaptionEntities             []MessageEntity `json:"caption_entities"`
	DisableContentTypeDetection bool            `json:"disable_content_type_detection"`
}

func (*InputMediaDocument) isInputMedia() {}

// InputMediaPhoto maps to Telegram Bot API type "InputMediaPhoto".
type InputMediaPhoto struct {
	Type                  string          `json:"type"`
	Media                 string          `json:"media"`
	Caption               string          `json:"caption"`
	ParseMode             string          `json:"parse_mode"`
	CaptionEntities       []MessageEntity `json:"caption_entities"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media"`
	HasSpoiler            bool            `json:"has_spoiler"`
}

func (*InputMediaPhoto) isInputMedia() {}

// InputMediaVideo maps to Telegram Bot API type "InputMediaVideo".
type InputMediaVideo struct {
	Type                  string          `json:"type"`
	Media                 string          `json:"media"`
	Thumbnail             string          `json:"thumbnail"`
	Cover                 string          `json:"cover"`
	StartTimestamp        int64           `json:"start_timestamp"`
	Caption               string          `json:"caption"`
	ParseMode             string          `json:"parse_mode"`
	CaptionEntities       []MessageEntity `json:"caption_entities"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media"`
	Width                 int64           `json:"width"`
	Height                int64           `json:"height"`
	Duration              int64           `json:"duration"`
	SupportsStreaming     bool            `json:"supports_streaming"`
	HasSpoiler            bool            `json:"has_spoiler"`
}

func (*InputMediaVideo) isInputMedia() {}

// InputMessageContent is a union type in Telegram Bot API.
type InputMessageContent interface {
	isInputMessageContent()
}

// InputPaidMedia is a union type in Telegram Bot API.
type InputPaidMedia interface {
	isInputPaidMedia()
}

// InputPaidMediaPhoto maps to Telegram Bot API type "InputPaidMediaPhoto".
type InputPaidMediaPhoto struct {
	Type  string `json:"type"`
	Media string `json:"media"`
}

func (*InputPaidMediaPhoto) isInputPaidMedia() {}

// InputPaidMediaVideo maps to Telegram Bot API type "InputPaidMediaVideo".
type InputPaidMediaVideo struct {
	Type              string `json:"type"`
	Media             string `json:"media"`
	Thumbnail         string `json:"thumbnail"`
	Cover             string `json:"cover"`
	StartTimestamp    int64  `json:"start_timestamp"`
	Width             int64  `json:"width"`
	Height            int64  `json:"height"`
	Duration          int64  `json:"duration"`
	SupportsStreaming bool   `json:"supports_streaming"`
}

func (*InputPaidMediaVideo) isInputPaidMedia() {}

// InputPollOption maps to Telegram Bot API type "InputPollOption".
type InputPollOption struct {
	Text          string          `json:"text"`
	TextParseMode string          `json:"text_parse_mode"`
	TextEntities  []MessageEntity `json:"text_entities"`
}

// InputProfilePhoto is a union type in Telegram Bot API.
type InputProfilePhoto interface {
	isInputProfilePhoto()
}

// InputProfilePhotoAnimated maps to Telegram Bot API type "InputProfilePhotoAnimated".
type InputProfilePhotoAnimated struct {
	Type               string  `json:"type"`
	Animation          string  `json:"animation"`
	MainFrameTimestamp float64 `json:"main_frame_timestamp"`
}

func (*InputProfilePhotoAnimated) isInputProfilePhoto() {}

// InputProfilePhotoStatic maps to Telegram Bot API type "InputProfilePhotoStatic".
type InputProfilePhotoStatic struct {
	Type  string `json:"type"`
	Photo string `json:"photo"`
}

func (*InputProfilePhotoStatic) isInputProfilePhoto() {}

// InputSticker maps to Telegram Bot API type "InputSticker".
type InputSticker struct {
	Sticker      string        `json:"sticker"`
	Format       string        `json:"format"`
	EmojiList    []string      `json:"emoji_list"`
	MaskPosition *MaskPosition `json:"mask_position"`
	Keywords     []string      `json:"keywords"`
}

// InputStoryContent is a union type in Telegram Bot API.
type InputStoryContent interface {
	isInputStoryContent()
}

// InputStoryContentPhoto maps to Telegram Bot API type "InputStoryContentPhoto".
type InputStoryContentPhoto struct {
	Type  string `json:"type"`
	Photo string `json:"photo"`
}

func (*InputStoryContentPhoto) isInputStoryContent() {}

// InputStoryContentVideo maps to Telegram Bot API type "InputStoryContentVideo".
type InputStoryContentVideo struct {
	Type                string  `json:"type"`
	Video               string  `json:"video"`
	Duration            float64 `json:"duration"`
	CoverFrameTimestamp float64 `json:"cover_frame_timestamp"`
	IsAnimation         bool    `json:"is_animation"`
}

func (*InputStoryContentVideo) isInputStoryContent() {}

// InputTextMessageContent maps to Telegram Bot API type "InputTextMessageContent".
type InputTextMessageContent struct {
	MessageText        string              `json:"message_text"`
	ParseMode          string              `json:"parse_mode"`
	Entities           []MessageEntity     `json:"entities"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options"`
}

func (*InputTextMessageContent) isInputMessageContent() {}

// InputVenueMessageContent maps to Telegram Bot API type "InputVenueMessageContent".
type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareID    string  `json:"foursquare_id"`
	FoursquareType  string  `json:"foursquare_type"`
	GooglePlaceID   string  `json:"google_place_id"`
	GooglePlaceType string  `json:"google_place_type"`
}

func (*InputVenueMessageContent) isInputMessageContent() {}

// Invoice maps to Telegram Bot API type "Invoice".
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int64  `json:"total_amount"`
}

// KeyboardButton maps to Telegram Bot API type "KeyboardButton".
type KeyboardButton struct {
	Text              string                      `json:"text"`
	IconCustomEmojiID string                      `json:"icon_custom_emoji_id"`
	Style             string                      `json:"style"`
	RequestUsers      *KeyboardButtonRequestUsers `json:"request_users"`
	RequestChat       *KeyboardButtonRequestChat  `json:"request_chat"`
	RequestContact    bool                        `json:"request_contact"`
	RequestLocation   bool                        `json:"request_location"`
	RequestPoll       *KeyboardButtonPollType     `json:"request_poll"`
	WebApp            *WebAppInfo                 `json:"web_app"`
}

// KeyboardButtonPollType maps to Telegram Bot API type "KeyboardButtonPollType".
type KeyboardButtonPollType struct {
	Type string `json:"type"`
}

// KeyboardButtonRequestChat maps to Telegram Bot API type "KeyboardButtonRequestChat".
type KeyboardButtonRequestChat struct {
	RequestID               int64                    `json:"request_id"`
	ChatIsChannel           bool                     `json:"chat_is_channel"`
	ChatIsForum             bool                     `json:"chat_is_forum"`
	ChatHasUsername         bool                     `json:"chat_has_username"`
	ChatIsCreated           bool                     `json:"chat_is_created"`
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights"`
	BotAdministratorRights  *ChatAdministratorRights `json:"bot_administrator_rights"`
	BotIsMember             bool                     `json:"bot_is_member"`
	RequestTitle            bool                     `json:"request_title"`
	RequestUsername         bool                     `json:"request_username"`
	RequestPhoto            bool                     `json:"request_photo"`
}

// KeyboardButtonRequestUsers maps to Telegram Bot API type "KeyboardButtonRequestUsers".
type KeyboardButtonRequestUsers struct {
	RequestID       int64 `json:"request_id"`
	UserIsBot       bool  `json:"user_is_bot"`
	UserIsPremium   bool  `json:"user_is_premium"`
	MaxQuantity     int64 `json:"max_quantity"`
	RequestName     bool  `json:"request_name"`
	RequestUsername bool  `json:"request_username"`
	RequestPhoto    bool  `json:"request_photo"`
}

// LabeledPrice maps to Telegram Bot API type "LabeledPrice".
type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int64  `json:"amount"`
}

// LinkPreviewOptions maps to Telegram Bot API type "LinkPreviewOptions".
type LinkPreviewOptions struct {
	IsDisabled       bool   `json:"is_disabled"`
	URL              string `json:"url"`
	PreferSmallMedia bool   `json:"prefer_small_media"`
	PreferLargeMedia bool   `json:"prefer_large_media"`
	ShowAboveText    bool   `json:"show_above_text"`
}

// Location maps to Telegram Bot API type "Location".
type Location struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy"`
	LivePeriod           int64   `json:"live_period"`
	Heading              int64   `json:"heading"`
	ProximityAlertRadius int64   `json:"proximity_alert_radius"`
}

// LocationAddress maps to Telegram Bot API type "LocationAddress".
type LocationAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	Street      string `json:"street"`
}

// LoginUrl maps to Telegram Bot API type "LoginUrl".
type LoginUrl struct {
	URL                string `json:"url"`
	ForwardText        string `json:"forward_text"`
	BotUsername        string `json:"bot_username"`
	RequestWriteAccess bool   `json:"request_write_access"`
}

// MaskPosition maps to Telegram Bot API type "MaskPosition".
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

// MaybeInaccessibleMessage is a union type in Telegram Bot API.
type MaybeInaccessibleMessage interface {
	isMaybeInaccessibleMessage()
}

// MenuButton is a union type in Telegram Bot API.
type MenuButton interface {
	isMenuButton()
}

// MenuButtonCommands maps to Telegram Bot API type "MenuButtonCommands".
type MenuButtonCommands struct {
	Type string `json:"type"`
}

func (*MenuButtonCommands) isMenuButton() {}

// MenuButtonDefault maps to Telegram Bot API type "MenuButtonDefault".
type MenuButtonDefault struct {
	Type string `json:"type"`
}

func (*MenuButtonDefault) isMenuButton() {}

// MenuButtonWebApp maps to Telegram Bot API type "MenuButtonWebApp".
type MenuButtonWebApp struct {
	Type   string      `json:"type"`
	Text   string      `json:"text"`
	WebApp *WebAppInfo `json:"web_app"`
}

func (*MenuButtonWebApp) isMenuButton() {}

// Message maps to Telegram Bot API type "Message".
type Message struct {
	MessageID                     int64                          `json:"message_id"`
	MessageThreadID               int64                          `json:"message_thread_id"`
	DirectMessagesTopic           *DirectMessagesTopic           `json:"direct_messages_topic"`
	From                          *User                          `json:"from"`
	SenderChat                    *Chat                          `json:"sender_chat"`
	SenderBoostCount              int64                          `json:"sender_boost_count"`
	SenderBusinessBot             *User                          `json:"sender_business_bot"`
	SenderTag                     string                         `json:"sender_tag"`
	Date                          int64                          `json:"date"`
	BusinessConnectionID          string                         `json:"business_connection_id"`
	Chat                          *Chat                          `json:"chat"`
	ForwardOrigin                 MessageOrigin                  `json:"forward_origin"`
	IsTopicMessage                bool                           `json:"is_topic_message"`
	IsAutomaticForward            bool                           `json:"is_automatic_forward"`
	ReplyToMessage                *Message                       `json:"reply_to_message"`
	ExternalReply                 *ExternalReplyInfo             `json:"external_reply"`
	Quote                         *TextQuote                     `json:"quote"`
	ReplyToStory                  *Story                         `json:"reply_to_story"`
	ReplyToChecklistTaskID        int64                          `json:"reply_to_checklist_task_id"`
	ViaBot                        *User                          `json:"via_bot"`
	EditDate                      int64                          `json:"edit_date"`
	HasProtectedContent           bool                           `json:"has_protected_content"`
	IsFromOffline                 bool                           `json:"is_from_offline"`
	IsPaidPost                    bool                           `json:"is_paid_post"`
	MediaGroupID                  string                         `json:"media_group_id"`
	AuthorSignature               string                         `json:"author_signature"`
	PaidStarCount                 int64                          `json:"paid_star_count"`
	Text                          string                         `json:"text"`
	Entities                      []MessageEntity                `json:"entities"`
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options"`
	SuggestedPostInfo             *SuggestedPostInfo             `json:"suggested_post_info"`
	EffectID                      string                         `json:"effect_id"`
	Animation                     *Animation                     `json:"animation"`
	Audio                         *Audio                         `json:"audio"`
	Document                      *Document                      `json:"document"`
	PaidMedia                     *PaidMediaInfo                 `json:"paid_media"`
	Photo                         []PhotoSize                    `json:"photo"`
	Sticker                       *Sticker                       `json:"sticker"`
	Story                         *Story                         `json:"story"`
	Video                         *Video                         `json:"video"`
	VideoNote                     *VideoNote                     `json:"video_note"`
	Voice                         *Voice                         `json:"voice"`
	Caption                       string                         `json:"caption"`
	CaptionEntities               []MessageEntity                `json:"caption_entities"`
	ShowCaptionAboveMedia         bool                           `json:"show_caption_above_media"`
	HasMediaSpoiler               bool                           `json:"has_media_spoiler"`
	Checklist                     *Checklist                     `json:"checklist"`
	Contact                       *Contact                       `json:"contact"`
	Dice                          *Dice                          `json:"dice"`
	Game                          *Game                          `json:"game"`
	Poll                          *Poll                          `json:"poll"`
	Venue                         *Venue                         `json:"venue"`
	Location                      *Location                      `json:"location"`
	NewChatMembers                []User                         `json:"new_chat_members"`
	LeftChatMember                *User                          `json:"left_chat_member"`
	ChatOwnerLeft                 *ChatOwnerLeft                 `json:"chat_owner_left"`
	ChatOwnerChanged              *ChatOwnerChanged              `json:"chat_owner_changed"`
	NewChatTitle                  string                         `json:"new_chat_title"`
	NewChatPhoto                  []PhotoSize                    `json:"new_chat_photo"`
	DeleteChatPhoto               bool                           `json:"delete_chat_photo"`
	GroupChatCreated              bool                           `json:"group_chat_created"`
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created"`
	ChannelChatCreated            bool                           `json:"channel_chat_created"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"`
	MigrateToChatID               int64                          `json:"migrate_to_chat_id"`
	MigrateFromChatID             int64                          `json:"migrate_from_chat_id"`
	PinnedMessage                 MaybeInaccessibleMessage       `json:"pinned_message"`
	Invoice                       *Invoice                       `json:"invoice"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment"`
	RefundedPayment               *RefundedPayment               `json:"refunded_payment"`
	UsersShared                   *UsersShared                   `json:"users_shared"`
	ChatShared                    *ChatShared                    `json:"chat_shared"`
	Gift                          *GiftInfo                      `json:"gift"`
	UniqueGift                    *UniqueGiftInfo                `json:"unique_gift"`
	GiftUpgradeSent               *GiftInfo                      `json:"gift_upgrade_sent"`
	ConnectedWebsite              string                         `json:"connected_website"`
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed"`
	PassportData                  *PassportData                  `json:"passport_data"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered"`
	BoostAdded                    *ChatBoostAdded                `json:"boost_added"`
	ChatBackgroundSet             *ChatBackground                `json:"chat_background_set"`
	ChecklistTasksDone            *ChecklistTasksDone            `json:"checklist_tasks_done"`
	ChecklistTasksAdded           *ChecklistTasksAdded           `json:"checklist_tasks_added"`
	DirectMessagePriceChanged     *DirectMessagePriceChanged     `json:"direct_message_price_changed"`
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created"`
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited"`
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed"`
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened"`
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden"`
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden"`
	GiveawayCreated               *GiveawayCreated               `json:"giveaway_created"`
	Giveaway                      *Giveaway                      `json:"giveaway"`
	GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners"`
	GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed"`
	PaidMessagePriceChanged       *PaidMessagePriceChanged       `json:"paid_message_price_changed"`
	SuggestedPostApproved         *SuggestedPostApproved         `json:"suggested_post_approved"`
	SuggestedPostApprovalFailed   *SuggestedPostApprovalFailed   `json:"suggested_post_approval_failed"`
	SuggestedPostDeclined         *SuggestedPostDeclined         `json:"suggested_post_declined"`
	SuggestedPostPaid             *SuggestedPostPaid             `json:"suggested_post_paid"`
	SuggestedPostRefunded         *SuggestedPostRefunded         `json:"suggested_post_refunded"`
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled"`
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started"`
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended"`
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited"`
	WebAppData                    *WebAppData                    `json:"web_app_data"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup"`
}

func (*Message) isMaybeInaccessibleMessage() {}

// MessageAutoDeleteTimerChanged maps to Telegram Bot API type "MessageAutoDeleteTimerChanged".
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int64 `json:"message_auto_delete_time"`
}

// MessageEntity maps to Telegram Bot API type "MessageEntity".
type MessageEntity struct {
	Type           string `json:"type"`
	Offset         int64  `json:"offset"`
	Length         int64  `json:"length"`
	URL            string `json:"url"`
	User           *User  `json:"user"`
	Language       string `json:"language"`
	CustomEmojiID  string `json:"custom_emoji_id"`
	UnixTime       int64  `json:"unix_time"`
	DateTimeFormat string `json:"date_time_format"`
}

// MessageId maps to Telegram Bot API type "MessageId".
type MessageId struct {
	MessageID int64 `json:"message_id"`
}

// MessageOrigin is a union type in Telegram Bot API.
type MessageOrigin interface {
	isMessageOrigin()
}

// MessageOriginChannel maps to Telegram Bot API type "MessageOriginChannel".
type MessageOriginChannel struct {
	Type            string `json:"type"`
	Date            int64  `json:"date"`
	Chat            *Chat  `json:"chat"`
	MessageID       int64  `json:"message_id"`
	AuthorSignature string `json:"author_signature"`
}

func (*MessageOriginChannel) isMessageOrigin() {}

// MessageOriginChat maps to Telegram Bot API type "MessageOriginChat".
type MessageOriginChat struct {
	Type            string `json:"type"`
	Date            int64  `json:"date"`
	SenderChat      *Chat  `json:"sender_chat"`
	AuthorSignature string `json:"author_signature"`
}

func (*MessageOriginChat) isMessageOrigin() {}

// MessageOriginHiddenUser maps to Telegram Bot API type "MessageOriginHiddenUser".
type MessageOriginHiddenUser struct {
	Type           string `json:"type"`
	Date           int64  `json:"date"`
	SenderUserName string `json:"sender_user_name"`
}

func (*MessageOriginHiddenUser) isMessageOrigin() {}

// MessageOriginUser maps to Telegram Bot API type "MessageOriginUser".
type MessageOriginUser struct {
	Type       string `json:"type"`
	Date       int64  `json:"date"`
	SenderUser *User  `json:"sender_user"`
}

func (*MessageOriginUser) isMessageOrigin() {}

// MessageReactionCountUpdated maps to Telegram Bot API type "MessageReactionCountUpdated".
type MessageReactionCountUpdated struct {
	Chat      *Chat           `json:"chat"`
	MessageID int64           `json:"message_id"`
	Date      int64           `json:"date"`
	Reactions []ReactionCount `json:"reactions"`
}

// MessageReactionUpdated maps to Telegram Bot API type "MessageReactionUpdated".
type MessageReactionUpdated struct {
	Chat        *Chat          `json:"chat"`
	MessageID   int64          `json:"message_id"`
	User        *User          `json:"user"`
	ActorChat   *Chat          `json:"actor_chat"`
	Date        int64          `json:"date"`
	OldReaction []ReactionType `json:"old_reaction"`
	NewReaction []ReactionType `json:"new_reaction"`
}

// OrderInfo maps to Telegram Bot API type "OrderInfo".
type OrderInfo struct {
	Name            string           `json:"name"`
	PhoneNumber     string           `json:"phone_number"`
	Email           string           `json:"email"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// OwnedGift is a union type in Telegram Bot API.
type OwnedGift interface {
	isOwnedGift()
}

// OwnedGiftRegular maps to Telegram Bot API type "OwnedGiftRegular".
type OwnedGiftRegular struct {
	Type                    string          `json:"type"`
	Gift                    *Gift           `json:"gift"`
	OwnedGiftID             string          `json:"owned_gift_id"`
	SenderUser              *User           `json:"sender_user"`
	SendDate                int64           `json:"send_date"`
	Text                    string          `json:"text"`
	Entities                []MessageEntity `json:"entities"`
	IsPrivate               bool            `json:"is_private"`
	IsSaved                 bool            `json:"is_saved"`
	CanBeUpgraded           bool            `json:"can_be_upgraded"`
	WasRefunded             bool            `json:"was_refunded"`
	ConvertStarCount        int64           `json:"convert_star_count"`
	PrepaidUpgradeStarCount int64           `json:"prepaid_upgrade_star_count"`
	IsUpgradeSeparate       bool            `json:"is_upgrade_separate"`
	UniqueGiftNumber        int64           `json:"unique_gift_number"`
}

func (*OwnedGiftRegular) isOwnedGift() {}

// OwnedGiftUnique maps to Telegram Bot API type "OwnedGiftUnique".
type OwnedGiftUnique struct {
	Type              string      `json:"type"`
	Gift              *UniqueGift `json:"gift"`
	OwnedGiftID       string      `json:"owned_gift_id"`
	SenderUser        *User       `json:"sender_user"`
	SendDate          int64       `json:"send_date"`
	IsSaved           bool        `json:"is_saved"`
	CanBeTransferred  bool        `json:"can_be_transferred"`
	TransferStarCount int64       `json:"transfer_star_count"`
	NextTransferDate  int64       `json:"next_transfer_date"`
}

func (*OwnedGiftUnique) isOwnedGift() {}

// OwnedGifts maps to Telegram Bot API type "OwnedGifts".
type OwnedGifts struct {
	TotalCount int64       `json:"total_count"`
	Gifts      []OwnedGift `json:"gifts"`
	NextOffset string      `json:"next_offset"`
}

// PaidMedia is a union type in Telegram Bot API.
type PaidMedia interface {
	isPaidMedia()
}

// PaidMediaInfo maps to Telegram Bot API type "PaidMediaInfo".
type PaidMediaInfo struct {
	StarCount int64       `json:"star_count"`
	PaidMedia []PaidMedia `json:"paid_media"`
}

// PaidMediaPhoto maps to Telegram Bot API type "PaidMediaPhoto".
type PaidMediaPhoto struct {
	Type  string      `json:"type"`
	Photo []PhotoSize `json:"photo"`
}

func (*PaidMediaPhoto) isPaidMedia() {}

// PaidMediaPreview maps to Telegram Bot API type "PaidMediaPreview".
type PaidMediaPreview struct {
	Type     string `json:"type"`
	Width    int64  `json:"width"`
	Height   int64  `json:"height"`
	Duration int64  `json:"duration"`
}

func (*PaidMediaPreview) isPaidMedia() {}

// PaidMediaPurchased maps to Telegram Bot API type "PaidMediaPurchased".
type PaidMediaPurchased struct {
	From             *User  `json:"from"`
	PaidMediaPayload string `json:"paid_media_payload"`
}

// PaidMediaVideo maps to Telegram Bot API type "PaidMediaVideo".
type PaidMediaVideo struct {
	Type  string `json:"type"`
	Video *Video `json:"video"`
}

func (*PaidMediaVideo) isPaidMedia() {}

// PaidMessagePriceChanged maps to Telegram Bot API type "PaidMessagePriceChanged".
type PaidMessagePriceChanged struct {
	PaidMessageStarCount int64 `json:"paid_message_star_count"`
}

// PassportData maps to Telegram Bot API type "PassportData".
type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials *EncryptedCredentials      `json:"credentials"`
}

// PassportElementError is a union type in Telegram Bot API.
type PassportElementError interface {
	isPassportElementError()
}

// PassportElementErrorDataField maps to Telegram Bot API type "PassportElementErrorDataField".
type PassportElementErrorDataField struct {
	Source    string `json:"source"`
	Type      string `json:"type"`
	FieldName string `json:"field_name"`
	DataHash  string `json:"data_hash"`
	Message   string `json:"message"`
}

func (*PassportElementErrorDataField) isPassportElementError() {}

// PassportElementErrorFile maps to Telegram Bot API type "PassportElementErrorFile".
type PassportElementErrorFile struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (*PassportElementErrorFile) isPassportElementError() {}

// PassportElementErrorFiles maps to Telegram Bot API type "PassportElementErrorFiles".
type PassportElementErrorFiles struct {
	Source     string   `json:"source"`
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

func (*PassportElementErrorFiles) isPassportElementError() {}

// PassportElementErrorFrontSide maps to Telegram Bot API type "PassportElementErrorFrontSide".
type PassportElementErrorFrontSide struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (*PassportElementErrorFrontSide) isPassportElementError() {}

// PassportElementErrorReverseSide maps to Telegram Bot API type "PassportElementErrorReverseSide".
type PassportElementErrorReverseSide struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (*PassportElementErrorReverseSide) isPassportElementError() {}

// PassportElementErrorSelfie maps to Telegram Bot API type "PassportElementErrorSelfie".
type PassportElementErrorSelfie struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (*PassportElementErrorSelfie) isPassportElementError() {}

// PassportElementErrorTranslationFile maps to Telegram Bot API type "PassportElementErrorTranslationFile".
type PassportElementErrorTranslationFile struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (*PassportElementErrorTranslationFile) isPassportElementError() {}

// PassportElementErrorTranslationFiles maps to Telegram Bot API type "PassportElementErrorTranslationFiles".
type PassportElementErrorTranslationFiles struct {
	Source     string   `json:"source"`
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

func (*PassportElementErrorTranslationFiles) isPassportElementError() {}

// PassportElementErrorUnspecified maps to Telegram Bot API type "PassportElementErrorUnspecified".
type PassportElementErrorUnspecified struct {
	Source      string `json:"source"`
	Type        string `json:"type"`
	ElementHash string `json:"element_hash"`
	Message     string `json:"message"`
}

func (*PassportElementErrorUnspecified) isPassportElementError() {}

// PassportFile maps to Telegram Bot API type "PassportFile".
type PassportFile struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	FileDate     int64  `json:"file_date"`
}

// PhotoSize maps to Telegram Bot API type "PhotoSize".
type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
	FileSize     int64  `json:"file_size"`
}

// Poll maps to Telegram Bot API type "Poll".
type Poll struct {
	ID                    string          `json:"id"`
	Question              string          `json:"question"`
	QuestionEntities      []MessageEntity `json:"question_entities"`
	Options               []PollOption    `json:"options"`
	TotalVoterCount       int64           `json:"total_voter_count"`
	IsClosed              bool            `json:"is_closed"`
	IsAnonymous           bool            `json:"is_anonymous"`
	Type                  string          `json:"type"`
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`
	CorrectOptionID       int64           `json:"correct_option_id"`
	Explanation           string          `json:"explanation"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities"`
	OpenPeriod            int64           `json:"open_period"`
	CloseDate             int64           `json:"close_date"`
}

// PollAnswer maps to Telegram Bot API type "PollAnswer".
type PollAnswer struct {
	PollID    string  `json:"poll_id"`
	VoterChat *Chat   `json:"voter_chat"`
	User      *User   `json:"user"`
	OptionIds []int64 `json:"option_ids"`
}

// PollOption maps to Telegram Bot API type "PollOption".
type PollOption struct {
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities"`
	VoterCount   int64           `json:"voter_count"`
}

// PreCheckoutQuery maps to Telegram Bot API type "PreCheckoutQuery".
type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             *User      `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int64      `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID string     `json:"shipping_option_id"`
	OrderInfo        *OrderInfo `json:"order_info"`
}

// PreparedInlineMessage maps to Telegram Bot API type "PreparedInlineMessage".
type PreparedInlineMessage struct {
	ID             string `json:"id"`
	ExpirationDate int64  `json:"expiration_date"`
}

// ProximityAlertTriggered maps to Telegram Bot API type "ProximityAlertTriggered".
type ProximityAlertTriggered struct {
	Traveler *User `json:"traveler"`
	Watcher  *User `json:"watcher"`
	Distance int64 `json:"distance"`
}

// ReactionCount maps to Telegram Bot API type "ReactionCount".
type ReactionCount struct {
	Type       ReactionType `json:"type"`
	TotalCount int64        `json:"total_count"`
}

// ReactionType is a union type in Telegram Bot API.
type ReactionType interface {
	isReactionType()
}

// ReactionTypeCustomEmoji maps to Telegram Bot API type "ReactionTypeCustomEmoji".
type ReactionTypeCustomEmoji struct {
	Type          string `json:"type"`
	CustomEmojiID string `json:"custom_emoji_id"`
}

func (*ReactionTypeCustomEmoji) isReactionType() {}

// ReactionTypeEmoji maps to Telegram Bot API type "ReactionTypeEmoji".
type ReactionTypeEmoji struct {
	Type  string `json:"type"`
	Emoji string `json:"emoji"`
}

func (*ReactionTypeEmoji) isReactionType() {}

// ReactionTypePaid maps to Telegram Bot API type "ReactionTypePaid".
type ReactionTypePaid struct {
	Type string `json:"type"`
}

func (*ReactionTypePaid) isReactionType() {}

// RefundedPayment maps to Telegram Bot API type "RefundedPayment".
type RefundedPayment struct {
	Currency                string `json:"currency"`
	TotalAmount             int64  `json:"total_amount"`
	InvoicePayload          string `json:"invoice_payload"`
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string `json:"provider_payment_charge_id"`
}

// ReplyKeyboardMarkup maps to Telegram Bot API type "ReplyKeyboardMarkup".
type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`
	IsPersistent          bool               `json:"is_persistent"`
	ResizeKeyboard        bool               `json:"resize_keyboard"`
	OneTimeKeyboard       bool               `json:"one_time_keyboard"`
	InputFieldPlaceholder string             `json:"input_field_placeholder"`
	Selective             bool               `json:"selective"`
}

// ReplyKeyboardRemove maps to Telegram Bot API type "ReplyKeyboardRemove".
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
}

// ReplyParameters maps to Telegram Bot API type "ReplyParameters".
type ReplyParameters struct {
	MessageID                int64           `json:"message_id"`
	ChatID                   any             `json:"chat_id"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	Quote                    string          `json:"quote"`
	QuoteParseMode           string          `json:"quote_parse_mode"`
	QuoteEntities            []MessageEntity `json:"quote_entities"`
	QuotePosition            int64           `json:"quote_position"`
	ChecklistTaskID          int64           `json:"checklist_task_id"`
}

// RevenueWithdrawalState is a union type in Telegram Bot API.
type RevenueWithdrawalState interface {
	isRevenueWithdrawalState()
}

// RevenueWithdrawalStateFailed maps to Telegram Bot API type "RevenueWithdrawalStateFailed".
type RevenueWithdrawalStateFailed struct {
	Type string `json:"type"`
}

func (*RevenueWithdrawalStateFailed) isRevenueWithdrawalState() {}

// RevenueWithdrawalStatePending maps to Telegram Bot API type "RevenueWithdrawalStatePending".
type RevenueWithdrawalStatePending struct {
	Type string `json:"type"`
}

func (*RevenueWithdrawalStatePending) isRevenueWithdrawalState() {}

// RevenueWithdrawalStateSucceeded maps to Telegram Bot API type "RevenueWithdrawalStateSucceeded".
type RevenueWithdrawalStateSucceeded struct {
	Type string `json:"type"`
	Date int64  `json:"date"`
	URL  string `json:"url"`
}

func (*RevenueWithdrawalStateSucceeded) isRevenueWithdrawalState() {}

// SentWebAppMessage maps to Telegram Bot API type "SentWebAppMessage".
type SentWebAppMessage struct {
	InlineMessageID string `json:"inline_message_id"`
}

// SharedUser maps to Telegram Bot API type "SharedUser".
type SharedUser struct {
	UserID    int64       `json:"user_id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Username  string      `json:"username"`
	Photo     []PhotoSize `json:"photo"`
}

// ShippingAddress maps to Telegram Bot API type "ShippingAddress".
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

// ShippingOption maps to Telegram Bot API type "ShippingOption".
type ShippingOption struct {
	ID     string         `json:"id"`
	Title  string         `json:"title"`
	Prices []LabeledPrice `json:"prices"`
}

// ShippingQuery maps to Telegram Bot API type "ShippingQuery".
type ShippingQuery struct {
	ID              string           `json:"id"`
	From            *User            `json:"from"`
	InvoicePayload  string           `json:"invoice_payload"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// StarAmount maps to Telegram Bot API type "StarAmount".
type StarAmount struct {
	Amount         int64 `json:"amount"`
	NanostarAmount int64 `json:"nanostar_amount"`
}

// StarTransaction maps to Telegram Bot API type "StarTransaction".
type StarTransaction struct {
	ID             string             `json:"id"`
	Amount         int64              `json:"amount"`
	NanostarAmount int64              `json:"nanostar_amount"`
	Date           int64              `json:"date"`
	Source         TransactionPartner `json:"source"`
	Receiver       TransactionPartner `json:"receiver"`
}

// StarTransactions maps to Telegram Bot API type "StarTransactions".
type StarTransactions struct {
	Transactions []StarTransaction `json:"transactions"`
}

// Sticker maps to Telegram Bot API type "Sticker".
type Sticker struct {
	FileID           string        `json:"file_id"`
	FileUniqueID     string        `json:"file_unique_id"`
	Type             string        `json:"type"`
	Width            int64         `json:"width"`
	Height           int64         `json:"height"`
	IsAnimated       bool          `json:"is_animated"`
	IsVideo          bool          `json:"is_video"`
	Thumbnail        *PhotoSize    `json:"thumbnail"`
	Emoji            string        `json:"emoji"`
	SetName          string        `json:"set_name"`
	PremiumAnimation *File         `json:"premium_animation"`
	MaskPosition     *MaskPosition `json:"mask_position"`
	CustomEmojiID    string        `json:"custom_emoji_id"`
	NeedsRepainting  bool          `json:"needs_repainting"`
	FileSize         int64         `json:"file_size"`
}

// StickerSet maps to Telegram Bot API type "StickerSet".
type StickerSet struct {
	Name        string     `json:"name"`
	Title       string     `json:"title"`
	StickerType string     `json:"sticker_type"`
	Stickers    []Sticker  `json:"stickers"`
	Thumbnail   *PhotoSize `json:"thumbnail"`
}

// Story maps to Telegram Bot API type "Story".
type Story struct {
	Chat *Chat `json:"chat"`
	ID   int64 `json:"id"`
}

// StoryArea maps to Telegram Bot API type "StoryArea".
type StoryArea struct {
	Position *StoryAreaPosition `json:"position"`
	Type     StoryAreaType      `json:"type"`
}

// StoryAreaPosition maps to Telegram Bot API type "StoryAreaPosition".
type StoryAreaPosition struct {
	XPercentage            float64 `json:"x_percentage"`
	YPercentage            float64 `json:"y_percentage"`
	WidthPercentage        float64 `json:"width_percentage"`
	HeightPercentage       float64 `json:"height_percentage"`
	RotationAngle          float64 `json:"rotation_angle"`
	CornerRadiusPercentage float64 `json:"corner_radius_percentage"`
}

// StoryAreaType is a union type in Telegram Bot API.
type StoryAreaType interface {
	isStoryAreaType()
}

// StoryAreaTypeLink maps to Telegram Bot API type "StoryAreaTypeLink".
type StoryAreaTypeLink struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

func (*StoryAreaTypeLink) isStoryAreaType() {}

// StoryAreaTypeLocation maps to Telegram Bot API type "StoryAreaTypeLocation".
type StoryAreaTypeLocation struct {
	Type      string           `json:"type"`
	Latitude  float64          `json:"latitude"`
	Longitude float64          `json:"longitude"`
	Address   *LocationAddress `json:"address"`
}

func (*StoryAreaTypeLocation) isStoryAreaType() {}

// StoryAreaTypeSuggestedReaction maps to Telegram Bot API type "StoryAreaTypeSuggestedReaction".
type StoryAreaTypeSuggestedReaction struct {
	Type         string       `json:"type"`
	ReactionType ReactionType `json:"reaction_type"`
	IsDark       bool         `json:"is_dark"`
	IsFlipped    bool         `json:"is_flipped"`
}

func (*StoryAreaTypeSuggestedReaction) isStoryAreaType() {}

// StoryAreaTypeUniqueGift maps to Telegram Bot API type "StoryAreaTypeUniqueGift".
type StoryAreaTypeUniqueGift struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

func (*StoryAreaTypeUniqueGift) isStoryAreaType() {}

// StoryAreaTypeWeather maps to Telegram Bot API type "StoryAreaTypeWeather".
type StoryAreaTypeWeather struct {
	Type            string  `json:"type"`
	Temperature     float64 `json:"temperature"`
	Emoji           string  `json:"emoji"`
	BackgroundColor int64   `json:"background_color"`
}

func (*StoryAreaTypeWeather) isStoryAreaType() {}

// SuccessfulPayment maps to Telegram Bot API type "SuccessfulPayment".
type SuccessfulPayment struct {
	Currency                   string     `json:"currency"`
	TotalAmount                int64      `json:"total_amount"`
	InvoicePayload             string     `json:"invoice_payload"`
	SubscriptionExpirationDate int64      `json:"subscription_expiration_date"`
	IsRecurring                bool       `json:"is_recurring"`
	IsFirstRecurring           bool       `json:"is_first_recurring"`
	ShippingOptionID           string     `json:"shipping_option_id"`
	OrderInfo                  *OrderInfo `json:"order_info"`
	TelegramPaymentChargeID    string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID    string     `json:"provider_payment_charge_id"`
}

// SuggestedPostApprovalFailed maps to Telegram Bot API type "SuggestedPostApprovalFailed".
type SuggestedPostApprovalFailed struct {
	SuggestedPostMessage *Message            `json:"suggested_post_message"`
	Price                *SuggestedPostPrice `json:"price"`
}

// SuggestedPostApproved maps to Telegram Bot API type "SuggestedPostApproved".
type SuggestedPostApproved struct {
	SuggestedPostMessage *Message            `json:"suggested_post_message"`
	Price                *SuggestedPostPrice `json:"price"`
	SendDate             int64               `json:"send_date"`
}

// SuggestedPostDeclined maps to Telegram Bot API type "SuggestedPostDeclined".
type SuggestedPostDeclined struct {
	SuggestedPostMessage *Message `json:"suggested_post_message"`
	Comment              string   `json:"comment"`
}

// SuggestedPostInfo maps to Telegram Bot API type "SuggestedPostInfo".
type SuggestedPostInfo struct {
	State    string              `json:"state"`
	Price    *SuggestedPostPrice `json:"price"`
	SendDate int64               `json:"send_date"`
}

// SuggestedPostPaid maps to Telegram Bot API type "SuggestedPostPaid".
type SuggestedPostPaid struct {
	SuggestedPostMessage *Message    `json:"suggested_post_message"`
	Currency             string      `json:"currency"`
	Amount               int64       `json:"amount"`
	StarAmount           *StarAmount `json:"star_amount"`
}

// SuggestedPostParameters maps to Telegram Bot API type "SuggestedPostParameters".
type SuggestedPostParameters struct {
	Price    *SuggestedPostPrice `json:"price"`
	SendDate int64               `json:"send_date"`
}

// SuggestedPostPrice maps to Telegram Bot API type "SuggestedPostPrice".
type SuggestedPostPrice struct {
	Currency string `json:"currency"`
	Amount   int64  `json:"amount"`
}

// SuggestedPostRefunded maps to Telegram Bot API type "SuggestedPostRefunded".
type SuggestedPostRefunded struct {
	SuggestedPostMessage *Message `json:"suggested_post_message"`
	Reason               string   `json:"reason"`
}

// SwitchInlineQueryChosenChat maps to Telegram Bot API type "SwitchInlineQueryChosenChat".
type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query"`
	AllowUserChats    bool   `json:"allow_user_chats"`
	AllowBotChats     bool   `json:"allow_bot_chats"`
	AllowGroupChats   bool   `json:"allow_group_chats"`
	AllowChannelChats bool   `json:"allow_channel_chats"`
}

// TextQuote maps to Telegram Bot API type "TextQuote".
type TextQuote struct {
	Text     string          `json:"text"`
	Entities []MessageEntity `json:"entities"`
	Position int64           `json:"position"`
	IsManual bool            `json:"is_manual"`
}

// TransactionPartner is a union type in Telegram Bot API.
type TransactionPartner interface {
	isTransactionPartner()
}

// TransactionPartnerAffiliateProgram maps to Telegram Bot API type "TransactionPartnerAffiliateProgram".
type TransactionPartnerAffiliateProgram struct {
	Type               string `json:"type"`
	SponsorUser        *User  `json:"sponsor_user"`
	CommissionPerMille int64  `json:"commission_per_mille"`
}

func (*TransactionPartnerAffiliateProgram) isTransactionPartner() {}

// TransactionPartnerChat maps to Telegram Bot API type "TransactionPartnerChat".
type TransactionPartnerChat struct {
	Type string `json:"type"`
	Chat *Chat  `json:"chat"`
	Gift *Gift  `json:"gift"`
}

func (*TransactionPartnerChat) isTransactionPartner() {}

// TransactionPartnerFragment maps to Telegram Bot API type "TransactionPartnerFragment".
type TransactionPartnerFragment struct {
	Type            string                 `json:"type"`
	WithdrawalState RevenueWithdrawalState `json:"withdrawal_state"`
}

func (*TransactionPartnerFragment) isTransactionPartner() {}

// TransactionPartnerOther maps to Telegram Bot API type "TransactionPartnerOther".
type TransactionPartnerOther struct {
	Type string `json:"type"`
}

func (*TransactionPartnerOther) isTransactionPartner() {}

// TransactionPartnerTelegramAds maps to Telegram Bot API type "TransactionPartnerTelegramAds".
type TransactionPartnerTelegramAds struct {
	Type string `json:"type"`
}

func (*TransactionPartnerTelegramAds) isTransactionPartner() {}

// TransactionPartnerTelegramApi maps to Telegram Bot API type "TransactionPartnerTelegramApi".
type TransactionPartnerTelegramApi struct {
	Type         string `json:"type"`
	RequestCount int64  `json:"request_count"`
}

func (*TransactionPartnerTelegramApi) isTransactionPartner() {}

// TransactionPartnerUser maps to Telegram Bot API type "TransactionPartnerUser".
type TransactionPartnerUser struct {
	Type                        string         `json:"type"`
	TransactionType             string         `json:"transaction_type"`
	User                        *User          `json:"user"`
	Affiliate                   *AffiliateInfo `json:"affiliate"`
	InvoicePayload              string         `json:"invoice_payload"`
	SubscriptionPeriod          int64          `json:"subscription_period"`
	PaidMedia                   []PaidMedia    `json:"paid_media"`
	PaidMediaPayload            string         `json:"paid_media_payload"`
	Gift                        *Gift          `json:"gift"`
	PremiumSubscriptionDuration int64          `json:"premium_subscription_duration"`
}

func (*TransactionPartnerUser) isTransactionPartner() {}

// UniqueGift maps to Telegram Bot API type "UniqueGift".
type UniqueGift struct {
	GiftID           string              `json:"gift_id"`
	BaseName         string              `json:"base_name"`
	Name             string              `json:"name"`
	Number           int64               `json:"number"`
	Model            *UniqueGiftModel    `json:"model"`
	Symbol           *UniqueGiftSymbol   `json:"symbol"`
	Backdrop         *UniqueGiftBackdrop `json:"backdrop"`
	IsPremium        bool                `json:"is_premium"`
	IsBurned         bool                `json:"is_burned"`
	IsFromBlockchain bool                `json:"is_from_blockchain"`
	Colors           *UniqueGiftColors   `json:"colors"`
	PublisherChat    *Chat               `json:"publisher_chat"`
}

// UniqueGiftBackdrop maps to Telegram Bot API type "UniqueGiftBackdrop".
type UniqueGiftBackdrop struct {
	Name           string                    `json:"name"`
	Colors         *UniqueGiftBackdropColors `json:"colors"`
	RarityPerMille int64                     `json:"rarity_per_mille"`
}

// UniqueGiftBackdropColors maps to Telegram Bot API type "UniqueGiftBackdropColors".
type UniqueGiftBackdropColors struct {
	CenterColor int64 `json:"center_color"`
	EdgeColor   int64 `json:"edge_color"`
	SymbolColor int64 `json:"symbol_color"`
	TextColor   int64 `json:"text_color"`
}

// UniqueGiftColors maps to Telegram Bot API type "UniqueGiftColors".
type UniqueGiftColors struct {
	ModelCustomEmojiID    string  `json:"model_custom_emoji_id"`
	SymbolCustomEmojiID   string  `json:"symbol_custom_emoji_id"`
	LightThemeMainColor   int64   `json:"light_theme_main_color"`
	LightThemeOtherColors []int64 `json:"light_theme_other_colors"`
	DarkThemeMainColor    int64   `json:"dark_theme_main_color"`
	DarkThemeOtherColors  []int64 `json:"dark_theme_other_colors"`
}

// UniqueGiftInfo maps to Telegram Bot API type "UniqueGiftInfo".
type UniqueGiftInfo struct {
	Gift               *UniqueGift `json:"gift"`
	Origin             string      `json:"origin"`
	LastResaleCurrency string      `json:"last_resale_currency"`
	LastResaleAmount   int64       `json:"last_resale_amount"`
	OwnedGiftID        string      `json:"owned_gift_id"`
	TransferStarCount  int64       `json:"transfer_star_count"`
	NextTransferDate   int64       `json:"next_transfer_date"`
}

// UniqueGiftModel maps to Telegram Bot API type "UniqueGiftModel".
type UniqueGiftModel struct {
	Name           string   `json:"name"`
	Sticker        *Sticker `json:"sticker"`
	RarityPerMille int64    `json:"rarity_per_mille"`
	Rarity         string   `json:"rarity"`
}

// UniqueGiftSymbol maps to Telegram Bot API type "UniqueGiftSymbol".
type UniqueGiftSymbol struct {
	Name           string   `json:"name"`
	Sticker        *Sticker `json:"sticker"`
	RarityPerMille int64    `json:"rarity_per_mille"`
}

// Update maps to Telegram Bot API type "Update".
type Update struct {
	UpdateID                int64                        `json:"update_id"`
	Message                 *Message                     `json:"message"`
	EditedMessage           *Message                     `json:"edited_message"`
	ChannelPost             *Message                     `json:"channel_post"`
	EditedChannelPost       *Message                     `json:"edited_channel_post"`
	BusinessConnection      *BusinessConnection          `json:"business_connection"`
	BusinessMessage         *Message                     `json:"business_message"`
	EditedBusinessMessage   *Message                     `json:"edited_business_message"`
	DeletedBusinessMessages *BusinessMessagesDeleted     `json:"deleted_business_messages"`
	MessageReaction         *MessageReactionUpdated      `json:"message_reaction"`
	MessageReactionCount    *MessageReactionCountUpdated `json:"message_reaction_count"`
	InlineQuery             *InlineQuery                 `json:"inline_query"`
	ChosenInlineResult      *ChosenInlineResult          `json:"chosen_inline_result"`
	CallbackQuery           *CallbackQuery               `json:"callback_query"`
	ShippingQuery           *ShippingQuery               `json:"shipping_query"`
	PreCheckoutQuery        *PreCheckoutQuery            `json:"pre_checkout_query"`
	PurchasedPaidMedia      *PaidMediaPurchased          `json:"purchased_paid_media"`
	Poll                    *Poll                        `json:"poll"`
	PollAnswer              *PollAnswer                  `json:"poll_answer"`
	MyChatMember            *ChatMemberUpdated           `json:"my_chat_member"`
	ChatMember              *ChatMemberUpdated           `json:"chat_member"`
	ChatJoinRequest         *ChatJoinRequest             `json:"chat_join_request"`
	ChatBoost               *ChatBoostUpdated            `json:"chat_boost"`
	RemovedChatBoost        *ChatBoostRemoved            `json:"removed_chat_boost"`
}

// User maps to Telegram Bot API type "User".
type User struct {
	ID                        int64  `json:"id"`
	IsBot                     bool   `json:"is_bot"`
	FirstName                 string `json:"first_name"`
	LastName                  string `json:"last_name"`
	Username                  string `json:"username"`
	LanguageCode              string `json:"language_code"`
	IsPremium                 bool   `json:"is_premium"`
	AddedToAttachmentMenu     bool   `json:"added_to_attachment_menu"`
	CanJoinGroups             bool   `json:"can_join_groups"`
	CanReadAllGroupMessages   bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries     bool   `json:"supports_inline_queries"`
	CanConnectToBusiness      bool   `json:"can_connect_to_business"`
	HasMainWebApp             bool   `json:"has_main_web_app"`
	HasTopicsEnabled          bool   `json:"has_topics_enabled"`
	AllowsUsersToCreateTopics bool   `json:"allows_users_to_create_topics"`
}

// UserChatBoosts maps to Telegram Bot API type "UserChatBoosts".
type UserChatBoosts struct {
	Boosts []ChatBoost `json:"boosts"`
}

// UserProfileAudios maps to Telegram Bot API type "UserProfileAudios".
type UserProfileAudios struct {
	TotalCount int64   `json:"total_count"`
	Audios     []Audio `json:"audios"`
}

// UserProfilePhotos maps to Telegram Bot API type "UserProfilePhotos".
type UserProfilePhotos struct {
	TotalCount int64         `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

// UserRating maps to Telegram Bot API type "UserRating".
type UserRating struct {
	Level              int64 `json:"level"`
	Rating             int64 `json:"rating"`
	CurrentLevelRating int64 `json:"current_level_rating"`
	NextLevelRating    int64 `json:"next_level_rating"`
}

// UsersShared maps to Telegram Bot API type "UsersShared".
type UsersShared struct {
	RequestID int64        `json:"request_id"`
	Users     []SharedUser `json:"users"`
}

// Venue maps to Telegram Bot API type "Venue".
type Venue struct {
	Location        *Location `json:"location"`
	Title           string    `json:"title"`
	Address         string    `json:"address"`
	FoursquareID    string    `json:"foursquare_id"`
	FoursquareType  string    `json:"foursquare_type"`
	GooglePlaceID   string    `json:"google_place_id"`
	GooglePlaceType string    `json:"google_place_type"`
}

// Video maps to Telegram Bot API type "Video".
type Video struct {
	FileID         string         `json:"file_id"`
	FileUniqueID   string         `json:"file_unique_id"`
	Width          int64          `json:"width"`
	Height         int64          `json:"height"`
	Duration       int64          `json:"duration"`
	Thumbnail      *PhotoSize     `json:"thumbnail"`
	Cover          []PhotoSize    `json:"cover"`
	StartTimestamp int64          `json:"start_timestamp"`
	Qualities      []VideoQuality `json:"qualities"`
	FileName       string         `json:"file_name"`
	MimeType       string         `json:"mime_type"`
	FileSize       int64          `json:"file_size"`
}

// VideoChatEnded maps to Telegram Bot API type "VideoChatEnded".
type VideoChatEnded struct {
	Duration int64 `json:"duration"`
}

// VideoChatParticipantsInvited maps to Telegram Bot API type "VideoChatParticipantsInvited".
type VideoChatParticipantsInvited struct {
	Users []User `json:"users"`
}

// VideoChatScheduled maps to Telegram Bot API type "VideoChatScheduled".
type VideoChatScheduled struct {
	StartDate int64 `json:"start_date"`
}

// VideoChatStarted maps to Telegram Bot API type "VideoChatStarted".
type VideoChatStarted struct {
}

// VideoNote maps to Telegram Bot API type "VideoNote".
type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int64      `json:"length"`
	Duration     int64      `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail"`
	FileSize     int64      `json:"file_size"`
}

// VideoQuality maps to Telegram Bot API type "VideoQuality".
type VideoQuality struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
	Codec        string `json:"codec"`
	FileSize     int64  `json:"file_size"`
}

// Voice maps to Telegram Bot API type "Voice".
type Voice struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Duration     int64  `json:"duration"`
	MimeType     string `json:"mime_type"`
	FileSize     int64  `json:"file_size"`
}

// WebAppData maps to Telegram Bot API type "WebAppData".
type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

// WebAppInfo maps to Telegram Bot API type "WebAppInfo".
type WebAppInfo struct {
	URL string `json:"url"`
}

// WebhookInfo maps to Telegram Bot API type "WebhookInfo".
type WebhookInfo struct {
	URL                          string   `json:"url"`
	HasCustomCertificate         bool     `json:"has_custom_certificate"`
	PendingUpdateCount           int64    `json:"pending_update_count"`
	IPAddress                    string   `json:"ip_address"`
	LastErrorDate                int64    `json:"last_error_date"`
	LastErrorMessage             string   `json:"last_error_message"`
	LastSynchronizationErrorDate int64    `json:"last_synchronization_error_date"`
	MaxConnections               int64    `json:"max_connections"`
	AllowedUpdates               []string `json:"allowed_updates"`
}

// WriteAccessAllowed maps to Telegram Bot API type "WriteAccessAllowed".
type WriteAccessAllowed struct {
	FromRequest        bool   `json:"from_request"`
	WebAppName         string `json:"web_app_name"`
	FromAttachmentMenu bool   `json:"from_attachment_menu"`
}
