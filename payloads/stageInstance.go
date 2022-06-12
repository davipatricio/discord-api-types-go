package payloads

import "github.com/denkylabs/discord-api-types-go/globals"

// https://discord.com/developers/docs/resources/stage-instance#stage-instance-object
type APIStageInstance struct {
	// The id of the stage instance
	Id globals.Snowflake `json:"id"`
	// The guild id of the associated stage channel
	GuildId globals.Snowflake `json:"guild_id"`
	// The channel id of the associated stage channel
	ChannelId globals.Snowflake `json:"channel_id"`
	// The topic of the stage instance (1-120 characters)
	Topic string `json:"topic"`
	// The privacy level of the stage instance
	// See https://discord.com/developers/docs/resources/stage-instance#stage-instance-object-privacy-level
	PrivacyLevel StageInstancePrivacyLevel `json:"privacy_level"`
	// Whether or not stage discovery is disabled
	DiscoverableDisabled bool `json:"discoverable_disabled"`
	// The id of the scheduled event for this stage instance
	GuildScheduledEventId globals.Snowflake `json:"guild_scheduled_event_id"`
}

// https://discord.com/developers/docs/resources/stage-instance#stage-instance-object-privacy-level
type StageInstancePrivacyLevel uint8

// https://discord.com/developers/docs/resources/stage-instance#stage-instance-object-privacy-level
const (
	// The stage instance is visible publicly, such as on stage discovery
	StageInstancePrivacyLevelPublic StageInstancePrivacyLevel = 1
	// The stage instance is visible to only guild members
	StageInstancePrivacyLevelGuildOnly StageInstancePrivacyLevel = 2
)

// https://discord.com/developers/docs/resources/invite#invite-stage-instance-object-invite-stage-instance-structure
type APIInviteStageInstance struct {
	// The topic of the stage instance (1-120 characters)
	Topic string `json:"topic"`
	// The number of users in the stage
	ParticipantCount uint32 `json:"participant_count"`
	// The number of users speaking in the stage
	SpeakerCount uint32 `json:"speaker_count"`
	// The members speaking in the stage
	// See https://discord.com/developers/docs/resources/guild#guild-member-object-guild-member-structure
	Members []APIGuildMember `json:"members"`
}
