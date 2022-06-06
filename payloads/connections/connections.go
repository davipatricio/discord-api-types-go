package connections

import (
	"github.com/denkylabs/discord-api-types-go/globals"
	"github.com/denkylabs/discord-api-types-go/payloads/guilds"
)

// https://discord.com/developers/docs/resources/user#connection-object
type APIConnection struct {
	// ID of the connection account
	Id globals.Snowflake `json:"id"`
	// The username of the connection account
	Name string `json:"name"`
	// The service of the connection
	Type string `json:"type"`
	// Whether the connection is revoked
	Revoked bool `json:"revoked"`
	// An array of partial server integrations
	// See https://discord.com/developers/docs/resources/guild#integration-object
	Integrations guilds.APIGuildIntegration `json:"integrations"` // TODO: APIGuildIntegration
	// Whether the connection is verified
	Verified bool `json:"verified"`
	// Whether friend sync is enabled for this connection
	FriendSync bool `json:"friend_sync"`
	// Whether activities related to this connection will be shown in presence updates
	ShowActivity bool `json:"show_activity"`
	// Visibility of this connection
	// See https://discord.com/developers/docs/resources/user#connection-object-visibility
	Visibility ConnectionVisibility `json:"visibility"`
}

// Visibility of this connection
// See https://discord.com/developers/docs/resources/user#connection-object-visibility
type ConnectionVisibility uint16

// Visibility of this connection
// See https://discord.com/developers/docs/resources/user#connection-object-visibility
const (
	// Invisible to everyone except the user themselves
	ConnectionVisibilityNone ConnectionVisibility = iota
	// Visible to everyone
	ConnectionVisibilityEveryone
)
