package oauth2

type OAuth2Scopes string

const (
	// For oauth2 bots, this puts the bot in the user's selected guild by default
	OAuth2ScopesBot OAuth2Scopes = "bot"
	// Allows [/users/@me/connections](https://discord.com/developers/docs/resources/user#get-user-connections) to return linked third-party accounts
	// See https://discord.com/developers/docs/resources/user#get-user-connections
	OAuth2ScopesConnections OAuth2Scopes = "connections"
	// Allows your app to see information about the user's DMs and group DMs - requires Discord approval
	OAuth2ScopesDMChannelsRead OAuth2Scopes = "dm_channels.read"
	// Enables [/users/@me](https://discord.com/developers/docs/resources/user#get-current-user) to return an `email`
	// See https://discord.com/developers/docs/resources/user#get-current-user
	OAuth2ScopesEmail OAuth2Scopes = "email"
	// Allows [/users/@me](https://discord.com/developers/docs/resources/user#get-current-user) without `email`
	// See https://discord.com/developers/docs/resources/user#get-current-user
	OAuth2ScopesIdentify OAuth2Scopes = "identify"
	// Allows [/users/@me/guilds](https://discord.com/developers/docs/resources/user#get-current-user-guilds) to return basic information about all of a user's guilds
	// See https://discord.com/developers/docs/resources/user#get-current-user-guilds
	OAuth2ScopesGuilds OAuth2Scopes = "guilds"
	// Allows [/guilds/{guild.id}/members/{user.id}](https://discord.com/developers/docs/resources/guild#add-guild-member) to be used for joining users to a guild
	// See https://discord.com/developers/docs/resources/guild#add-guild-member
	OAuth2ScopesGuildsJoin OAuth2Scopes = "guilds.join"
	// Allows /users/@me/guilds/{guild.id}/member to return a user's member information in a guild
	// See https://discord.com/developers/docs/resources/user#get-current-user-guild-member
	OAuth2ScopesGuildsMembersRead OAuth2Scopes = "guilds.members.read"
	// Allows your app to join users to a group dm
	// See https://discord.com/developers/docs/resources/channel#group-dm-add-recipient
	OAuth2ScopesGroupDMJoins OAuth2Scopes = "gdm.join"
	// For local rpc server api access, this allows you to read messages from all client channels (otherwise restricted to channels/guilds your app creates)
	OAuth2ScopesMessagesRead OAuth2Scopes = "messages.read"
	// For local rpc server access, this allows you to control a user's local Discord client - requires Discord approval
	OAuth2ScopesRPC OAuth2Scopes = "rpc"
	// For local rpc server api access, this allows you to receive notifications pushed out to the user - requires Discord approval
	OAuth2ScopesRPCNotificationsRead OAuth2Scopes = "rpc.notifications.read"
	// This generates a webhook that is returned in the oauth token response for authorization code grants
	OAuth2ScopesWebhookIncoming OAuth2Scopes = "webhook.incoming"
	// Allows your app to connect to voice on user's behalf and see all the voice members - requires Discord approval
	OAuth2ScopesVoice OAuth2Scopes = "voice"
	// Allows your app to upload/update builds for a user's applications - requires Discord approval
	OAuth2ScopesApplicationsBuildsUpload OAuth2Scopes = "applications.builds.upload"
	// Allows your app to read build data for a user's applications
	OAuth2ScopesApplicationsBuildsRead OAuth2Scopes = "applications.builds.read"
	// Allows your app to read and update store data (SKUs, store listings, achievements, etc.) for a user's applications
	OAuth2ScopesApplicationsStoreUpdate OAuth2Scopes = "applications.store.update"
	// Allows your app to read entitlements for a user's applications
	OAuth2ScopesApplicationsEntitlements OAuth2Scopes = "applications.entitlements"
	// Allows your app to know a user's friends and implicit relationships - requires Discord approval
	RelationshipsRead OAuth2Scopes = "relationships.read"
	// Allows your app to fetch data from a user's "Now Playing/Recently Played" list - requires Discord approval
	OAuth2ScopesActivitiesRead OAuth2Scopes = "activities.read"
	// Allows your app to update a user's activity - requires Discord approval (NOT REQUIRED FOR GAMESDK ACTIVITY MANAGER)
	// See https://discord.com/developers/docs/game-sdk/activities
	OAuth2ScopesActivitiesWrite OAuth2Scopes = "activities.write"
	// Allows your app to use Application Commands in a guild
	// See https://discord.com/developers/docs/interactions/application-commands
	OAuth2ScopesApplicationsCommands OAuth2Scopes = "applications.commands"
	// Allows your app to update its Application Commands via this bearer token - client credentials grant only
	// See https://discord.com/developers/docs/interactions/application-commands
	OAuth2ScopesApplicationsCommandsUpdate OAuth2Scopes = "applications.commands.update"
	// Allows your app to update permissions for its commands using a Bearer token - client credentials grant only
	// See https://discord.com/developers/docs/interactions/application-commands
	OAuth2ScopesApplicationCommandsPermissionsUpdate OAuth2Scopes = "applications.commands.permissions.update"
)
