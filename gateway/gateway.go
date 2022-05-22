package gateway

const GatewayVersion string = "10"

// https://discord.com/developers/docs/topics/opcodes-and-status-codes#gateway-gateway-opcodes
const (
	// A event was dispatched
	GatewayOpcodesDispatch = 0
	// A heartbeat was received
	GatewayOpcodesHeartbeat = 1
	// Starts a new session during the initial handshake
	GatewayOpcodesIdentify = 2
	// Update the client's presence
	GatewayOpcodesPresenceUpdate = 3
	// Used to join/leave or move between voice channels
	GatewayOpcodesVoiceStateUpdate = 4
	// Resume a previous session that was disconnected
	GatewayOpcodesResume = 6
	// You should attempt to reconnect and resume immediately
	GatewayOpcodesReconnect = 7
	// Request information about offline guild members in a large guild
	GatewayOpcodesRequestGuildMembers = 8
	// The session has been invalidated. You should reconnect and identify/resume accordingly
	GatewayOpcodesInvalidSession = 9
	// Sent immediately after connecting, contains the heartbeat_interval to use
	GatewayOpcodesHello = 10
	// 	Sent in response to receiving a heartbeat to acknowledge that it has been received
	GatewayOpcodesHeartbeatACK = 11
)

//https://discord.com/developers/docs/topics/opcodes-and-status-codes#gateway-gateway-close-event-codes
var (
	// We're not sure what went wrong. Try reconnecting?
	GatewayCloseCodesUnknownError = 4000
	// You sent an invalid Gateway opcode or an invalid payload for an opcode. Don't do that!
	// See https://discord.com/developers/docs/topics/gateway#payloads-and-opcodes
	GatewayCloseCodesUnknownOpcode = 4001
	// You sent an invalid payload to us. Don't do that!
	// See https://discord.com/developers/docs/topics/gateway#sending-payloads
	GatewayCloseCodesDecodeError = 4002
	// You sent us a payload prior to identifying
	// See https://discord.com/developers/docs/topics/gateway#identify
	GatewayCloseCodesNotAuthenticated = 4003
	// The account token sent with your identify payload is incorrect
	// See https://discord.com/developers/docs/topics/gateway#identify
	GatewayCloseCodesAuthenticationFailed = 4004
	// You sent more than one identify payload. Don't do that!
	GatewayCloseCodesAlreadyAuthenticated = 4005
	// The sequence sent when resuming the session was invalid. Reconnect and start a new session
	// See https://discord.com/developers/docs/topics/gateway#resume
	GatewayCloseCodesInvalidSeq = 4007
	// Woah nelly! You're sending payloads to us too quickly. Slow it down! You will be disconnected on receiving this
	GatewayCloseCodesRateLimited = 4008
	// Your session timed out. Reconnect and start a new one
	GatewayCloseCodesSessionTimedOut = 4009
	// You sent us an invalid shard when identifying
	// See https://discord.com/developers/docs/topics/gateway#sharding
	GatewayCloseCodesInvalidShard = 4010
	// The session would have handled too many guilds - you are required to shard your connection in order to connect
	// See https://discord.com/developers/docs/topics/gateway#sharding
	GatewayCloseCodesShardingRequired = 4011
	// You sent an invalid version for the gateway
	GatewayCloseCodesInvalidAPIVersion = 4012
	// You sent an invalid intent for a Gateway Intent. You may have incorrectly calculated the bitwise value
	GatewayCloseCodesInvalidIntents = 4013
	// You sent a disallowed intent for a Gateway Intent. You may have tried to specify an intent that you have not
	// enabled or are not whitelisted for
	GatewayCloseCodesDisallowedIntents = 4014
)

// https://discord.com/developers/docs/topics/gateway#list-of-intents
var (
	Guilds                 = 1 << 0
	GuildMembers           = 1 << 1
	GuildBans              = 1 << 2
	GuildEmojisAndStickers = 1 << 3
	GuildIntegrations      = 1 << 4
	GuildWebhooks          = 1 << 5
	GuildInvites           = 1 << 6
	GuildVoiceStates       = 1 << 7
	GuildPresences         = 1 << 8
	GuildMessages          = 1 << 9
	GuildMessageReactions  = 1 << 10
	GuildMessageTyping     = 1 << 11
	DirectMessages         = 1 << 12
	DirectMessageReactions = 1 << 13
	DirectMessageTyping    = 1 << 14
	MessageContent         = 1 << 15
	GuildScheduledEvents   = 1 << 16
)

// https://discord.com/developers/docs/topics/gateway#commands-and-events-gateway-events
var (
	GatewayDispatchEventsApplicationCommandPermissionsUpdate = "APPLICATION_COMMAND_PERMISSIONS_UPDATE"
	GatewayDispatchEventsChannelCreate                       = "CHANNEL_CREATE"
	GatewayDispatchEventsChannelDelete                       = "CHANNEL_DELETE"
	GatewayDispatchEventsChannelPinsUpdate                   = "CHANNEL_PINS_UPDATE"
	GatewayDispatchEventsChannelUpdate                       = "CHANNEL_UPDATE"
	GatewayDispatchEventsGuildBanAdd                         = "GUILD_BAN_ADD"
	GatewayDispatchEventsGuildBanRemove                      = "GUILD_BAN_REMOVE"
	GatewayDispatchEventsGuildCreate                         = "GUILD_CREATE"
	GatewayDispatchEventsGuildDelete                         = "GUILD_DELETE"
	GatewayDispatchEventsGuildEmojisUpdate                   = "GUILD_EMOJIS_UPDATE"
	GatewayDispatchEventsGuildIntegrationsUpdate             = "GUILD_INTEGRATIONS_UPDATE"
	GatewayDispatchEventsGuildMemberAdd                      = "GUILD_MEMBER_ADD"
	GatewayDispatchEventsGuildMemberRemove                   = "GUILD_MEMBER_REMOVE"
	GatewayDispatchEventsGuildMembersChunk                   = "GUILD_MEMBERS_CHUNK"
	GatewayDispatchEventsGuildMemberUpdate                   = "GUILD_MEMBER_UPDATE"
	GatewayDispatchEventsGuildRoleCreate                     = "GUILD_ROLE_CREATE"
	GatewayDispatchEventsGuildRoleDelete                     = "GUILD_ROLE_DELETE"
	GatewayDispatchEventsGuildRoleUpdate                     = "GUILD_ROLE_UPDATE"
	GatewayDispatchEventsGuildStickersUpdate                 = "GUILD_STICKERS_UPDATE"
	GatewayDispatchEventsGuildUpdate                         = "GUILD_UPDATE"
	GatewayDispatchEventsIntegrationCreate                   = "INTEGRATION_CREATE"
	GatewayDispatchEventsIntegrationDelete                   = "INTEGRATION_DELETE"
	GatewayDispatchEventsIntegrationUpdate                   = "INTEGRATION_UPDATE"
	GatewayDispatchEventsInteractionCreate                   = "INTERACTION_CREATE"
	GatewayDispatchEventsInviteCreate                        = "INVITE_CREATE"
	GatewayDispatchEventsInviteDelete                        = "INVITE_DELETE"
	GatewayDispatchEventsMessageCreate                       = "MESSAGE_CREATE"
	GatewayDispatchEventsMessageDelete                       = "MESSAGE_DELETE"
	GatewayDispatchEventsMessageDeleteBulk                   = "MESSAGE_DELETE_BULK"
	GatewayDispatchEventsMessageReactionAdd                  = "MESSAGE_REACTION_ADD"
	GatewayDispatchEventsMessageReactionRemove               = "MESSAGE_REACTION_REMOVE"
	GatewayDispatchEventsMessageReactionRemoveAll            = "MESSAGE_REACTION_REMOVE_ALL"
	GatewayDispatchEventsMessageReactionRemoveEmoji          = "MESSAGE_REACTION_REMOVE_EMOJI"
	GatewayDispatchEventsMessageUpdate                       = "MESSAGE_UPDATE"
	GatewayDispatchEventsPresenceUpdate                      = "PRESENCE_UPDATE"
	GatewayDispatchEventsStageInstanceCreate                 = "STAGE_INSTANCE_CREATE"
	GatewayDispatchEventsStageInstanceDelete                 = "STAGE_INSTANCE_DELETE"
	GatewayDispatchEventsStageInstanceUpdate                 = "STAGE_INSTANCE_UPDATE"
	GatewayDispatchEventsReady                               = "READY"
	GatewayDispatchEventsResumed                             = "RESUMED"
	GatewayDispatchEventsThreadCreate                        = "THREAD_CREATE"
	GatewayDispatchEventsThreadDelete                        = "THREAD_DELETE"
	GatewayDispatchEventsThreadListSync                      = "THREAD_LIST_SYNC"
	GatewayDispatchEventsThreadMembersUpdate                 = "THREAD_MEMBERS_UPDATE"
	GatewayDispatchEventsThreadMemberUpdate                  = "THREAD_MEMBER_UPDATE"
	GatewayDispatchEventsThreadUpdate                        = "THREAD_UPDATE"
	GatewayDispatchEventsTypingStart                         = "TYPING_START"
	GatewayDispatchEventsUserUpdate                          = "USER_UPDATE"
	GatewayDispatchEventsVoiceServerUpdate                   = "VOICE_SERVER_UPDATE"
	GatewayDispatchEventsVoiceStateUpdate                    = "VOICE_STATE_UPDATE"
	GatewayDispatchEventsWebhooksUpdate                      = "WEBHOOKS_UPDATE"
	GatewayDispatchEventsGuildScheduledEventCreate           = "GUILD_SCHEDULED_EVENT_CREATE"
	GatewayDispatchEventsGuildScheduledEventUpdate           = "GUILD_SCHEDULED_EVENT_UPDATE"
	GatewayDispatchEventsGuildScheduledEventDelete           = "GUILD_SCHEDULED_EVENT_DELETE"
	GatewayDispatchEventsGuildScheduledEventUserAdd          = "GUILD_SCHEDULED_EVENT_USER_ADD"
	GatewayDispatchEventsGuildScheduledEventUserRemove       = "GUILD_SCHEDULED_EVENT_USER_REMOVE"
)

// Gateway Dispatch Payloads

// https://discord.com/developers/docs/topics/gateway#hello
type GatewayHello struct {
	Op int `json:"op"`
	// The data of the event
	D GatewayHelloData `json:"d"`
	// Sequence number, used for resuming sessions and heartbeats
	S int `json:"s"`
}

// https://discord.com/developers/docs/topics/gateway#hello
type GatewayHelloData struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

// https://discord.com/developers/docs/topics/gateway#heartbeating
type GatewayHeartbeatRequest struct {
	Op int `json:"op"`
	// Sequence number, used for resuming sessions and heartbeats
	S int `json:"s"`
}

// https://discord.com/developers/docs/topics/gateway#heartbeating-example-gateway-heartbeat-ack
type GatewayHeartbeatAck struct {
	Op int `json:"op"`
	// Sequence number, used for resuming sessions and heartbeats
	S int `json:"s"`
}

// https://discord.com/developers/docs/topics/gateway#invalid-session
type GatewayInvalidSession struct {
	Op int `json:"op"`
	// The data of the event
	D GatewayInvalidSessionData `json:"d"`
	// Sequence number, used for resuming sessions and heartbeats
	S int `json:"s"`
}

// https://discord.com/developers/docs/topics/gateway#invalid-session
type GatewayInvalidSessionData bool

// https://discord.com/developers/docs/topics/gateway#reconnect
type GatewayReconnect struct {
	Op int `json:"op"`
}

// https://discord.com/developers/docs/topics/gateway#reconnect
type GatewayReadyDispatch struct {
	Op int `json:"op"`
	// The name of the event
	T string `json:"t"`
	// The data of the event
	D GatewayReadyDispatchData `json:"d"`
}

// Data for the READY dispatch. https://discord.com/developers/docs/topics/gateway#reconnect
type GatewayReadyDispatchData struct {
	// Gateway version.
	// See https://discord.com/developers/docs/topics/gateway#gateways-gateway-versions
	V int
	// Information about the user including email.
	// See https://discord.com/developers/docs/resources/user#user-object
	User map[string]interface{} // TODO: Implement APIUser type
	// The guilds the user is in. See https://discord.com/developers/docs/resources/guild#unavailable-guild-object
	Guilds []map[string]interface{} // TODO: Implement APIUnavailableGuild type
	// Used for resuming connections
	SessionId string
	// The shard information associated with this session, if sent when identifying. See https://discord.com/developers/docs/topics/gateway#sharding
	Shard []GatewayShard
	// Contains 'id' and 'flags' properties
	Application map[string]interface{} // TODO: Implement APIApplication type
}

// https://discord.com/developers/docs/topics/gateway#resumed
type GatewayResumedDispatch struct {
	Op int `json:"op"`
	// The name of the event
	T string `json:"t"`
}

type GatewayChannelModifyDispatch struct {
	Op int `json:"op"`
	// The name of the event
	T string `json:"t"`
	// The data of the event
	D GatewayChannelModifyDispatchData `json:"d"`
}

type GatewayChannelModifyDispatchData interface{} // TODO: Implement APIChannel type

// https://discord.com/developers/docs/topics/gateway#channel-update
type GatewayChannelCreateDispatch GatewayChannelModifyDispatch

// https://discord.com/developers/docs/topics/gateway#channel-update
type GatewayChannelCreateDispatchData GatewayChannelModifyDispatchData

// https://discord.com/developers/docs/topics/gateway#channel-update
type GatewayChannelUpdateDispatch GatewayChannelModifyDispatch

// https://discord.com/developers/docs/topics/gateway#channel-update
type GatewayChannelUpdateDispatchData GatewayChannelModifyDispatchData

// https://discord.com/developers/docs/topics/gateway#channel-update
type GatewayChannelDeleteDispatch GatewayChannelModifyDispatch

// https://discord.com/developers/docs/topics/gateway#channel-update
type GatewayChannelDeleteDispatchData GatewayChannelModifyDispatchData

// https://discord.com/developers/docs/topics/gateway#channel-pins-update
type GatewayChannelPinsUpdateDispatch struct {
	Op int `json:"op"`
	// The name of the event
	T string `json:"t"`
	// The data of the event
	D GatewayChannelPinsUpdateDispatchData `json:"d"`
	// Sequence number, used for resuming sessions and heartbeats
	S int `json:"s"`
}

// https://discord.com/developers/docs/topics/gateway#channel-pins-update
type GatewayChannelPinsUpdateDispatchData struct {
	// The id of the guild
	GuildId string `json:"guild_id"`
	// The id of the channel
	ChannelId string `json:"channel_id"`
	// The time at which the most recent pinned message was pinned
	LastPinTimestamp int `json:"last_pin_timestamp"`
}

// https://discord.com/developers/docs/topics/gateway#guild-create
// https://discord.com/developers/docs/topics/gateway#guild-update
type GatewayGuildModifyDispatch struct {
	Op int `json:"op"`
	// The name of the event
	T string `json:"t"`
	// The data of the event
	D GatewayChannelModifyDispatchData `json:"d"`
	// Sequence number, used for resuming sessions and heartbeats
	S int `json:"s"`
}

// https://discord.com/developers/docs/topics/gateway#guild-update
type GatewayGuildModifyDispatchData interface{} // TODO: Implement APIGuild type

// https://discord.com/developers/docs/topics/gateway#guild-create
type GatewayGuildCreateDispatch GatewayGuildModifyDispatch

// https://discord.com/developers/docs/topics/gateway#guild-create
// https://discord.com/developers/docs/topics/gateway#guild-create-guild-create-extra-fields
type GatewayGuildCreateDispatchData interface{} // TODO: Implement APIGuild type

// https://discord.com/developers/docs/topics/gateway#guild-update
type GatewayGuildUpdateDispatch GatewayGuildModifyDispatch

// https://discord.com/developers/docs/topics/gateway#guild-update
type GatewayGuildUpdateDispatchData interface{} // TODO: Implement APIGuild type

// https://discord.com/developers/docs/topics/gateway#guild-delete
type GatewayGuildDeleteDispatch struct {
	Op int `json:"op"`
	// The name of the event
	T string `json:"t"`
	// The data of the event
	D GatewayGuildDeleteDispatchData `json:"d"`
	// Sequence number, used for resuming sessions and heartbeats
	S int `json:"s"`
}

// https://discord.com/developers/docs/topics/gateway#guild-delete
type GatewayGuildDeleteDispatchData interface{} // TODO: Implement APIUnavailableGuild type

// https://discord.com/developers/docs/topics/gateway#guild-ban-add
// https://discord.com/developers/docs/topics/gateway#guild-ban-remove
type GatewayGuildBanModifyDispatch struct {
	Op int `json:"op"`
	// The name of the event
	T string `json:"t"`
	// The data of the event
	D GatewayGuildBanModifyDispatchData `json:"d"`
	// Sequence number, used for resuming sessions and heartbeats
	S int `json:"s"`
}

// https://discord.com/developers/docs/topics/gateway#guild-ban-add
// https://discord.com/developers/docs/topics/gateway#guild-ban-remove
type GatewayGuildBanModifyDispatchData struct {
	// ID of the guild
	GuildId string `json:"guild_id"`
	// The banned user. See https://discord.com/developers/docs/resources/user#user-object
	User interface{} `json:"user"` // TODO: Implement APIUser type
}

// https://discord.com/developers/docs/topics/gateway#guild-ban-add
type GatewayGuildBanAddDispatch GatewayGuildBanModifyDispatch

// https://discord.com/developers/docs/topics/gateway#guild-ban-add
type GatewayGuildBanAddDispatchData GatewayGuildBanModifyDispatchData

// https://discord.com/developers/docs/topics/gateway#guild-ban-remove
type GatewayGuildBanRemoveDispatch GatewayGuildBanModifyDispatchData

// https://discord.com/developers/docs/topics/gateway#guild-ban-remove
type GatewayGuildBanRemoveDispatchData GatewayGuildBanModifyDispatchData

// https://discord.com/developers/docs/topics/gateway#guild-emojis-update
type GatewayGuildEmojisUpdateDispatch struct {
	Op int `json:"op"`
	// The name of the event
	T string `json:"t"`
	// The data of the event
	D GatewayGuildEmojisUpdateDispatchData `json:"d"`
	// Sequence number, used for resuming sessions and heartbeats
	S int `json:"s"`
}

// https://discord.com/developers/docs/topics/gateway#guild-emojis-update
type GatewayGuildEmojisUpdateDispatchData struct {
	// ID of the guild
	GuildId string `json:"guild_id"`
	// Array of emojis. See https://discord.com/developers/docs/resources/emoji#emoji-object
	Emojis []interface{} `json:"emojis"` // TODO: Implement APIEmoji type
}

// https://discord.com/developers/docs/topics/gateway#guild-stickers-update
type GatewayGuildStickersUpdateDispatch struct {
	Op int `json:"op"`
	// The name of the event
	T string `json:"t"`
	// The data of the event
	D GatewayGuildStickersUpdateDispatchData `json:"d"`
	// Sequence number, used for resuming sessions and heartbeats
	S int `json:"s"`
}

// https://discord.com/developers/docs/topics/gateway#guild-stickers-update
type GatewayGuildStickersUpdateDispatchData struct {
	// ID of the guild
	GuildId string `json:"guild_id"`
	// Array of stickers. See https://discord.com/developers/docs/resources/sticker#sticker-object
	Stickers []interface{} `json:"stickers"` // TODO: Implement APISticker type
}

// https://discord.com/developers/docs/topics/gateway#guild-integrations-update
type GatewayGuildIntegrationsUpdateDispatch struct {
	Op int `json:"op"`
	// The name of the event
	T string `json:"t"`
	// The data of the event
	D GatewayGuildIntegrationsUpdateDispatchData `json:"d"`
	// Sequence number, used for resuming sessions and heartbeats
	S int `json:"s"`
}

// https://discord.com/developers/docs/topics/gateway#guild-integrations-update
type GatewayGuildIntegrationsUpdateDispatchData struct {
	/// ID of the guild whose integrations were updated
	GuildId string `json:"guild_id"`
}

// Used for Guild Sharding. See https://discord.com/developers/docs/topics/gateway#sharding
type GatewayShard struct {
	ShardId    int `json:"shard_id"`
	ShardCount int `json:"shard_count"`
}
