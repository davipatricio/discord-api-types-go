package gateway

import (
	"github.com/denkylabs/discord-api-types-go/globals"
	"github.com/denkylabs/discord-api-types-go/payloads/emojis"
	"github.com/denkylabs/discord-api-types-go/payloads/user"
)

// https://discord.com/developers/docs/topics/gateway#get-gateway
type APIGatewayInfo struct {
	// The WSS URL that can be used for connecting to the gateway
	Url string `json:"url"`
}

// https://discord.com/developers/docs/topics/gateway#get-gateway-bot
type APIGatewayBotInfo struct {
	// The WSS URL that can be used for connecting to the gateway
	Url string `json:"url"`
	// The recommended number of shards to use when connecting
	// See https://discord.com/developers/docs/topics/gateway#sharding
	Shards uint32 `json:"shards"`
	// Information on the current session start limit
	// See https://discord.com/developers/docs/topics/gateway#session-start-limit-object
	SessionStartLimit APIGatewaySessionStartLimit `json:"session_start_limit"`
}

// https://discord.com/developers/docs/topics/gateway#session-start-limit-object
type APIGatewaySessionStartLimit struct {
	// The total number of session starts the current user is allowed
	Total uint32 `json:"total"`
	// The number of session starts the current user has left
	Remaining uint32 `json:"remaining"`
	// The number of milliseconds after which the limit resets
	ResetAfter uint32 `json:"reset_after"`
	// The number of identify requests allowed per 5 seconds
	MaxConcurrency uint32 `json:"max_concurrency"`
}

// https://discord.com/developers/docs/topics/gateway#presence-update-presence-update-event-fields
type GatewayPresenceUpdate struct {
	// The user presence is being updated for
	// The user object within this event can be partial, the only field which must be sent is the `id` field,
	// everything else is optional.
	// See https://discord.com/developers/docs/resources/user#user-object
	User user.APIUser `json:"user"`
	// ID of the guild
	GuildId globals.Snowflake `json:"guild_id"`
	// Either "idle", "dnd", "online", or "offline"
	Status PresenceUpdateStatus `json:"status"`
	// User's current activities
	// See See https://discord.com/developers/docs/topics/gateway#activity-object
	Activities []GatewayActivity `json:"activities"`
	// User's platform-dependent status
	// See https://discord.com/developers/docs/topics/gateway#client-status-object
	ClientStatus GatewayPresenceClientStatus `json:"client_status"`
}

type PresenceUpdateStatus string

const (
	PresenceUpdateStatusOnline       PresenceUpdateStatus = "online"
	PresenceUpdateStatusDoNotDisturb PresenceUpdateStatus = "dnd"
	PresenceUpdateStatusIdle         PresenceUpdateStatus = "idle"
	// Invisible and shown as offline
	PresenceUpdateStatusInvisible PresenceUpdateStatus = "invisible"
	PresenceUpdateStatusOffline   PresenceUpdateStatus = "offline"
)

//https://discord.com/developers/docs/topics/gateway#client-status-object
type GatewayPresenceClientStatus struct {
}

//https://discord.com/developers/docs/topics/gateway#activity-object-activity-structure
type GatewayActivity struct {
	// The activity's id
	Id globals.Snowflake `json:"id"`
	// The activity's name
	Name string `json:"name"`
	// Activity type
	// See See https://discord.com/developers/docs/topics/gateway#activity-object-activity-types
	Type GatewayActivityType `json:"type"`
	// Stream url, is validated when type is `1`
	Url string `json:"url"`
	// Unix timestamp of when the activity was added to the user's session
	CreatedAt uint64 `json:"created_at"`
	// Unix timestamps for start and/or end of the game
	Timestamps GatewayActivityTimestamps `json:"timestamps"`
	// The Spotify song id
	SyncId string `json:"sync_id"`
	// The platform this activity is being done on
	Platform string `json:"platform"`
	// Application id for the game
	ApplicationId globals.Snowflake `json:"application_id"`
	// What the player is currently doing
	Details string `json:"details"`
	// The user's current party status
	State string `json:"state"`
	// The emoji used for a custom status
	// See https://discord.com/developers/docs/topics/gateway#activity-object-activity-emoji
	Emoji GatewayActivityEmoji `json:"emoji"`
	SessionId string `json:"session_id"`
	// Information for the current party of the player
	// See https://discord.com/developers/docs/topics/gateway#activity-object-activity-party
	Party GatewayActivityParty `json:"party"`
	// Images for the presence and their hover texts
	// See https://discord.com/developers/docs/topics/gateway#activity-object-activity-assets
	Assets GatewayActivityAssets `json:"assets"`
	// Secrets for Rich Presence joining and spectating
	// See https://discord.com/developers/docs/topics/gateway#activity-object-activity-secrets
	Secrets GatewayActivitySecrets `json:"secrets"`
	// Whether or not the activity is an instanced game session
	Instance bool `json:"instance"`
	// Activity flags `OR`d together, describes what the payload includes
	// See https://discord.com/developers/docs/topics/gateway#activity-object-activity-flags
	// See https://en.wikipedia.org/wiki/Bit_field
	Flags ActivityFlags `json:"flags"`
	// The custom buttons shown in the Rich Presence (max 2)
	// This field is a interface{} because Discord can sometimes send []GatewayActivityButton or []string
	Buttons []interface{} `json:"buttons"`
}

// This enum is currently not documented by Discord but has known values which we will try to keep up to date.
// Values might be added or removed without a major version bump.
type ActivityPlatform uint16

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-types
type GatewayActivityType uint16

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-timestamps
type GatewayActivityTimestamps struct {
	// Unix time (in milliseconds) of when the activity started
	Start uint64 `json:"start"`
	// Unix time (in milliseconds) of when the activity ended
	End uint64 `json:"end"`
}

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-emoji
type GatewayActivityEmoji emojis.APIEmoji

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-party
type GatewayActivityParty struct {
	// The id of the party
	Id globals.Snowflake `json:"id"`
	// Used to show the party's current and maximum size
	//  [current_size: uint32, max_size: uint32]
	Size []uint32 `json:"size"`
}

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-assets
// Map keys can be: "large_image", "large_text", "small_image" and "small_text"
type GatewayActivityAssets map[string]string

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-secrets
// Map keys can be: "join", "spectate" and "match"
type GatewayActivitySecrets map[string]string

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-flags
type ActivityFlags uint32

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-flags
const (
	ActivityFlagsInstance                 ActivityFlags = 1 << 0
	ActivityFlagsJoin                     ActivityFlags = 1 << 1
	ActivityFlagsSpectate                 ActivityFlags = 1 << 2
	ActivityFlagsJoinRequest              ActivityFlags = 1 << 3
	ActivityFlagsSync                     ActivityFlags = 1 << 4
	ActivityFlagsPlay                     ActivityFlags = 1 << 5
	ActivityFlagsPartyPrivacyFriends      ActivityFlags = 1 << 6
	ActivityFlagsPartyPrivacyVoiceChannel ActivityFlags = 1 << 7
	ActivityFlagsEmbedded                 ActivityFlags = 1 << 8
)

type GatewayActivityButton struct {
	// The text shown on the button (1-32 characters)
	Label string `json:"label"`
	// The url opened when clicking the button (1-512 characters)
	Url string `json:"url"`
}

// https://discord.com/developers/docs/topics/gateway#thread-list-sync-thread-list-sync-event-fields
type GatewayThreadListSync struct {
	// ID of the guild
	GuildId globals.Snowflake `json:"guild_id"`
	// The ids of all the parent channels whose threads are being synced, otherwise the entire guild
	ChannelIds []globals.Snowflake `json:"channel_ids"`
	// Array of the synced threads
	Threads []interface{} `json:"threads"` // TODO: APIChannel
	// The member objects for the client user in each joined thread that was synced
	Members []interface{} `json:"members"` // TODO: APIThreadMember
}

// https://discord.com/developers/docs/topics/gateway#thread-members-update-thread-members-update-event-fields
type GatewayThreadMembersUpdate struct {
	// The id of the thread for which members are being synced
	Id globals.Snowflake `json:"id"`
	// The id of the guild that the thread is in
	GuildId globals.Snowflake `json:"guild_id"`
	// The approximate member count of the thread, does not count above 50 even if there are more members
	MemberCount uint8 `json:"member_count"`
	// The members that were added to the thread
	AddedMembers []interface{} `json:"added_members"` // TODO: APIThreadMember
	// The members that were removed from the thread
	RemovedMembersIds []globals.Snowflake `json:"removed_member_ids"`
}
