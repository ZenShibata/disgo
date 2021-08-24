package events

import (
	"github.com/DisgoOrg/disgo/core"
	"github.com/DisgoOrg/disgo/discord"
)

// GenericGuildChannelEvent is called upon receiving GuildChannelCreateEvent, GuildChannelUpdateEvent or GuildChannelDeleteEvent
type GenericGuildChannelEvent struct {
	*GenericChannelEvent
	GuildID      discord.Snowflake
	GuildChannel core.GuildChannel
}

// Guild returns the cached api.Guild the event happened in
func (e GenericGuildChannelEvent) Guild() *core.Guild {
	return e.Disgo().Cache().GuildCache().Get(e.GuildID)
}

// GuildChannelCreateEvent indicates that a new api.GetGuildChannel got created in an api.Guild
type GuildChannelCreateEvent struct {
	*GenericGuildChannelEvent
}

// GuildChannelUpdateEvent indicates that an api.GetGuildChannel got updated in an api.Guild
type GuildChannelUpdateEvent struct {
	*GenericGuildChannelEvent
	OldGuildChannel core.GuildChannel
}

type GuildChannelPinsUpdateEvent struct {
	*GenericGuildChannelEvent
	OldLastPinTimestamp *discord.Time
	NewLastPinTimestamp *discord.Time
}

// GuildChannelDeleteEvent indicates that an api.GetGuildChannel got deleted in an api.Guild
type GuildChannelDeleteEvent struct {
	*GenericGuildChannelEvent
}
