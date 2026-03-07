package ext

import (
	"encoding/json"
	"fmt"
	"io"

	tg "github.com/cloudapp3/tgbot"
)

// UpdateType is Telegram update discriminator used by the application dispatcher.
type UpdateType string

const (
	UpdateTypeUnknown                 UpdateType = ""
	UpdateTypeMessage                 UpdateType = "message"
	UpdateTypeEditedMessage           UpdateType = "edited_message"
	UpdateTypeChannelPost             UpdateType = "channel_post"
	UpdateTypeEditedChannelPost       UpdateType = "edited_channel_post"
	UpdateTypeBusinessConnection      UpdateType = "business_connection"
	UpdateTypeBusinessMessage         UpdateType = "business_message"
	UpdateTypeEditedBusinessMessage   UpdateType = "edited_business_message"
	UpdateTypeDeletedBusinessMessages UpdateType = "deleted_business_messages"
	UpdateTypeMessageReaction         UpdateType = "message_reaction"
	UpdateTypeMessageReactionCount    UpdateType = "message_reaction_count"
	UpdateTypeInlineQuery             UpdateType = "inline_query"
	UpdateTypeChosenInlineResult      UpdateType = "chosen_inline_result"
	UpdateTypeCallbackQuery           UpdateType = "callback_query"
	UpdateTypeShippingQuery           UpdateType = "shipping_query"
	UpdateTypePreCheckoutQuery        UpdateType = "pre_checkout_query"
	UpdateTypePurchasedPaidMedia      UpdateType = "purchased_paid_media"
	UpdateTypePoll                    UpdateType = "poll"
	UpdateTypePollAnswer              UpdateType = "poll_answer"
	UpdateTypeMyChatMember            UpdateType = "my_chat_member"
	UpdateTypeChatMember              UpdateType = "chat_member"
	UpdateTypeChatJoinRequest         UpdateType = "chat_join_request"
	UpdateTypeChatBoost               UpdateType = "chat_boost"
	UpdateTypeRemovedChatBoost        UpdateType = "removed_chat_boost"
)

// AllUpdateTypes lists every currently supported Telegram update type.
var AllUpdateTypes = []UpdateType{
	UpdateTypeMessage,
	UpdateTypeEditedMessage,
	UpdateTypeChannelPost,
	UpdateTypeEditedChannelPost,
	UpdateTypeBusinessConnection,
	UpdateTypeBusinessMessage,
	UpdateTypeEditedBusinessMessage,
	UpdateTypeDeletedBusinessMessages,
	UpdateTypeMessageReaction,
	UpdateTypeMessageReactionCount,
	UpdateTypeInlineQuery,
	UpdateTypeChosenInlineResult,
	UpdateTypeCallbackQuery,
	UpdateTypeShippingQuery,
	UpdateTypePreCheckoutQuery,
	UpdateTypePurchasedPaidMedia,
	UpdateTypePoll,
	UpdateTypePollAnswer,
	UpdateTypeMyChatMember,
	UpdateTypeChatMember,
	UpdateTypeChatJoinRequest,
	UpdateTypeChatBoost,
	UpdateTypeRemovedChatBoost,
}

// Update is the complete Telegram update envelope used by the application layer.
type Update struct {
	UpdateID                int64                           `json:"update_id"`
	Message                 *tg.Message                     `json:"message,omitempty"`
	EditedMessage           *tg.Message                     `json:"edited_message,omitempty"`
	ChannelPost             *tg.Message                     `json:"channel_post,omitempty"`
	EditedChannelPost       *tg.Message                     `json:"edited_channel_post,omitempty"`
	BusinessConnection      *tg.BusinessConnection          `json:"business_connection,omitempty"`
	BusinessMessage         *tg.Message                     `json:"business_message,omitempty"`
	EditedBusinessMessage   *tg.Message                     `json:"edited_business_message,omitempty"`
	DeletedBusinessMessages *tg.BusinessMessagesDeleted     `json:"deleted_business_messages,omitempty"`
	MessageReaction         *tg.MessageReactionUpdated      `json:"message_reaction,omitempty"`
	MessageReactionCount    *tg.MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`
	InlineQuery             *tg.InlineQuery                 `json:"inline_query,omitempty"`
	ChosenInlineResult      *tg.ChosenInlineResult          `json:"chosen_inline_result,omitempty"`
	CallbackQuery           *tg.CallbackQuery               `json:"callback_query,omitempty"`
	ShippingQuery           *tg.ShippingQuery               `json:"shipping_query,omitempty"`
	PreCheckoutQuery        *tg.PreCheckoutQuery            `json:"pre_checkout_query,omitempty"`
	PurchasedPaidMedia      *tg.PaidMediaPurchased          `json:"purchased_paid_media,omitempty"`
	Poll                    *tg.Poll                        `json:"poll,omitempty"`
	PollAnswer              *tg.PollAnswer                  `json:"poll_answer,omitempty"`
	MyChatMember            *tg.ChatMemberUpdated           `json:"my_chat_member,omitempty"`
	ChatMember              *tg.ChatMemberUpdated           `json:"chat_member,omitempty"`
	ChatJoinRequest         *tg.ChatJoinRequest             `json:"chat_join_request,omitempty"`
	ChatBoost               *tg.ChatBoostUpdated            `json:"chat_boost,omitempty"`
	RemovedChatBoost        *tg.ChatBoostRemoved            `json:"removed_chat_boost,omitempty"`
}

// WrapUpdate converts a root tgbot.Update into an ext.Update.
func WrapUpdate(update tg.Update) *Update {
	return &Update{
		UpdateID:                update.UpdateID,
		Message:                 update.Message,
		EditedMessage:           update.EditedMessage,
		ChannelPost:             update.ChannelPost,
		EditedChannelPost:       update.EditedChannelPost,
		BusinessConnection:      update.BusinessConnection,
		BusinessMessage:         update.BusinessMessage,
		EditedBusinessMessage:   update.EditedBusinessMessage,
		DeletedBusinessMessages: update.DeletedBusinessMessages,
		MessageReaction:         update.MessageReaction,
		MessageReactionCount:    update.MessageReactionCount,
		InlineQuery:             update.InlineQuery,
		ChosenInlineResult:      update.ChosenInlineResult,
		CallbackQuery:           update.CallbackQuery,
		ShippingQuery:           update.ShippingQuery,
		PreCheckoutQuery:        update.PreCheckoutQuery,
		PurchasedPaidMedia:      update.PurchasedPaidMedia,
		Poll:                    update.Poll,
		PollAnswer:              update.PollAnswer,
		MyChatMember:            update.MyChatMember,
		ChatMember:              update.ChatMember,
		ChatJoinRequest:         update.ChatJoinRequest,
		ChatBoost:               update.ChatBoost,
		RemovedChatBoost:        update.RemovedChatBoost,
	}
}

// DecodeUpdate parses a Telegram update payload.
func DecodeUpdate(data []byte) (*Update, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("telegram update payload is empty")
	}

	var update Update
	if err := json.Unmarshal(data, &update); err != nil {
		return nil, fmt.Errorf("decode telegram update: %w", err)
	}
	return &update, nil
}

// DecodeUpdateFromReader parses a Telegram update payload from a reader.
func DecodeUpdateFromReader(reader io.Reader) (*Update, error) {
	if reader == nil {
		return nil, fmt.Errorf("telegram update reader is nil")
	}
	var update Update
	if err := json.NewDecoder(reader).Decode(&update); err != nil {
		return nil, fmt.Errorf("decode telegram update: %w", err)
	}
	return &update, nil
}

// Type returns the concrete type of the update.
func (update *Update) Type() UpdateType {
	if update == nil {
		return UpdateTypeUnknown
	}
	switch {
	case update.Message != nil:
		return UpdateTypeMessage
	case update.EditedMessage != nil:
		return UpdateTypeEditedMessage
	case update.ChannelPost != nil:
		return UpdateTypeChannelPost
	case update.EditedChannelPost != nil:
		return UpdateTypeEditedChannelPost
	case update.BusinessConnection != nil:
		return UpdateTypeBusinessConnection
	case update.BusinessMessage != nil:
		return UpdateTypeBusinessMessage
	case update.EditedBusinessMessage != nil:
		return UpdateTypeEditedBusinessMessage
	case update.DeletedBusinessMessages != nil:
		return UpdateTypeDeletedBusinessMessages
	case update.MessageReaction != nil:
		return UpdateTypeMessageReaction
	case update.MessageReactionCount != nil:
		return UpdateTypeMessageReactionCount
	case update.InlineQuery != nil:
		return UpdateTypeInlineQuery
	case update.ChosenInlineResult != nil:
		return UpdateTypeChosenInlineResult
	case update.CallbackQuery != nil:
		return UpdateTypeCallbackQuery
	case update.ShippingQuery != nil:
		return UpdateTypeShippingQuery
	case update.PreCheckoutQuery != nil:
		return UpdateTypePreCheckoutQuery
	case update.PurchasedPaidMedia != nil:
		return UpdateTypePurchasedPaidMedia
	case update.Poll != nil:
		return UpdateTypePoll
	case update.PollAnswer != nil:
		return UpdateTypePollAnswer
	case update.MyChatMember != nil:
		return UpdateTypeMyChatMember
	case update.ChatMember != nil:
		return UpdateTypeChatMember
	case update.ChatJoinRequest != nil:
		return UpdateTypeChatJoinRequest
	case update.ChatBoost != nil:
		return UpdateTypeChatBoost
	case update.RemovedChatBoost != nil:
		return UpdateTypeRemovedChatBoost
	default:
		return UpdateTypeUnknown
	}
}

// Payload returns the typed payload matching the update type.
func (update *Update) Payload() any {
	if update == nil {
		return nil
	}
	switch update.Type() {
	case UpdateTypeMessage:
		return update.Message
	case UpdateTypeEditedMessage:
		return update.EditedMessage
	case UpdateTypeChannelPost:
		return update.ChannelPost
	case UpdateTypeEditedChannelPost:
		return update.EditedChannelPost
	case UpdateTypeBusinessConnection:
		return update.BusinessConnection
	case UpdateTypeBusinessMessage:
		return update.BusinessMessage
	case UpdateTypeEditedBusinessMessage:
		return update.EditedBusinessMessage
	case UpdateTypeDeletedBusinessMessages:
		return update.DeletedBusinessMessages
	case UpdateTypeMessageReaction:
		return update.MessageReaction
	case UpdateTypeMessageReactionCount:
		return update.MessageReactionCount
	case UpdateTypeInlineQuery:
		return update.InlineQuery
	case UpdateTypeChosenInlineResult:
		return update.ChosenInlineResult
	case UpdateTypeCallbackQuery:
		return update.CallbackQuery
	case UpdateTypeShippingQuery:
		return update.ShippingQuery
	case UpdateTypePreCheckoutQuery:
		return update.PreCheckoutQuery
	case UpdateTypePurchasedPaidMedia:
		return update.PurchasedPaidMedia
	case UpdateTypePoll:
		return update.Poll
	case UpdateTypePollAnswer:
		return update.PollAnswer
	case UpdateTypeMyChatMember:
		return update.MyChatMember
	case UpdateTypeChatMember:
		return update.ChatMember
	case UpdateTypeChatJoinRequest:
		return update.ChatJoinRequest
	case UpdateTypeChatBoost:
		return update.ChatBoost
	case UpdateTypeRemovedChatBoost:
		return update.RemovedChatBoost
	default:
		return nil
	}
}

// EffectiveMessage returns the first message-like payload for routing helpers.
func (update *Update) EffectiveMessage() *tg.Message {
	if update == nil {
		return nil
	}
	switch {
	case update.Message != nil:
		return update.Message
	case update.EditedMessage != nil:
		return update.EditedMessage
	case update.ChannelPost != nil:
		return update.ChannelPost
	case update.EditedChannelPost != nil:
		return update.EditedChannelPost
	case update.BusinessMessage != nil:
		return update.BusinessMessage
	case update.EditedBusinessMessage != nil:
		return update.EditedBusinessMessage
	default:
		return nil
	}
}

// Command extracts command and args from the effective message text.
func (update *Update) Command() (string, string, bool) {
	message := update.EffectiveMessage()
	if message == nil {
		return "", "", false
	}
	return tg.ParseCommand(message.Text)
}
