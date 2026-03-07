package ext

import (
	"regexp"
	"strings"
)

// Filter decides whether a message-like update should be handled.
type Filter interface {
	Match(*Context) bool
}

// FilterFunc is an adapter for inline filter callbacks.
type FilterFunc func(*Context) bool

// Match evaluates the filter.
func (fn FilterFunc) Match(ctx *Context) bool {
	if fn == nil {
		return false
	}
	return fn(ctx)
}

// AnyFilter matches every context.
func AnyFilter() Filter {
	return FilterFunc(func(*Context) bool { return true })
}

// TextFilter matches updates whose effective message has non-empty text.
func TextFilter() Filter {
	return FilterFunc(func(ctx *Context) bool {
		if ctx == nil {
			return false
		}
		message := ctx.EffectiveMessage()
		return message != nil && strings.TrimSpace(message.Text) != ""
	})
}

// CommandFilter matches updates whose effective message contains a command.
func CommandFilter() Filter {
	return FilterFunc(func(ctx *Context) bool {
		if ctx == nil {
			return false
		}
		_, _, ok := ctx.Command()
		return ok
	})
}

// RegexFilter matches updates whose effective message text matches the regexp.
func RegexFilter(pattern *regexp.Regexp) Filter {
	return FilterFunc(func(ctx *Context) bool {
		if pattern == nil || ctx == nil {
			return false
		}
		message := ctx.EffectiveMessage()
		if message == nil {
			return false
		}
		return pattern.MatchString(message.Text)
	})
}

// UpdateTypeFilter matches the provided update types.
func UpdateTypeFilter(updateTypes ...UpdateType) Filter {
	set := make(map[UpdateType]struct{}, len(updateTypes))
	for _, updateType := range updateTypes {
		if updateType == UpdateTypeUnknown {
			continue
		}
		set[updateType] = struct{}{}
	}
	return FilterFunc(func(ctx *Context) bool {
		if ctx == nil || len(set) == 0 {
			return false
		}
		_, ok := set[ctx.UpdateType()]
		return ok
	})
}

// And matches when all provided filters match.
func And(filters ...Filter) Filter {
	return FilterFunc(func(ctx *Context) bool {
		if len(filters) == 0 {
			return false
		}
		for _, filter := range filters {
			if filter == nil || !filter.Match(ctx) {
				return false
			}
		}
		return true
	})
}

// Or matches when any provided filter matches.
func Or(filters ...Filter) Filter {
	return FilterFunc(func(ctx *Context) bool {
		for _, filter := range filters {
			if filter != nil && filter.Match(ctx) {
				return true
			}
		}
		return false
	})
}

// Not negates a filter.
func Not(filter Filter) Filter {
	return FilterFunc(func(ctx *Context) bool {
		if filter == nil {
			return true
		}
		return !filter.Match(ctx)
	})
}
