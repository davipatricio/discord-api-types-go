package payloads

import "github.com/denkylabs/discord-api-types-go/globals"

// https://discord.com/developers/docs/resources/template#template-object
type APITemplate struct {
	// The template code (unique ID)
	Code string `json:"code"`
	// Template name
	Name string `json:"name"`
	// The description for the template
	Description string `json:"description"`
	// Number of times this template has been used
	UsageCount uint32 `json:"usage_count"`
	// The ID of the user who created the template
	CreatorId globals.Snowflake `json:"creator_id"`
	// The user who created the template
	// See https://discord.com/developers/docs/resources/user#user-object
	Creator APIUser `json:"creator"`
	// When this template was created
	CreatedAt string `json:"created_at"`
	// When this template was last synced to the source guild
	UpdatedAt string `json:"updated_at"`
	// The ID of the guild this template is based on
	SourceGuildId globals.Snowflake `json:"source_guild_id"`
	// The guild snapshot this template contains
	SerializedSourceGuild APIGuild `json:"serialized_source_guild"`
	// Whether the template has unsynced changes
	IsDirty bool `json:"is_dirty"`
}
