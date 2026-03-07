// Telegram Bot API method wrappers aligned with the official docs.
// Source: https://core.telegram.org/bots/api (Bot API 9.5, March 1, 2026)

package tgbot

import "context"

// AddStickerToSetParams contains params for Telegram method "addStickerToSet".
type AddStickerToSetParams struct {
	UserID  int64        `json:"user_id"`
	Name    string       `json:"name"`
	Sticker InputSticker `json:"sticker"`
}

// AddStickerToSet calls Telegram method "addStickerToSet".
// Doc: https://core.telegram.org/bots/api#addstickertoset
func (bot *Bot) AddStickerToSet(ctx context.Context, params *AddStickerToSetParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "addStickerToSet", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// AnswerCallbackQueryParams contains params for Telegram method "answerCallbackQuery".
type AnswerCallbackQueryParams struct {
	CallbackQueryID string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	URL             string `json:"url,omitempty"`
	CacheTime       int64  `json:"cache_time,omitempty"`
}

// AnswerCallbackQuery calls Telegram method "answerCallbackQuery".
// Doc: https://core.telegram.org/bots/api#answercallbackquery
func (bot *Bot) AnswerCallbackQuery(ctx context.Context, params *AnswerCallbackQueryParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "answerCallbackQuery", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// AnswerInlineQueryParams contains params for Telegram method "answerInlineQuery".
type AnswerInlineQueryParams struct {
	InlineQueryID string                   `json:"inline_query_id"`
	Results       []InlineQueryResult      `json:"results"`
	CacheTime     int64                    `json:"cache_time,omitempty"`
	IsPersonal    bool                     `json:"is_personal,omitempty"`
	NextOffset    string                   `json:"next_offset,omitempty"`
	Button        InlineQueryResultsButton `json:"button,omitempty"`
}

// AnswerInlineQuery calls Telegram method "answerInlineQuery".
// Doc: https://core.telegram.org/bots/api#answerinlinequery
func (bot *Bot) AnswerInlineQuery(ctx context.Context, params *AnswerInlineQueryParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "answerInlineQuery", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// AnswerPreCheckoutQueryParams contains params for Telegram method "answerPreCheckoutQuery".
type AnswerPreCheckoutQueryParams struct {
	PreCheckoutQueryID string `json:"pre_checkout_query_id"`
	Ok                 bool   `json:"ok"`
	ErrorMessage       string `json:"error_message,omitempty"`
}

// AnswerPreCheckoutQuery calls Telegram method "answerPreCheckoutQuery".
// Doc: https://core.telegram.org/bots/api#answerprecheckoutquery
func (bot *Bot) AnswerPreCheckoutQuery(ctx context.Context, params *AnswerPreCheckoutQueryParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "answerPreCheckoutQuery", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// AnswerShippingQueryParams contains params for Telegram method "answerShippingQuery".
type AnswerShippingQueryParams struct {
	ShippingQueryID string           `json:"shipping_query_id"`
	Ok              bool             `json:"ok"`
	ShippingOptions []ShippingOption `json:"shipping_options,omitempty"`
	ErrorMessage    string           `json:"error_message,omitempty"`
}

// AnswerShippingQuery calls Telegram method "answerShippingQuery".
// Doc: https://core.telegram.org/bots/api#answershippingquery
func (bot *Bot) AnswerShippingQuery(ctx context.Context, params *AnswerShippingQueryParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "answerShippingQuery", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// AnswerWebAppQueryParams contains params for Telegram method "answerWebAppQuery".
type AnswerWebAppQueryParams struct {
	WebAppQueryID string            `json:"web_app_query_id"`
	Result        InlineQueryResult `json:"result"`
}

// AnswerWebAppQuery calls Telegram method "answerWebAppQuery".
// Doc: https://core.telegram.org/bots/api#answerwebappquery
func (bot *Bot) AnswerWebAppQuery(ctx context.Context, params *AnswerWebAppQueryParams) (SentWebAppMessage, error) {
	var result SentWebAppMessage
	if err := bot.call(ctx, "answerWebAppQuery", params, &result); err != nil {
		return SentWebAppMessage{}, err
	}
	return result, nil
}

// ApproveChatJoinRequestParams contains params for Telegram method "approveChatJoinRequest".
type ApproveChatJoinRequestParams struct {
	ChatID any   `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

// ApproveChatJoinRequest calls Telegram method "approveChatJoinRequest".
// Doc: https://core.telegram.org/bots/api#approvechatjoinrequest
func (bot *Bot) ApproveChatJoinRequest(ctx context.Context, params *ApproveChatJoinRequestParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "approveChatJoinRequest", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// ApproveSuggestedPostParams contains params for Telegram method "approveSuggestedPost".
type ApproveSuggestedPostParams struct {
	ChatID    int64 `json:"chat_id"`
	MessageID int64 `json:"message_id"`
	SendDate  int64 `json:"send_date,omitempty"`
}

// ApproveSuggestedPost calls Telegram method "approveSuggestedPost".
// Doc: https://core.telegram.org/bots/api#approvesuggestedpost
func (bot *Bot) ApproveSuggestedPost(ctx context.Context, params *ApproveSuggestedPostParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "approveSuggestedPost", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// BanChatMemberParams contains params for Telegram method "banChatMember".
type BanChatMemberParams struct {
	ChatID         any   `json:"chat_id"`
	UserID         int64 `json:"user_id"`
	UntilDate      int64 `json:"until_date,omitempty"`
	RevokeMessages bool  `json:"revoke_messages,omitempty"`
}

// BanChatMember calls Telegram method "banChatMember".
// Doc: https://core.telegram.org/bots/api#banchatmember
func (bot *Bot) BanChatMember(ctx context.Context, params *BanChatMemberParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "banChatMember", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// BanChatSenderChatParams contains params for Telegram method "banChatSenderChat".
type BanChatSenderChatParams struct {
	ChatID       any   `json:"chat_id"`
	SenderChatID int64 `json:"sender_chat_id"`
}

// BanChatSenderChat calls Telegram method "banChatSenderChat".
// Doc: https://core.telegram.org/bots/api#banchatsenderchat
func (bot *Bot) BanChatSenderChat(ctx context.Context, params *BanChatSenderChatParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "banChatSenderChat", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// CloseParams contains params for Telegram method "close".
type CloseParams struct {
}

// Close calls Telegram method "close".
// Doc: https://core.telegram.org/bots/api#close
func (bot *Bot) Close(ctx context.Context, params *CloseParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "close", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// CloseForumTopicParams contains params for Telegram method "closeForumTopic".
type CloseForumTopicParams struct {
	ChatID          any   `json:"chat_id"`
	MessageThreadID int64 `json:"message_thread_id"`
}

// CloseForumTopic calls Telegram method "closeForumTopic".
// Doc: https://core.telegram.org/bots/api#closeforumtopic
func (bot *Bot) CloseForumTopic(ctx context.Context, params *CloseForumTopicParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "closeForumTopic", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// CloseGeneralForumTopicParams contains params for Telegram method "closeGeneralForumTopic".
type CloseGeneralForumTopicParams struct {
	ChatID any `json:"chat_id"`
}

// CloseGeneralForumTopic calls Telegram method "closeGeneralForumTopic".
// Doc: https://core.telegram.org/bots/api#closegeneralforumtopic
func (bot *Bot) CloseGeneralForumTopic(ctx context.Context, params *CloseGeneralForumTopicParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "closeGeneralForumTopic", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// ConvertGiftToStarsParams contains params for Telegram method "convertGiftToStars".
type ConvertGiftToStarsParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	OwnedGiftID          string `json:"owned_gift_id"`
}

// ConvertGiftToStars calls Telegram method "convertGiftToStars".
// Doc: https://core.telegram.org/bots/api#convertgifttostars
func (bot *Bot) ConvertGiftToStars(ctx context.Context, params *ConvertGiftToStarsParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "convertGiftToStars", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// CopyMessageParams contains params for Telegram method "copyMessage".
type CopyMessageParams struct {
	ChatID                  any                     `json:"chat_id"`
	MessageThreadID         int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int64                   `json:"direct_messages_topic_id,omitempty"`
	FromChatID              any                     `json:"from_chat_id"`
	MessageID               int64                   `json:"message_id"`
	VideoStartTimestamp     int64                   `json:"video_start_timestamp,omitempty"`
	Caption                 string                  `json:"caption,omitempty"`
	ParseMode               string                  `json:"parse_mode,omitempty"`
	CaptionEntities         []MessageEntity         `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                    `json:"show_caption_above_media,omitempty"`
	DisableNotification     bool                    `json:"disable_notification,omitempty"`
	ProtectContent          bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                     `json:"reply_markup,omitempty"`
}

// CopyMessage calls Telegram method "copyMessage".
// Doc: https://core.telegram.org/bots/api#copymessage
func (bot *Bot) CopyMessage(ctx context.Context, params *CopyMessageParams) (MessageId, error) {
	var result MessageId
	if err := bot.call(ctx, "copyMessage", params, &result); err != nil {
		return MessageId{}, err
	}
	return result, nil
}

// CopyMessagesParams contains params for Telegram method "copyMessages".
type CopyMessagesParams struct {
	ChatID                any     `json:"chat_id"`
	MessageThreadID       int64   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID int64   `json:"direct_messages_topic_id,omitempty"`
	FromChatID            any     `json:"from_chat_id"`
	MessageIds            []int64 `json:"message_ids"`
	DisableNotification   bool    `json:"disable_notification,omitempty"`
	ProtectContent        bool    `json:"protect_content,omitempty"`
	RemoveCaption         bool    `json:"remove_caption,omitempty"`
}

// CopyMessages calls Telegram method "copyMessages".
// Doc: https://core.telegram.org/bots/api#copymessages
func (bot *Bot) CopyMessages(ctx context.Context, params *CopyMessagesParams) (MessageId, error) {
	var result MessageId
	if err := bot.call(ctx, "copyMessages", params, &result); err != nil {
		return MessageId{}, err
	}
	return result, nil
}

// CreateChatInviteLinkParams contains params for Telegram method "createChatInviteLink".
type CreateChatInviteLinkParams struct {
	ChatID             any    `json:"chat_id"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int64  `json:"expire_date,omitempty"`
	MemberLimit        int64  `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

// CreateChatInviteLink calls Telegram method "createChatInviteLink".
// Doc: https://core.telegram.org/bots/api#createchatinvitelink
func (bot *Bot) CreateChatInviteLink(ctx context.Context, params *CreateChatInviteLinkParams) (ChatInviteLink, error) {
	var result ChatInviteLink
	if err := bot.call(ctx, "createChatInviteLink", params, &result); err != nil {
		return ChatInviteLink{}, err
	}
	return result, nil
}

// CreateChatSubscriptionInviteLinkParams contains params for Telegram method "createChatSubscriptionInviteLink".
type CreateChatSubscriptionInviteLinkParams struct {
	ChatID             any    `json:"chat_id"`
	Name               string `json:"name,omitempty"`
	SubscriptionPeriod int64  `json:"subscription_period"`
	SubscriptionPrice  int64  `json:"subscription_price"`
}

// CreateChatSubscriptionInviteLink calls Telegram method "createChatSubscriptionInviteLink".
// Doc: https://core.telegram.org/bots/api#createchatsubscriptioninvitelink
func (bot *Bot) CreateChatSubscriptionInviteLink(ctx context.Context, params *CreateChatSubscriptionInviteLinkParams) (ChatInviteLink, error) {
	var result ChatInviteLink
	if err := bot.call(ctx, "createChatSubscriptionInviteLink", params, &result); err != nil {
		return ChatInviteLink{}, err
	}
	return result, nil
}

// CreateForumTopicParams contains params for Telegram method "createForumTopic".
type CreateForumTopicParams struct {
	ChatID            any    `json:"chat_id"`
	Name              string `json:"name"`
	IconColor         int64  `json:"icon_color,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// CreateForumTopic calls Telegram method "createForumTopic".
// Doc: https://core.telegram.org/bots/api#createforumtopic
func (bot *Bot) CreateForumTopic(ctx context.Context, params *CreateForumTopicParams) (ForumTopic, error) {
	var result ForumTopic
	if err := bot.call(ctx, "createForumTopic", params, &result); err != nil {
		return ForumTopic{}, err
	}
	return result, nil
}

// CreateInvoiceLinkParams contains params for Telegram method "createInvoiceLink".
type CreateInvoiceLinkParams struct {
	BusinessConnectionID      string         `json:"business_connection_id,omitempty"`
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             string         `json:"provider_token,omitempty"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	SubscriptionPeriod        int64          `json:"subscription_period,omitempty"`
	MaxTipAmount              int64          `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int64        `json:"suggested_tip_amounts,omitempty"`
	ProviderData              string         `json:"provider_data,omitempty"`
	PhotoURL                  string         `json:"photo_url,omitempty"`
	PhotoSize                 int64          `json:"photo_size,omitempty"`
	PhotoWidth                int64          `json:"photo_width,omitempty"`
	PhotoHeight               int64          `json:"photo_height,omitempty"`
	NeedName                  bool           `json:"need_name,omitempty"`
	NeedPhoneNumber           bool           `json:"need_phone_number,omitempty"`
	NeedEmail                 bool           `json:"need_email,omitempty"`
	NeedShippingAddress       bool           `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool           `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool           `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool           `json:"is_flexible,omitempty"`
}

// CreateInvoiceLink calls Telegram method "createInvoiceLink".
// Doc: https://core.telegram.org/bots/api#createinvoicelink
func (bot *Bot) CreateInvoiceLink(ctx context.Context, params *CreateInvoiceLinkParams) (string, error) {
	var result string
	if err := bot.call(ctx, "createInvoiceLink", params, &result); err != nil {
		return "", err
	}
	return result, nil
}

// CreateNewStickerSetParams contains params for Telegram method "createNewStickerSet".
type CreateNewStickerSetParams struct {
	UserID          int64          `json:"user_id"`
	Name            string         `json:"name"`
	Title           string         `json:"title"`
	Stickers        []InputSticker `json:"stickers"`
	StickerType     string         `json:"sticker_type,omitempty"`
	NeedsRepainting bool           `json:"needs_repainting,omitempty"`
}

// CreateNewStickerSet calls Telegram method "createNewStickerSet".
// Doc: https://core.telegram.org/bots/api#createnewstickerset
func (bot *Bot) CreateNewStickerSet(ctx context.Context, params *CreateNewStickerSetParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "createNewStickerSet", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// DeclineChatJoinRequestParams contains params for Telegram method "declineChatJoinRequest".
type DeclineChatJoinRequestParams struct {
	ChatID any   `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

// DeclineChatJoinRequest calls Telegram method "declineChatJoinRequest".
// Doc: https://core.telegram.org/bots/api#declinechatjoinrequest
func (bot *Bot) DeclineChatJoinRequest(ctx context.Context, params *DeclineChatJoinRequestParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "declineChatJoinRequest", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// DeclineSuggestedPostParams contains params for Telegram method "declineSuggestedPost".
type DeclineSuggestedPostParams struct {
	ChatID    int64  `json:"chat_id"`
	MessageID int64  `json:"message_id"`
	Comment   string `json:"comment,omitempty"`
}

// DeclineSuggestedPost calls Telegram method "declineSuggestedPost".
// Doc: https://core.telegram.org/bots/api#declinesuggestedpost
func (bot *Bot) DeclineSuggestedPost(ctx context.Context, params *DeclineSuggestedPostParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "declineSuggestedPost", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// DeleteBusinessMessagesParams contains params for Telegram method "deleteBusinessMessages".
type DeleteBusinessMessagesParams struct {
	BusinessConnectionID string  `json:"business_connection_id"`
	MessageIds           []int64 `json:"message_ids"`
}

// DeleteBusinessMessages calls Telegram method "deleteBusinessMessages".
// Doc: https://core.telegram.org/bots/api#deletebusinessmessages
func (bot *Bot) DeleteBusinessMessages(ctx context.Context, params *DeleteBusinessMessagesParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "deleteBusinessMessages", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// DeleteChatPhotoParams contains params for Telegram method "deleteChatPhoto".
type DeleteChatPhotoParams struct {
	ChatID any `json:"chat_id"`
}

// DeleteChatPhoto calls Telegram method "deleteChatPhoto".
// Doc: https://core.telegram.org/bots/api#deletechatphoto
func (bot *Bot) DeleteChatPhoto(ctx context.Context, params *DeleteChatPhotoParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "deleteChatPhoto", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// DeleteChatStickerSetParams contains params for Telegram method "deleteChatStickerSet".
type DeleteChatStickerSetParams struct {
	ChatID any `json:"chat_id"`
}

// DeleteChatStickerSet calls Telegram method "deleteChatStickerSet".
// Doc: https://core.telegram.org/bots/api#deletechatstickerset
func (bot *Bot) DeleteChatStickerSet(ctx context.Context, params *DeleteChatStickerSetParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "deleteChatStickerSet", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// DeleteForumTopicParams contains params for Telegram method "deleteForumTopic".
type DeleteForumTopicParams struct {
	ChatID          any   `json:"chat_id"`
	MessageThreadID int64 `json:"message_thread_id"`
}

// DeleteForumTopic calls Telegram method "deleteForumTopic".
// Doc: https://core.telegram.org/bots/api#deleteforumtopic
func (bot *Bot) DeleteForumTopic(ctx context.Context, params *DeleteForumTopicParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "deleteForumTopic", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// DeleteMessageParams contains params for Telegram method "deleteMessage".
type DeleteMessageParams struct {
	ChatID    any   `json:"chat_id"`
	MessageID int64 `json:"message_id"`
}

// DeleteMessage calls Telegram method "deleteMessage".
// Doc: https://core.telegram.org/bots/api#deletemessage
func (bot *Bot) DeleteMessage(ctx context.Context, params *DeleteMessageParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "deleteMessage", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// DeleteMessagesParams contains params for Telegram method "deleteMessages".
type DeleteMessagesParams struct {
	ChatID     any     `json:"chat_id"`
	MessageIds []int64 `json:"message_ids"`
}

// DeleteMessages calls Telegram method "deleteMessages".
// Doc: https://core.telegram.org/bots/api#deletemessages
func (bot *Bot) DeleteMessages(ctx context.Context, params *DeleteMessagesParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "deleteMessages", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// DeleteMyCommandsParams contains params for Telegram method "deleteMyCommands".
type DeleteMyCommandsParams struct {
	Scope        BotCommandScope `json:"scope,omitempty"`
	LanguageCode string          `json:"language_code,omitempty"`
}

// DeleteMyCommands calls Telegram method "deleteMyCommands".
// Doc: https://core.telegram.org/bots/api#deletemycommands
func (bot *Bot) DeleteMyCommands(ctx context.Context, params *DeleteMyCommandsParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "deleteMyCommands", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// DeleteStickerFromSetParams contains params for Telegram method "deleteStickerFromSet".
type DeleteStickerFromSetParams struct {
	Sticker string `json:"sticker"`
}

// DeleteStickerFromSet calls Telegram method "deleteStickerFromSet".
// Doc: https://core.telegram.org/bots/api#deletestickerfromset
func (bot *Bot) DeleteStickerFromSet(ctx context.Context, params *DeleteStickerFromSetParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "deleteStickerFromSet", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// DeleteStickerSetParams contains params for Telegram method "deleteStickerSet".
type DeleteStickerSetParams struct {
	Name string `json:"name"`
}

// DeleteStickerSet calls Telegram method "deleteStickerSet".
// Doc: https://core.telegram.org/bots/api#deletestickerset
func (bot *Bot) DeleteStickerSet(ctx context.Context, params *DeleteStickerSetParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "deleteStickerSet", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// DeleteStoryParams contains params for Telegram method "deleteStory".
type DeleteStoryParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	StoryID              int64  `json:"story_id"`
}

// DeleteStory calls Telegram method "deleteStory".
// Doc: https://core.telegram.org/bots/api#deletestory
func (bot *Bot) DeleteStory(ctx context.Context, params *DeleteStoryParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "deleteStory", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// DeleteWebhookParams contains params for Telegram method "deleteWebhook".
type DeleteWebhookParams struct {
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

// DeleteWebhook calls Telegram method "deleteWebhook".
// Doc: https://core.telegram.org/bots/api#deletewebhook
func (bot *Bot) DeleteWebhook(ctx context.Context, params *DeleteWebhookParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "deleteWebhook", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// EditChatInviteLinkParams contains params for Telegram method "editChatInviteLink".
type EditChatInviteLinkParams struct {
	ChatID             any    `json:"chat_id"`
	InviteLink         string `json:"invite_link"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int64  `json:"expire_date,omitempty"`
	MemberLimit        int64  `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

// EditChatInviteLink calls Telegram method "editChatInviteLink".
// Doc: https://core.telegram.org/bots/api#editchatinvitelink
func (bot *Bot) EditChatInviteLink(ctx context.Context, params *EditChatInviteLinkParams) (ChatInviteLink, error) {
	var result ChatInviteLink
	if err := bot.call(ctx, "editChatInviteLink", params, &result); err != nil {
		return ChatInviteLink{}, err
	}
	return result, nil
}

// EditChatSubscriptionInviteLinkParams contains params for Telegram method "editChatSubscriptionInviteLink".
type EditChatSubscriptionInviteLinkParams struct {
	ChatID     any    `json:"chat_id"`
	InviteLink string `json:"invite_link"`
	Name       string `json:"name,omitempty"`
}

// EditChatSubscriptionInviteLink calls Telegram method "editChatSubscriptionInviteLink".
// Doc: https://core.telegram.org/bots/api#editchatsubscriptioninvitelink
func (bot *Bot) EditChatSubscriptionInviteLink(ctx context.Context, params *EditChatSubscriptionInviteLinkParams) (ChatInviteLink, error) {
	var result ChatInviteLink
	if err := bot.call(ctx, "editChatSubscriptionInviteLink", params, &result); err != nil {
		return ChatInviteLink{}, err
	}
	return result, nil
}

// EditForumTopicParams contains params for Telegram method "editForumTopic".
type EditForumTopicParams struct {
	ChatID            any    `json:"chat_id"`
	MessageThreadID   int64  `json:"message_thread_id"`
	Name              string `json:"name,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// EditForumTopic calls Telegram method "editForumTopic".
// Doc: https://core.telegram.org/bots/api#editforumtopic
func (bot *Bot) EditForumTopic(ctx context.Context, params *EditForumTopicParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "editForumTopic", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// EditGeneralForumTopicParams contains params for Telegram method "editGeneralForumTopic".
type EditGeneralForumTopicParams struct {
	ChatID any    `json:"chat_id"`
	Name   string `json:"name"`
}

// EditGeneralForumTopic calls Telegram method "editGeneralForumTopic".
// Doc: https://core.telegram.org/bots/api#editgeneralforumtopic
func (bot *Bot) EditGeneralForumTopic(ctx context.Context, params *EditGeneralForumTopicParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "editGeneralForumTopic", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// EditMessageCaptionParams contains params for Telegram method "editMessageCaption".
type EditMessageCaptionParams struct {
	BusinessConnectionID  string               `json:"business_connection_id,omitempty"`
	ChatID                any                  `json:"chat_id,omitempty"`
	MessageID             int64                `json:"message_id,omitempty"`
	InlineMessageID       string               `json:"inline_message_id,omitempty"`
	Caption               string               `json:"caption,omitempty"`
	ParseMode             string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity      `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageCaption calls Telegram method "editMessageCaption".
// Doc: https://core.telegram.org/bots/api#editmessagecaption
func (bot *Bot) EditMessageCaption(ctx context.Context, params *EditMessageCaptionParams) (any, error) {
	var result any
	if err := bot.call(ctx, "editMessageCaption", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// EditMessageChecklistParams contains params for Telegram method "editMessageChecklist".
type EditMessageChecklistParams struct {
	BusinessConnectionID string               `json:"business_connection_id"`
	ChatID               int64                `json:"chat_id"`
	MessageID            int64                `json:"message_id"`
	Checklist            InputChecklist       `json:"checklist"`
	ReplyMarkup          InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageChecklist calls Telegram method "editMessageChecklist".
// Doc: https://core.telegram.org/bots/api#editmessagechecklist
func (bot *Bot) EditMessageChecklist(ctx context.Context, params *EditMessageChecklistParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "editMessageChecklist", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// EditMessageLiveLocationParams contains params for Telegram method "editMessageLiveLocation".
type EditMessageLiveLocationParams struct {
	BusinessConnectionID string               `json:"business_connection_id,omitempty"`
	ChatID               any                  `json:"chat_id,omitempty"`
	MessageID            int64                `json:"message_id,omitempty"`
	InlineMessageID      string               `json:"inline_message_id,omitempty"`
	Latitude             float64              `json:"latitude"`
	Longitude            float64              `json:"longitude"`
	LivePeriod           int64                `json:"live_period,omitempty"`
	HorizontalAccuracy   float64              `json:"horizontal_accuracy,omitempty"`
	Heading              int64                `json:"heading,omitempty"`
	ProximityAlertRadius int64                `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageLiveLocation calls Telegram method "editMessageLiveLocation".
// Doc: https://core.telegram.org/bots/api#editmessagelivelocation
func (bot *Bot) EditMessageLiveLocation(ctx context.Context, params *EditMessageLiveLocationParams) (any, error) {
	var result any
	if err := bot.call(ctx, "editMessageLiveLocation", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// EditMessageMediaParams contains params for Telegram method "editMessageMedia".
type EditMessageMediaParams struct {
	BusinessConnectionID string               `json:"business_connection_id,omitempty"`
	ChatID               any                  `json:"chat_id,omitempty"`
	MessageID            int64                `json:"message_id,omitempty"`
	InlineMessageID      string               `json:"inline_message_id,omitempty"`
	Media                InputMedia           `json:"media"`
	ReplyMarkup          InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageMedia calls Telegram method "editMessageMedia".
// Doc: https://core.telegram.org/bots/api#editmessagemedia
func (bot *Bot) EditMessageMedia(ctx context.Context, params *EditMessageMediaParams) (any, error) {
	var result any
	if err := bot.call(ctx, "editMessageMedia", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// EditMessageReplyMarkupParams contains params for Telegram method "editMessageReplyMarkup".
type EditMessageReplyMarkupParams struct {
	BusinessConnectionID string               `json:"business_connection_id,omitempty"`
	ChatID               any                  `json:"chat_id,omitempty"`
	MessageID            int64                `json:"message_id,omitempty"`
	InlineMessageID      string               `json:"inline_message_id,omitempty"`
	ReplyMarkup          InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageReplyMarkup calls Telegram method "editMessageReplyMarkup".
// Doc: https://core.telegram.org/bots/api#editmessagereplymarkup
func (bot *Bot) EditMessageReplyMarkup(ctx context.Context, params *EditMessageReplyMarkupParams) (any, error) {
	var result any
	if err := bot.call(ctx, "editMessageReplyMarkup", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// EditMessageTextParams contains params for Telegram method "editMessageText".
type EditMessageTextParams struct {
	BusinessConnectionID string               `json:"business_connection_id,omitempty"`
	ChatID               any                  `json:"chat_id,omitempty"`
	MessageID            int64                `json:"message_id,omitempty"`
	InlineMessageID      string               `json:"inline_message_id,omitempty"`
	Text                 string               `json:"text"`
	ParseMode            string               `json:"parse_mode,omitempty"`
	Entities             []MessageEntity      `json:"entities,omitempty"`
	LinkPreviewOptions   LinkPreviewOptions   `json:"link_preview_options,omitempty"`
	ReplyMarkup          InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageText calls Telegram method "editMessageText".
// Doc: https://core.telegram.org/bots/api#editmessagetext
func (bot *Bot) EditMessageText(ctx context.Context, params *EditMessageTextParams) (any, error) {
	var result any
	if err := bot.call(ctx, "editMessageText", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// EditStoryParams contains params for Telegram method "editStory".
type EditStoryParams struct {
	BusinessConnectionID string            `json:"business_connection_id"`
	StoryID              int64             `json:"story_id"`
	Content              InputStoryContent `json:"content"`
	Caption              string            `json:"caption,omitempty"`
	ParseMode            string            `json:"parse_mode,omitempty"`
	CaptionEntities      []MessageEntity   `json:"caption_entities,omitempty"`
	Areas                []StoryArea       `json:"areas,omitempty"`
}

// EditStory calls Telegram method "editStory".
// Doc: https://core.telegram.org/bots/api#editstory
func (bot *Bot) EditStory(ctx context.Context, params *EditStoryParams) (Story, error) {
	var result Story
	if err := bot.call(ctx, "editStory", params, &result); err != nil {
		return Story{}, err
	}
	return result, nil
}

// EditUserStarSubscriptionParams contains params for Telegram method "editUserStarSubscription".
type EditUserStarSubscriptionParams struct {
	UserID                  int64  `json:"user_id"`
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
	IsCanceled              bool   `json:"is_canceled"`
}

// EditUserStarSubscription calls Telegram method "editUserStarSubscription".
// Doc: https://core.telegram.org/bots/api#edituserstarsubscription
func (bot *Bot) EditUserStarSubscription(ctx context.Context, params *EditUserStarSubscriptionParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "editUserStarSubscription", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// ExportChatInviteLinkParams contains params for Telegram method "exportChatInviteLink".
type ExportChatInviteLinkParams struct {
	ChatID any `json:"chat_id"`
}

// ExportChatInviteLink calls Telegram method "exportChatInviteLink".
// Doc: https://core.telegram.org/bots/api#exportchatinvitelink
func (bot *Bot) ExportChatInviteLink(ctx context.Context, params *ExportChatInviteLinkParams) (string, error) {
	var result string
	if err := bot.call(ctx, "exportChatInviteLink", params, &result); err != nil {
		return "", err
	}
	return result, nil
}

// ForwardMessageParams contains params for Telegram method "forwardMessage".
type ForwardMessageParams struct {
	ChatID                  any                     `json:"chat_id"`
	MessageThreadID         int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int64                   `json:"direct_messages_topic_id,omitempty"`
	FromChatID              any                     `json:"from_chat_id"`
	VideoStartTimestamp     int64                   `json:"video_start_timestamp,omitempty"`
	DisableNotification     bool                    `json:"disable_notification,omitempty"`
	ProtectContent          bool                    `json:"protect_content,omitempty"`
	MessageEffectID         string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	MessageID               int64                   `json:"message_id"`
}

// ForwardMessage calls Telegram method "forwardMessage".
// Doc: https://core.telegram.org/bots/api#forwardmessage
func (bot *Bot) ForwardMessage(ctx context.Context, params *ForwardMessageParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "forwardMessage", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// ForwardMessagesParams contains params for Telegram method "forwardMessages".
type ForwardMessagesParams struct {
	ChatID                any     `json:"chat_id"`
	MessageThreadID       int64   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID int64   `json:"direct_messages_topic_id,omitempty"`
	FromChatID            any     `json:"from_chat_id"`
	MessageIds            []int64 `json:"message_ids"`
	DisableNotification   bool    `json:"disable_notification,omitempty"`
	ProtectContent        bool    `json:"protect_content,omitempty"`
}

// ForwardMessages calls Telegram method "forwardMessages".
// Doc: https://core.telegram.org/bots/api#forwardmessages
func (bot *Bot) ForwardMessages(ctx context.Context, params *ForwardMessagesParams) (MessageId, error) {
	var result MessageId
	if err := bot.call(ctx, "forwardMessages", params, &result); err != nil {
		return MessageId{}, err
	}
	return result, nil
}

// GetAvailableGiftsParams contains params for Telegram method "getAvailableGifts".
type GetAvailableGiftsParams struct {
}

// GetAvailableGifts calls Telegram method "getAvailableGifts".
// Doc: https://core.telegram.org/bots/api#getavailablegifts
func (bot *Bot) GetAvailableGifts(ctx context.Context, params *GetAvailableGiftsParams) (Gifts, error) {
	var result Gifts
	if err := bot.call(ctx, "getAvailableGifts", params, &result); err != nil {
		return Gifts{}, err
	}
	return result, nil
}

// GetBusinessAccountGiftsParams contains params for Telegram method "getBusinessAccountGifts".
type GetBusinessAccountGiftsParams struct {
	BusinessConnectionID        string `json:"business_connection_id"`
	ExcludeUnsaved              bool   `json:"exclude_unsaved,omitempty"`
	ExcludeSaved                bool   `json:"exclude_saved,omitempty"`
	ExcludeUnlimited            bool   `json:"exclude_unlimited,omitempty"`
	ExcludeLimitedUpgradable    bool   `json:"exclude_limited_upgradable,omitempty"`
	ExcludeLimitedNonUpgradable bool   `json:"exclude_limited_non_upgradable,omitempty"`
	ExcludeUnique               bool   `json:"exclude_unique,omitempty"`
	ExcludeFromBlockchain       bool   `json:"exclude_from_blockchain,omitempty"`
	SortByPrice                 bool   `json:"sort_by_price,omitempty"`
	Offset                      string `json:"offset,omitempty"`
	Limit                       int64  `json:"limit,omitempty"`
}

// GetBusinessAccountGifts calls Telegram method "getBusinessAccountGifts".
// Doc: https://core.telegram.org/bots/api#getbusinessaccountgifts
func (bot *Bot) GetBusinessAccountGifts(ctx context.Context, params *GetBusinessAccountGiftsParams) (OwnedGifts, error) {
	var result OwnedGifts
	if err := bot.call(ctx, "getBusinessAccountGifts", params, &result); err != nil {
		return OwnedGifts{}, err
	}
	return result, nil
}

// GetBusinessAccountStarBalanceParams contains params for Telegram method "getBusinessAccountStarBalance".
type GetBusinessAccountStarBalanceParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
}

// GetBusinessAccountStarBalance calls Telegram method "getBusinessAccountStarBalance".
// Doc: https://core.telegram.org/bots/api#getbusinessaccountstarbalance
func (bot *Bot) GetBusinessAccountStarBalance(ctx context.Context, params *GetBusinessAccountStarBalanceParams) (StarAmount, error) {
	var result StarAmount
	if err := bot.call(ctx, "getBusinessAccountStarBalance", params, &result); err != nil {
		return StarAmount{}, err
	}
	return result, nil
}

// GetBusinessConnectionParams contains params for Telegram method "getBusinessConnection".
type GetBusinessConnectionParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
}

// GetBusinessConnection calls Telegram method "getBusinessConnection".
// Doc: https://core.telegram.org/bots/api#getbusinessconnection
func (bot *Bot) GetBusinessConnection(ctx context.Context, params *GetBusinessConnectionParams) (BusinessConnection, error) {
	var result BusinessConnection
	if err := bot.call(ctx, "getBusinessConnection", params, &result); err != nil {
		return BusinessConnection{}, err
	}
	return result, nil
}

// GetChatParams contains params for Telegram method "getChat".
type GetChatParams struct {
	ChatID any `json:"chat_id"`
}

// GetChat calls Telegram method "getChat".
// Doc: https://core.telegram.org/bots/api#getchat
func (bot *Bot) GetChat(ctx context.Context, params *GetChatParams) (ChatFullInfo, error) {
	var result ChatFullInfo
	if err := bot.call(ctx, "getChat", params, &result); err != nil {
		return ChatFullInfo{}, err
	}
	return result, nil
}

// GetChatAdministratorsParams contains params for Telegram method "getChatAdministrators".
type GetChatAdministratorsParams struct {
	ChatID any `json:"chat_id"`
}

// GetChatAdministrators calls Telegram method "getChatAdministrators".
// Doc: https://core.telegram.org/bots/api#getchatadministrators
func (bot *Bot) GetChatAdministrators(ctx context.Context, params *GetChatAdministratorsParams) ([]ChatMember, error) {
	var result []ChatMember
	if err := bot.call(ctx, "getChatAdministrators", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetChatGiftsParams contains params for Telegram method "getChatGifts".
type GetChatGiftsParams struct {
	ChatID                      any    `json:"chat_id"`
	ExcludeUnsaved              bool   `json:"exclude_unsaved,omitempty"`
	ExcludeSaved                bool   `json:"exclude_saved,omitempty"`
	ExcludeUnlimited            bool   `json:"exclude_unlimited,omitempty"`
	ExcludeLimitedUpgradable    bool   `json:"exclude_limited_upgradable,omitempty"`
	ExcludeLimitedNonUpgradable bool   `json:"exclude_limited_non_upgradable,omitempty"`
	ExcludeFromBlockchain       bool   `json:"exclude_from_blockchain,omitempty"`
	ExcludeUnique               bool   `json:"exclude_unique,omitempty"`
	SortByPrice                 bool   `json:"sort_by_price,omitempty"`
	Offset                      string `json:"offset,omitempty"`
	Limit                       int64  `json:"limit,omitempty"`
}

// GetChatGifts calls Telegram method "getChatGifts".
// Doc: https://core.telegram.org/bots/api#getchatgifts
func (bot *Bot) GetChatGifts(ctx context.Context, params *GetChatGiftsParams) (OwnedGifts, error) {
	var result OwnedGifts
	if err := bot.call(ctx, "getChatGifts", params, &result); err != nil {
		return OwnedGifts{}, err
	}
	return result, nil
}

// GetChatMemberParams contains params for Telegram method "getChatMember".
type GetChatMemberParams struct {
	ChatID any   `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

// GetChatMember calls Telegram method "getChatMember".
// Doc: https://core.telegram.org/bots/api#getchatmember
func (bot *Bot) GetChatMember(ctx context.Context, params *GetChatMemberParams) (ChatMember, error) {
	var result ChatMember
	if err := bot.call(ctx, "getChatMember", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetChatMemberCountParams contains params for Telegram method "getChatMemberCount".
type GetChatMemberCountParams struct {
	ChatID any `json:"chat_id"`
}

// GetChatMemberCount calls Telegram method "getChatMemberCount".
// Doc: https://core.telegram.org/bots/api#getchatmembercount
func (bot *Bot) GetChatMemberCount(ctx context.Context, params *GetChatMemberCountParams) (int64, error) {
	var result int64
	if err := bot.call(ctx, "getChatMemberCount", params, &result); err != nil {
		return 0, err
	}
	return result, nil
}

// GetChatMenuButtonParams contains params for Telegram method "getChatMenuButton".
type GetChatMenuButtonParams struct {
	ChatID int64 `json:"chat_id,omitempty"`
}

// GetChatMenuButton calls Telegram method "getChatMenuButton".
// Doc: https://core.telegram.org/bots/api#getchatmenubutton
func (bot *Bot) GetChatMenuButton(ctx context.Context, params *GetChatMenuButtonParams) (MenuButton, error) {
	var result MenuButton
	if err := bot.call(ctx, "getChatMenuButton", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetCustomEmojiStickersParams contains params for Telegram method "getCustomEmojiStickers".
type GetCustomEmojiStickersParams struct {
	CustomEmojiIds []string `json:"custom_emoji_ids"`
}

// GetCustomEmojiStickers calls Telegram method "getCustomEmojiStickers".
// Doc: https://core.telegram.org/bots/api#getcustomemojistickers
func (bot *Bot) GetCustomEmojiStickers(ctx context.Context, params *GetCustomEmojiStickersParams) ([]Sticker, error) {
	var result []Sticker
	if err := bot.call(ctx, "getCustomEmojiStickers", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetFileParams contains params for Telegram method "getFile".
type GetFileParams struct {
	FileID string `json:"file_id"`
}

// GetFile calls Telegram method "getFile".
// Doc: https://core.telegram.org/bots/api#getfile
func (bot *Bot) GetFile(ctx context.Context, params *GetFileParams) (File, error) {
	var result File
	if err := bot.call(ctx, "getFile", params, &result); err != nil {
		return File{}, err
	}
	return result, nil
}

// GetForumTopicIconStickersParams contains params for Telegram method "getForumTopicIconStickers".
type GetForumTopicIconStickersParams struct {
}

// GetForumTopicIconStickers calls Telegram method "getForumTopicIconStickers".
// Doc: https://core.telegram.org/bots/api#getforumtopiciconstickers
func (bot *Bot) GetForumTopicIconStickers(ctx context.Context, params *GetForumTopicIconStickersParams) ([]Sticker, error) {
	var result []Sticker
	if err := bot.call(ctx, "getForumTopicIconStickers", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetGameHighScoresParams contains params for Telegram method "getGameHighScores".
type GetGameHighScoresParams struct {
	UserID          int64  `json:"user_id"`
	ChatID          int64  `json:"chat_id,omitempty"`
	MessageID       int64  `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// GetGameHighScores calls Telegram method "getGameHighScores".
// Doc: https://core.telegram.org/bots/api#getgamehighscores
func (bot *Bot) GetGameHighScores(ctx context.Context, params *GetGameHighScoresParams) ([]GameHighScore, error) {
	var result []GameHighScore
	if err := bot.call(ctx, "getGameHighScores", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetMeParams contains params for Telegram method "getMe".
type GetMeParams struct {
}

// GetMe calls Telegram method "getMe".
// Doc: https://core.telegram.org/bots/api#getme
func (bot *Bot) GetMe(ctx context.Context, params *GetMeParams) (User, error) {
	var result User
	if err := bot.call(ctx, "getMe", params, &result); err != nil {
		return User{}, err
	}
	return result, nil
}

// GetMyCommandsParams contains params for Telegram method "getMyCommands".
type GetMyCommandsParams struct {
	Scope        BotCommandScope `json:"scope,omitempty"`
	LanguageCode string          `json:"language_code,omitempty"`
}

// GetMyCommands calls Telegram method "getMyCommands".
// Doc: https://core.telegram.org/bots/api#getmycommands
func (bot *Bot) GetMyCommands(ctx context.Context, params *GetMyCommandsParams) ([]BotCommand, error) {
	var result []BotCommand
	if err := bot.call(ctx, "getMyCommands", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetMyDefaultAdministratorRightsParams contains params for Telegram method "getMyDefaultAdministratorRights".
type GetMyDefaultAdministratorRightsParams struct {
	ForChannels bool `json:"for_channels,omitempty"`
}

// GetMyDefaultAdministratorRights calls Telegram method "getMyDefaultAdministratorRights".
// Doc: https://core.telegram.org/bots/api#getmydefaultadministratorrights
func (bot *Bot) GetMyDefaultAdministratorRights(ctx context.Context, params *GetMyDefaultAdministratorRightsParams) (ChatAdministratorRights, error) {
	var result ChatAdministratorRights
	if err := bot.call(ctx, "getMyDefaultAdministratorRights", params, &result); err != nil {
		return ChatAdministratorRights{}, err
	}
	return result, nil
}

// GetMyDescriptionParams contains params for Telegram method "getMyDescription".
type GetMyDescriptionParams struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// GetMyDescription calls Telegram method "getMyDescription".
// Doc: https://core.telegram.org/bots/api#getmydescription
func (bot *Bot) GetMyDescription(ctx context.Context, params *GetMyDescriptionParams) (BotDescription, error) {
	var result BotDescription
	if err := bot.call(ctx, "getMyDescription", params, &result); err != nil {
		return BotDescription{}, err
	}
	return result, nil
}

// GetMyNameParams contains params for Telegram method "getMyName".
type GetMyNameParams struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// GetMyName calls Telegram method "getMyName".
// Doc: https://core.telegram.org/bots/api#getmyname
func (bot *Bot) GetMyName(ctx context.Context, params *GetMyNameParams) (BotName, error) {
	var result BotName
	if err := bot.call(ctx, "getMyName", params, &result); err != nil {
		return BotName{}, err
	}
	return result, nil
}

// GetMyShortDescriptionParams contains params for Telegram method "getMyShortDescription".
type GetMyShortDescriptionParams struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// GetMyShortDescription calls Telegram method "getMyShortDescription".
// Doc: https://core.telegram.org/bots/api#getmyshortdescription
func (bot *Bot) GetMyShortDescription(ctx context.Context, params *GetMyShortDescriptionParams) (BotShortDescription, error) {
	var result BotShortDescription
	if err := bot.call(ctx, "getMyShortDescription", params, &result); err != nil {
		return BotShortDescription{}, err
	}
	return result, nil
}

// GetMyStarBalanceParams contains params for Telegram method "getMyStarBalance".
type GetMyStarBalanceParams struct {
}

// GetMyStarBalance calls Telegram method "getMyStarBalance".
// Doc: https://core.telegram.org/bots/api#getmystarbalance
func (bot *Bot) GetMyStarBalance(ctx context.Context, params *GetMyStarBalanceParams) (StarAmount, error) {
	var result StarAmount
	if err := bot.call(ctx, "getMyStarBalance", params, &result); err != nil {
		return StarAmount{}, err
	}
	return result, nil
}

// GetStarTransactionsParams contains params for Telegram method "getStarTransactions".
type GetStarTransactionsParams struct {
	Offset int64 `json:"offset,omitempty"`
	Limit  int64 `json:"limit,omitempty"`
}

// GetStarTransactions calls Telegram method "getStarTransactions".
// Doc: https://core.telegram.org/bots/api#getstartransactions
func (bot *Bot) GetStarTransactions(ctx context.Context, params *GetStarTransactionsParams) (StarTransactions, error) {
	var result StarTransactions
	if err := bot.call(ctx, "getStarTransactions", params, &result); err != nil {
		return StarTransactions{}, err
	}
	return result, nil
}

// GetStickerSetParams contains params for Telegram method "getStickerSet".
type GetStickerSetParams struct {
	Name string `json:"name"`
}

// GetStickerSet calls Telegram method "getStickerSet".
// Doc: https://core.telegram.org/bots/api#getstickerset
func (bot *Bot) GetStickerSet(ctx context.Context, params *GetStickerSetParams) (StickerSet, error) {
	var result StickerSet
	if err := bot.call(ctx, "getStickerSet", params, &result); err != nil {
		return StickerSet{}, err
	}
	return result, nil
}

// GetUpdatesParams contains params for Telegram method "getUpdates".
type GetUpdatesParams struct {
	Offset         int64    `json:"offset,omitempty"`
	Limit          int64    `json:"limit,omitempty"`
	Timeout        int64    `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// GetUpdates calls Telegram method "getUpdates".
// Doc: https://core.telegram.org/bots/api#getupdates
func (bot *Bot) GetUpdates(ctx context.Context, params *GetUpdatesParams) ([]Update, error) {
	var result []Update
	if err := bot.call(ctx, "getUpdates", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserChatBoostsParams contains params for Telegram method "getUserChatBoosts".
type GetUserChatBoostsParams struct {
	ChatID any   `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

// GetUserChatBoosts calls Telegram method "getUserChatBoosts".
// Doc: https://core.telegram.org/bots/api#getuserchatboosts
func (bot *Bot) GetUserChatBoosts(ctx context.Context, params *GetUserChatBoostsParams) (UserChatBoosts, error) {
	var result UserChatBoosts
	if err := bot.call(ctx, "getUserChatBoosts", params, &result); err != nil {
		return UserChatBoosts{}, err
	}
	return result, nil
}

// GetUserGiftsParams contains params for Telegram method "getUserGifts".
type GetUserGiftsParams struct {
	UserID                      int64  `json:"user_id"`
	ExcludeUnlimited            bool   `json:"exclude_unlimited,omitempty"`
	ExcludeLimitedUpgradable    bool   `json:"exclude_limited_upgradable,omitempty"`
	ExcludeLimitedNonUpgradable bool   `json:"exclude_limited_non_upgradable,omitempty"`
	ExcludeFromBlockchain       bool   `json:"exclude_from_blockchain,omitempty"`
	ExcludeUnique               bool   `json:"exclude_unique,omitempty"`
	SortByPrice                 bool   `json:"sort_by_price,omitempty"`
	Offset                      string `json:"offset,omitempty"`
	Limit                       int64  `json:"limit,omitempty"`
}

// GetUserGifts calls Telegram method "getUserGifts".
// Doc: https://core.telegram.org/bots/api#getusergifts
func (bot *Bot) GetUserGifts(ctx context.Context, params *GetUserGiftsParams) (OwnedGifts, error) {
	var result OwnedGifts
	if err := bot.call(ctx, "getUserGifts", params, &result); err != nil {
		return OwnedGifts{}, err
	}
	return result, nil
}

// GetUserProfileAudiosParams contains params for Telegram method "getUserProfileAudios".
type GetUserProfileAudiosParams struct {
	UserID int64 `json:"user_id"`
	Offset int64 `json:"offset,omitempty"`
	Limit  int64 `json:"limit,omitempty"`
}

// GetUserProfileAudios calls Telegram method "getUserProfileAudios".
// Doc: https://core.telegram.org/bots/api#getuserprofileaudios
func (bot *Bot) GetUserProfileAudios(ctx context.Context, params *GetUserProfileAudiosParams) (UserProfileAudios, error) {
	var result UserProfileAudios
	if err := bot.call(ctx, "getUserProfileAudios", params, &result); err != nil {
		return UserProfileAudios{}, err
	}
	return result, nil
}

// GetUserProfilePhotosParams contains params for Telegram method "getUserProfilePhotos".
type GetUserProfilePhotosParams struct {
	UserID int64 `json:"user_id"`
	Offset int64 `json:"offset,omitempty"`
	Limit  int64 `json:"limit,omitempty"`
}

// GetUserProfilePhotos calls Telegram method "getUserProfilePhotos".
// Doc: https://core.telegram.org/bots/api#getuserprofilephotos
func (bot *Bot) GetUserProfilePhotos(ctx context.Context, params *GetUserProfilePhotosParams) (UserProfilePhotos, error) {
	var result UserProfilePhotos
	if err := bot.call(ctx, "getUserProfilePhotos", params, &result); err != nil {
		return UserProfilePhotos{}, err
	}
	return result, nil
}

// GetWebhookInfoParams contains params for Telegram method "getWebhookInfo".
type GetWebhookInfoParams struct {
}

// GetWebhookInfo calls Telegram method "getWebhookInfo".
// Doc: https://core.telegram.org/bots/api#getwebhookinfo
func (bot *Bot) GetWebhookInfo(ctx context.Context, params *GetWebhookInfoParams) (WebhookInfo, error) {
	var result WebhookInfo
	if err := bot.call(ctx, "getWebhookInfo", params, &result); err != nil {
		return WebhookInfo{}, err
	}
	return result, nil
}

// GiftPremiumSubscriptionParams contains params for Telegram method "giftPremiumSubscription".
type GiftPremiumSubscriptionParams struct {
	UserID        int64           `json:"user_id"`
	MonthCount    int64           `json:"month_count"`
	StarCount     int64           `json:"star_count"`
	Text          string          `json:"text,omitempty"`
	TextParseMode string          `json:"text_parse_mode,omitempty"`
	TextEntities  []MessageEntity `json:"text_entities,omitempty"`
}

// GiftPremiumSubscription calls Telegram method "giftPremiumSubscription".
// Doc: https://core.telegram.org/bots/api#giftpremiumsubscription
func (bot *Bot) GiftPremiumSubscription(ctx context.Context, params *GiftPremiumSubscriptionParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "giftPremiumSubscription", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// HideGeneralForumTopicParams contains params for Telegram method "hideGeneralForumTopic".
type HideGeneralForumTopicParams struct {
	ChatID any `json:"chat_id"`
}

// HideGeneralForumTopic calls Telegram method "hideGeneralForumTopic".
// Doc: https://core.telegram.org/bots/api#hidegeneralforumtopic
func (bot *Bot) HideGeneralForumTopic(ctx context.Context, params *HideGeneralForumTopicParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "hideGeneralForumTopic", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// LeaveChatParams contains params for Telegram method "leaveChat".
type LeaveChatParams struct {
	ChatID any `json:"chat_id"`
}

// LeaveChat calls Telegram method "leaveChat".
// Doc: https://core.telegram.org/bots/api#leavechat
func (bot *Bot) LeaveChat(ctx context.Context, params *LeaveChatParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "leaveChat", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// LogOutParams contains params for Telegram method "logOut".
type LogOutParams struct {
}

// LogOut calls Telegram method "logOut".
// Doc: https://core.telegram.org/bots/api#logout
func (bot *Bot) LogOut(ctx context.Context, params *LogOutParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "logOut", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// PinChatMessageParams contains params for Telegram method "pinChatMessage".
type PinChatMessageParams struct {
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	ChatID               any    `json:"chat_id"`
	MessageID            int64  `json:"message_id"`
	DisableNotification  bool   `json:"disable_notification,omitempty"`
}

// PinChatMessage calls Telegram method "pinChatMessage".
// Doc: https://core.telegram.org/bots/api#pinchatmessage
func (bot *Bot) PinChatMessage(ctx context.Context, params *PinChatMessageParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "pinChatMessage", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// PostStoryParams contains params for Telegram method "postStory".
type PostStoryParams struct {
	BusinessConnectionID string            `json:"business_connection_id"`
	Content              InputStoryContent `json:"content"`
	ActivePeriod         int64             `json:"active_period"`
	Caption              string            `json:"caption,omitempty"`
	ParseMode            string            `json:"parse_mode,omitempty"`
	CaptionEntities      []MessageEntity   `json:"caption_entities,omitempty"`
	Areas                []StoryArea       `json:"areas,omitempty"`
	PostToChatPage       bool              `json:"post_to_chat_page,omitempty"`
	ProtectContent       bool              `json:"protect_content,omitempty"`
}

// PostStory calls Telegram method "postStory".
// Doc: https://core.telegram.org/bots/api#poststory
func (bot *Bot) PostStory(ctx context.Context, params *PostStoryParams) (Story, error) {
	var result Story
	if err := bot.call(ctx, "postStory", params, &result); err != nil {
		return Story{}, err
	}
	return result, nil
}

// PromoteChatMemberParams contains params for Telegram method "promoteChatMember".
type PromoteChatMemberParams struct {
	ChatID                  any   `json:"chat_id"`
	UserID                  int64 `json:"user_id"`
	IsAnonymous             bool  `json:"is_anonymous,omitempty"`
	CanManageChat           bool  `json:"can_manage_chat,omitempty"`
	CanDeleteMessages       bool  `json:"can_delete_messages,omitempty"`
	CanManageVideoChats     bool  `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers      bool  `json:"can_restrict_members,omitempty"`
	CanPromoteMembers       bool  `json:"can_promote_members,omitempty"`
	CanChangeInfo           bool  `json:"can_change_info,omitempty"`
	CanInviteUsers          bool  `json:"can_invite_users,omitempty"`
	CanPostStories          bool  `json:"can_post_stories,omitempty"`
	CanEditStories          bool  `json:"can_edit_stories,omitempty"`
	CanDeleteStories        bool  `json:"can_delete_stories,omitempty"`
	CanPostMessages         bool  `json:"can_post_messages,omitempty"`
	CanEditMessages         bool  `json:"can_edit_messages,omitempty"`
	CanPinMessages          bool  `json:"can_pin_messages,omitempty"`
	CanManageTopics         bool  `json:"can_manage_topics,omitempty"`
	CanManageDirectMessages bool  `json:"can_manage_direct_messages,omitempty"`
	CanManageTags           bool  `json:"can_manage_tags,omitempty"`
}

// PromoteChatMember calls Telegram method "promoteChatMember".
// Doc: https://core.telegram.org/bots/api#promotechatmember
func (bot *Bot) PromoteChatMember(ctx context.Context, params *PromoteChatMemberParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "promoteChatMember", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// ReadBusinessMessageParams contains params for Telegram method "readBusinessMessage".
type ReadBusinessMessageParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	ChatID               int64  `json:"chat_id"`
	MessageID            int64  `json:"message_id"`
}

// ReadBusinessMessage calls Telegram method "readBusinessMessage".
// Doc: https://core.telegram.org/bots/api#readbusinessmessage
func (bot *Bot) ReadBusinessMessage(ctx context.Context, params *ReadBusinessMessageParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "readBusinessMessage", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// RefundStarPaymentParams contains params for Telegram method "refundStarPayment".
type RefundStarPaymentParams struct {
	UserID                  int64  `json:"user_id"`
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
}

// RefundStarPayment calls Telegram method "refundStarPayment".
// Doc: https://core.telegram.org/bots/api#refundstarpayment
func (bot *Bot) RefundStarPayment(ctx context.Context, params *RefundStarPaymentParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "refundStarPayment", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// RemoveBusinessAccountProfilePhotoParams contains params for Telegram method "removeBusinessAccountProfilePhoto".
type RemoveBusinessAccountProfilePhotoParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	IsPublic             bool   `json:"is_public,omitempty"`
}

// RemoveBusinessAccountProfilePhoto calls Telegram method "removeBusinessAccountProfilePhoto".
// Doc: https://core.telegram.org/bots/api#removebusinessaccountprofilephoto
func (bot *Bot) RemoveBusinessAccountProfilePhoto(ctx context.Context, params *RemoveBusinessAccountProfilePhotoParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "removeBusinessAccountProfilePhoto", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// RemoveChatVerificationParams contains params for Telegram method "removeChatVerification".
type RemoveChatVerificationParams struct {
	ChatID any `json:"chat_id"`
}

// RemoveChatVerification calls Telegram method "removeChatVerification".
// Doc: https://core.telegram.org/bots/api#removechatverification
func (bot *Bot) RemoveChatVerification(ctx context.Context, params *RemoveChatVerificationParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "removeChatVerification", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// RemoveMyProfilePhotoParams contains params for Telegram method "removeMyProfilePhoto".
type RemoveMyProfilePhotoParams struct {
}

// RemoveMyProfilePhoto calls Telegram method "removeMyProfilePhoto".
// Doc: https://core.telegram.org/bots/api#removemyprofilephoto
func (bot *Bot) RemoveMyProfilePhoto(ctx context.Context, params *RemoveMyProfilePhotoParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "removeMyProfilePhoto", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// RemoveUserVerificationParams contains params for Telegram method "removeUserVerification".
type RemoveUserVerificationParams struct {
	UserID int64 `json:"user_id"`
}

// RemoveUserVerification calls Telegram method "removeUserVerification".
// Doc: https://core.telegram.org/bots/api#removeuserverification
func (bot *Bot) RemoveUserVerification(ctx context.Context, params *RemoveUserVerificationParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "removeUserVerification", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// ReopenForumTopicParams contains params for Telegram method "reopenForumTopic".
type ReopenForumTopicParams struct {
	ChatID          any   `json:"chat_id"`
	MessageThreadID int64 `json:"message_thread_id"`
}

// ReopenForumTopic calls Telegram method "reopenForumTopic".
// Doc: https://core.telegram.org/bots/api#reopenforumtopic
func (bot *Bot) ReopenForumTopic(ctx context.Context, params *ReopenForumTopicParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "reopenForumTopic", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// ReopenGeneralForumTopicParams contains params for Telegram method "reopenGeneralForumTopic".
type ReopenGeneralForumTopicParams struct {
	ChatID any `json:"chat_id"`
}

// ReopenGeneralForumTopic calls Telegram method "reopenGeneralForumTopic".
// Doc: https://core.telegram.org/bots/api#reopengeneralforumtopic
func (bot *Bot) ReopenGeneralForumTopic(ctx context.Context, params *ReopenGeneralForumTopicParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "reopenGeneralForumTopic", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// ReplaceStickerInSetParams contains params for Telegram method "replaceStickerInSet".
type ReplaceStickerInSetParams struct {
	UserID     int64        `json:"user_id"`
	Name       string       `json:"name"`
	OldSticker string       `json:"old_sticker"`
	Sticker    InputSticker `json:"sticker"`
}

// ReplaceStickerInSet calls Telegram method "replaceStickerInSet".
// Doc: https://core.telegram.org/bots/api#replacestickerinset
func (bot *Bot) ReplaceStickerInSet(ctx context.Context, params *ReplaceStickerInSetParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "replaceStickerInSet", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// RepostStoryParams contains params for Telegram method "repostStory".
type RepostStoryParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	FromChatID           int64  `json:"from_chat_id"`
	FromStoryID          int64  `json:"from_story_id"`
	ActivePeriod         int64  `json:"active_period"`
	PostToChatPage       bool   `json:"post_to_chat_page,omitempty"`
	ProtectContent       bool   `json:"protect_content,omitempty"`
}

// RepostStory calls Telegram method "repostStory".
// Doc: https://core.telegram.org/bots/api#repoststory
func (bot *Bot) RepostStory(ctx context.Context, params *RepostStoryParams) (Story, error) {
	var result Story
	if err := bot.call(ctx, "repostStory", params, &result); err != nil {
		return Story{}, err
	}
	return result, nil
}

// RestrictChatMemberParams contains params for Telegram method "restrictChatMember".
type RestrictChatMemberParams struct {
	ChatID                        any             `json:"chat_id"`
	UserID                        int64           `json:"user_id"`
	Permissions                   ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions bool            `json:"use_independent_chat_permissions,omitempty"`
	UntilDate                     int64           `json:"until_date,omitempty"`
}

// RestrictChatMember calls Telegram method "restrictChatMember".
// Doc: https://core.telegram.org/bots/api#restrictchatmember
func (bot *Bot) RestrictChatMember(ctx context.Context, params *RestrictChatMemberParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "restrictChatMember", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// RevokeChatInviteLinkParams contains params for Telegram method "revokeChatInviteLink".
type RevokeChatInviteLinkParams struct {
	ChatID     any    `json:"chat_id"`
	InviteLink string `json:"invite_link"`
}

// RevokeChatInviteLink calls Telegram method "revokeChatInviteLink".
// Doc: https://core.telegram.org/bots/api#revokechatinvitelink
func (bot *Bot) RevokeChatInviteLink(ctx context.Context, params *RevokeChatInviteLinkParams) (ChatInviteLink, error) {
	var result ChatInviteLink
	if err := bot.call(ctx, "revokeChatInviteLink", params, &result); err != nil {
		return ChatInviteLink{}, err
	}
	return result, nil
}

// SavePreparedInlineMessageParams contains params for Telegram method "savePreparedInlineMessage".
type SavePreparedInlineMessageParams struct {
	UserID            int64             `json:"user_id"`
	Result            InlineQueryResult `json:"result"`
	AllowUserChats    bool              `json:"allow_user_chats,omitempty"`
	AllowBotChats     bool              `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   bool              `json:"allow_group_chats,omitempty"`
	AllowChannelChats bool              `json:"allow_channel_chats,omitempty"`
}

// SavePreparedInlineMessage calls Telegram method "savePreparedInlineMessage".
// Doc: https://core.telegram.org/bots/api#savepreparedinlinemessage
func (bot *Bot) SavePreparedInlineMessage(ctx context.Context, params *SavePreparedInlineMessageParams) (PreparedInlineMessage, error) {
	var result PreparedInlineMessage
	if err := bot.call(ctx, "savePreparedInlineMessage", params, &result); err != nil {
		return PreparedInlineMessage{}, err
	}
	return result, nil
}

// SendAnimationParams contains params for Telegram method "sendAnimation".
type SendAnimationParams struct {
	BusinessConnectionID    string                  `json:"business_connection_id,omitempty"`
	ChatID                  any                     `json:"chat_id"`
	MessageThreadID         int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int64                   `json:"direct_messages_topic_id,omitempty"`
	Animation               any                     `json:"animation"`
	Duration                int64                   `json:"duration,omitempty"`
	Width                   int64                   `json:"width,omitempty"`
	Height                  int64                   `json:"height,omitempty"`
	Thumbnail               any                     `json:"thumbnail,omitempty"`
	Caption                 string                  `json:"caption,omitempty"`
	ParseMode               string                  `json:"parse_mode,omitempty"`
	CaptionEntities         []MessageEntity         `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                    `json:"show_caption_above_media,omitempty"`
	HasSpoiler              bool                    `json:"has_spoiler,omitempty"`
	DisableNotification     bool                    `json:"disable_notification,omitempty"`
	ProtectContent          bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                     `json:"reply_markup,omitempty"`
}

// SendAnimation calls Telegram method "sendAnimation".
// Doc: https://core.telegram.org/bots/api#sendanimation
func (bot *Bot) SendAnimation(ctx context.Context, params *SendAnimationParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "sendAnimation", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// SendAudioParams contains params for Telegram method "sendAudio".
type SendAudioParams struct {
	BusinessConnectionID    string                  `json:"business_connection_id,omitempty"`
	ChatID                  any                     `json:"chat_id"`
	MessageThreadID         int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int64                   `json:"direct_messages_topic_id,omitempty"`
	Audio                   any                     `json:"audio"`
	Caption                 string                  `json:"caption,omitempty"`
	ParseMode               string                  `json:"parse_mode,omitempty"`
	CaptionEntities         []MessageEntity         `json:"caption_entities,omitempty"`
	Duration                int64                   `json:"duration,omitempty"`
	Performer               string                  `json:"performer,omitempty"`
	Title                   string                  `json:"title,omitempty"`
	Thumbnail               any                     `json:"thumbnail,omitempty"`
	DisableNotification     bool                    `json:"disable_notification,omitempty"`
	ProtectContent          bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                     `json:"reply_markup,omitempty"`
}

// SendAudio calls Telegram method "sendAudio".
// Doc: https://core.telegram.org/bots/api#sendaudio
func (bot *Bot) SendAudio(ctx context.Context, params *SendAudioParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "sendAudio", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// SendChatActionParams contains params for Telegram method "sendChatAction".
type SendChatActionParams struct {
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	ChatID               any    `json:"chat_id"`
	MessageThreadID      int64  `json:"message_thread_id,omitempty"`
	Action               string `json:"action"`
}

// SendChatAction calls Telegram method "sendChatAction".
// Doc: https://core.telegram.org/bots/api#sendchataction
func (bot *Bot) SendChatAction(ctx context.Context, params *SendChatActionParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "sendChatAction", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SendChecklistParams contains params for Telegram method "sendChecklist".
type SendChecklistParams struct {
	BusinessConnectionID string               `json:"business_connection_id"`
	ChatID               int64                `json:"chat_id"`
	Checklist            InputChecklist       `json:"checklist"`
	DisableNotification  bool                 `json:"disable_notification,omitempty"`
	ProtectContent       bool                 `json:"protect_content,omitempty"`
	MessageEffectID      string               `json:"message_effect_id,omitempty"`
	ReplyParameters      ReplyParameters      `json:"reply_parameters,omitempty"`
	ReplyMarkup          InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// SendChecklist calls Telegram method "sendChecklist".
// Doc: https://core.telegram.org/bots/api#sendchecklist
func (bot *Bot) SendChecklist(ctx context.Context, params *SendChecklistParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "sendChecklist", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// SendContactParams contains params for Telegram method "sendContact".
type SendContactParams struct {
	BusinessConnectionID    string                  `json:"business_connection_id,omitempty"`
	ChatID                  any                     `json:"chat_id"`
	MessageThreadID         int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int64                   `json:"direct_messages_topic_id,omitempty"`
	PhoneNumber             string                  `json:"phone_number"`
	FirstName               string                  `json:"first_name"`
	LastName                string                  `json:"last_name,omitempty"`
	Vcard                   string                  `json:"vcard,omitempty"`
	DisableNotification     bool                    `json:"disable_notification,omitempty"`
	ProtectContent          bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                     `json:"reply_markup,omitempty"`
}

// SendContact calls Telegram method "sendContact".
// Doc: https://core.telegram.org/bots/api#sendcontact
func (bot *Bot) SendContact(ctx context.Context, params *SendContactParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "sendContact", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// SendDiceParams contains params for Telegram method "sendDice".
type SendDiceParams struct {
	BusinessConnectionID    string                  `json:"business_connection_id,omitempty"`
	ChatID                  any                     `json:"chat_id"`
	MessageThreadID         int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int64                   `json:"direct_messages_topic_id,omitempty"`
	Emoji                   string                  `json:"emoji,omitempty"`
	DisableNotification     bool                    `json:"disable_notification,omitempty"`
	ProtectContent          bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                     `json:"reply_markup,omitempty"`
}

// SendDice calls Telegram method "sendDice".
// Doc: https://core.telegram.org/bots/api#senddice
func (bot *Bot) SendDice(ctx context.Context, params *SendDiceParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "sendDice", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// SendDocumentParams contains params for Telegram method "sendDocument".
type SendDocumentParams struct {
	BusinessConnectionID        string                  `json:"business_connection_id,omitempty"`
	ChatID                      any                     `json:"chat_id"`
	MessageThreadID             int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID       int64                   `json:"direct_messages_topic_id,omitempty"`
	Document                    any                     `json:"document"`
	Thumbnail                   any                     `json:"thumbnail,omitempty"`
	Caption                     string                  `json:"caption,omitempty"`
	ParseMode                   string                  `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity         `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                    `json:"disable_content_type_detection,omitempty"`
	DisableNotification         bool                    `json:"disable_notification,omitempty"`
	ProtectContent              bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast          bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID             string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters     SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters             ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup                 any                     `json:"reply_markup,omitempty"`
}

// SendDocument calls Telegram method "sendDocument".
// Doc: https://core.telegram.org/bots/api#senddocument
func (bot *Bot) SendDocument(ctx context.Context, params *SendDocumentParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "sendDocument", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// SendGameParams contains params for Telegram method "sendGame".
type SendGameParams struct {
	BusinessConnectionID string               `json:"business_connection_id,omitempty"`
	ChatID               int64                `json:"chat_id"`
	MessageThreadID      int64                `json:"message_thread_id,omitempty"`
	GameShortName        string               `json:"game_short_name"`
	DisableNotification  bool                 `json:"disable_notification,omitempty"`
	ProtectContent       bool                 `json:"protect_content,omitempty"`
	AllowPaidBroadcast   bool                 `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID      string               `json:"message_effect_id,omitempty"`
	ReplyParameters      ReplyParameters      `json:"reply_parameters,omitempty"`
	ReplyMarkup          InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// SendGame calls Telegram method "sendGame".
// Doc: https://core.telegram.org/bots/api#sendgame
func (bot *Bot) SendGame(ctx context.Context, params *SendGameParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "sendGame", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// SendGiftParams contains params for Telegram method "sendGift".
type SendGiftParams struct {
	UserID        int64           `json:"user_id,omitempty"`
	ChatID        any             `json:"chat_id,omitempty"`
	GiftID        string          `json:"gift_id"`
	PayForUpgrade bool            `json:"pay_for_upgrade,omitempty"`
	Text          string          `json:"text,omitempty"`
	TextParseMode string          `json:"text_parse_mode,omitempty"`
	TextEntities  []MessageEntity `json:"text_entities,omitempty"`
}

// SendGift calls Telegram method "sendGift".
// Doc: https://core.telegram.org/bots/api#sendgift
func (bot *Bot) SendGift(ctx context.Context, params *SendGiftParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "sendGift", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SendInvoiceParams contains params for Telegram method "sendInvoice".
type SendInvoiceParams struct {
	ChatID                    any                     `json:"chat_id"`
	MessageThreadID           int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID     int64                   `json:"direct_messages_topic_id,omitempty"`
	Title                     string                  `json:"title"`
	Description               string                  `json:"description"`
	Payload                   string                  `json:"payload"`
	ProviderToken             string                  `json:"provider_token,omitempty"`
	Currency                  string                  `json:"currency"`
	Prices                    []LabeledPrice          `json:"prices"`
	MaxTipAmount              int64                   `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int64                 `json:"suggested_tip_amounts,omitempty"`
	StartParameter            string                  `json:"start_parameter,omitempty"`
	ProviderData              string                  `json:"provider_data,omitempty"`
	PhotoURL                  string                  `json:"photo_url,omitempty"`
	PhotoSize                 int64                   `json:"photo_size,omitempty"`
	PhotoWidth                int64                   `json:"photo_width,omitempty"`
	PhotoHeight               int64                   `json:"photo_height,omitempty"`
	NeedName                  bool                    `json:"need_name,omitempty"`
	NeedPhoneNumber           bool                    `json:"need_phone_number,omitempty"`
	NeedEmail                 bool                    `json:"need_email,omitempty"`
	NeedShippingAddress       bool                    `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool                    `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool                    `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool                    `json:"is_flexible,omitempty"`
	DisableNotification       bool                    `json:"disable_notification,omitempty"`
	ProtectContent            bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast        bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID           string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters   SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters           ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup               InlineKeyboardMarkup    `json:"reply_markup,omitempty"`
}

// SendInvoice calls Telegram method "sendInvoice".
// Doc: https://core.telegram.org/bots/api#sendinvoice
func (bot *Bot) SendInvoice(ctx context.Context, params *SendInvoiceParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "sendInvoice", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// SendLocationParams contains params for Telegram method "sendLocation".
type SendLocationParams struct {
	BusinessConnectionID    string                  `json:"business_connection_id,omitempty"`
	ChatID                  any                     `json:"chat_id"`
	MessageThreadID         int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int64                   `json:"direct_messages_topic_id,omitempty"`
	Latitude                float64                 `json:"latitude"`
	Longitude               float64                 `json:"longitude"`
	HorizontalAccuracy      float64                 `json:"horizontal_accuracy,omitempty"`
	LivePeriod              int64                   `json:"live_period,omitempty"`
	Heading                 int64                   `json:"heading,omitempty"`
	ProximityAlertRadius    int64                   `json:"proximity_alert_radius,omitempty"`
	DisableNotification     bool                    `json:"disable_notification,omitempty"`
	ProtectContent          bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                     `json:"reply_markup,omitempty"`
}

// SendLocation calls Telegram method "sendLocation".
// Doc: https://core.telegram.org/bots/api#sendlocation
func (bot *Bot) SendLocation(ctx context.Context, params *SendLocationParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "sendLocation", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// SendMediaGroupParams contains params for Telegram method "sendMediaGroup".
type SendMediaGroupParams struct {
	BusinessConnectionID  string          `json:"business_connection_id,omitempty"`
	ChatID                any             `json:"chat_id"`
	MessageThreadID       int64           `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID int64           `json:"direct_messages_topic_id,omitempty"`
	Media                 []InputMedia    `json:"media"`
	DisableNotification   bool            `json:"disable_notification,omitempty"`
	ProtectContent        bool            `json:"protect_content,omitempty"`
	AllowPaidBroadcast    bool            `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID       string          `json:"message_effect_id,omitempty"`
	ReplyParameters       ReplyParameters `json:"reply_parameters,omitempty"`
}

// SendMediaGroup calls Telegram method "sendMediaGroup".
// Doc: https://core.telegram.org/bots/api#sendmediagroup
func (bot *Bot) SendMediaGroup(ctx context.Context, params *SendMediaGroupParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "sendMediaGroup", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// SendMessageParams contains params for Telegram method "sendMessage".
type SendMessageParams struct {
	BusinessConnectionID    string                  `json:"business_connection_id,omitempty"`
	ChatID                  any                     `json:"chat_id"`
	MessageThreadID         int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int64                   `json:"direct_messages_topic_id,omitempty"`
	Text                    string                  `json:"text"`
	ParseMode               string                  `json:"parse_mode,omitempty"`
	Entities                []MessageEntity         `json:"entities,omitempty"`
	LinkPreviewOptions      LinkPreviewOptions      `json:"link_preview_options,omitempty"`
	DisableNotification     bool                    `json:"disable_notification,omitempty"`
	ProtectContent          bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                     `json:"reply_markup,omitempty"`
}

// SendMessage calls Telegram method "sendMessage".
// Doc: https://core.telegram.org/bots/api#sendmessage
func (bot *Bot) SendMessage(ctx context.Context, params *SendMessageParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "sendMessage", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// SendMessageDraftParams contains params for Telegram method "sendMessageDraft".
type SendMessageDraftParams struct {
	ChatID          int64           `json:"chat_id"`
	MessageThreadID int64           `json:"message_thread_id,omitempty"`
	DraftID         int64           `json:"draft_id"`
	Text            string          `json:"text"`
	ParseMode       string          `json:"parse_mode,omitempty"`
	Entities        []MessageEntity `json:"entities,omitempty"`
}

// SendMessageDraft calls Telegram method "sendMessageDraft".
// Doc: https://core.telegram.org/bots/api#sendmessagedraft
func (bot *Bot) SendMessageDraft(ctx context.Context, params *SendMessageDraftParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "sendMessageDraft", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SendPaidMediaParams contains params for Telegram method "sendPaidMedia".
type SendPaidMediaParams struct {
	BusinessConnectionID    string                  `json:"business_connection_id,omitempty"`
	ChatID                  any                     `json:"chat_id"`
	MessageThreadID         int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int64                   `json:"direct_messages_topic_id,omitempty"`
	StarCount               int64                   `json:"star_count"`
	Media                   []InputPaidMedia        `json:"media"`
	Payload                 string                  `json:"payload,omitempty"`
	Caption                 string                  `json:"caption,omitempty"`
	ParseMode               string                  `json:"parse_mode,omitempty"`
	CaptionEntities         []MessageEntity         `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                    `json:"show_caption_above_media,omitempty"`
	DisableNotification     bool                    `json:"disable_notification,omitempty"`
	ProtectContent          bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                    `json:"allow_paid_broadcast,omitempty"`
	SuggestedPostParameters SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                     `json:"reply_markup,omitempty"`
}

// SendPaidMedia calls Telegram method "sendPaidMedia".
// Doc: https://core.telegram.org/bots/api#sendpaidmedia
func (bot *Bot) SendPaidMedia(ctx context.Context, params *SendPaidMediaParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "sendPaidMedia", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// SendPhotoParams contains params for Telegram method "sendPhoto".
type SendPhotoParams struct {
	BusinessConnectionID    string                  `json:"business_connection_id,omitempty"`
	ChatID                  any                     `json:"chat_id"`
	MessageThreadID         int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int64                   `json:"direct_messages_topic_id,omitempty"`
	Photo                   any                     `json:"photo"`
	Caption                 string                  `json:"caption,omitempty"`
	ParseMode               string                  `json:"parse_mode,omitempty"`
	CaptionEntities         []MessageEntity         `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                    `json:"show_caption_above_media,omitempty"`
	HasSpoiler              bool                    `json:"has_spoiler,omitempty"`
	DisableNotification     bool                    `json:"disable_notification,omitempty"`
	ProtectContent          bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                     `json:"reply_markup,omitempty"`
}

// SendPhoto calls Telegram method "sendPhoto".
// Doc: https://core.telegram.org/bots/api#sendphoto
func (bot *Bot) SendPhoto(ctx context.Context, params *SendPhotoParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "sendPhoto", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// SendPollParams contains params for Telegram method "sendPoll".
type SendPollParams struct {
	BusinessConnectionID  string            `json:"business_connection_id,omitempty"`
	ChatID                any               `json:"chat_id"`
	MessageThreadID       int64             `json:"message_thread_id,omitempty"`
	Question              string            `json:"question"`
	QuestionParseMode     string            `json:"question_parse_mode,omitempty"`
	QuestionEntities      []MessageEntity   `json:"question_entities,omitempty"`
	Options               []InputPollOption `json:"options"`
	IsAnonymous           bool              `json:"is_anonymous,omitempty"`
	Type                  string            `json:"type,omitempty"`
	AllowsMultipleAnswers bool              `json:"allows_multiple_answers,omitempty"`
	CorrectOptionID       int64             `json:"correct_option_id,omitempty"`
	Explanation           string            `json:"explanation,omitempty"`
	ExplanationParseMode  string            `json:"explanation_parse_mode,omitempty"`
	ExplanationEntities   []MessageEntity   `json:"explanation_entities,omitempty"`
	OpenPeriod            int64             `json:"open_period,omitempty"`
	CloseDate             int64             `json:"close_date,omitempty"`
	IsClosed              bool              `json:"is_closed,omitempty"`
	DisableNotification   bool              `json:"disable_notification,omitempty"`
	ProtectContent        bool              `json:"protect_content,omitempty"`
	AllowPaidBroadcast    bool              `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID       string            `json:"message_effect_id,omitempty"`
	ReplyParameters       ReplyParameters   `json:"reply_parameters,omitempty"`
	ReplyMarkup           any               `json:"reply_markup,omitempty"`
}

// SendPoll calls Telegram method "sendPoll".
// Doc: https://core.telegram.org/bots/api#sendpoll
func (bot *Bot) SendPoll(ctx context.Context, params *SendPollParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "sendPoll", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// SendStickerParams contains params for Telegram method "sendSticker".
type SendStickerParams struct {
	BusinessConnectionID    string                  `json:"business_connection_id,omitempty"`
	ChatID                  any                     `json:"chat_id"`
	MessageThreadID         int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int64                   `json:"direct_messages_topic_id,omitempty"`
	Sticker                 any                     `json:"sticker"`
	Emoji                   string                  `json:"emoji,omitempty"`
	DisableNotification     bool                    `json:"disable_notification,omitempty"`
	ProtectContent          bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                     `json:"reply_markup,omitempty"`
}

// SendSticker calls Telegram method "sendSticker".
// Doc: https://core.telegram.org/bots/api#sendsticker
func (bot *Bot) SendSticker(ctx context.Context, params *SendStickerParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "sendSticker", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// SendVenueParams contains params for Telegram method "sendVenue".
type SendVenueParams struct {
	BusinessConnectionID    string                  `json:"business_connection_id,omitempty"`
	ChatID                  any                     `json:"chat_id"`
	MessageThreadID         int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int64                   `json:"direct_messages_topic_id,omitempty"`
	Latitude                float64                 `json:"latitude"`
	Longitude               float64                 `json:"longitude"`
	Title                   string                  `json:"title"`
	Address                 string                  `json:"address"`
	FoursquareID            string                  `json:"foursquare_id,omitempty"`
	FoursquareType          string                  `json:"foursquare_type,omitempty"`
	GooglePlaceID           string                  `json:"google_place_id,omitempty"`
	GooglePlaceType         string                  `json:"google_place_type,omitempty"`
	DisableNotification     bool                    `json:"disable_notification,omitempty"`
	ProtectContent          bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                     `json:"reply_markup,omitempty"`
}

// SendVenue calls Telegram method "sendVenue".
// Doc: https://core.telegram.org/bots/api#sendvenue
func (bot *Bot) SendVenue(ctx context.Context, params *SendVenueParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "sendVenue", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// SendVideoParams contains params for Telegram method "sendVideo".
type SendVideoParams struct {
	BusinessConnectionID    string                  `json:"business_connection_id,omitempty"`
	ChatID                  any                     `json:"chat_id"`
	MessageThreadID         int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int64                   `json:"direct_messages_topic_id,omitempty"`
	Video                   any                     `json:"video"`
	Duration                int64                   `json:"duration,omitempty"`
	Width                   int64                   `json:"width,omitempty"`
	Height                  int64                   `json:"height,omitempty"`
	Thumbnail               any                     `json:"thumbnail,omitempty"`
	Cover                   any                     `json:"cover,omitempty"`
	StartTimestamp          int64                   `json:"start_timestamp,omitempty"`
	Caption                 string                  `json:"caption,omitempty"`
	ParseMode               string                  `json:"parse_mode,omitempty"`
	CaptionEntities         []MessageEntity         `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                    `json:"show_caption_above_media,omitempty"`
	HasSpoiler              bool                    `json:"has_spoiler,omitempty"`
	SupportsStreaming       bool                    `json:"supports_streaming,omitempty"`
	DisableNotification     bool                    `json:"disable_notification,omitempty"`
	ProtectContent          bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                     `json:"reply_markup,omitempty"`
}

// SendVideo calls Telegram method "sendVideo".
// Doc: https://core.telegram.org/bots/api#sendvideo
func (bot *Bot) SendVideo(ctx context.Context, params *SendVideoParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "sendVideo", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// SendVideoNoteParams contains params for Telegram method "sendVideoNote".
type SendVideoNoteParams struct {
	BusinessConnectionID    string                  `json:"business_connection_id,omitempty"`
	ChatID                  any                     `json:"chat_id"`
	MessageThreadID         int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int64                   `json:"direct_messages_topic_id,omitempty"`
	VideoNote               any                     `json:"video_note"`
	Duration                int64                   `json:"duration,omitempty"`
	Length                  int64                   `json:"length,omitempty"`
	Thumbnail               any                     `json:"thumbnail,omitempty"`
	DisableNotification     bool                    `json:"disable_notification,omitempty"`
	ProtectContent          bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                     `json:"reply_markup,omitempty"`
}

// SendVideoNote calls Telegram method "sendVideoNote".
// Doc: https://core.telegram.org/bots/api#sendvideonote
func (bot *Bot) SendVideoNote(ctx context.Context, params *SendVideoNoteParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "sendVideoNote", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// SendVoiceParams contains params for Telegram method "sendVoice".
type SendVoiceParams struct {
	BusinessConnectionID    string                  `json:"business_connection_id,omitempty"`
	ChatID                  any                     `json:"chat_id"`
	MessageThreadID         int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int64                   `json:"direct_messages_topic_id,omitempty"`
	Voice                   any                     `json:"voice"`
	Caption                 string                  `json:"caption,omitempty"`
	ParseMode               string                  `json:"parse_mode,omitempty"`
	CaptionEntities         []MessageEntity         `json:"caption_entities,omitempty"`
	Duration                int64                   `json:"duration,omitempty"`
	DisableNotification     bool                    `json:"disable_notification,omitempty"`
	ProtectContent          bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                     `json:"reply_markup,omitempty"`
}

// SendVoice calls Telegram method "sendVoice".
// Doc: https://core.telegram.org/bots/api#sendvoice
func (bot *Bot) SendVoice(ctx context.Context, params *SendVoiceParams) (Message, error) {
	var result Message
	if err := bot.call(ctx, "sendVoice", params, &result); err != nil {
		return Message{}, err
	}
	return result, nil
}

// SetBusinessAccountBioParams contains params for Telegram method "setBusinessAccountBio".
type SetBusinessAccountBioParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	Bio                  string `json:"bio,omitempty"`
}

// SetBusinessAccountBio calls Telegram method "setBusinessAccountBio".
// Doc: https://core.telegram.org/bots/api#setbusinessaccountbio
func (bot *Bot) SetBusinessAccountBio(ctx context.Context, params *SetBusinessAccountBioParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setBusinessAccountBio", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetBusinessAccountGiftSettingsParams contains params for Telegram method "setBusinessAccountGiftSettings".
type SetBusinessAccountGiftSettingsParams struct {
	BusinessConnectionID string            `json:"business_connection_id"`
	ShowGiftButton       bool              `json:"show_gift_button"`
	AcceptedGiftTypes    AcceptedGiftTypes `json:"accepted_gift_types"`
}

// SetBusinessAccountGiftSettings calls Telegram method "setBusinessAccountGiftSettings".
// Doc: https://core.telegram.org/bots/api#setbusinessaccountgiftsettings
func (bot *Bot) SetBusinessAccountGiftSettings(ctx context.Context, params *SetBusinessAccountGiftSettingsParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setBusinessAccountGiftSettings", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetBusinessAccountNameParams contains params for Telegram method "setBusinessAccountName".
type SetBusinessAccountNameParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name,omitempty"`
}

// SetBusinessAccountName calls Telegram method "setBusinessAccountName".
// Doc: https://core.telegram.org/bots/api#setbusinessaccountname
func (bot *Bot) SetBusinessAccountName(ctx context.Context, params *SetBusinessAccountNameParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setBusinessAccountName", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetBusinessAccountProfilePhotoParams contains params for Telegram method "setBusinessAccountProfilePhoto".
type SetBusinessAccountProfilePhotoParams struct {
	BusinessConnectionID string            `json:"business_connection_id"`
	Photo                InputProfilePhoto `json:"photo"`
	IsPublic             bool              `json:"is_public,omitempty"`
}

// SetBusinessAccountProfilePhoto calls Telegram method "setBusinessAccountProfilePhoto".
// Doc: https://core.telegram.org/bots/api#setbusinessaccountprofilephoto
func (bot *Bot) SetBusinessAccountProfilePhoto(ctx context.Context, params *SetBusinessAccountProfilePhotoParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setBusinessAccountProfilePhoto", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetBusinessAccountUsernameParams contains params for Telegram method "setBusinessAccountUsername".
type SetBusinessAccountUsernameParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	Username             string `json:"username,omitempty"`
}

// SetBusinessAccountUsername calls Telegram method "setBusinessAccountUsername".
// Doc: https://core.telegram.org/bots/api#setbusinessaccountusername
func (bot *Bot) SetBusinessAccountUsername(ctx context.Context, params *SetBusinessAccountUsernameParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setBusinessAccountUsername", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetChatAdministratorCustomTitleParams contains params for Telegram method "setChatAdministratorCustomTitle".
type SetChatAdministratorCustomTitleParams struct {
	ChatID      any    `json:"chat_id"`
	UserID      int64  `json:"user_id"`
	CustomTitle string `json:"custom_title"`
}

// SetChatAdministratorCustomTitle calls Telegram method "setChatAdministratorCustomTitle".
// Doc: https://core.telegram.org/bots/api#setchatadministratorcustomtitle
func (bot *Bot) SetChatAdministratorCustomTitle(ctx context.Context, params *SetChatAdministratorCustomTitleParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setChatAdministratorCustomTitle", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetChatDescriptionParams contains params for Telegram method "setChatDescription".
type SetChatDescriptionParams struct {
	ChatID      any    `json:"chat_id"`
	Description string `json:"description,omitempty"`
}

// SetChatDescription calls Telegram method "setChatDescription".
// Doc: https://core.telegram.org/bots/api#setchatdescription
func (bot *Bot) SetChatDescription(ctx context.Context, params *SetChatDescriptionParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setChatDescription", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetChatMemberTagParams contains params for Telegram method "setChatMemberTag".
type SetChatMemberTagParams struct {
	ChatID any    `json:"chat_id"`
	UserID int64  `json:"user_id"`
	Tag    string `json:"tag,omitempty"`
}

// SetChatMemberTag calls Telegram method "setChatMemberTag".
// Doc: https://core.telegram.org/bots/api#setchatmembertag
func (bot *Bot) SetChatMemberTag(ctx context.Context, params *SetChatMemberTagParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setChatMemberTag", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetChatMenuButtonParams contains params for Telegram method "setChatMenuButton".
type SetChatMenuButtonParams struct {
	ChatID     int64      `json:"chat_id,omitempty"`
	MenuButton MenuButton `json:"menu_button,omitempty"`
}

// SetChatMenuButton calls Telegram method "setChatMenuButton".
// Doc: https://core.telegram.org/bots/api#setchatmenubutton
func (bot *Bot) SetChatMenuButton(ctx context.Context, params *SetChatMenuButtonParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setChatMenuButton", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetChatPermissionsParams contains params for Telegram method "setChatPermissions".
type SetChatPermissionsParams struct {
	ChatID                        any             `json:"chat_id"`
	Permissions                   ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions bool            `json:"use_independent_chat_permissions,omitempty"`
}

// SetChatPermissions calls Telegram method "setChatPermissions".
// Doc: https://core.telegram.org/bots/api#setchatpermissions
func (bot *Bot) SetChatPermissions(ctx context.Context, params *SetChatPermissionsParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setChatPermissions", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetChatPhotoParams contains params for Telegram method "setChatPhoto".
type SetChatPhotoParams struct {
	ChatID any       `json:"chat_id"`
	Photo  InputFile `json:"photo"`
}

// SetChatPhoto calls Telegram method "setChatPhoto".
// Doc: https://core.telegram.org/bots/api#setchatphoto
func (bot *Bot) SetChatPhoto(ctx context.Context, params *SetChatPhotoParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setChatPhoto", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetChatStickerSetParams contains params for Telegram method "setChatStickerSet".
type SetChatStickerSetParams struct {
	ChatID         any    `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
}

// SetChatStickerSet calls Telegram method "setChatStickerSet".
// Doc: https://core.telegram.org/bots/api#setchatstickerset
func (bot *Bot) SetChatStickerSet(ctx context.Context, params *SetChatStickerSetParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setChatStickerSet", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetChatTitleParams contains params for Telegram method "setChatTitle".
type SetChatTitleParams struct {
	ChatID any    `json:"chat_id"`
	Title  string `json:"title"`
}

// SetChatTitle calls Telegram method "setChatTitle".
// Doc: https://core.telegram.org/bots/api#setchattitle
func (bot *Bot) SetChatTitle(ctx context.Context, params *SetChatTitleParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setChatTitle", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetCustomEmojiStickerSetThumbnailParams contains params for Telegram method "setCustomEmojiStickerSetThumbnail".
type SetCustomEmojiStickerSetThumbnailParams struct {
	Name          string `json:"name"`
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

// SetCustomEmojiStickerSetThumbnail calls Telegram method "setCustomEmojiStickerSetThumbnail".
// Doc: https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
func (bot *Bot) SetCustomEmojiStickerSetThumbnail(ctx context.Context, params *SetCustomEmojiStickerSetThumbnailParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setCustomEmojiStickerSetThumbnail", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetGameScoreParams contains params for Telegram method "setGameScore".
type SetGameScoreParams struct {
	UserID             int64  `json:"user_id"`
	Score              int64  `json:"score"`
	Force              bool   `json:"force,omitempty"`
	DisableEditMessage bool   `json:"disable_edit_message,omitempty"`
	ChatID             int64  `json:"chat_id,omitempty"`
	MessageID          int64  `json:"message_id,omitempty"`
	InlineMessageID    string `json:"inline_message_id,omitempty"`
}

// SetGameScore calls Telegram method "setGameScore".
// Doc: https://core.telegram.org/bots/api#setgamescore
func (bot *Bot) SetGameScore(ctx context.Context, params *SetGameScoreParams) (any, error) {
	var result any
	if err := bot.call(ctx, "setGameScore", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetMessageReactionParams contains params for Telegram method "setMessageReaction".
type SetMessageReactionParams struct {
	ChatID    any            `json:"chat_id"`
	MessageID int64          `json:"message_id"`
	Reaction  []ReactionType `json:"reaction,omitempty"`
	IsBig     bool           `json:"is_big,omitempty"`
}

// SetMessageReaction calls Telegram method "setMessageReaction".
// Doc: https://core.telegram.org/bots/api#setmessagereaction
func (bot *Bot) SetMessageReaction(ctx context.Context, params *SetMessageReactionParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setMessageReaction", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetMyCommandsParams contains params for Telegram method "setMyCommands".
type SetMyCommandsParams struct {
	Commands     []BotCommand    `json:"commands"`
	Scope        BotCommandScope `json:"scope,omitempty"`
	LanguageCode string          `json:"language_code,omitempty"`
}

// SetMyCommands calls Telegram method "setMyCommands".
// Doc: https://core.telegram.org/bots/api#setmycommands
func (bot *Bot) SetMyCommands(ctx context.Context, params *SetMyCommandsParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setMyCommands", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetMyDefaultAdministratorRightsParams contains params for Telegram method "setMyDefaultAdministratorRights".
type SetMyDefaultAdministratorRightsParams struct {
	Rights      ChatAdministratorRights `json:"rights,omitempty"`
	ForChannels bool                    `json:"for_channels,omitempty"`
}

// SetMyDefaultAdministratorRights calls Telegram method "setMyDefaultAdministratorRights".
// Doc: https://core.telegram.org/bots/api#setmydefaultadministratorrights
func (bot *Bot) SetMyDefaultAdministratorRights(ctx context.Context, params *SetMyDefaultAdministratorRightsParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setMyDefaultAdministratorRights", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetMyDescriptionParams contains params for Telegram method "setMyDescription".
type SetMyDescriptionParams struct {
	Description  string `json:"description,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

// SetMyDescription calls Telegram method "setMyDescription".
// Doc: https://core.telegram.org/bots/api#setmydescription
func (bot *Bot) SetMyDescription(ctx context.Context, params *SetMyDescriptionParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setMyDescription", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetMyNameParams contains params for Telegram method "setMyName".
type SetMyNameParams struct {
	Name         string `json:"name,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

// SetMyName calls Telegram method "setMyName".
// Doc: https://core.telegram.org/bots/api#setmyname
func (bot *Bot) SetMyName(ctx context.Context, params *SetMyNameParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setMyName", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetMyProfilePhotoParams contains params for Telegram method "setMyProfilePhoto".
type SetMyProfilePhotoParams struct {
	Photo InputProfilePhoto `json:"photo"`
}

// SetMyProfilePhoto calls Telegram method "setMyProfilePhoto".
// Doc: https://core.telegram.org/bots/api#setmyprofilephoto
func (bot *Bot) SetMyProfilePhoto(ctx context.Context, params *SetMyProfilePhotoParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setMyProfilePhoto", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetMyShortDescriptionParams contains params for Telegram method "setMyShortDescription".
type SetMyShortDescriptionParams struct {
	ShortDescription string `json:"short_description,omitempty"`
	LanguageCode     string `json:"language_code,omitempty"`
}

// SetMyShortDescription calls Telegram method "setMyShortDescription".
// Doc: https://core.telegram.org/bots/api#setmyshortdescription
func (bot *Bot) SetMyShortDescription(ctx context.Context, params *SetMyShortDescriptionParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setMyShortDescription", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetPassportDataErrorsParams contains params for Telegram method "setPassportDataErrors".
type SetPassportDataErrorsParams struct {
	UserID int64                  `json:"user_id"`
	Errors []PassportElementError `json:"errors"`
}

// SetPassportDataErrors calls Telegram method "setPassportDataErrors".
// Doc: https://core.telegram.org/bots/api#setpassportdataerrors
func (bot *Bot) SetPassportDataErrors(ctx context.Context, params *SetPassportDataErrorsParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setPassportDataErrors", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetStickerEmojiListParams contains params for Telegram method "setStickerEmojiList".
type SetStickerEmojiListParams struct {
	Sticker   string   `json:"sticker"`
	EmojiList []string `json:"emoji_list"`
}

// SetStickerEmojiList calls Telegram method "setStickerEmojiList".
// Doc: https://core.telegram.org/bots/api#setstickeremojilist
func (bot *Bot) SetStickerEmojiList(ctx context.Context, params *SetStickerEmojiListParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setStickerEmojiList", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetStickerKeywordsParams contains params for Telegram method "setStickerKeywords".
type SetStickerKeywordsParams struct {
	Sticker  string   `json:"sticker"`
	Keywords []string `json:"keywords,omitempty"`
}

// SetStickerKeywords calls Telegram method "setStickerKeywords".
// Doc: https://core.telegram.org/bots/api#setstickerkeywords
func (bot *Bot) SetStickerKeywords(ctx context.Context, params *SetStickerKeywordsParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setStickerKeywords", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetStickerMaskPositionParams contains params for Telegram method "setStickerMaskPosition".
type SetStickerMaskPositionParams struct {
	Sticker      string       `json:"sticker"`
	MaskPosition MaskPosition `json:"mask_position,omitempty"`
}

// SetStickerMaskPosition calls Telegram method "setStickerMaskPosition".
// Doc: https://core.telegram.org/bots/api#setstickermaskposition
func (bot *Bot) SetStickerMaskPosition(ctx context.Context, params *SetStickerMaskPositionParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setStickerMaskPosition", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetStickerPositionInSetParams contains params for Telegram method "setStickerPositionInSet".
type SetStickerPositionInSetParams struct {
	Sticker  string `json:"sticker"`
	Position int64  `json:"position"`
}

// SetStickerPositionInSet calls Telegram method "setStickerPositionInSet".
// Doc: https://core.telegram.org/bots/api#setstickerpositioninset
func (bot *Bot) SetStickerPositionInSet(ctx context.Context, params *SetStickerPositionInSetParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setStickerPositionInSet", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetStickerSetThumbnailParams contains params for Telegram method "setStickerSetThumbnail".
type SetStickerSetThumbnailParams struct {
	Name      string `json:"name"`
	UserID    int64  `json:"user_id"`
	Thumbnail any    `json:"thumbnail,omitempty"`
	Format    string `json:"format"`
}

// SetStickerSetThumbnail calls Telegram method "setStickerSetThumbnail".
// Doc: https://core.telegram.org/bots/api#setstickersetthumbnail
func (bot *Bot) SetStickerSetThumbnail(ctx context.Context, params *SetStickerSetThumbnailParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setStickerSetThumbnail", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetStickerSetTitleParams contains params for Telegram method "setStickerSetTitle".
type SetStickerSetTitleParams struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

// SetStickerSetTitle calls Telegram method "setStickerSetTitle".
// Doc: https://core.telegram.org/bots/api#setstickersettitle
func (bot *Bot) SetStickerSetTitle(ctx context.Context, params *SetStickerSetTitleParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setStickerSetTitle", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetUserEmojiStatusParams contains params for Telegram method "setUserEmojiStatus".
type SetUserEmojiStatusParams struct {
	UserID                    int64  `json:"user_id"`
	EmojiStatusCustomEmojiID  string `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate int64  `json:"emoji_status_expiration_date,omitempty"`
}

// SetUserEmojiStatus calls Telegram method "setUserEmojiStatus".
// Doc: https://core.telegram.org/bots/api#setuseremojistatus
func (bot *Bot) SetUserEmojiStatus(ctx context.Context, params *SetUserEmojiStatusParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setUserEmojiStatus", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// SetWebhookParams contains params for Telegram method "setWebhook".
type SetWebhookParams struct {
	URL                string    `json:"url"`
	Certificate        InputFile `json:"certificate,omitempty"`
	IPAddress          string    `json:"ip_address,omitempty"`
	MaxConnections     int64     `json:"max_connections,omitempty"`
	AllowedUpdates     []string  `json:"allowed_updates,omitempty"`
	DropPendingUpdates bool      `json:"drop_pending_updates,omitempty"`
	SecretToken        string    `json:"secret_token,omitempty"`
}

// SetWebhook calls Telegram method "setWebhook".
// Doc: https://core.telegram.org/bots/api#setwebhook
func (bot *Bot) SetWebhook(ctx context.Context, params *SetWebhookParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "setWebhook", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// StopMessageLiveLocationParams contains params for Telegram method "stopMessageLiveLocation".
type StopMessageLiveLocationParams struct {
	BusinessConnectionID string               `json:"business_connection_id,omitempty"`
	ChatID               any                  `json:"chat_id,omitempty"`
	MessageID            int64                `json:"message_id,omitempty"`
	InlineMessageID      string               `json:"inline_message_id,omitempty"`
	ReplyMarkup          InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// StopMessageLiveLocation calls Telegram method "stopMessageLiveLocation".
// Doc: https://core.telegram.org/bots/api#stopmessagelivelocation
func (bot *Bot) StopMessageLiveLocation(ctx context.Context, params *StopMessageLiveLocationParams) (any, error) {
	var result any
	if err := bot.call(ctx, "stopMessageLiveLocation", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// StopPollParams contains params for Telegram method "stopPoll".
type StopPollParams struct {
	BusinessConnectionID string               `json:"business_connection_id,omitempty"`
	ChatID               any                  `json:"chat_id"`
	MessageID            int64                `json:"message_id"`
	ReplyMarkup          InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// StopPoll calls Telegram method "stopPoll".
// Doc: https://core.telegram.org/bots/api#stoppoll
func (bot *Bot) StopPoll(ctx context.Context, params *StopPollParams) (Poll, error) {
	var result Poll
	if err := bot.call(ctx, "stopPoll", params, &result); err != nil {
		return Poll{}, err
	}
	return result, nil
}

// TransferBusinessAccountStarsParams contains params for Telegram method "transferBusinessAccountStars".
type TransferBusinessAccountStarsParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	StarCount            int64  `json:"star_count"`
}

// TransferBusinessAccountStars calls Telegram method "transferBusinessAccountStars".
// Doc: https://core.telegram.org/bots/api#transferbusinessaccountstars
func (bot *Bot) TransferBusinessAccountStars(ctx context.Context, params *TransferBusinessAccountStarsParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "transferBusinessAccountStars", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// TransferGiftParams contains params for Telegram method "transferGift".
type TransferGiftParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	OwnedGiftID          string `json:"owned_gift_id"`
	NewOwnerChatID       int64  `json:"new_owner_chat_id"`
	StarCount            int64  `json:"star_count,omitempty"`
}

// TransferGift calls Telegram method "transferGift".
// Doc: https://core.telegram.org/bots/api#transfergift
func (bot *Bot) TransferGift(ctx context.Context, params *TransferGiftParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "transferGift", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// UnbanChatMemberParams contains params for Telegram method "unbanChatMember".
type UnbanChatMemberParams struct {
	ChatID       any   `json:"chat_id"`
	UserID       int64 `json:"user_id"`
	OnlyIfBanned bool  `json:"only_if_banned,omitempty"`
}

// UnbanChatMember calls Telegram method "unbanChatMember".
// Doc: https://core.telegram.org/bots/api#unbanchatmember
func (bot *Bot) UnbanChatMember(ctx context.Context, params *UnbanChatMemberParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "unbanChatMember", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// UnbanChatSenderChatParams contains params for Telegram method "unbanChatSenderChat".
type UnbanChatSenderChatParams struct {
	ChatID       any   `json:"chat_id"`
	SenderChatID int64 `json:"sender_chat_id"`
}

// UnbanChatSenderChat calls Telegram method "unbanChatSenderChat".
// Doc: https://core.telegram.org/bots/api#unbanchatsenderchat
func (bot *Bot) UnbanChatSenderChat(ctx context.Context, params *UnbanChatSenderChatParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "unbanChatSenderChat", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// UnhideGeneralForumTopicParams contains params for Telegram method "unhideGeneralForumTopic".
type UnhideGeneralForumTopicParams struct {
	ChatID any `json:"chat_id"`
}

// UnhideGeneralForumTopic calls Telegram method "unhideGeneralForumTopic".
// Doc: https://core.telegram.org/bots/api#unhidegeneralforumtopic
func (bot *Bot) UnhideGeneralForumTopic(ctx context.Context, params *UnhideGeneralForumTopicParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "unhideGeneralForumTopic", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// UnpinAllChatMessagesParams contains params for Telegram method "unpinAllChatMessages".
type UnpinAllChatMessagesParams struct {
	ChatID any `json:"chat_id"`
}

// UnpinAllChatMessages calls Telegram method "unpinAllChatMessages".
// Doc: https://core.telegram.org/bots/api#unpinallchatmessages
func (bot *Bot) UnpinAllChatMessages(ctx context.Context, params *UnpinAllChatMessagesParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "unpinAllChatMessages", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// UnpinAllForumTopicMessagesParams contains params for Telegram method "unpinAllForumTopicMessages".
type UnpinAllForumTopicMessagesParams struct {
	ChatID          any   `json:"chat_id"`
	MessageThreadID int64 `json:"message_thread_id"`
}

// UnpinAllForumTopicMessages calls Telegram method "unpinAllForumTopicMessages".
// Doc: https://core.telegram.org/bots/api#unpinallforumtopicmessages
func (bot *Bot) UnpinAllForumTopicMessages(ctx context.Context, params *UnpinAllForumTopicMessagesParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "unpinAllForumTopicMessages", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// UnpinAllGeneralForumTopicMessagesParams contains params for Telegram method "unpinAllGeneralForumTopicMessages".
type UnpinAllGeneralForumTopicMessagesParams struct {
	ChatID any `json:"chat_id"`
}

// UnpinAllGeneralForumTopicMessages calls Telegram method "unpinAllGeneralForumTopicMessages".
// Doc: https://core.telegram.org/bots/api#unpinallgeneralforumtopicmessages
func (bot *Bot) UnpinAllGeneralForumTopicMessages(ctx context.Context, params *UnpinAllGeneralForumTopicMessagesParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "unpinAllGeneralForumTopicMessages", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// UnpinChatMessageParams contains params for Telegram method "unpinChatMessage".
type UnpinChatMessageParams struct {
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	ChatID               any    `json:"chat_id"`
	MessageID            int64  `json:"message_id,omitempty"`
}

// UnpinChatMessage calls Telegram method "unpinChatMessage".
// Doc: https://core.telegram.org/bots/api#unpinchatmessage
func (bot *Bot) UnpinChatMessage(ctx context.Context, params *UnpinChatMessageParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "unpinChatMessage", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// UpgradeGiftParams contains params for Telegram method "upgradeGift".
type UpgradeGiftParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	OwnedGiftID          string `json:"owned_gift_id"`
	KeepOriginalDetails  bool   `json:"keep_original_details,omitempty"`
	StarCount            int64  `json:"star_count,omitempty"`
}

// UpgradeGift calls Telegram method "upgradeGift".
// Doc: https://core.telegram.org/bots/api#upgradegift
func (bot *Bot) UpgradeGift(ctx context.Context, params *UpgradeGiftParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "upgradeGift", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// UploadStickerFileParams contains params for Telegram method "uploadStickerFile".
type UploadStickerFileParams struct {
	UserID        int64     `json:"user_id"`
	Sticker       InputFile `json:"sticker"`
	StickerFormat string    `json:"sticker_format"`
}

// UploadStickerFile calls Telegram method "uploadStickerFile".
// Doc: https://core.telegram.org/bots/api#uploadstickerfile
func (bot *Bot) UploadStickerFile(ctx context.Context, params *UploadStickerFileParams) (File, error) {
	var result File
	if err := bot.call(ctx, "uploadStickerFile", params, &result); err != nil {
		return File{}, err
	}
	return result, nil
}

// VerifyChatParams contains params for Telegram method "verifyChat".
type VerifyChatParams struct {
	ChatID            any    `json:"chat_id"`
	CustomDescription string `json:"custom_description,omitempty"`
}

// VerifyChat calls Telegram method "verifyChat".
// Doc: https://core.telegram.org/bots/api#verifychat
func (bot *Bot) VerifyChat(ctx context.Context, params *VerifyChatParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "verifyChat", params, &result); err != nil {
		return false, err
	}
	return result, nil
}

// VerifyUserParams contains params for Telegram method "verifyUser".
type VerifyUserParams struct {
	UserID            int64  `json:"user_id"`
	CustomDescription string `json:"custom_description,omitempty"`
}

// VerifyUser calls Telegram method "verifyUser".
// Doc: https://core.telegram.org/bots/api#verifyuser
func (bot *Bot) VerifyUser(ctx context.Context, params *VerifyUserParams) (bool, error) {
	var result bool
	if err := bot.call(ctx, "verifyUser", params, &result); err != nil {
		return false, err
	}
	return result, nil
}
