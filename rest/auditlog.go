package rest

import "github.com/denkylabs/discord-api-types-go/globals"

// https://discord.com/developers/docs/resources/audit-log#get-guild-audit-log
type RESTGetAPIAuditLogQuery struct {
	// Filter the log for actions made by a user
	UserId globals.Snowflake `json:"user_id"`
	// The type of audit log events
	ActionType interface{} `json:"action_type"` // TODO: AuditLogEvent
	// Filter the log before a certain entry ID
	Before globals.Snowflake `json:"before"`
	//  How many entries are returned (default 50, minimum 1, maximum 100). Defaults to 50
	Limit uint8 `json:"limit"`
}

type RESTGetAPIAuditLogResult interface{} // TODO: APIAuditLog