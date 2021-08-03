package events

import (
	"github.com/DisgoOrg/disgo/api"
)

// GenericTextChannelEvent is called upon receiving TextChannelCreateEvent, TextChannelUpdateEvent or TextChannelDeleteEvent
type GenericTextChannelEvent struct {
	*GenericGuildChannelEvent
	TextChannel *api.TextChannel
}

// TextChannelCreateEvent indicates that a new api.TextChannel got created in an api.Guild
type TextChannelCreateEvent struct {
	*GenericTextChannelEvent
}

// TextChannelUpdateEvent indicates that an api.TextChannel got updated in an api.Guild
type TextChannelUpdateEvent struct {
	*GenericTextChannelEvent
	OldTextChannel *api.TextChannel
}

// TextChannelDeleteEvent indicates that an api.TextChannel got deleted in an api.Guild
type TextChannelDeleteEvent struct {
	*GenericTextChannelEvent
}

// WebhooksUpdateEvent indicates that an api.Webhook updated in this api.TextChannel
type WebhooksUpdateEvent struct {
	*GenericTextChannelEvent
}
