package invite

import (
	"github.com/denkylabs/discord-api-types-go/globals"
	"github.com/denkylabs/discord-api-types-go/payloads/application"
	"github.com/denkylabs/discord-api-types-go/payloads/guilds"
	"github.com/denkylabs/discord-api-types-go/payloads/user"
)

type APIInviteGuild struct {
	Id                       globals.Snowflake             `json:"id"`
	Name                     string                        `json:"name"`
	Splash                   string                        `json:"splash"`
	Banner                   string                        `json:"banner"`
	Icon                     string                        `json:"icon"`
	VanityURLCode            string                        `json:"vanity_url_code"`
	Description              string                        `json:"description"`
	Features                 []guilds.GuildFeature         `json:"features"`
	VerificationLevel        guilds.GuildVerificationLevel `json:"verification_level"`
	NSFWLevel                guilds.GuildNSFWLevel         `json:"nsfw_level"`
	PremiumSubscriptionCount uint32                        `json:"premium_subscription_count"`
}

// https://discord.com/developers/docs/resources/invite#invite-object
type APIInvite struct {
	// The invite code (unique ID)
	Code string `json:"code"`
	// The guild this invite is for
	// See https://discord.com/developers/docs/resources/guild#guild-object
	Guild APIInviteGuild `json:"guild"`
	// The channel this invite is for
	// See https://discord.com/developers/docs/resources/channel#channel-object
	Channel interface{} `json:"channel"` // TODO: APIChannel
	// The user who created the invite
	// See https://discord.com/developers/docs/resources/user#user-object
	Inviter user.APIUser `json:"inviter"`
	// The type of target for this voice channel invite
	// See https://discord.com/developers/docs/resources/invite#invite-object-target-user-types
	TargetType InviteTargetType `json:"target_type"`
	// The user whose stream to display for this voice channel stream invite
	// See https://discord.com/developers/docs/resources/user#user-object
	TargetUser user.APIUser `json:"target_user"`
	// The embedded application to open for this voice channel embedded application invite
	// See https://discord.com/developers/docs/topics/oauth2#application
	TargetApplication application.APIApplication `json:"target_application"`
	// Approximate count of online members, returned from the `GET /invites/<code>` endpoint when `with_counts` is `true`
	ApproximatePresenceCount uint32 `json:"approximate_presence_count"`
	// Approximate count of members, returned from the `GET /invites/<code>` endpoint when `with_counts` is `true`
	ApproximateMemberCount uint32 `json:"approximate_member_count"`
	// The expiration date of this invite, returned from the `GET /invites/<code>` endpoint when `with_expiration` is `true`
	ExpiresAt string `json:"expires_at"`
	// The stage instance data if there is a public stage instance in the stage channel this invite is for
	StageInstance interface{} `json:"-"` // TODO: APIStageInstance
	// The guild scheduled event data, returned from the `GET /invites/<code>` endpoint when `guild_scheduled_event_id` is a valid guild scheduled event id
	ScheduledEvent interface{} `json:"scheduled_event"` // TODO: APIScheduledEvent
}

// https://discord.com/developers/docs/resources/invite#invite-object-invite-target-types
type InviteTargetType uint8

// https://discord.com/developers/docs/resources/invite#invite-object-invite-target-types
const (
	InviteTargetTypeStream              InviteTargetType = 1
	InviteTargetTypeEmbeddedApplication InviteTargetType = 2
)

// https://discord.com/developers/docs/resources/invite#invite-metadata-object
type APIExtendedInvite struct {
	// The invite code (unique ID)
	Code string `json:"code"`
	// The guild this invite is for
	// See https://discord.com/developers/docs/resources/guild#guild-object
	Guild APIInviteGuild `json:"guild"`
	// The channel this invite is for
	// See https://discord.com/developers/docs/resources/channel#channel-object
	Channel interface{} `json:"channel"` // TODO: APIChannel
	// The user who created the invite
	// See https://discord.com/developers/docs/resources/user#user-object
	Inviter user.APIUser `json:"inviter"`
	// The type of target for this voice channel invite
	// See https://discord.com/developers/docs/resources/invite#invite-object-target-user-types
	TargetType InviteTargetType `json:"target_type"`
	// The user whose stream to display for this voice channel stream invite
	// See https://discord.com/developers/docs/resources/user#user-object
	TargetUser user.APIUser `json:"target_user"`
	// The embedded application to open for this voice channel embedded application invite
	// See https://discord.com/developers/docs/topics/oauth2#application
	TargetApplication application.APIApplication `json:"target_application"`
	// Approximate count of online members, returned from the `GET /invites/<code>` endpoint when `with_counts` is `true`
	ApproximatePresenceCount uint32 `json:"approximate_presence_count"`
	// Approximate count of members, returned from the `GET /invites/<code>` endpoint when `with_counts` is `true`
	ApproximateMemberCount uint32 `json:"approximate_member_count"`
	// The expiration date of this invite, returned from the `GET /invites/<code>` endpoint when `with_expiration` is `true`
	ExpiresAt string `json:"expires_at"`
	// The stage instance data if there is a public stage instance in the stage channel this invite is for
	StageInstance interface{} `json:"-"` // TODO: APIStageInstance
	// The guild scheduled event data, returned from the `GET /invites/<code>` endpoint when `guild_scheduled_event_id` is a valid guild scheduled event id
	ScheduledEvent interface{} `json:"scheduled_event"` // TODO: APIScheduledEvent
	// Number of times this invite has been used
	Uses uint32 `json:"uses"`
	// Max number of times this invite can be used
	MaxUses uint32 `json:"max_uses"`
	// Duration (in seconds) after which the invite expires
	MaxAge uint32 `json:"max_age"`
	// Whether this invite only grants temporary membership
	Temporary bool `json:"temporary"`
	// When this invite was created
	CreatedAt string `json:"created_at"`
}
