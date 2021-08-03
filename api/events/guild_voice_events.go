package events

import (
	"github.com/DisgoOrg/disgo/api"
)

// GenericGuildVoiceEvent is called upon receiving GuildVoiceJoinEvent, GuildVoiceUpdateEvent, GuildVoiceLeaveEvent
type GenericGuildVoiceEvent struct {
	*GenericGuildMemberEvent
	VoiceState *api.VoiceState
}

// GuildVoiceJoinEvent indicates that an api.Member joined an api.VoiceChannel(requires api.GatewayIntentsGuildVoiceStates)
type GuildVoiceJoinEvent struct {
	*GenericGuildVoiceEvent
}

// GuildVoiceUpdateEvent indicates that an api.Member moved an api.VoiceChannel(requires api.GatewayIntentsGuildVoiceStates)
type GuildVoiceUpdateEvent struct {
	*GenericGuildVoiceEvent
	OldVoiceState *api.VoiceState
}

// GuildVoiceLeaveEvent indicates that an api.Member left an api.VoiceChannel(requires api.GatewayIntentsGuildVoiceStates)
type GuildVoiceLeaveEvent struct {
	*GenericGuildVoiceEvent
}
