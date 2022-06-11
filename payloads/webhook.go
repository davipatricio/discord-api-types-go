package payloads

import "github.com/denkylabs/discord-api-types-go/globals"

// https://discord.com/developers/docs/resources/webhook#webhook-object
type APIWebhook struct {
	//The id of the webhook
	Id globals.Snowflake `json:"id"`
	//The type of the webhook
	// See https://discord.com/developers/docs/resources/webhook#webhook-object-webhook-types
	Type WebhookType `json:"type"`
	//The guild id this webhook is for
	GuildId globals.Snowflake `json:"guild_id"`
	//The channel id this webhook is for
	ChannelId globals.Snowflake `json:"channel_id"`
	//The user this webhook was created by (not returned when getting a webhook with its token)
	// See https://discord.com/developers/docs/resources/user#user-object
	User APIUser `json:"user"`
	//The default name of the webhook
	Name string `json:"name"`
	//The default avatar of the webhook
	Avatar string `json:"avatar"`
	//The secure token of the webhook (returned for Incoming Webhooks)
	Token string `json:"token"`
	//The bot/OAuth2 application that created this webhook
	ApplicationId globals.Snowflake `json:"application_id"`
	//The guild of the channel that this webhook is following (returned for Channel Follower Webhooks)
	SourceGuild APIPartialGuild `json:"source_guild"`
	//The channel that this webhook is following (returned for Channel Follower Webhooks)
	SourceChannel interface{} `json:"source_channel"` // TODO: APIPartialChannel
	//The url used for executing the webhook (returned by the webhooks OAuth2 flow)
	Url string `json:"url"`
}

type WebhookType uint8

const (
	// Incoming Webhooks can post messages to channels with a generated token
	WebhookTypeIncoming        WebhookType = 1
	// Channel Follower Webhooks are internal webhooks used with Channel Following to post new messages into channels
	WebhookTypeChannelFollower WebhookType = 2
	// Application webhooks are webhooks used with Interactions
	WebhookTypeApplication     WebhookType = 3
)
