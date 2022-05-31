package user

import "github.com/denkylabs/discord-api-types-go/globals"

// https://discord.com/developers/docs/resources/user#user-object
type APIUser struct {
	// The user's id
	Id globals.Snowflake `json:"id"`
	// The user's username, not unique across the platform
	Username string `json:"username"`
	// The user's 4-digit discord-tag
	Discriminator string `json:"discriminator"`
	// The user's avatar hash
	// See https://discord.com/developers/docs/reference#image-formatting
	Avatar string `json:"avatar"`
	// Whether the user belongs to an OAuth2 application
	Bot bool `json:"bot"`
	// Whether the user is an Official Discord System user (part of the urgent message system)
	System bool `json:"system"`
	// Whether the user has two factor enabled on their account
	MfaEnabled bool `json:"mfa_enabled"`
	// The user's banner hash
	// See https://discord.com/developers/docs/reference#image-formatting
	Banner string `json:"banner"`
	// The user's banner color encoded as an integer representation of hexadecimal color code
	AccentColor int `json:"accent_color"`
	// Whether the email on this account has been verified
	Verified bool `json:"verified"`
	// The user's email
	Email string `json:"email"`
	// The flags on a user's account
	// See https://discord.com/developers/docs/resources/user#user-object-user-flags
	Flags int `json:"flags"`
	// The type of Nitro subscription on a user's account
	// See https://discord.com/developers/docs/resources/user#user-object-premium-types
	PremiumType int `json:"premium_type"`
	// The public flags on a user's account
	// See https://discord.com/developers/docs/resources/user#user-object-user-flags
	PublicFlags int `json:"public_flags"`
}

// https://discord.com/developers/docs/resources/user#user-object-user-flags
const (
	// Discord Employee
	UserFlagsStaff int = 1 << 0
	// Partnered Server Owner
	UserFlagsPartner int = 1 << 1
	// HypeSquad Events Member
	UserFlagsHypesquad int = 1 << 2
	// Bug Hunter Level 1
	UserFlagsBugHunterLevel1 int = 1 << 3
	// House Bravery Member
	UserFlagsHypeSquadOnlineHouse1 int = 1 << 6
	// House Brilliance Member
	UserFlagsHypeSquadOnlineHouse2 int = 1 << 7
	// House Balance Member
	UserFlagsHypeSquadOnlineHouse3 int = 1 << 8
	// Early Nitro Supporter
	UserFlagsPremiumEarlySupporter int = 1 << 9
	// User is a [team](https://discord.com/developers/docs/topics/teams)
	UserFlagsTeamPseudoUser int = 1 << 10
	// Bug Hunter Level 2
	BugHunterLevel2 int = 1 << 14
	// Verified Bot
	VerifiedBot int = 1 << 16
	// Early Verified Bot Developer
	VerifiedDeveloper int = 1 << 17
	// Discord Certified Moderator
	CertifiedModerator int = 1 << 18
	// Bot uses only [HTTP interactions](https://discord.com/developers/docs/interactions/receiving-and-responding#receiving-an-interaction) and is shown in the online member list
	BotHTTPInteractions int = 1 << 19
	// User has been identified as spammer
	Spammer int = 1 << 20
)

// https://discord.com/developers/docs/resources/user#user-object-premium-types
const (
	UserPremiumTypeNone int = iota
	UserPremiumTypeNitroClassic
	UserPremiumTypeNitro
)

// https://discord.com/developers/docs/resources/user#connection-object
type APIConnetion struct {
	// ID of the connection account
	Id string `json:"id"`
	// The username of the connection account
	Name string `json:"name"`
	// The service of the connection
	Type string `json:"type"`
	// Whether the connection is revoked
	Revoked bool `json:"revoked"`
	// An array of partial server integrations
	// See https://discord.com/developers/docs/resources/guild#integration-object
	Integrations []interface{} `json:"integrations"` // TODO: APIGuildIntegration
	// Whether the connection is verified
	Verified bool `json:"verified"`
	// Whether friend sync is enabled for this connection
	FriendSync bool `json:"friend_sync"`
	// Whether activities related to this connection will be shown in presence updates
	ShowActivity bool `json:"show_activity"`
	// Visibility of this connection
	// See https://discord.com/developers/docs/resources/user#connection-object-visibility
	Visibility int `json:"visibility"`
}

const (
	// Invisible to everyone except the user themselves
	ConnectionVisibilityNone int = iota
	// Visible to everyone
	ConnectionVisibilityEveryone
)
