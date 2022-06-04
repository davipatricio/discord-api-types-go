package permissions

import "github.com/denkylabs/discord-api-types-go/globals"

// https://discord.com/developers/docs/topics/permissions#role-object

type APIRole struct {
	// Role id
	Id globals.Snowflake `json:"id"`
	// Role name
	Name string `json:"name"`
	// Integer representation of hexadecimal color code
	Color int `json:"color"`
	// If this role is pinned in the user listing
	Hoist bool `json:"hoist"`
	// The role icon hash
	Icon string `json:"icon"`
	// The role unicode emoji as a standard emoji
	UnicodeEmoji string `json:"unicode_emoji"`
	// Position of this role
	Position uint16 `json:"position"`
	// Permission bit set
	// See https://en.wikipedia.org/wiki/Bit_field
	Permissions globals.Permissions `json:"permissions"`
	// Whether this role is managed by an integration
	Managed bool `json:"managed"`
	// Whether this role is mentionable
	Mentionable bool `json:"mentionable"`
	// The tags this role has
	Tags APIRoleTags `json:"tags"`
}

// https://discord.com/developers/docs/topics/permissions#role-object-role-tags-structure
type APIRoleTags struct {
	// The id of the bot this role belongs to
	BotId globals.Snowflake `json:"bot_id"`
	// Whether this is the guild's premium subscriber role
	PremiumSubscriber interface{} `json:"premium_subscriber"`
	// The id of the integration this role belongs to
	IntegrationId globals.Snowflake `json:"integration_id"`
}
