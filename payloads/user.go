package payloads

import (
	"github.com/denkylabs/discord-api-types-go/globals"
)

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
	Flags UserFlags `json:"flags"`
	// The type of Nitro subscription on a user's account
	// See https://discord.com/developers/docs/resources/user#user-object-premium-types
	PremiumType UserPremiumType `json:"premium_type"`
	// The public flags on a user's account
	// See https://discord.com/developers/docs/resources/user#user-object-user-flags
	PublicFlags UserFlags `json:"public_flags"`
}

// https://discord.com/developers/docs/resources/user#user-object-user-flags
type UserFlags uint64

// https://discord.com/developers/docs/resources/user#user-object-user-flags
const (
	// Discord Employee
	UserFlagsStaff UserFlags = 1 << 0
	// Partnered Server Owner
	UserFlagsPartner UserFlags = 1 << 1
	// HypeSquad Events Member
	UserFlagsHypesquad UserFlags = 1 << 2
	// Bug Hunter Level 1
	UserFlagsBugHunterLevel1 UserFlags = 1 << 3
	// House Bravery Member
	UserFlagsHypeSquadOnlineHouse1 UserFlags = 1 << 6
	// House Brilliance Member
	UserFlagsHypeSquadOnlineHouse2 UserFlags = 1 << 7
	// House Balance Member
	UserFlagsHypeSquadOnlineHouse3 UserFlags = 1 << 8
	// Early Nitro Supporter
	UserFlagsPremiumEarlySupporter UserFlags = 1 << 9
	// User is a team
	// See https://discord.com/developers/docs/topics/teams
	UserFlagsTeamPseudoUser UserFlags = 1 << 10
	// Bug Hunter Level 2
	UserFlagsBugHunterLevel2 UserFlags = 1 << 14
	// Verified Bot
	UserFlagsVerifiedBot UserFlags = 1 << 16
	// Early Verified Bot Developer
	UserFlagsVerifiedDeveloper UserFlags = 1 << 17
	// Discord Certified Moderator
	UserFlagsCertifiedModerator UserFlags = 1 << 18
	// Bot uses only HTTP interactions and is shown in the online member list
	// See https://discord.com/developers/docs/interactions/receiving-and-responding#receiving-an-interaction
	UserFlagsBotHTTPInteractions UserFlags = 1 << 19
	// User has been identified as spammer
	UserFlagsSpammer UserFlags = 1 << 20
	// User's account has been quarantined based on recent activity
	// This user flag is currently not documented by Discord but has a known value which we will try to keep up to date
	UserFlagsQuarantined UserFlags = 1 << 44
)

// https://discord.com/developers/docs/resources/user#user-object-premium-types
type UserPremiumType uint16

// https://discord.com/developers/docs/resources/user#user-object-premium-types
const (
	UserPremiumTypeNone UserPremiumType = iota
	UserPremiumTypeNitroClassic
	UserPremiumTypeNitro
)
