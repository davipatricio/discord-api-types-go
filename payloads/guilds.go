package payloads

import (
	"github.com/denkylabs/discord-api-types-go/globals"
)

// https://discord.com/developers/docs/resources/guild#unavailable-guild-object
type APIUnavailableGuild struct {
	// Guild id
	Id globals.Snowflake `json:"id"`
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
	Features []GuildFeature `json:"features"`
	// Verification level required for the guild
	// See https://discord.com/developers/docs/resources/guild#guild-object-verification-level
	VerificationLevel GuildVerificationLevel `json:"verification_level"`
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
	// Guild name (2-100 characters, excluding trailing and leading whitespace)
	Name string `json:"name"`
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
	OwnerId globals.Snowflake `json:"owner_id"`
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
	AfkTimeout uint32 `json:"afk_timeout"`
	// `true` if the guild widget is enabled
	WidgetEnabled bool `json:"widget_enabled"`
	// The channel id that the widget will generate an invite to, or `null` if set to no invite
	WidgetChannelId globals.Snowflake `json:"widget_channel_id"`
	// Verification level required for the guild
	// See https://discord.com/developers/docs/resources/guild#guild-object-verification-level
	VerificationLevel GuildVerificationLevel `json:"verification_level"`
	// Default message notifications level
	// See https://discord.com/developers/docs/resources/guild#guild-object-default-message-notification-level
	DefaultMessageNotifications GuildDefaultMessageNotifications `json:"default_message_notifications"`
	// Explicit content filter level
	// See https://discord.com/developers/docs/resources/guild#guild-object-explicit-content-filter-level
	ExplicitContentFilter GuildExplicitContentFilterLevel `json:"explicit_content_filter"`
	// Roles in the guild
	// See https://discord.com/developers/docs/topics/permissions#role-object
	Roles []APIRole `json:"roles"`
	// Custom guild emojis
	// See https://discord.com/developers/docs/resources/emoji#emoji-object
	Emojis []APIEmoji `json:"emojis"`
	// Enabled guild features
	// See https://discord.com/developers/docs/resources/guild#guild-object-guild-features
	Features []GuildFeature `json:"features"`
	// Required MFA level for the guild
	// See https://discord.com/developers/docs/resources/guild#guild-object-mfa-level
	MFALevel GuildMFALevel `json:"mfa_level"`
	// Application id of the guild creator if it is bot-created
	ApplicationId globals.Snowflake `json:"application_id"`
	// The id of the channel where guild notices such as welcome messages and boost events are posted
	SystemChannelId globals.Snowflake `json:"system_channel_id"`
	// System channel flags
	// See https://discord.com/developers/docs/resources/guild#guild-object-system-channel-flags
	SystemChannelFlags GuildSystemChannelFlags `json:"system_channel_flags"`
	// The id of the channel where Community guilds can display rules and/or guidelines
	RulesChannelId globals.Snowflake `json:"rules_channel_id"`
	// The maximum number of presences for the guild (`null` is always returned, apart from the largest of guilds)
	MaxPresences uint32 `json:"max_presences"`
	// The maximum number of members for the guild
	MaxMembers uint32 `json:"max_members"`
	// The vanity url code for the guild
	VanityURLCode string `json:"vanity_url_code"`
	// Banner hash
	// See https://discord.com/developers/docs/reference#image-formatting
	Banner string `json:"banner"`
	// Premium tier (Server Boost level)
	// See https://discord.com/developers/docs/resources/guild#guild-object-premium-tier
	PremiumTier GuildPremiumTier `json:"premium_tier"`
	// The number of boosts this guild currently has
	PremiumSubscriptionCount uint32 `json:"premium_subscription_count"`
	// The preferred locale of a Community guild; used in guild discovery and notices from Discord; defaults to "en-US"
	PreferredLocale string `json:"preferred_locale"`
	// The id of the channel where admins and moderators of Community guilds receive notices from Discord
	PublicUpdatesChannelId globals.Snowflake `json:"public_updates_channel_id"`
	// The maximum amount of users in a video channel
	MaxVideoChannelUsers uint32 `json:"max_video_channel_users"`
	// **This field is only received from https://discord.com/developers/docs/resources/guild#get-guild with the `with_counts` query parameter set to `true`**
	ApproximateMemberCount uint32 `json:"approximate_member_count"`
	// **This field is only received from https://discord.com/developers/docs/resources/guild#get-guild with the `with_counts` query parameter set to `true`**
	ApproximatePresenceCount uint32 `json:"approximate_presence_count"`
	// The welcome screen of a Community guild, shown to new members
	// Returned in the invite object
	WelcomeScreen APIGuildWelcomeScreen `json:"welcome_screen"`
	// The nsfw level of the guild
	// See https://discord.com/developers/docs/resources/guild#guild-object-guild-nsfw-level
	NSFWLevel GuildNSFWLevel `json:"nsfw_level"`
	// Custom guild stickers
	// See https://discord.com/developers/docs/resources/sticker#sticker-object
	Stickers []APISticker `json:"stickers"`
	// Whether the guild has the boost progress bar enabled
	PremiumProgressBarEnabled bool `json:"premium_progress_bar_enabled"`
	// The type of Student Hub the guild is
	HubType GuildHubType `json:"hub_type"`
}

// https://discord.com/developers/docs/resources/guild#guild-object-default-message-notification-level
type GuildDefaultMessageNotifications uint16

// https://discord.com/developers/docs/resources/guild#guild-object-default-message-notification-level
const (
	GuildDefaultMessageNotificationsAllMessages GuildDefaultMessageNotifications = iota
	GuildDefaultMessageNotificationsOnlyMentions
)

// https://discord.com/developers/docs/resources/guild#guild-object-explicit-content-filter-level
type GuildExplicitContentFilterLevel uint16

// https://discord.com/developers/docs/resources/guild#guild-object-explicit-content-filter-level
const (
	GuildExplicitContentFilterDisabled GuildExplicitContentFilterLevel = iota
	GuildExplicitContentFilterMembersWithoutRoles
	GuildExplicitContentFilterAllMembers
)

// https://discord.com/developers/docs/resources/guild#guild-object-guild-nsfw-level
type GuildNSFWLevel uint16

// https://discord.com/developers/docs/resources/guild#guild-object-guild-nsfw-level
const (
	GuildNSFWLevelDefault GuildNSFWLevel = iota
	GuildNSFWLevelExplicit
	GuildNSFWLevelSafe
	GuildNSFWLevelAgeRestricted
)

// https://discord.com/developers/docs/resources/guild#guild-object-verification-level
type GuildVerificationLevel uint16

// https://discord.com/developers/docs/resources/guild#guild-object-verification-level
const (
	// Unrestricted
	GuildVerificationLevelNone GuildVerificationLevel = iota
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
type GuildMFALevel uint16

// https://discord.com/developers/docs/resources/guild#guild-object-mfa-level
const (
	GuildMFALevelNone GuildMFALevel = iota
	GuildMFALevelElevated
)

// https://discord.com/developers/docs/resources/guild#guild-object-premium-tier
type GuildPremiumTier uint16

// https://discord.com/developers/docs/resources/guild#guild-object-premium-tier
const (
	GuildPremiumTierNone GuildPremiumTier = iota
	GuildPremiumTierTier1
	GuildPremiumTierTier2
	GuildPremiumTierTier3
)

type GuildHubType uint8

const (
	GuildHubTypeDefault GuildHubType = iota
	GuildHubTypeHighSchool
	GuildHubTypeCollege
)

// https://discord.com/developers/docs/resources/guild#guild-object-system-channel-flags
type GuildSystemChannelFlags uint16

// https://discord.com/developers/docs/resources/guild#guild-object-system-channel-flags
const (
	// Suppress member join notifications
	GuildSystemChannelFlagsSuppressJoinNotifications GuildSystemChannelFlags = 1 << 0
	// Suppress server boost notifications
	GuildSystemChannelFlagsSuppressPremiumSubscriptions GuildSystemChannelFlags = 1 << 1
	// Suppress server setup tips
	GuildSystemChannelFlagsSuppressGuildReminderNotifications GuildSystemChannelFlags = 1 << 2
	// Hide member join sticker reply buttons
	GuildSystemChannelFlagsSuppressJoinNotificationReplies GuildSystemChannelFlags = 1 << 3
)

// https://discord.com/developers/docs/resources/guild#guild-object-guild-features
type GuildFeature string

// https://discord.com/developers/docs/resources/guild#guild-object-guild-features
const (
	// Guild has access to set an animated guild banner image
	GuildFeatureAnimatedBanner GuildFeature = "ANIMATED_BANNER"
	// Guild has access to set an animated guild icon
	GuildFeatureAnimatedIcon GuildFeature = "ANIMATED_ICON"
	//  Guild has access to set a guild banner image
	GuildFeatureBanner GuildFeature = "BANNER"
	// Guild has access to use commerce features (i.e. create store channels)
	GuildFeatureCommerce GuildFeature = "COMMERCE"
	// Guild can enable welcome screen, Membership Screening and discovery, and receives community updates
	GuildFeatureCommunity GuildFeature = "COMMUNITY"
	// Guild is able to be discovered in the directory
	GuildFeatureDiscoverable GuildFeature = "DISCOVERABLE"
	// Guild is able to be featured in the directory
	GuildFeatureFeaturable GuildFeature = "FEATURABLE"
	// Guild is listed in a directory channel
	GuildFeatureHasDirectoryEntry GuildFeature = "HAS_DIRECTORY_ENTRY"
	// Guild is a Student Hub
	// See https://support.discord.com/hc/en-us/articles/4406046651927-Discord-Student-Hubs-FAQ
	GuildFeatureHub GuildFeature = "HUB"
	// Guild has access to set an invite splash background
	GuildFeatureInviteSplash GuildFeature = "INVITE_SPLASH"
	// Guild is in a Student Hub
	// See https://support.discord.com/hc/en-us/articles/4406046651927-Discord-Student-Hubs-FAQ
	GuildFeatureLinkedToHub GuildFeature = "LINKED_TO_HUB"
	// Guild has enabled Membership Screening
	GuildFeatureMemberVerificationGateEnabled GuildFeature = "MEMBER_VERIFICATION_GATE_ENABLED"
	// Guild has enabled monetization
	GuildFeatureMonetizationEnabled GuildFeature = "MONETIZATION_ENABLED"
	// Guild has increased custom sticker slots
	GuildFeatureMoreStickers GuildFeature = "MORE_STICKERS"
	// Guild has access to create news channels
	GuildFeatureNews GuildFeature = "NEWS"
	// Guild is partnered
	GuildFeaturePartnered GuildFeature = "PARTNERED"
	// Guild can be previewed before joining via Membership Screening or the directory
	GuildFeaturePreviewEnabled GuildFeature = "PREVIEW_ENABLED"
	// Guild has access to create private threads
	GuildFeaturePrivateThreads GuildFeature = "PRIVATE_THREADS"
	GuildFeatureRelayEnabled   GuildFeature = "RELAY_ENABLED"
	// Guild is able to set role icons
	GuildFeatureRoleIcons GuildFeature = "ROLE_ICONS"
	// Guild has enabled ticketed events
	GuildFeatureTicketedEventsEnabled GuildFeature = "TICKETED_EVENTS_ENABLED"
	// Guild has access to set a vanity URL
	GuildFeatureVanityURL GuildFeature = "VANITY_URL"
	// Guild is verified
	GuildFeatureVerified GuildFeature = "VERIFIED"
	// Guild has access to set 384kbps bitrate in voice (previously VIP voice servers)
	GuildFeatureVIPRegions GuildFeature = "VIP_REGIONS"
	// Guild has enabled the welcome screen
	GuildFeatureWelcomeScreenEnabled GuildFeature = "WELCOME_SCREEN_ENABLED"
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
	Emojis APIEmoji `json:"emojis"`
	// Enabled guild features
	// See https://discord.com/developers/docs/resources/guild#guild-object-guild-features
	Features []GuildFeature `json:"features"`
	// Approximate number of members in this guild
	ApproximateMemberCount uint32 `json:"approximate_member_count"`
	// Approximate number of online members in this guild
	ApproximatePresenceCount uint32 `json:"approximate_presence_count"`
	// The description for the guild
	Description string `json:"description"`
	// Custom guild stickers
	Stickers []APISticker `json:"stickers"`
}

// https://discord.com/developers/docs/resources/guild#guild-widget-object
type APIGuildWidgetSettings struct {
	// Whether the widget is enabled
	Enabled bool `json:"enabled"`
	// The widget channel id
	ChannelId globals.Snowflake `json:"channel_id"`
}

// TODO: APIGuildMember
// https://discord.com/developers/docs/resources/guild#guild-member-object
type APIGuildMember struct{}

// https://discord.com/developers/docs/resources/guild#integration-object
type APIGuildIntegration struct {
	// Integration id
	Id globals.Snowflake `json:"id"`
	// Integration name
	Name string `json:"name"`
	// Integration type
	Type APIGuildIntegrationType `json:"type"`
	// Is this integration enabled
	Enabled bool `json:"enabled"`
	// Is this integration syncing
	// This field is not provided for `discord` bot integrations.
	Syncing bool `json:"syncing"`
	// ID that this integration uses for "subscribers"
	// This field is not provided for `discord` bot integrations.
	RoleId globals.Snowflake `json:"role_id"`
	// Whether emoticons should be synced for this integration (`twitch` only currently)
	// This field is not provided for `discord` bot integrations.
	EnableEmoticons bool `json:"enable_emoticons"`
	// The behavior of expiring subscribers
	//This field is not provided for `discord` bot integrations.
	// See https://discord.com/developers/docs/resources/guild#integration-object-integration-expire-behaviors
	ExpireBehavior IntegrationExpireBehavior `json:"expire_behavior"`
	// The grace period (in days) before expiring subscribers
	// This field is not provided for `discord` bot integrations.
	ExpireGracePeriod uint16 `json:"expire_grace_period"`
	// User for this integration
	// This field is not provided for `discord` bot integrations.
	// See https://discord.com/developers/docs/resources/user#user-object
	User APIUser `json:"user"`
	// Integration account information
	// See https://discord.com/developers/docs/resources/guild#integration-account-object
	Account APIIntegrationAccount `json:"account"`
	// When this integration was last synced
	// This field is not provided for `discord` bot integrations.
	SyncedAt string `json:"synced_at"`
	// How many subscribers this integration has
	// This field is not provided for `discord` bot integrations.
	SubscriberCount uint32 `json:"subscriber_count"`
	// Has this integration been revoked
	// This field is not provided for `discord` bot integrations.
	Revoked bool `json:"revoked"`
	// The bot/OAuth2 application for discord integrations
	// See https://discord.com/developers/docs/resources/guild#integration-application-object
	// This field is not provided for `discord` bot integrations.
	Application APIGuildIntegrationApplication `json:"application"`
}

type APIGuildIntegrationType string

const (
	APIGuildIntegrationTypeTwitch  APIGuildIntegrationType = "twitch"
	APIGuildIntegrationTypeYouTube APIGuildIntegrationType = "youtube"
	APIGuildIntegrationTypeDiscord APIGuildIntegrationType = "discord"
)

// https://discord.com/developers/docs/resources/guild#integration-object-integration-expire-behaviors
type IntegrationExpireBehavior uint8

// https://discord.com/developers/docs/resources/guild#integration-object-integration-expire-behaviors
const (
	IntegrationExpireBehaviorRemoveRole IntegrationExpireBehavior = iota
	IntegrationExpireBehaviorKick
)

// https://discord.com/developers/docs/resources/guild#integration-account-object
type APIIntegrationAccount struct {
	// ID of the account
	Id string `json:"id"`
	// Name of the account
	Name string `json:"name"`
}

type APIGuildIntegrationApplication struct {
	// The id of the app
	Id globals.Snowflake `json:"id"`
	// The name of the app
	Name string `json:"name"`
	// The icon hash of the app
	// See https://discord.com/developers/docs/reference#image-formatting
	Icon string `json:"icon"`
	// The description of the app
	Description string `json:"description"`
	// The bot associated with this application
	// See https://discord.com/developers/docs/resources/user#user-object
	Bot APIUser `json:"bot"`
}

// https://discord.com/developers/docs/resources/guild#ban-object
type APIBan struct {
	// The reason for the ban
	Reason string `json:"reason"`
	// The banned user
	User APIUser `json:"user"`
}

// https://discord.com/developers/docs/resources/guild#get-guild-widget-example-get-guild-widget
type APIGuildWidget struct {
	Id            globals.Snowflake       `json:"id"`
	Name          string                  `json:"name"`
	InstantInvite string                  `json:"instant_invite"`
	Channels      []APIGuildWidgetChannel `json:"channels"`
	Members       []APIGuildWidgetMember  `json:"members"`
	PresenceCount uint32                  `json:"presence_count"`
}

// https://discord.com/developers/docs/resources/guild#get-guild-widget-example-get-guild-widget
type APIGuildWidgetChannel struct {
	Id       globals.Snowflake `json:"id"`
	Name     string            `json:"name"`
	Position uint32            `json:"position"`
}

// https://discord.com/developers/docs/resources/guild#get-guild-widget-example-get-guild-widget
type APIGuildWidgetMember struct {
	Id            globals.Snowflake    `json:"id"`
	Username      string               `json:"username"`
	Discriminator string               `json:"discriminator"`
	Avatar        string               `json:"avatar"`
	Status        PresenceUpdateStatus `json:"status"`
	Activity      struct {
		Name string `json:"name"`
	} `json:"activity"`
	AvatarURL string `json:"avatar_url"`
}

// https://discord.com/developers/docs/resources/guild#get-guild-widget-image-widget-style-options
type GuildWidgetStyle string

// https://discord.com/developers/docs/resources/guild#get-guild-widget-image-widget-style-options
const (
	// Shield style widget with Discord icon and guild members online count
	GuildWidgetStyleShield GuildWidgetStyle = "shield"
	// Large image with guild icon, name and online count. "POWERED BY DISCORD" as the footer of the widget
	GuildWidgetStyleBanner1 GuildWidgetStyle = "banner1"
	// Smaller widget style with guild icon, name and online count. Split on the right with Discord logo
	GuildWidgetStyleBanner2 GuildWidgetStyle = "banner2"
	// Large image with guild icon, name and online count. In the footer, Discord logo on the left and "Chat Now" on the right
	GuildWidgetStyleBanner3 GuildWidgetStyle = "banner3"
	// Large Discord logo at the top of the widget. Guild icon, name and online count in the middle portion of the widget
	// and a "JOIN MY SERVER" button at the bottom
	GuildWidgetStyleBanner4 GuildWidgetStyle = "banner4"
)

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

type APIGuildMembershipScreening struct {
	// When the fields were last updated
	Version string `json:"version"`
	// The steps in the screening form
	FormFields []APIGuildMembershipScreeningField `json:"form_fields"`
	// The server description shown in the screening form
	Description string `json:"description"`
}

type APIGuildMembershipScreeningField struct {
	// The type of the field
	Type MembershipScreeningFieldType `json:"type"`
	// The title of the field
	Label string `json:"label"`
	// The list of rules
	Rules []string `json:"rules"`
	// Whether the user has to fill out this field
	Required bool `json:"required"`
}

type MembershipScreeningFieldType string

const (
	// Server Rules
	MembershipScreeningFieldTypeTerms MembershipScreeningFieldType = "TERMS"
)
