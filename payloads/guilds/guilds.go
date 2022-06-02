package guilds

import (
	"github.com/denkylabs/discord-api-types-go/globals"
	"github.com/denkylabs/discord-api-types-go/payloads/emojis"
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
	Banner string `json:"banner"`
	// The description for the guild
	Description string `json:"description"`
	// Enabled guild features
	// See https://discord.com/developers/docs/resources/guild#guild-object-guild-features
	Features []string `json:"features"`
	// Verification level required for the guild
	// See https://discord.com/developers/docs/resources/guild#guild-object-verification-level
	VerificationLevel int `json:"verification_level"`
	// The vanity url code for the guild
	VanityURLCode string `json:"vanity_url_code"`
	// `true` if this guild is unavailable due to an outage
	Unavailable bool `json:"unavailable"`
	// The welcome screen of a Community guild, shown to new members
	// Returned in the invite object
	WelcomeScreen APIGuildWelcomeScreen `json:"welcome_screen"`
}

// https://discord.com/developers/docs/resources/guild#guild-object-guild-structure
type APIGuild struct {
	// Icon hash, returned when in the template object
	// See https://discord.com/developers/docs/reference#image-formatting
	IconHash string `json:"icon_hash"`
	// Discovery splash hash; only present for guilds with the "DISCOVERABLE" feature
	// See https://discord.com/developers/docs/reference#image-formatting
	DiscoverySplash string `json:"discovery_splash"`
	// `true` if the user is the owner of the guild
	// This field is only received from https://discord.com/developers/docs/resources/user#get-current-user-guilds
	Owner bool `json:"owner"`
	// ID of owner
	OwnerID globals.Snowflake `json:"owner_id"`
	// Total permissions for the user in the guild (excludes overrides)
	// This field is only received from https://discord.com/developers/docs/resources/user#get-current-user-guilds
	// See https://en.wikipedia.org/wiki/Bit_field
	Permissions globals.Permissions `json:"permissions"`
	// Voice region id for the guild
	// See https://discord.com/developers/docs/resources/voice#voice-region-object
	// @deprecated This field has been deprecated in favor of `rtc_region` on the channel.
	Region string `json:"region"`
	// ID of afk channel
	AfkChannelId globals.Snowflake `json:"afk_channel_id"`
	// AFK timeout in seconds
	AfkTimeout int `json:"afk_timeout"`
	// `true` if the guild widget is enabled
	WidgetEnabled bool `json:"widget_enabled"`
	// The channel id that the widget will generate an invite to, or `null` if set to no invite
	WidgetChannelId globals.Snowflake `json:"widget_channel_id"`
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
	// Custom guild emojis
	// See https://discord.com/developers/docs/resources/emoji#emoji-object
	Emojis []emojis.APIEmoji `json:"emojis"`
	// Enabled guild features
	// See https://discord.com/developers/docs/resources/guild#guild-object-guild-features
	Features []string `json:"features"`
	// Required MFA level for the guild
	// See https://discord.com/developers/docs/resources/guild#guild-object-mfa-level
	MFALevel int `json:"mfa_level"`
	// Application id of the guild creator if it is bot-created
	ApplicationId globals.Snowflake `json:"application_id"`
	// The id of the channel where guild notices such as welcome messages and boost events are posted
	SystemChannelId globals.Snowflake `json:"system_channel_id"`
	// System channel flags
	// See https://discord.com/developers/docs/resources/guild#guild-object-system-channel-flags
	SystemChannelFlags int `json:"system_channel_flags"`
	// The id of the channel where Community guilds can display rules and/or guidelines
	RulesChannelId globals.Snowflake `json:"rules_channel_id"`
	// The maximum number of presences for the guild (`null` is always returned, apart from the largest of guilds)
	MaxPresences int `json:"max_presences"`
	// The maximum number of members for the guild
	MaxMembers int `json:"max_members"`
	// The vanity url code for the guild
	VanityURLCode string `json:"vanity_url_code"`
	// Banner hash
	// See https://discord.com/developers/docs/reference#image-formatting
	Banner string `json:"banner"`
	// Premium tier (Server Boost level)
	// See https://discord.com/developers/docs/resources/guild#guild-object-premium-tier
	PremiumTier int `json:"premium_tier"`
	// The number of boosts this guild currently has
	PremiumSubscriptionCount int `json:"premium_subscription_count"`
	// The preferred locale of a Community guild; used in guild discovery and notices from Discord; defaults to "en-US"
	PreferredLocale string `json:"preferred_locale"`
	// The id of the channel where admins and moderators of Community guilds receive notices from Discord
	PublicUpdatesChannelId globals.Snowflake `json:"public_updates_channel_id"`
	// The maximum amount of users in a video channel
	MaxVideoChannelUsers int `json:"max_video_channel_users"`
	// **This field is only received from https://discord.com/developers/docs/resources/guild#get-guild with the `with_counts` query parameter set to `true`**
	ApproximateMemberCount int `json:"approximate_member_count"`
	// **This field is only received from https://discord.com/developers/docs/resources/guild#get-guild with the `with_counts` query parameter set to `true`**
	ApproximatePresenceCount int `json:"approximate_presence_count"`
	// The welcome screen of a Community guild, shown to new members
	// Returned in the invite object
	WelcomeScreen APIGuildWelcomeScreen `json:"welcome_screen"`
	// The nsfw level of the guild
	// See https://discord.com/developers/docs/resources/guild#guild-object-guild-nsfw-level
	NSFWLevel int `json:"nsfw_level"`
	// Custom guild stickers
	// See https://discord.com/developers/docs/resources/sticker#sticker-object
	Stickers []interface{} `json:"stickers"` // TODO: APISticker
	// Whether the guild has the boost progress bar enabled
	PremiumProgressBarEnabled bool `json:"premium_progress_bar_enabled"`
	// The type of Student Hub the guild is
	HubType int `json:"hub_type"`
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

// https://discord.com/developers/docs/resources/guild#get-guild-widget-image-widget-style-options
const (
	// Shield style widget with Discord icon and guild members online count
	GuildWidgetStyleShield = "shield"
	// Large image with guild icon, name and online count. "POWERED BY DISCORD" as the footer of the widget
	GuildWidgetStyleBanner1 = "banner1"
	// Smaller widget style with guild icon, name and online count. Split on the right with Discord logo
	GuildWidgetStyleBanner2 = "banner2"
	// Large image with guild icon, name and online count. In the footer, Discord logo on the left and "Chat Now" on the right
	GuildWidgetStyleBanner3 = "banner3"
	// Large Discord logo at the top of the widget. Guild icon, name and online count in the middle portion of the widget
	// and a "JOIN MY SERVER" button at the bottom
	GuildWidgetStyleBanner4 = "banner4"
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

// https://discord.com/developers/docs/resources/guild#guild-object-guild-features
const (
	// Guild has access to set an animated guild banner image
	GuildFeatureAnimatedBanner = "ANIMATED_BANNER"
	// Guild has access to set an animated guild icon
	GuildFeatureAnimatedIcon = "ANIMATED_ICON"
	//  Guild has access to set a guild banner image
	GuildFeatureBanner = "BANNER"
	// Guild has access to use commerce features (i.e. create store channels)
	GuildFeatureCommerce = "COMMERCE"
	// Guild can enable welcome screen, Membership Screening and discovery, and receives community updates
	GuildFeatureCommunity = "COMMUNITY"
	// Guild is able to be discovered in the directory
	GuildFeatureDiscoverable = "DISCOVERABLE"
	// Guild is able to be featured in the directory
	GuildFeatureFeaturable = "FEATURABLE"
	// Guild is listed in a directory channel
	GuildFeatureHasDirectoryEntry = "HAS_DIRECTORY_ENTRY"
	// Guild is a Student Hub
	// See https://support.discord.com/hc/en-us/articles/4406046651927-Discord-Student-Hubs-FAQ
	GuildFeatureHub = "HUB"
	// Guild has access to set an invite splash background
	GuildFeatureInviteSplash = "INVITE_SPLASH"
	// Guild is in a Student Hub
	// See https://support.discord.com/hc/en-us/articles/4406046651927-Discord-Student-Hubs-FAQ
	GuildFeatureLinkedToHub = "LINKED_TO_HUB"
	// Guild has enabled Membership Screening
	GuildFeatureMemberVerificationGateEnabled = "MEMBER_VERIFICATION_GATE_ENABLED"
	// Guild has enabled monetization
	GuildFeatureMonetizationEnabled = "MONETIZATION_ENABLED"
	// Guild has increased custom sticker slots
	GuildFeatureMoreStickers = "MORE_STICKERS"
	// Guild has access to create news channels
	GuildFeatureNews = "NEWS"
	// Guild is partnered
	GuildFeaturePartnered = "PARTNERED"
	// Guild can be previewed before joining via Membership Screening or the directory
	GuildFeaturePreviewEnabled = "PREVIEW_ENABLED"
	// Guild has access to create private threads
	GuildFeaturePrivateThreads = "PRIVATE_THREADS"
	GuildFeatureRelayEnabled   = "RELAY_ENABLED"
	// Guild is able to set role icons
	GuildFeatureRoleIcons = "ROLE_ICONS"
	// Guild has enabled ticketed events
	GuildFeatureTicketedEventsEnabled = "TICKETED_EVENTS_ENABLED"
	// Guild has access to set a vanity URL
	GuildFeatureVanityURL = "VANITY_URL"
	// Guild is verified
	GuildFeatureVerified = "VERIFIED"
	// Guild has access to set 384kbps bitrate in voice (previously VIP voice servers)
	GuildFeatureVIPRegions = "VIP_REGIONS"
	// Guild has enabled the welcome screen
	GuildFeatureWelcomeScreenEnabled = "WELCOME_SCREEN_ENABLED"
)

// https://discord.com/developers/docs/resources/guild#guild-preview-object
type APIGuildPreview struct {
	// Guild id
	Id globals.Snowflake `json:"id"`
	// Guild name (2-100 characters)
	Name string `json:"name"`
	// Icon hash
	// See https://discord.com/developers/docs/reference#image-formatting
	Icon string `json:"icon"`
	// Splash hash
	// See https://discord.com/developers/docs/reference#image-formatting
	Splash string `json:"splash"`
	// Discovery splash hash; only present for guilds with the "DISCOVERABLE" feature
	// See https://discord.com/developers/docs/reference#image-formatting
	DiscoverySplash string `json:"discovery_splash"`
	// Custom guild emojis
	// See https://discord.com/developers/docs/resources/emoji#emoji-object
	Emojis []interface{} `json:"emojis"` // TODO: APIEmoji
	// Enabled guild features
	// See https://discord.com/developers/docs/resources/guild#guild-object-guild-features
	Features []string `json:"features"`
	// Approximate number of members in this guild
	ApproximateMemberCount int `json:"approximate_member_count"`
	// Approximate number of online members in this guild
	ApproximatePresenceCount int `json:"approximate_presence_count"`
	// The description for the guild
	Description string `json:"description"`
	// Custom guild stickers
	Stickers []interface{} `json:"stickers"` // TODO: APISticker
}

type APIGuildWelcomeScreen struct {
	// The welcome screen short message
	Description string `json:"description"`
	// Array of suggested channels
	WelcomeChannels []APIGuildWelcomeScreenChannel `json:"welcome_channels"`
}

type APIGuildWelcomeScreenChannel struct {
	// The channel id that is suggested
	ChannelId globals.Snowflake `json:"channel_id"`
	/// The description shown for the channel
	Description string `json:"description"`
	// The emoji id of the emoji that is shown on the left of the channel
	EmojiId globals.Snowflake `json:"emoji_id"`
	// The emoji name of the emoji that is shown on the left of the channel
	EmojiName string `json:"emoji_name"`
}
