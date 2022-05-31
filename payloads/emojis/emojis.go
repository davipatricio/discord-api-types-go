package emojis

import "github.com/denkylabs/discord-api-types-go/globals"

// Not documented but mentioned
type APIPartialEmoji struct {
	// Emoji id
	Id globals.Snowflake `json:"id"`
	// Emoji name (can be null only in reaction emoji objects)
	Name string `json:"name"`
	// Whether this emoji is animated
	Animated bool `json:"animated"`
}

// https://discord.com/developers/docs/resources/emoji#emoji-object-emoji-structure
type APIEmoji struct {
	// Roles this emoji is whitelisted to
	Roles []globals.Snowflake `json:"roles"`
	// User that created this emoji
	User interface{} `json:"user"` // TODO: APIUser
	// Whether this emoji must be wrapped in colons
	RequiresColons bool `json:"requires_colons"`
	// Whether this emoji is managed
	Managed bool `json:"managed"`
	// Whether this emoji can be used, may be false due to loss of Server Boosts
	Available bool `json:"available"`
}
