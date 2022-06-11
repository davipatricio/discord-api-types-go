package payloads

import "github.com/denkylabs/discord-api-types-go/globals"

// https://discord.com/developers/docs/resources/audit-log#audit-log-object-audit-log-structure
type APIAuditLog struct {
	// List of application commands found in the audit log
	// See https://discord.com/developers/docs/interactions/application-commands#application-command-object
	ApplicationCommands []interface{} `json:"application_commands"` // TODO: ApplicationCommand
	// Webhooks found in the audit log
	// See https://discord.com/developers/docs/resources/webhook#webhook-object
	Webhooks []APIWebhook `json:"webhooks"`
	// Users found in the audit log
	// See https://discord.com/developers/docs/resources/user#user-object
	Users []APIUser `json:"users"`
	// Audit log entries
	// See https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object
	AuditLogEntries []APIAuditLogEntry `json:"audit_log_entries"`
	// Partial integration objects
	// See https://discord.com/developers/docs/resources/guild#integration-object
	Integrations []APIGuildIntegration `json:"integrations"`
	// Threads found in the audit log
	// Threads referenced in THREAD_CREATE and THREAD_UPDATE events are included in the threads map, since archived threads might not be kept in memory by clients.
	// See https://discord.com/developers/docs/resources/channel#channel-object
	Threads []interface{} `json:"threads"` // TODO: APIChannel
	// The guild scheduled events in the audit log
	// See https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object
	GuildScheduledEvents []interface{} `json:"guild_scheduled_events"` // TODO: APIGuildScheduledEvent
}

// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-audit-log-entry-structure
type APIAuditLogEntry struct {
	// ID of the affected entity (webhook, user, role, etc.)
	TargetId globals.Snowflake `json:"target_id"`
	// Changes made to the `target_id`
	// See https://discord.com/developers/docs/resources/audit-log#audit-log-change-object
	Changes []APIAuditLogChange `json:"changes"`
	// The user who made the changes
	// This can be `null` in some cases (webhooks deleting themselves by using their own token, for example)
	UserId globals.Snowflake `json:"user_id"`
	// ID of the entry
	Id globals.Snowflake `json:"id"`
	// Type of action that occurred
	// See https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-audit-log-events
	ActionType AuditLogEvent `json:"action_type"`
	// Additional info for certain action types
	// See https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-optional-audit-entry-info
	Options APIAuditLogOptions `json:"options"`
	// The reason for the change (0-512 characters)
	Reason string `json:"reason"`
}

// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-audit-log-events
type AuditLogEvent uint16

// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-audit-log-events
const (
	AuditLogEventGuildUpdate AuditLogEvent = 1

	AuditLogEventChannelCreate          AuditLogEvent = 10
	AuditLogEventChannelUpdate          AuditLogEvent = 11
	AuditLogEventChannelDelete          AuditLogEvent = 12
	AuditLogEventChannelOverwriteCreate AuditLogEvent = 13
	AuditLogEventChannelOverwriteUpdate AuditLogEvent = 14
	AuditLogEventChannelOverwriteDelete AuditLogEvent = 15

	AuditLogEventMemberKick       AuditLogEvent = 20
	AuditLogEventMemberPrune      AuditLogEvent = 21
	AuditLogEventMemberBanAdd     AuditLogEvent = 22
	AuditLogEventMemberBanRemove  AuditLogEvent = 23
	AuditLogEventMemberUpdate     AuditLogEvent = 24
	AuditLogEventMemberRoleUpdate AuditLogEvent = 25
	AuditLogEventMemberMove       AuditLogEvent = 26
	AuditLogEventMemberDisconnect AuditLogEvent = 27
	AuditLogEventBotAdd           AuditLogEvent = 28

	AuditLogEventRoleCreate AuditLogEvent = 30
	AuditLogEventRoleUpdate AuditLogEvent = 31
	AuditLogEventRoleDelete AuditLogEvent = 32

	AuditLogEventInviteCreate AuditLogEvent = 40
	AuditLogEventInviteUpdate AuditLogEvent = 41
	AuditLogEventInviteDelete AuditLogEvent = 42

	AuditLogEventWebhookCreate AuditLogEvent = 50
	AuditLogEventWebhookUpdate AuditLogEvent = 51
	AuditLogEventWebhookDelete AuditLogEvent = 52

	AuditLogEventEmojiCreate AuditLogEvent = 60
	AuditLogEventEmojiUpdate AuditLogEvent = 61
	AuditLogEventEmojiDelete AuditLogEvent = 62

	AuditLogEventMessageDelete     AuditLogEvent = 72
	AuditLogEventMessageBulkDelete AuditLogEvent = 73
	AuditLogEventMessagePin        AuditLogEvent = 74
	AuditLogEventMessageUnpin      AuditLogEvent = 75

	AuditLogEventIntegrationCreate   AuditLogEvent = 80
	AuditLogEventIntegrationUpdate   AuditLogEvent = 81
	AuditLogEventIntegrationDelete   AuditLogEvent = 82
	AuditLogEventStageInstanceCreate AuditLogEvent = 83
	AuditLogEventStageInstanceUpdate AuditLogEvent = 84
	AuditLogEventStageInstanceDelete AuditLogEvent = 85

	AuditLogEventStickerCreate AuditLogEvent = 90
	AuditLogEventStickerUpdate AuditLogEvent = 91
	AuditLogEventStickerDelete AuditLogEvent = 92

	AuditLogEventGuildScheduledEventCreate AuditLogEvent = 100
	AuditLogEventGuildScheduledEventUpdate AuditLogEvent = 101
	AuditLogEventGuildScheduledEventDelete AuditLogEvent = 102

	AuditLogEventThreadCreate AuditLogEvent = 110
	AuditLogEventThreadUpdate AuditLogEvent = 111
	AuditLogEventThreadDelete AuditLogEvent = 112

	AuditLogEventApplicationCommandPermissionUpdate AuditLogEvent = 121
)

// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-optional-audit-entry-info
type APIAuditLogOptions struct {
	// Number of days after which inactive members were kicked
	//
	// Present from:
	// - MEMBER_PRUNE
	DeleteMemberDays string `json:"delete_member_days"`
	// Number of members removed by the prune
	//
	// Present from:
	// - MEMBER_PRUNE
	MembersRemoved string `json:"members_removed"`
	// Channel in which the entities were targeted
	//
	// Present from:
	// - MEMBER_MOVE
	// - MESSAGE_PIN
	// - MESSAGE_UNPIN
	// - MESSAGE_DELETE
	// - STAGE_INSTANCE_CREATE
	// - STAGE_INSTANCE_UPDATE
	// - STAGE_INSTANCE_DELETE
	ChannelId globals.Snowflake `json:"channel_id"`
	// ID of the message that was targeted
	//
	// Present from:
	// - MESSAGE_PIN
	// - MESSAGE_UNPIN
	MessageId globals.Snowflake `json:"message_id"`
	// Number of entities that were targeted
	//
	// Present from:
	// - MESSAGE_DELETE
	// - MESSAGE_BULK_DELETE
	// - MEMBER_DISCONNECT
	// - MEMBER_MOVE
	Count string `json:"count"`
	// ID of the overwritten entity
	//
	// Present from:
	// - CHANNEL_OVERWRITE_CREATE
	// - CHANNEL_OVERWRITE_UPDATE
	// - CHANNEL_OVERWRITE_DELETE
	Id string `json:"id"`
	// Type of overwritten entity - "0" for "role" or "1" for "member"
	//
	// Present from:
	// - CHANNEL_OVERWRITE_CREATE
	// - CHANNEL_OVERWRITE_UPDATE
	// - CHANNEL_OVERWRITE_DELETE
	Type AuditLogOptionsType `json:"type"`
	// Name of the role
	//
	// Present from:
	// - CHANNEL_OVERWRITE_CREATE
	// - CHANNEL_OVERWRITE_UPDATE
	// - CHANNEL_OVERWRITE_DELETE
	//
	// Present only if the APIAuditLogOptions.type entry type is "0"
	RoleName string `json:"role_name"`
}

type AuditLogOptionsType string

const (
	AuditLogOptionsTypeRole   AuditLogOptionsType = "0"
	AuditLogOptionsTypeMember AuditLogOptionsType = "1"
)

// https://discord.com/developers/docs/resources/audit-log#audit-log-change-object-audit-log-change-structure
type APIAuditLogChange AuditLogChangeData[string, any]

type AuditLogChangeData[K string, D any] struct {
	Key K `json:"key"`
	/**
	 * The new value
	 *
	 * If `new_value` is not present in the change object, while `old_value` is,
	 * that means the property that was changed has been reset, or set to `null`
	 */
	NewValue D `json:"new_value"`
	OldValue D `json:"old_value"`
}
