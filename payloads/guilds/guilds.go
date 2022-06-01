package guilds

import (
	"github.com/denkylabs/discord-api-types-go/globals"
	"github.com/denkylabs/discord-api-types-go/payloads/permissions"
)

// https://discord.com/developers/docs/resources/guild#unavailable-guild-object
type APIUnavailableGuild struct {
	// Guild id
	ID string `json:"id"`
	// `true` if this guild is unavailable due to an outage
	Unavailable bool `json:"unavailable"`
}

// https://discord.com/developers/docs/resources/guild#guild-object-guild-structure
type APIPartialGuild struct {
	// Guild name (2-100 characters, excluding trailing and leading whitespace)
	Name string `json:"name"`
	// Icon hash
	// See https://discord.com/developers/docs/reference#image-formatting
	Icon string `json:"icon"`
	// Splash hash
	// See https://discord.com/developers/docs/reference#image-formatting
	Splash string `json:"splash"`
	// Banner hash
	// See https://discord.com/developers/docs/reference#image-formatting
	Banner string `json:"banner,omitempty"`
	// The description for the guild
	Description string `json:"description,omitempty"`
	// Enabled guild features
	// See https://discord.com/developers/docs/resources/guild#guild-object-guild-features
	Features []interface{} `json:"features,omitempty"` // TODO: Guild features
	// Verification level required for the guild
	// See https://discord.com/developers/docs/resources/guild#guild-object-verification-level
	VerificationLevel int `json:"verification_level,omitempty"`
	// The vanity url code for the guild
	VanityURLCode string `json:"vanity_url_code,omitempty"`
	// `true` if this guild is unavailable due to an outage
	Unavailable bool `json:"unavailable,omitempty"`
	// HACK: This should be the same as APIGuild
	WelcomeScreen APIGuildWelcomeScreen `json:"welcome_screen,omitempty"`
}

// https://discord.com/developers/docs/resources/guild#guild-object-guild-structure
type APIGuild struct {
	// Icon hash, returned when in the template object
	// See https://discord.com/developers/docs/reference#image-formatting
	IconHash string `json:"icon_hash,omitempty"`
	// Discovery splash hash; only present for guilds with the "DISCOVERABLE" feature
	// See https://discord.com/developers/docs/reference#image-formatting
	DiscoverySplash string `json:"discovery_splash,omitempty"`
	// `true` if the user is the owner of the guild
	// This field is only received from https://discord.com/developers/docs/resources/user#get-current-user-guilds
	Owner bool `json:"owner,omitempty"`
	// ID of owner
	OwnerID globals.Snowflake `json:"owner_id"`
	// Total permissions for the user in the guild (excludes overrides)
	// This field is only received from https://discord.com/developers/docs/resources/user#get-current-user-guilds
	// See https://en.wikipedia.org/wiki/Bit_field
	Permissions globals.Permissions `json:"permissions,omitempty"`
	// Voice region id for the guild
	// See https://discord.com/developers/docs/resources/voice#voice-region-object
	// @deprecated This field has been deprecated in favor of `rtc_region` on the channel.
	Region string `json:"region"`
	// ID of afk channel
	AfkChannelId globals.Snowflake `json:"afk_channel_id,omitempty"`
	// AFK timeout in seconds
	AfkTimeout int `json:"afk_timeout"`
	// `true` if the guild widget is enabled
	WidgetEnabled bool `json:"widget_enabled,omitempty"`
	// The channel id that the widget will generate an invite to, or `null` if set to no invite
	WidgetChannelId globals.Snowflake `json:"widget_channel_id,omitempty"`
	// Verification level required for the guild
	// See https://discord.com/developers/docs/resources/guild#guild-object-verification-level
	VerificationLevel int `json:"verification_level"`
	// Default message notifications level
	// See https://discord.com/developers/docs/resources/guild#guild-object-default-message-notification-level
	DefaultMessageNotifications int `json:"default_message_notifications"`
	// Explicit content filter level
	// See https://discord.com/developers/docs/resources/guild#guild-object-explicit-content-filter-level
	ExplicitContentFilter int `json:"explicit_content_filter"`
	// Roles in the guild
	// See https://discord.com/developers/docs/topics/permissions#role-object
	Roles permissions.APIRole `json:"roles"`
	// The welcome screen of a Community guild, shown to new members
	// Returned in the invite object
	WelcomeScreen APIGuildWelcomeScreen `json:"welcome_screen,omitempty"`
}

// https://discord.com/developers/docs/resources/guild#guild-object-default-message-notification-level
const (
	GuildDefaultMessageNotificationsAllMessages int = iota
	GuildDefaultMessageNotificationsOnlyMentions
)

// https://discord.com/developers/docs/resources/guild#guild-object-explicit-content-filter-level
const (
	GuildExplicitContentFilterDisabled int = iota
	GuildExplicitContentFilterMembersWithoutRoles
	GuildExplicitContentFilterAllMembers
)

// https://discord.com/developers/docs/resources/guild#guild-object-guild-nsfw-level
const (
	GuildNSFWLevelDefault int = iota
	GuildNSFWLevelExplicit
	GuildNSFWLevelSafe
	GuildNSFWLevelAgeRestricted
)

// https://discord.com/developers/docs/resources/guild#guild-object-verification-level
const (
	// Unrestricted
	GuildVerificationLevelNone uint16 = iota
	// Must have verified email on account
	GuildVerificationLevelLow
	// Must be registered on Discord for longer than 5 minutes
	GuildVerificationLevelMedium
	// Must be a member of the guild for longer than 10 minutes
	GuildVerificationLevelHigh
	// Must have a verified phone number
	GuildVerificationLevelVeryHigh
)

// https://discord.com/developers/docs/resources/guild#guild-object-mfa-level
const (
	GuildMFALevelNone int = iota
	GuildMFALevelElevated
)

// https://discord.com/developers/docs/resources/guild#guild-object-premium-tier
const (
	GuildPremiumTierNone int = iota
	GuildPremiumTierTier1
	GuildPremiumTierTier2
	GuildPremiumTierTier3
)

const (
	GuildHubTypeDefault int = iota
	GuildHubTypeHighSchool
	GuildHubTypeCollege
)

// https://discord.com/developers/docs/resources/guild#guild-object-system-channel-flags
const (
	// Suppress member join notifications
	GuildSystemChannelFlagsSuppressJoinNotifications int = 1 << 0
	// Suppress server boost notifications
	GuildSystemChannelFlagsSuppressPremiumSubscriptions int = 1 << 1
	// Suppress server setup tips
	GuildSystemChannelFlagsSuppressGuildReminderNotifications int = 1 << 2
	// Hide member join sticker reply buttons
	GuildSystemChannelFlagsSuppressJoinNotificationReplies int = 1 << 3
)

type APIGuildWelcomeScreen struct {
	// The welcome screen short message
	Description string `json:"description,omitempty"`
	// Array of suggested channels
	WelcomeChannels []APIGuildWelcomeScreenChannel `json:"welcome_channels"`
}

type APIGuildWelcomeScreenChannel struct {
	// The channel id that is suggested
	ChannelId globals.Snowflake `json:"channel_id"`
	/// The description shown for the channel
	Description string `json:"description"`
	// The emoji id of the emoji that is shown on the left of the channel
	EmojiId globals.Snowflake `json:"emoji_id,omitempty"`
	// The emoji name of the emoji that is shown on the left of the channel
	EmojiName string `json:"emoji_name,omitempty"`
}
