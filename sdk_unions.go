// Telegram Bot API union decoding helpers aligned with the official docs.
// Source: https://core.telegram.org/bots/api (Bot API 9.5, March 1, 2026)

package tgbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func decodeIntoKnownUnionResult(target any, raw json.RawMessage) (bool, error) {
	switch result := target.(type) {
	case *BackgroundFill:
		value, err := decodeBackgroundFill(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *[]BackgroundFill:
		value, err := decodeBackgroundFillSlice(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *BackgroundType:
		value, err := decodeBackgroundType(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *[]BackgroundType:
		value, err := decodeBackgroundTypeSlice(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *ChatBoostSource:
		value, err := decodeChatBoostSource(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *[]ChatBoostSource:
		value, err := decodeChatBoostSourceSlice(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *ChatMember:
		value, err := decodeChatMember(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *[]ChatMember:
		value, err := decodeChatMemberSlice(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *InputMessageContent:
		value, err := decodeInputMessageContent(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *[]InputMessageContent:
		value, err := decodeInputMessageContentSlice(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *MaybeInaccessibleMessage:
		value, err := decodeMaybeInaccessibleMessage(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *[]MaybeInaccessibleMessage:
		value, err := decodeMaybeInaccessibleMessageSlice(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *MenuButton:
		value, err := decodeMenuButton(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *[]MenuButton:
		value, err := decodeMenuButtonSlice(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *MessageOrigin:
		value, err := decodeMessageOrigin(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *[]MessageOrigin:
		value, err := decodeMessageOriginSlice(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *OwnedGift:
		value, err := decodeOwnedGift(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *[]OwnedGift:
		value, err := decodeOwnedGiftSlice(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *PaidMedia:
		value, err := decodePaidMedia(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *[]PaidMedia:
		value, err := decodePaidMediaSlice(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *ReactionType:
		value, err := decodeReactionType(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *[]ReactionType:
		value, err := decodeReactionTypeSlice(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *RevenueWithdrawalState:
		value, err := decodeRevenueWithdrawalState(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *[]RevenueWithdrawalState:
		value, err := decodeRevenueWithdrawalStateSlice(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *StoryAreaType:
		value, err := decodeStoryAreaType(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *[]StoryAreaType:
		value, err := decodeStoryAreaTypeSlice(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *TransactionPartner:
		value, err := decodeTransactionPartner(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	case *[]TransactionPartner:
		value, err := decodeTransactionPartnerSlice(raw)
		if err != nil {
			return true, err
		}
		*result = value
		return true, nil
	default:
		return false, nil
	}
}

func splitUnionFields(data []byte, fields ...string) (map[string]json.RawMessage, []byte, error) {
	if len(data) == 0 {
		return nil, []byte("{}"), nil
	}

	allFields := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &allFields); err != nil {
		return nil, nil, err
	}

	union := make(map[string]json.RawMessage, len(fields))
	for _, field := range fields {
		if raw, ok := allFields[field]; ok {
			union[field] = raw
			delete(allFields, field)
		}
	}

	base, err := json.Marshal(allFields)
	if err != nil {
		return nil, nil, err
	}
	if len(base) == 0 {
		base = []byte("{}")
	}
	return union, base, nil
}

func decodeStringTag(raw json.RawMessage, field string) (string, error) {
	probe := map[string]json.RawMessage{}
	if err := json.Unmarshal(raw, &probe); err != nil {
		return "", err
	}

	value, ok := probe[field]
	if !ok {
		return "", fmt.Errorf("missing discriminator %q", field)
	}

	var result string
	if err := json.Unmarshal(value, &result); err != nil {
		return "", err
	}
	return strings.TrimSpace(result), nil
}

func isJSONNull(raw json.RawMessage) bool {
	trimmed := bytes.TrimSpace(raw)
	return len(trimmed) == 0 || bytes.Equal(trimmed, []byte("null"))
}

func decodeRawSlice(raw json.RawMessage) ([]json.RawMessage, error) {
	if isJSONNull(raw) {
		return nil, nil
	}
	var values []json.RawMessage
	if err := json.Unmarshal(raw, &values); err != nil {
		return nil, err
	}
	return values, nil
}

func decodeBackgroundFill(raw json.RawMessage) (BackgroundFill, error) {
	if isJSONNull(raw) {
		return nil, nil
	}
	typeTag, err := decodeStringTag(raw, "type")
	if err != nil {
		return nil, err
	}
	switch typeTag {
	case "solid":
		var value BackgroundFillSolid
		return &value, json.Unmarshal(raw, &value)
	case "gradient":
		var value BackgroundFillGradient
		return &value, json.Unmarshal(raw, &value)
	case "freeform_gradient":
		var value BackgroundFillFreeformGradient
		return &value, json.Unmarshal(raw, &value)
	default:
		return nil, fmt.Errorf("unknown BackgroundFill discriminator %q", typeTag)
	}
}

func decodeBackgroundFillSlice(raw json.RawMessage) ([]BackgroundFill, error) {
	rawItems, err := decodeRawSlice(raw)
	if err != nil {
		return nil, err
	}
	result := make([]BackgroundFill, 0, len(rawItems))
	for _, item := range rawItems {
		value, err := decodeBackgroundFill(item)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}

func decodeBackgroundType(raw json.RawMessage) (BackgroundType, error) {
	if isJSONNull(raw) {
		return nil, nil
	}
	typeTag, err := decodeStringTag(raw, "type")
	if err != nil {
		return nil, err
	}
	switch typeTag {
	case "fill":
		var value BackgroundTypeFill
		return &value, json.Unmarshal(raw, &value)
	case "wallpaper":
		var value BackgroundTypeWallpaper
		return &value, json.Unmarshal(raw, &value)
	case "pattern":
		var value BackgroundTypePattern
		return &value, json.Unmarshal(raw, &value)
	case "chat_theme":
		var value BackgroundTypeChatTheme
		return &value, json.Unmarshal(raw, &value)
	default:
		return nil, fmt.Errorf("unknown BackgroundType discriminator %q", typeTag)
	}
}

func decodeBackgroundTypeSlice(raw json.RawMessage) ([]BackgroundType, error) {
	rawItems, err := decodeRawSlice(raw)
	if err != nil {
		return nil, err
	}
	result := make([]BackgroundType, 0, len(rawItems))
	for _, item := range rawItems {
		value, err := decodeBackgroundType(item)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}

func decodeChatBoostSource(raw json.RawMessage) (ChatBoostSource, error) {
	if isJSONNull(raw) {
		return nil, nil
	}
	typeTag, err := decodeStringTag(raw, "source")
	if err != nil {
		return nil, err
	}
	switch typeTag {
	case "premium":
		var value ChatBoostSourcePremium
		return &value, json.Unmarshal(raw, &value)
	case "gift_code":
		var value ChatBoostSourceGiftCode
		return &value, json.Unmarshal(raw, &value)
	case "giveaway":
		var value ChatBoostSourceGiveaway
		return &value, json.Unmarshal(raw, &value)
	default:
		return nil, fmt.Errorf("unknown ChatBoostSource discriminator %q", typeTag)
	}
}

func decodeChatBoostSourceSlice(raw json.RawMessage) ([]ChatBoostSource, error) {
	rawItems, err := decodeRawSlice(raw)
	if err != nil {
		return nil, err
	}
	result := make([]ChatBoostSource, 0, len(rawItems))
	for _, item := range rawItems {
		value, err := decodeChatBoostSource(item)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}

func decodeChatMember(raw json.RawMessage) (ChatMember, error) {
	if isJSONNull(raw) {
		return nil, nil
	}
	typeTag, err := decodeStringTag(raw, "status")
	if err != nil {
		return nil, err
	}
	switch typeTag {
	case "creator":
		var value ChatMemberOwner
		return &value, json.Unmarshal(raw, &value)
	case "administrator":
		var value ChatMemberAdministrator
		return &value, json.Unmarshal(raw, &value)
	case "member":
		var value ChatMemberMember
		return &value, json.Unmarshal(raw, &value)
	case "restricted":
		var value ChatMemberRestricted
		return &value, json.Unmarshal(raw, &value)
	case "left":
		var value ChatMemberLeft
		return &value, json.Unmarshal(raw, &value)
	case "kicked":
		var value ChatMemberBanned
		return &value, json.Unmarshal(raw, &value)
	default:
		return nil, fmt.Errorf("unknown ChatMember discriminator %q", typeTag)
	}
}

func decodeChatMemberSlice(raw json.RawMessage) ([]ChatMember, error) {
	rawItems, err := decodeRawSlice(raw)
	if err != nil {
		return nil, err
	}
	result := make([]ChatMember, 0, len(rawItems))
	for _, item := range rawItems {
		value, err := decodeChatMember(item)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}

func decodeInputMessageContent(raw json.RawMessage) (InputMessageContent, error) {
	if isJSONNull(raw) {
		return nil, nil
	}
	probe := map[string]json.RawMessage{}
	if err := json.Unmarshal(raw, &probe); err != nil {
		return nil, err
	}
	if _, ok := probe["message_text"]; ok {
		var value InputTextMessageContent
		return &value, json.Unmarshal(raw, &value)
	}
	if _, ok := probe["phone_number"]; ok {
		var value InputContactMessageContent
		return &value, json.Unmarshal(raw, &value)
	}
	if _, ok := probe["payload"]; ok {
		var value InputInvoiceMessageContent
		return &value, json.Unmarshal(raw, &value)
	}
	if _, ok := probe["prices"]; ok {
		var value InputInvoiceMessageContent
		return &value, json.Unmarshal(raw, &value)
	}
	if _, hasLatitude := probe["latitude"]; hasLatitude {
		if _, hasLongitude := probe["longitude"]; hasLongitude {
			if _, hasAddress := probe["address"]; hasAddress {
				var value InputVenueMessageContent
				return &value, json.Unmarshal(raw, &value)
			}
			var value InputLocationMessageContent
			return &value, json.Unmarshal(raw, &value)
		}
	}
	return nil, fmt.Errorf("unknown input message content shape")
}

func decodeInputMessageContentSlice(raw json.RawMessage) ([]InputMessageContent, error) {
	rawItems, err := decodeRawSlice(raw)
	if err != nil {
		return nil, err
	}
	result := make([]InputMessageContent, 0, len(rawItems))
	for _, item := range rawItems {
		value, err := decodeInputMessageContent(item)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}

func decodeMaybeInaccessibleMessage(raw json.RawMessage) (MaybeInaccessibleMessage, error) {
	if isJSONNull(raw) {
		return nil, nil
	}
	probe := struct {
		Date int64 `json:"date"`
	}{}
	if err := json.Unmarshal(raw, &probe); err != nil {
		return nil, err
	}
	if probe.Date == 0 {
		var value InaccessibleMessage
		return &value, json.Unmarshal(raw, &value)
	}
	var value Message
	return &value, json.Unmarshal(raw, &value)
}

func decodeMaybeInaccessibleMessageSlice(raw json.RawMessage) ([]MaybeInaccessibleMessage, error) {
	rawItems, err := decodeRawSlice(raw)
	if err != nil {
		return nil, err
	}
	result := make([]MaybeInaccessibleMessage, 0, len(rawItems))
	for _, item := range rawItems {
		value, err := decodeMaybeInaccessibleMessage(item)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}

func decodeMenuButton(raw json.RawMessage) (MenuButton, error) {
	if isJSONNull(raw) {
		return nil, nil
	}
	typeTag, err := decodeStringTag(raw, "type")
	if err != nil {
		return nil, err
	}
	switch typeTag {
	case "commands":
		var value MenuButtonCommands
		return &value, json.Unmarshal(raw, &value)
	case "web_app":
		var value MenuButtonWebApp
		return &value, json.Unmarshal(raw, &value)
	case "default":
		var value MenuButtonDefault
		return &value, json.Unmarshal(raw, &value)
	default:
		return nil, fmt.Errorf("unknown MenuButton discriminator %q", typeTag)
	}
}

func decodeMenuButtonSlice(raw json.RawMessage) ([]MenuButton, error) {
	rawItems, err := decodeRawSlice(raw)
	if err != nil {
		return nil, err
	}
	result := make([]MenuButton, 0, len(rawItems))
	for _, item := range rawItems {
		value, err := decodeMenuButton(item)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}

func decodeMessageOrigin(raw json.RawMessage) (MessageOrigin, error) {
	if isJSONNull(raw) {
		return nil, nil
	}
	typeTag, err := decodeStringTag(raw, "type")
	if err != nil {
		return nil, err
	}
	switch typeTag {
	case "user":
		var value MessageOriginUser
		return &value, json.Unmarshal(raw, &value)
	case "hidden_user":
		var value MessageOriginHiddenUser
		return &value, json.Unmarshal(raw, &value)
	case "chat":
		var value MessageOriginChat
		return &value, json.Unmarshal(raw, &value)
	case "channel":
		var value MessageOriginChannel
		return &value, json.Unmarshal(raw, &value)
	default:
		return nil, fmt.Errorf("unknown MessageOrigin discriminator %q", typeTag)
	}
}

func decodeMessageOriginSlice(raw json.RawMessage) ([]MessageOrigin, error) {
	rawItems, err := decodeRawSlice(raw)
	if err != nil {
		return nil, err
	}
	result := make([]MessageOrigin, 0, len(rawItems))
	for _, item := range rawItems {
		value, err := decodeMessageOrigin(item)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}

func decodeOwnedGift(raw json.RawMessage) (OwnedGift, error) {
	if isJSONNull(raw) {
		return nil, nil
	}
	typeTag, err := decodeStringTag(raw, "type")
	if err != nil {
		return nil, err
	}
	switch typeTag {
	case "regular":
		var value OwnedGiftRegular
		return &value, json.Unmarshal(raw, &value)
	case "unique":
		var value OwnedGiftUnique
		return &value, json.Unmarshal(raw, &value)
	default:
		return nil, fmt.Errorf("unknown OwnedGift discriminator %q", typeTag)
	}
}

func decodeOwnedGiftSlice(raw json.RawMessage) ([]OwnedGift, error) {
	rawItems, err := decodeRawSlice(raw)
	if err != nil {
		return nil, err
	}
	result := make([]OwnedGift, 0, len(rawItems))
	for _, item := range rawItems {
		value, err := decodeOwnedGift(item)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}

func decodePaidMedia(raw json.RawMessage) (PaidMedia, error) {
	if isJSONNull(raw) {
		return nil, nil
	}
	typeTag, err := decodeStringTag(raw, "type")
	if err != nil {
		return nil, err
	}
	switch typeTag {
	case "preview":
		var value PaidMediaPreview
		return &value, json.Unmarshal(raw, &value)
	case "photo":
		var value PaidMediaPhoto
		return &value, json.Unmarshal(raw, &value)
	case "video":
		var value PaidMediaVideo
		return &value, json.Unmarshal(raw, &value)
	default:
		return nil, fmt.Errorf("unknown PaidMedia discriminator %q", typeTag)
	}
}

func decodePaidMediaSlice(raw json.RawMessage) ([]PaidMedia, error) {
	rawItems, err := decodeRawSlice(raw)
	if err != nil {
		return nil, err
	}
	result := make([]PaidMedia, 0, len(rawItems))
	for _, item := range rawItems {
		value, err := decodePaidMedia(item)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}

func decodeReactionType(raw json.RawMessage) (ReactionType, error) {
	if isJSONNull(raw) {
		return nil, nil
	}
	typeTag, err := decodeStringTag(raw, "type")
	if err != nil {
		return nil, err
	}
	switch typeTag {
	case "emoji":
		var value ReactionTypeEmoji
		return &value, json.Unmarshal(raw, &value)
	case "custom_emoji":
		var value ReactionTypeCustomEmoji
		return &value, json.Unmarshal(raw, &value)
	case "paid":
		var value ReactionTypePaid
		return &value, json.Unmarshal(raw, &value)
	default:
		return nil, fmt.Errorf("unknown ReactionType discriminator %q", typeTag)
	}
}

func decodeReactionTypeSlice(raw json.RawMessage) ([]ReactionType, error) {
	rawItems, err := decodeRawSlice(raw)
	if err != nil {
		return nil, err
	}
	result := make([]ReactionType, 0, len(rawItems))
	for _, item := range rawItems {
		value, err := decodeReactionType(item)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}

func decodeRevenueWithdrawalState(raw json.RawMessage) (RevenueWithdrawalState, error) {
	if isJSONNull(raw) {
		return nil, nil
	}
	typeTag, err := decodeStringTag(raw, "type")
	if err != nil {
		return nil, err
	}
	switch typeTag {
	case "pending":
		var value RevenueWithdrawalStatePending
		return &value, json.Unmarshal(raw, &value)
	case "succeeded":
		var value RevenueWithdrawalStateSucceeded
		return &value, json.Unmarshal(raw, &value)
	case "failed":
		var value RevenueWithdrawalStateFailed
		return &value, json.Unmarshal(raw, &value)
	default:
		return nil, fmt.Errorf("unknown RevenueWithdrawalState discriminator %q", typeTag)
	}
}

func decodeRevenueWithdrawalStateSlice(raw json.RawMessage) ([]RevenueWithdrawalState, error) {
	rawItems, err := decodeRawSlice(raw)
	if err != nil {
		return nil, err
	}
	result := make([]RevenueWithdrawalState, 0, len(rawItems))
	for _, item := range rawItems {
		value, err := decodeRevenueWithdrawalState(item)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}

func decodeStoryAreaType(raw json.RawMessage) (StoryAreaType, error) {
	if isJSONNull(raw) {
		return nil, nil
	}
	typeTag, err := decodeStringTag(raw, "type")
	if err != nil {
		return nil, err
	}
	switch typeTag {
	case "location":
		var value StoryAreaTypeLocation
		return &value, json.Unmarshal(raw, &value)
	case "suggested_reaction":
		var value StoryAreaTypeSuggestedReaction
		return &value, json.Unmarshal(raw, &value)
	case "link":
		var value StoryAreaTypeLink
		return &value, json.Unmarshal(raw, &value)
	case "weather":
		var value StoryAreaTypeWeather
		return &value, json.Unmarshal(raw, &value)
	case "unique_gift":
		var value StoryAreaTypeUniqueGift
		return &value, json.Unmarshal(raw, &value)
	default:
		return nil, fmt.Errorf("unknown StoryAreaType discriminator %q", typeTag)
	}
}

func decodeStoryAreaTypeSlice(raw json.RawMessage) ([]StoryAreaType, error) {
	rawItems, err := decodeRawSlice(raw)
	if err != nil {
		return nil, err
	}
	result := make([]StoryAreaType, 0, len(rawItems))
	for _, item := range rawItems {
		value, err := decodeStoryAreaType(item)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}

func decodeTransactionPartner(raw json.RawMessage) (TransactionPartner, error) {
	if isJSONNull(raw) {
		return nil, nil
	}
	typeTag, err := decodeStringTag(raw, "type")
	if err != nil {
		return nil, err
	}
	switch typeTag {
	case "user":
		var value TransactionPartnerUser
		return &value, json.Unmarshal(raw, &value)
	case "chat":
		var value TransactionPartnerChat
		return &value, json.Unmarshal(raw, &value)
	case "affiliate_program":
		var value TransactionPartnerAffiliateProgram
		return &value, json.Unmarshal(raw, &value)
	case "fragment":
		var value TransactionPartnerFragment
		return &value, json.Unmarshal(raw, &value)
	case "telegram_ads":
		var value TransactionPartnerTelegramAds
		return &value, json.Unmarshal(raw, &value)
	case "telegram_api":
		var value TransactionPartnerTelegramApi
		return &value, json.Unmarshal(raw, &value)
	case "other":
		var value TransactionPartnerOther
		return &value, json.Unmarshal(raw, &value)
	default:
		return nil, fmt.Errorf("unknown TransactionPartner discriminator %q", typeTag)
	}
}

func decodeTransactionPartnerSlice(raw json.RawMessage) ([]TransactionPartner, error) {
	rawItems, err := decodeRawSlice(raw)
	if err != nil {
		return nil, err
	}
	result := make([]TransactionPartner, 0, len(rawItems))
	for _, item := range rawItems {
		value, err := decodeTransactionPartner(item)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}

func (value *BackgroundTypeFill) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "fill")
	if err != nil {
		return err
	}
	type alias BackgroundTypeFill
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = BackgroundTypeFill(temp)
	if raw, ok := union["fill"]; ok {
		decoded, err := decodeBackgroundFill(raw)
		if err != nil {
			return err
		}
		value.Fill = decoded
	}
	return nil
}

func (value *BackgroundTypePattern) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "fill")
	if err != nil {
		return err
	}
	type alias BackgroundTypePattern
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = BackgroundTypePattern(temp)
	if raw, ok := union["fill"]; ok {
		decoded, err := decodeBackgroundFill(raw)
		if err != nil {
			return err
		}
		value.Fill = decoded
	}
	return nil
}

func (value *CallbackQuery) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "message")
	if err != nil {
		return err
	}
	type alias CallbackQuery
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = CallbackQuery(temp)
	if raw, ok := union["message"]; ok {
		decoded, err := decodeMaybeInaccessibleMessage(raw)
		if err != nil {
			return err
		}
		value.Message = decoded
	}
	return nil
}

func (value *ChatBackground) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "type")
	if err != nil {
		return err
	}
	type alias ChatBackground
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = ChatBackground(temp)
	if raw, ok := union["type"]; ok {
		decoded, err := decodeBackgroundType(raw)
		if err != nil {
			return err
		}
		value.Type = decoded
	}
	return nil
}

func (value *ChatBoost) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "source")
	if err != nil {
		return err
	}
	type alias ChatBoost
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = ChatBoost(temp)
	if raw, ok := union["source"]; ok {
		decoded, err := decodeChatBoostSource(raw)
		if err != nil {
			return err
		}
		value.Source = decoded
	}
	return nil
}

func (value *ChatBoostRemoved) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "source")
	if err != nil {
		return err
	}
	type alias ChatBoostRemoved
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = ChatBoostRemoved(temp)
	if raw, ok := union["source"]; ok {
		decoded, err := decodeChatBoostSource(raw)
		if err != nil {
			return err
		}
		value.Source = decoded
	}
	return nil
}

func (value *ChatFullInfo) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "available_reactions")
	if err != nil {
		return err
	}
	type alias ChatFullInfo
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = ChatFullInfo(temp)
	if raw, ok := union["available_reactions"]; ok {
		decoded, err := decodeReactionTypeSlice(raw)
		if err != nil {
			return err
		}
		value.AvailableReactions = decoded
	}
	return nil
}

func (value *ChatMemberUpdated) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "old_chat_member", "new_chat_member")
	if err != nil {
		return err
	}
	type alias ChatMemberUpdated
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = ChatMemberUpdated(temp)
	if raw, ok := union["old_chat_member"]; ok {
		decoded, err := decodeChatMember(raw)
		if err != nil {
			return err
		}
		value.OldChatMember = decoded
	}
	if raw, ok := union["new_chat_member"]; ok {
		decoded, err := decodeChatMember(raw)
		if err != nil {
			return err
		}
		value.NewChatMember = decoded
	}
	return nil
}

func (value *ExternalReplyInfo) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "origin")
	if err != nil {
		return err
	}
	type alias ExternalReplyInfo
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = ExternalReplyInfo(temp)
	if raw, ok := union["origin"]; ok {
		decoded, err := decodeMessageOrigin(raw)
		if err != nil {
			return err
		}
		value.Origin = decoded
	}
	return nil
}

func (value *InlineQueryResultArticle) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "input_message_content")
	if err != nil {
		return err
	}
	type alias InlineQueryResultArticle
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = InlineQueryResultArticle(temp)
	if raw, ok := union["input_message_content"]; ok {
		decoded, err := decodeInputMessageContent(raw)
		if err != nil {
			return err
		}
		value.InputMessageContent = decoded
	}
	return nil
}

func (value *InlineQueryResultAudio) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "input_message_content")
	if err != nil {
		return err
	}
	type alias InlineQueryResultAudio
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = InlineQueryResultAudio(temp)
	if raw, ok := union["input_message_content"]; ok {
		decoded, err := decodeInputMessageContent(raw)
		if err != nil {
			return err
		}
		value.InputMessageContent = decoded
	}
	return nil
}

func (value *InlineQueryResultCachedAudio) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "input_message_content")
	if err != nil {
		return err
	}
	type alias InlineQueryResultCachedAudio
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = InlineQueryResultCachedAudio(temp)
	if raw, ok := union["input_message_content"]; ok {
		decoded, err := decodeInputMessageContent(raw)
		if err != nil {
			return err
		}
		value.InputMessageContent = decoded
	}
	return nil
}

func (value *InlineQueryResultCachedDocument) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "input_message_content")
	if err != nil {
		return err
	}
	type alias InlineQueryResultCachedDocument
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = InlineQueryResultCachedDocument(temp)
	if raw, ok := union["input_message_content"]; ok {
		decoded, err := decodeInputMessageContent(raw)
		if err != nil {
			return err
		}
		value.InputMessageContent = decoded
	}
	return nil
}

func (value *InlineQueryResultCachedGif) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "input_message_content")
	if err != nil {
		return err
	}
	type alias InlineQueryResultCachedGif
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = InlineQueryResultCachedGif(temp)
	if raw, ok := union["input_message_content"]; ok {
		decoded, err := decodeInputMessageContent(raw)
		if err != nil {
			return err
		}
		value.InputMessageContent = decoded
	}
	return nil
}

func (value *InlineQueryResultCachedMpeg4Gif) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "input_message_content")
	if err != nil {
		return err
	}
	type alias InlineQueryResultCachedMpeg4Gif
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = InlineQueryResultCachedMpeg4Gif(temp)
	if raw, ok := union["input_message_content"]; ok {
		decoded, err := decodeInputMessageContent(raw)
		if err != nil {
			return err
		}
		value.InputMessageContent = decoded
	}
	return nil
}

func (value *InlineQueryResultCachedPhoto) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "input_message_content")
	if err != nil {
		return err
	}
	type alias InlineQueryResultCachedPhoto
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = InlineQueryResultCachedPhoto(temp)
	if raw, ok := union["input_message_content"]; ok {
		decoded, err := decodeInputMessageContent(raw)
		if err != nil {
			return err
		}
		value.InputMessageContent = decoded
	}
	return nil
}

func (value *InlineQueryResultCachedSticker) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "input_message_content")
	if err != nil {
		return err
	}
	type alias InlineQueryResultCachedSticker
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = InlineQueryResultCachedSticker(temp)
	if raw, ok := union["input_message_content"]; ok {
		decoded, err := decodeInputMessageContent(raw)
		if err != nil {
			return err
		}
		value.InputMessageContent = decoded
	}
	return nil
}

func (value *InlineQueryResultCachedVideo) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "input_message_content")
	if err != nil {
		return err
	}
	type alias InlineQueryResultCachedVideo
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = InlineQueryResultCachedVideo(temp)
	if raw, ok := union["input_message_content"]; ok {
		decoded, err := decodeInputMessageContent(raw)
		if err != nil {
			return err
		}
		value.InputMessageContent = decoded
	}
	return nil
}

func (value *InlineQueryResultCachedVoice) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "input_message_content")
	if err != nil {
		return err
	}
	type alias InlineQueryResultCachedVoice
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = InlineQueryResultCachedVoice(temp)
	if raw, ok := union["input_message_content"]; ok {
		decoded, err := decodeInputMessageContent(raw)
		if err != nil {
			return err
		}
		value.InputMessageContent = decoded
	}
	return nil
}

func (value *InlineQueryResultContact) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "input_message_content")
	if err != nil {
		return err
	}
	type alias InlineQueryResultContact
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = InlineQueryResultContact(temp)
	if raw, ok := union["input_message_content"]; ok {
		decoded, err := decodeInputMessageContent(raw)
		if err != nil {
			return err
		}
		value.InputMessageContent = decoded
	}
	return nil
}

func (value *InlineQueryResultDocument) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "input_message_content")
	if err != nil {
		return err
	}
	type alias InlineQueryResultDocument
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = InlineQueryResultDocument(temp)
	if raw, ok := union["input_message_content"]; ok {
		decoded, err := decodeInputMessageContent(raw)
		if err != nil {
			return err
		}
		value.InputMessageContent = decoded
	}
	return nil
}

func (value *InlineQueryResultGif) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "input_message_content")
	if err != nil {
		return err
	}
	type alias InlineQueryResultGif
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = InlineQueryResultGif(temp)
	if raw, ok := union["input_message_content"]; ok {
		decoded, err := decodeInputMessageContent(raw)
		if err != nil {
			return err
		}
		value.InputMessageContent = decoded
	}
	return nil
}

func (value *InlineQueryResultLocation) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "input_message_content")
	if err != nil {
		return err
	}
	type alias InlineQueryResultLocation
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = InlineQueryResultLocation(temp)
	if raw, ok := union["input_message_content"]; ok {
		decoded, err := decodeInputMessageContent(raw)
		if err != nil {
			return err
		}
		value.InputMessageContent = decoded
	}
	return nil
}

func (value *InlineQueryResultMpeg4Gif) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "input_message_content")
	if err != nil {
		return err
	}
	type alias InlineQueryResultMpeg4Gif
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = InlineQueryResultMpeg4Gif(temp)
	if raw, ok := union["input_message_content"]; ok {
		decoded, err := decodeInputMessageContent(raw)
		if err != nil {
			return err
		}
		value.InputMessageContent = decoded
	}
	return nil
}

func (value *InlineQueryResultPhoto) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "input_message_content")
	if err != nil {
		return err
	}
	type alias InlineQueryResultPhoto
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = InlineQueryResultPhoto(temp)
	if raw, ok := union["input_message_content"]; ok {
		decoded, err := decodeInputMessageContent(raw)
		if err != nil {
			return err
		}
		value.InputMessageContent = decoded
	}
	return nil
}

func (value *InlineQueryResultVenue) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "input_message_content")
	if err != nil {
		return err
	}
	type alias InlineQueryResultVenue
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = InlineQueryResultVenue(temp)
	if raw, ok := union["input_message_content"]; ok {
		decoded, err := decodeInputMessageContent(raw)
		if err != nil {
			return err
		}
		value.InputMessageContent = decoded
	}
	return nil
}

func (value *InlineQueryResultVideo) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "input_message_content")
	if err != nil {
		return err
	}
	type alias InlineQueryResultVideo
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = InlineQueryResultVideo(temp)
	if raw, ok := union["input_message_content"]; ok {
		decoded, err := decodeInputMessageContent(raw)
		if err != nil {
			return err
		}
		value.InputMessageContent = decoded
	}
	return nil
}

func (value *InlineQueryResultVoice) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "input_message_content")
	if err != nil {
		return err
	}
	type alias InlineQueryResultVoice
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = InlineQueryResultVoice(temp)
	if raw, ok := union["input_message_content"]; ok {
		decoded, err := decodeInputMessageContent(raw)
		if err != nil {
			return err
		}
		value.InputMessageContent = decoded
	}
	return nil
}

func (value *Message) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "forward_origin", "pinned_message")
	if err != nil {
		return err
	}
	type alias Message
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = Message(temp)
	if raw, ok := union["forward_origin"]; ok {
		decoded, err := decodeMessageOrigin(raw)
		if err != nil {
			return err
		}
		value.ForwardOrigin = decoded
	}
	if raw, ok := union["pinned_message"]; ok {
		decoded, err := decodeMaybeInaccessibleMessage(raw)
		if err != nil {
			return err
		}
		value.PinnedMessage = decoded
	}
	return nil
}

func (value *MessageReactionUpdated) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "old_reaction", "new_reaction")
	if err != nil {
		return err
	}
	type alias MessageReactionUpdated
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = MessageReactionUpdated(temp)
	if raw, ok := union["old_reaction"]; ok {
		decoded, err := decodeReactionTypeSlice(raw)
		if err != nil {
			return err
		}
		value.OldReaction = decoded
	}
	if raw, ok := union["new_reaction"]; ok {
		decoded, err := decodeReactionTypeSlice(raw)
		if err != nil {
			return err
		}
		value.NewReaction = decoded
	}
	return nil
}

func (value *OwnedGifts) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "gifts")
	if err != nil {
		return err
	}
	type alias OwnedGifts
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = OwnedGifts(temp)
	if raw, ok := union["gifts"]; ok {
		decoded, err := decodeOwnedGiftSlice(raw)
		if err != nil {
			return err
		}
		value.Gifts = decoded
	}
	return nil
}

func (value *PaidMediaInfo) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "paid_media")
	if err != nil {
		return err
	}
	type alias PaidMediaInfo
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = PaidMediaInfo(temp)
	if raw, ok := union["paid_media"]; ok {
		decoded, err := decodePaidMediaSlice(raw)
		if err != nil {
			return err
		}
		value.PaidMedia = decoded
	}
	return nil
}

func (value *ReactionCount) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "type")
	if err != nil {
		return err
	}
	type alias ReactionCount
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = ReactionCount(temp)
	if raw, ok := union["type"]; ok {
		decoded, err := decodeReactionType(raw)
		if err != nil {
			return err
		}
		value.Type = decoded
	}
	return nil
}

func (value *StarTransaction) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "source", "receiver")
	if err != nil {
		return err
	}
	type alias StarTransaction
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = StarTransaction(temp)
	if raw, ok := union["source"]; ok {
		decoded, err := decodeTransactionPartner(raw)
		if err != nil {
			return err
		}
		value.Source = decoded
	}
	if raw, ok := union["receiver"]; ok {
		decoded, err := decodeTransactionPartner(raw)
		if err != nil {
			return err
		}
		value.Receiver = decoded
	}
	return nil
}

func (value *StoryArea) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "type")
	if err != nil {
		return err
	}
	type alias StoryArea
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = StoryArea(temp)
	if raw, ok := union["type"]; ok {
		decoded, err := decodeStoryAreaType(raw)
		if err != nil {
			return err
		}
		value.Type = decoded
	}
	return nil
}

func (value *StoryAreaTypeSuggestedReaction) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "reaction_type")
	if err != nil {
		return err
	}
	type alias StoryAreaTypeSuggestedReaction
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = StoryAreaTypeSuggestedReaction(temp)
	if raw, ok := union["reaction_type"]; ok {
		decoded, err := decodeReactionType(raw)
		if err != nil {
			return err
		}
		value.ReactionType = decoded
	}
	return nil
}

func (value *TransactionPartnerFragment) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "withdrawal_state")
	if err != nil {
		return err
	}
	type alias TransactionPartnerFragment
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = TransactionPartnerFragment(temp)
	if raw, ok := union["withdrawal_state"]; ok {
		decoded, err := decodeRevenueWithdrawalState(raw)
		if err != nil {
			return err
		}
		value.WithdrawalState = decoded
	}
	return nil
}

func (value *TransactionPartnerUser) UnmarshalJSON(data []byte) error {
	union, base, err := splitUnionFields(data, "paid_media")
	if err != nil {
		return err
	}
	type alias TransactionPartnerUser
	var temp alias
	if err := json.Unmarshal(base, &temp); err != nil {
		return err
	}
	*value = TransactionPartnerUser(temp)
	if raw, ok := union["paid_media"]; ok {
		decoded, err := decodePaidMediaSlice(raw)
		if err != nil {
			return err
		}
		value.PaidMedia = decoded
	}
	return nil
}
