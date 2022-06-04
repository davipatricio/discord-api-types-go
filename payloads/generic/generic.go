package generic

// https://discord.com/developers/docs/topics/permissions#permissions-bitwise-permission-flags
type PermissionFlagsBits uint64

// https://discord.com/developers/docs/topics/permissions#permissions-bitwise-permission-flags
const (
	PermissionFlagsBitsCreateInstantInvite     PermissionFlagsBits = 1 << 0
	PermissionFlagsBitsKickMembers             PermissionFlagsBits = 1 << 1
	PermissionFlagsBitsBanMembers              PermissionFlagsBits = 1 << 2
	PermissionFlagsBitsAdministrator           PermissionFlagsBits = 1 << 3
	PermissionFlagsBitsManageChannels          PermissionFlagsBits = 1 << 4
	PermissionFlagsBitsManageGuild             PermissionFlagsBits = 1 << 5
	PermissionFlagsBitsAddReactions            PermissionFlagsBits = 1 << 6
	PermissionFlagsBitsViewAuditLog            PermissionFlagsBits = 1 << 7
	PermissionFlagsBitsPrioritySpeaker         PermissionFlagsBits = 1 << 8
	PermissionFlagsBitsStream                  PermissionFlagsBits = 1 << 9
	PermissionFlagsBitsViewChannel             PermissionFlagsBits = 1 << 10
	PermissionFlagsBitsSendMessages            PermissionFlagsBits = 1 << 11
	PermissionFlagsBitsSendTTSMessages         PermissionFlagsBits = 1 << 12
	PermissionFlagsBitsManageMessages          PermissionFlagsBits = 1 << 13
	PermissionFlagsBitsEmbedLinks              PermissionFlagsBits = 1 << 14
	PermissionFlagsBitsAttachFiles             PermissionFlagsBits = 1 << 15
	PermissionFlagsBitsReadMessageHistory      PermissionFlagsBits = 1 << 16
	PermissionFlagsBitsMentionEveryone         PermissionFlagsBits = 1 << 17
	PermissionFlagsBitsUseExternalEmojis       PermissionFlagsBits = 1 << 18
	PermissionFlagsBitsViewGuildInsights       PermissionFlagsBits = 1 << 19
	PermissionFlagsBitsConnect                 PermissionFlagsBits = 1 << 20
	PermissionFlagsBitsSpeak                   PermissionFlagsBits = 1 << 21
	PermissionFlagsBitsMuteMembers             PermissionFlagsBits = 1 << 22
	PermissionFlagsBitsDeafenMembers           PermissionFlagsBits = 1 << 23
	PermissionFlagsBitsMoveMembers             PermissionFlagsBits = 1 << 24
	PermissionFlagsBitsUseVAD                  PermissionFlagsBits = 1 << 25
	PermissionFlagsBitsChangeNickname          PermissionFlagsBits = 1 << 26
	PermissionFlagsBitsManageNicknames         PermissionFlagsBits = 1 << 27
	PermissionFlagsBitsManageRoles             PermissionFlagsBits = 1 << 28
	PermissionFlagsBitsManageWebhooks          PermissionFlagsBits = 1 << 29
	PermissionFlagsBitsManageEmojisAndStickers PermissionFlagsBits = 1 << 30
	PermissionFlagsBitsUseApplicationCommands  PermissionFlagsBits = 1 << 31
	PermissionFlagsBitsRequestToSpeak          PermissionFlagsBits = 1 << 32
	PermissionFlagsBitsManageEvents            PermissionFlagsBits = 1 << 33
	PermissionFlagsBitsManageThreads           PermissionFlagsBits = 1 << 34
	PermissionFlagsBitsCreatePublicThreads     PermissionFlagsBits = 1 << 35
	PermissionFlagsBitsCreatePrivateThreads    PermissionFlagsBits = 1 << 36
	PermissionFlagsBitsUseExternalStickers     PermissionFlagsBits = 1 << 37
	PermissionFlagsBitsSendMessagesInThreads   PermissionFlagsBits = 1 << 38
	PermissionFlagsBitsUseEmbeddedActivities   PermissionFlagsBits = 1 << 39
	PermissionFlagsBitsModerateMembers         PermissionFlagsBits = 1 << 40
)
