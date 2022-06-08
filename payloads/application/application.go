package application

import (
	"github.com/denkylabs/discord-api-types-go/globals"
	"github.com/denkylabs/discord-api-types-go/payloads/user"
)

// https://discord.com/developers/docs/resources/application#application-object
type APIApplication struct {
	// The id of the app
	Id globals.Snowflake `json:"id"`
	// The name of the app
	Name string `json:"name"`
	// The icon hash of the app
	Icon string `json:"icon"`
	// The description of the app
	Description string `json:"description"`
	// An array of rpc origin urls, if rpc is enabled
	RpcOrigins []string `json:"rpc_origins"`
	// When `false` only app owner can join the app's bot to guilds
	BotPublic bool `json:"bot_public"`
	// When `true` the app's bot will only join upon completion of the full oauth2 code grant flow
	BotRequireCodeGrant bool `json:"bot_require_code_grant"`
	// The url of the application's terms of service
	TermsOfServiceUrl string `json:"terms_of_service_url"`
	// The url of the application's privacy policy
	PrivacyPolicyUrl string `json:"privacy_policy_url"`
	// Partial user object containing info on the owner of the application
	// See https://discord.com/developers/docs/resources/user#user-object
	Owner user.APIUser `json:"owner"`
	// The hexadecimal encoded key for verification in interactions and the GameSDK's GetTicket function
	// See https://discord.com/developers/docs/game-sdk/applications#get-ticket
	VerifyKey string `json:"verify_key"`
	// The team this application belongs to
	// See https://discord.com/developers/docs/topics/teams#data-models-team-object
	Team interface{} `json:"team"` // TODO: APITeam
	// If this application is a game sold on Discord, this field will be the guild to which it has been linked
	GuildId globals.Snowflake `json:"guild_id"`
	// If this application is a game sold on Discord, this field will be the id of the "Game SKU" that is created, if exists
	PrimarySkuId globals.Snowflake `json:"primary_sku_id"`
	// If this application is a game sold on Discord, this field will be the URL slug that links to the store page
	Slug string `json:"slug"`
	// If this application is a game sold on Discord, this field will be the hash of the image on store embeds
	CoverImage string `json:"cover_image"`
	// The application's public flags
	// See https://discord.com/developers/docs/resources/application#application-object-application-flags
	Flags []ApplicationFlags `json:"flags"`
	// Up to 5 tags describing the content and functionality of the application
	Tags []string `json:"tags"`
	// Settings for the application's default in-app authorization link, if enabled
	InstallParams APIApplicationInstallParams `json:"install_parameters"`
	// The application's default custom authorization link, if enabled
	CustomInstallUrl string `json:"custom_install_url"`
}

type APIApplicationInstallParams struct {
	Scopes      []interface{}       `json:"scopes"` // TODO: OAuth2Scopes
	Permissions globals.Permissions `json:"permissions"`
}

// https://discord.com/developers/docs/resources/application#application-object-application-flags
type ApplicationFlags uint32

// https://discord.com/developers/docs/resources/application#application-object-application-flags
const (
	ApplicationFlagsEmbeddedReleased              ApplicationFlags = 1 << 1
	ApplicationFlagsManagedEmoji                  ApplicationFlags = 1 << 2
	ApplicationFlagsGroupDMCreate                 ApplicationFlags = 1 << 4
	ApplicationFlagsRPCHasConnected               ApplicationFlags = 1 << 11
	ApplicationFlagsGatewayPresence               ApplicationFlags = 1 << 12
	ApplicationFlagsGatewayPresenceLimited        ApplicationFlags = 1 << 13
	ApplicationFlagsGatewayGuildMembers           ApplicationFlags = 1 << 14
	ApplicationFlagsGatewayGuildMembersLimited    ApplicationFlags = 1 << 15
	ApplicationFlagsVerificationPendingGuildLimit ApplicationFlags = 1 << 16
	ApplicationFlagsEmbedded                      ApplicationFlags = 1 << 17
	ApplicationFlagsGatewayMessageContent         ApplicationFlags = 1 << 18
	ApplicationFlagsGatewayMessageContentLimited  ApplicationFlags = 1 << 19
	ApplicationFlagsEmbeddedFirstParty            ApplicationFlags = 1 << 20
)
