package tgbot

import "strings"

// ParseCommand parses Telegram-style command text like "/start arg".
func ParseCommand(text string) (command string, args string, ok bool) {
	trimmed := strings.TrimSpace(text)
	if !strings.HasPrefix(trimmed, "/") {
		return "", "", false
	}

	token := trimmed
	args = ""
	if split := strings.IndexAny(trimmed, " \t\n\r"); split >= 0 {
		token = trimmed[:split]
		args = strings.TrimSpace(trimmed[split+1:])
	}

	token = strings.TrimPrefix(strings.TrimSpace(token), "/")
	if token == "" {
		return "", "", false
	}
	if mention := strings.Index(token, "@"); mention >= 0 {
		token = token[:mention]
	}
	token = strings.ToLower(strings.TrimSpace(token))
	if token == "" {
		return "", "", false
	}
	return token, args, true
}
