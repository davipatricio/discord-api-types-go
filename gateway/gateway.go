package gateway

import (
	"github.com/denkylabs/discord-api-types-go/payloads/emojis"
	"github.com/denkylabs/discord-api-types-go/payloads/guilds"
	"github.com/denkylabs/discord-api-types-go/payloads/user"
)

const GatewayVersion string = "10"

// https://discord.com/developers/docs/topics/opcodes-and-status-codes#gateway-gateway-opcodes
const (
	// A event was dispatched
	GatewayOpcodesDispatch uint16 = iota
	// A heartbeat was received
	GatewayOpcodesHeartbeat
	// Starts a new session during the initial handshake
	GatewayOpcodesIdentify
	// Update the client's presence
	GatewayOpcodesPresenceUpdate
	// Used to join/leave or move between voice channels
	GatewayOpcodesVoiceStateUpdate
	// Resume a previous session that was disconnected
	GatewayOpcodesResume uint16 = 6
	// You should attempt to reconnect and resume immediately
	GatewayOpcodesReconnect uint16 = 7
	// Request information about offline guild members in a large guild
	GatewayOpcodesRequestGuildMembers uint16 = 8
	// The session has been invalidated. You should reconnect and identify/resume accordingly
	GatewayOpcodesInvalidSession uint16 = 9
	// Sent immediately after connecting, contains the heartbeat_interval to use
	GatewayOpcodesHello uint16 = 10
	// 	Sent in response to receiving a heartbeat to acknowledge that it has been received
	GatewayOpcodesHeartbeatACK uint16 = 11
)

//https://discord.com/developers/docs/topics/opcodes-and-status-codes#gateway-gateway-close-event-codes
const (
	// We're not sure what went wrong. Try reconnecting?
	GatewayCloseCodesUnknownError int = 4000
	// You sent an invalid Gateway opcode or an invalid payload for an opcode. Don't do that!
	// See https://discord.com/developers/docs/topics/gateway#payloads-and-opcodes
	GatewayCloseCodesUnknownOpcode int = 4001
	// You sent an invalid payload to us. Don't do that!
	// See https://discord.com/developers/docs/topics/gateway#sending-payloads
	GatewayCloseCodesDecodeError int = 4002
	// You sent us a payload prior to identifying
	// See https://discord.com/developers/docs/topics/gateway#identify
	GatewayCloseCodesNotAuthenticated int = 4003
	// The account token sent with your identify payload is incorrect
	// See https://discord.com/developers/docs/topics/gateway#identify
	GatewayCloseCodesAuthenticationFailed int = 4004
	// You sent more than one identify payload. Don't do that!
	GatewayCloseCodesAlreadyAuthenticated int = 4005
	// The sequence sent when resuming the session was invalid. Reconnect and start a new session
	// See https://discord.com/developers/docs/topics/gateway#resume
	GatewayCloseCodesInvalidSeq int = 4007
	// Woah nelly! You're sending payloads to us too quickly. Slow it down! You will be disconnected on receiving this
	GatewayCloseCodesRateLimited int = 4008
	// Your session timed out. Reconnect and start a new one
	GatewayCloseCodesSessionTimedOut int = 4009
	// You sent us an invalid shard when identifying
	// See https://discord.com/developers/docs/topics/gateway#sharding
	GatewayCloseCodesInvalidShard int = 4010
	// The session would have handled too many guilds - you are required to shard your connection in order to connect
	// See https://discord.com/developers/docs/topics/gateway#sharding
	GatewayCloseCodesShardingRequired int = 4011
	// You sent an invalid version for the gateway
	GatewayCloseCodesInvalidAPIVersion int = 4012
	// You sent an invalid intent for a Gateway Intent. You may have incorrectly calculated the bitwise value
	GatewayCloseCodesInvalidIntents int = 4013
	// You sent a disallowed intent for a Gateway Intent. You may have tried to specify an intent that you have not
	// enabled or are not whitelisted for
	GatewayCloseCodesDisallowedIntents int = 4014
)

// https://discord.com/developers/docs/topics/gateway#list-of-intents
const (
	Guilds                 int = 1 << 0
	GuildMembers           int = 1 << 1
	GuildBans              int = 1 << 2
	GuildEmojisAndStickers int = 1 << 3
	GuildIntegrations      int = 1 << 4
	GuildWebhooks          int = 1 << 5
	GuildInvites           int = 1 << 6
	GuildVoiceStates       int = 1 << 7
	GuildPresences         int = 1 << 8
	GuildMessages          int = 1 << 9
	GuildMessageReactions  int = 1 << 10
	GuildMessageTyping     int = 1 << 11
	DirectMessages         int = 1 << 12
	DirectMessageReactions int = 1 << 13
	DirectMessageTyping    int = 1 << 14
	MessageContent         int = 1 << 15
	GuildScheduledEvents   int = 1 << 16
)

// https://discord.com/developers/docs/topics/gateway#commands-and-events-gateway-events
const (
	GatewayDispatchEventsApplicationCommandPermissionsUpdate string = "APPLICATION_COMMAND_PERMISSIONS_UPDATE"
	GatewayDispatchEventsChannelCreate                       string = "CHANNEL_CREATE"
	GatewayDispatchEventsChannelDelete                       string = "CHANNEL_DELETE"
	GatewayDispatchEventsChannelPinsUpdate                   string = "CHANNEL_PINS_UPDATE"
	GatewayDispatchEventsChannelUpdate                       string = "CHANNEL_UPDATE"
	GatewayDispatchEventsGuildBanAdd                         string = "GUILD_BAN_ADD"
	GatewayDispatchEventsGuildBanRemove                      string = "GUILD_BAN_REMOVE"
	GatewayDispatchEventsGuildCreate                         string = "GUILD_CREATE"
	GatewayDispatchEventsGuildDelete                         string = "GUILD_DELETE"
	GatewayDispatchEventsGuildEmojisUpdate                   string = "GUILD_EMOJIS_UPDATE"
	GatewayDispatchEventsGuildIntegrationsUpdate             string = "GUILD_INTEGRATIONS_UPDATE"
	GatewayDispatchEventsGuildMemberAdd                      string = "GUILD_MEMBER_ADD"
	GatewayDispatchEventsGuildMemberRemove                   string = "GUILD_MEMBER_REMOVE"
	GatewayDispatchEventsGuildMembersChunk                   string = "GUILD_MEMBERS_CHUNK"
	GatewayDispatchEventsGuildMemberUpdate                   string = "GUILD_MEMBER_UPDATE"
	GatewayDispatchEventsGuildRoleCreate                     string = "GUILD_ROLE_CREATE"
	GatewayDispatchEventsGuildRoleDelete                     string = "GUILD_ROLE_DELETE"
	GatewayDispatchEventsGuildRoleUpdate                     string = "GUILD_ROLE_UPDATE"
	GatewayDispatchEventsGuildStickersUpdate                 string = "GUILD_STICKERS_UPDATE"
	GatewayDispatchEventsGuildUpdate                         string = "GUILD_UPDATE"
	GatewayDispatchEventsIntegrationCreate                   string = "INTEGRATION_CREATE"
	GatewayDispatchEventsIntegrationDelete                   string = "INTEGRATION_DELETE"
	GatewayDispatchEventsIntegrationUpdate                   string = "INTEGRATION_UPDATE"
	GatewayDispatchEventsInteractionCreate                   string = "INTERACTION_CREATE"
	GatewayDispatchEventsInviteCreate                        string = "INVITE_CREATE"
	GatewayDispatchEventsInviteDelete                        string = "INVITE_DELETE"
	GatewayDispatchEventsMessageCreate                       string = "MESSAGE_CREATE"
	GatewayDispatchEventsMessageDelete                       string = "MESSAGE_DELETE"
	GatewayDispatchEventsMessageDeleteBulk                   string = "MESSAGE_DELETE_BULK"
	GatewayDispatchEventsMessageReactionAdd                  string = "MESSAGE_REACTION_ADD"
	GatewayDispatchEventsMessageReactionRemove               string = "MESSAGE_REACTION_REMOVE"
	GatewayDispatchEventsMessageReactionRemoveAll            string = "MESSAGE_REACTION_REMOVE_ALL"
	GatewayDispatchEventsMessageReactionRemoveEmoji          string = "MESSAGE_REACTION_REMOVE_EMOJI"
	GatewayDispatchEventsMessageUpdate                       string = "MESSAGE_UPDATE"
	GatewayDispatchEventsPresenceUpdate                      string = "PRESENCE_UPDATE"
	GatewayDispatchEventsStageInstanceCreate                 string = "STAGE_INSTANCE_CREATE"
	GatewayDispatchEventsStageInstanceDelete                 string = "STAGE_INSTANCE_DELETE"
	GatewayDispatchEventsStageInstanceUpdate                 string = "STAGE_INSTANCE_UPDATE"
	GatewayDispatchEventsReady                               string = "READY"
	GatewayDispatchEventsResumed                             string = "RESUMED"
	GatewayDispatchEventsThreadCreate                        string = "THREAD_CREATE"
	GatewayDispatchEventsThreadDelete                        string = "THREAD_DELETE"
	GatewayDispatchEventsThreadListSync                      string = "THREAD_LIST_SYNC"
	GatewayDispatchEventsThreadMembersUpdate                 string = "THREAD_MEMBERS_UPDATE"
	GatewayDispatchEventsThreadMemberUpdate                  string = "THREAD_MEMBER_UPDATE"
	GatewayDispatchEventsThreadUpdate                        string = "THREAD_UPDATE"
	GatewayDispatchEventsTypingStart                         string = "TYPING_START"
	GatewayDispatchEventsUserUpdate                          string = "USER_UPDATE"
	GatewayDispatchEventsVoiceServerUpdate                   string = "VOICE_SERVER_UPDATE"
	GatewayDispatchEventsVoiceStateUpdate                    string = "VOICE_STATE_UPDATE"
	GatewayDispatchEventsWebhooksUpdate                      string = "WEBHOOKS_UPDATE"
	GatewayDispatchEventsGuildScheduledEventCreate           string = "GUILD_SCHEDULED_EVENT_CREATE"
	GatewayDispatchEventsGuildScheduledEventUpdate           string = "GUILD_SCHEDULED_EVENT_UPDATE"
	GatewayDispatchEventsGuildScheduledEventDelete           string = "GUILD_SCHEDULED_EVENT_DELETE"
	GatewayDispatchEventsGuildScheduledEventUserAdd          string = "GUILD_SCHEDULED_EVENT_USER_ADD"
	GatewayDispatchEventsGuildScheduledEventUserRemove       string = "GUILD_SCHEDULED_EVENT_USER_REMOVE"
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
	User user.APIUser // TODO: Implement APIUser type
	// The guilds the user is in. See https://discord.com/developers/docs/resources/guild#unavailable-guild-object
	Guilds guilds.APIUnavailableGuild
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
type GatewayGuildModifyDispatchData guilds.APIGuild

// https://discord.com/developers/docs/topics/gateway#guild-create
type GatewayGuildCreateDispatch GatewayGuildModifyDispatch

// https://discord.com/developers/docs/topics/gateway#guild-create
// https://discord.com/developers/docs/topics/gateway#guild-create-guild-create-extra-fields
type GatewayGuildCreateDispatchData guilds.APIGuild

// https://discord.com/developers/docs/topics/gateway#guild-update
type GatewayGuildUpdateDispatch GatewayGuildModifyDispatch

// https://discord.com/developers/docs/topics/gateway#guild-update
type GatewayGuildUpdateDispatchData guilds.APIGuild

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
type GatewayGuildDeleteDispatchData guilds.APIUnavailableGuild

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
	User user.APIUser `json:"user"`
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
	Emojis []emojis.APIEmoji `json:"emojis"`
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

// https://discord.com/developers/docs/topics/gateway#guild-member-add
type GatewayGuildMemberAddDispatch struct {
	Op int `json:"op"`
	// The name of the event
	T string `json:"t"`
	// The data of the event
	D GatewayGuildMemberAddDispatchData `json:"d"`
	// Sequence number, used for resuming sessions and heartbeats
	S int `json:"s"`
}

// https://discord.com/developers/docs/topics/gateway#guild-member-add
type GatewayGuildMemberAddDispatchData struct { // TODO: Implement APIGuildMember type
	// The id of the guild
	GuildId string `json:"guild_id"`
}

// https://discord.com/developers/docs/topics/gateway#guild-member-remove
type GatewayGuildMemberRemoveDispatch struct {
	Op int `json:"op"`
	// The name of the event
	T string `json:"t"`
	// The data of the event
	D GatewayGuildMemberRemoveDispatchData `json:"d"`
	// Sequence number, used for resuming sessions and heartbeats
	S int `json:"s"`
}

// https://discord.com/developers/docs/topics/gateway#guild-member-remove
type GatewayGuildMemberRemoveDispatchData struct { // TODO: Implement APIGuildMember type
	// The id of the guild
	GuildId string `json:"guild_id"`
	// The user who was removed
	User user.APIUser `json:"user"`
}

// Used for Guild Sharding. See https://discord.com/developers/docs/topics/gateway#sharding
type GatewayShard struct {
	ShardId    int `json:"shard_id"`
	ShardCount int `json:"shard_count"`
}
