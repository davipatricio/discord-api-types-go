package rest

// https://discord.com/developers/docs/topics/opcodes-and-status-codes#json-json-error-codes
type RESTJSONErrorCodes uint32

// https://discord.com/developers/docs/topics/opcodes-and-status-codes#json-json-error-codes
const (
	RESTJSONErrorCodesGeneralError RESTJSONErrorCodes = iota

	RESTJSONErrorCodesUnknownAccount             RESTJSONErrorCodes = 10001
	RESTJSONErrorCodesUnknownApplication         RESTJSONErrorCodes = 10002
	RESTJSONErrorCodesUnknownChannel             RESTJSONErrorCodes = 10003
	RESTJSONErrorCodesUnknownGuild               RESTJSONErrorCodes = 10004
	RESTJSONErrorCodesUnknownIntegration         RESTJSONErrorCodes = 10005
	RESTJSONErrorCodesUnknownInvite              RESTJSONErrorCodes = 10006
	RESTJSONErrorCodesUnknownMember              RESTJSONErrorCodes = 10007
	RESTJSONErrorCodesUnknownMessage             RESTJSONErrorCodes = 10008
	RESTJSONErrorCodesUnknownPermissionOverwrite RESTJSONErrorCodes = 10009
	RESTJSONErrorCodesUnknownProvider            RESTJSONErrorCodes = 10010
	RESTJSONErrorCodesUnknownRole                RESTJSONErrorCodes = 10011
	RESTJSONErrorCodesUnknownToken               RESTJSONErrorCodes = 10012
	RESTJSONErrorCodesUnknownUser                RESTJSONErrorCodes = 10013
	RESTJSONErrorCodesUnknownEmoji               RESTJSONErrorCodes = 10014
	RESTJSONErrorCodesUnknownWebhook             RESTJSONErrorCodes = 10015
	RESTJSONErrorCodesUnknownWebhookService      RESTJSONErrorCodes = 10016

	RESTJSONErrorCodesUnknownSession RESTJSONErrorCodes = 10020

	RESTJSONErrorCodesUnknownBan                  RESTJSONErrorCodes = 10026
	RESTJSONErrorCodesUnknownSKU                  RESTJSONErrorCodes = 10027
	RESTJSONErrorCodesUnknownStoreListing         RESTJSONErrorCodes = 10028
	RESTJSONErrorCodesUnknownEntitlement          RESTJSONErrorCodes = 10029
	RESTJSONErrorCodesUnknownBuild                RESTJSONErrorCodes = 10030
	RESTJSONErrorCodesUnknownLobby                RESTJSONErrorCodes = 10031
	RESTJSONErrorCodesUnknownBranch               RESTJSONErrorCodes = 10032
	RESTJSONErrorCodesUnknownStoreDirectoryLayout RESTJSONErrorCodes = 10033

	RESTJSONErrorCodesUnknownRedistributable RESTJSONErrorCodes = 10036

	RESTJSONErrorCodesUnknownGiftCode RESTJSONErrorCodes = 10038

	RESTJSONErrorCodesUnknownStream                         RESTJSONErrorCodes = 10049
	RESTJSONErrorCodesUnknownPremiumServerSubscribeCooldown RESTJSONErrorCodes = 10050

	RESTJSONErrorCodesUnknownGuildTemplate RESTJSONErrorCodes = 10057

	RESTJSONErrorCodesUnknownDiscoverableServerCategory RESTJSONErrorCodes = 10059
	RESTJSONErrorCodesUnknownSticker                    RESTJSONErrorCodes = 10060

	RESTJSONErrorCodesUnknownInteraction        RESTJSONErrorCodes = 10062
	RESTJSONErrorCodesUnknownApplicationCommand RESTJSONErrorCodes = 10063

	RESTJSONErrorCodesUnknownVoiceState                    RESTJSONErrorCodes = 10065
	RESTJSONErrorCodesUnknownApplicationCommandPermissions RESTJSONErrorCodes = 10066
	RESTJSONErrorCodesUnknownStageInstance                 RESTJSONErrorCodes = 10067
	RESTJSONErrorCodesUnknownGuildMemberVerificationForm   RESTJSONErrorCodes = 10068
	RESTJSONErrorCodesUnknownGuildWelcomeScreen            RESTJSONErrorCodes = 10069
	RESTJSONErrorCodesUnknownGuildScheduledEvent           RESTJSONErrorCodes = 10070
	RESTJSONErrorCodesUnknownGuildScheduledEventUser       RESTJSONErrorCodes = 10071

	RESTJSONErrorCodesUnknownTag RESTJSONErrorCodes = 10087

	RESTJSONErrorCodesBotsCannotUseThisEndpoint  RESTJSONErrorCodes = 20001
	RESTJSONErrorCodesOnlyBotsCanUseThisEndpoint RESTJSONErrorCodes = 20002

	RESTJSONErrorCodesExplicitContentCannotBeSentToTheDesiredRecipient RESTJSONErrorCodes = 20009

	RESTJSONErrorCodesNotAuthorizedToPerformThisActionOnThisApplication RESTJSONErrorCodes = 20012

	RESTJSONErrorCodesActionCannotBePerformedDueToSlowmodeRateLimit RESTJSONErrorCodes = 20016
	RESTJSONErrorCodesTheMazeIsntMeantForYou                        RESTJSONErrorCodes = 20017
	RESTJSONErrorCodesOnlyTheOwnerOfThisAccountCanPerformThisAction RESTJSONErrorCodes = 20018

	RESTJSONErrorCodesAnnouncementEditLimitExceeded RESTJSONErrorCodes = 20022

	RESTJSONErrorCodesUnderMinimumAge RESTJSONErrorCodes = 20024

	RESTJSONErrorCodesChannelSendRateLimit RESTJSONErrorCodes = 20028
	RESTJSONErrorCodesServerSendRateLimit  RESTJSONErrorCodes = 20029

	RESTJSONErrorCodesStageTopicServerNameServerDescriptionOrChannelNamesContainDisallowedWords = 20031

	RESTJSONErrorCodesGuildPremiumSubscriptionLevelTooLow = 20035

	RESTJSONErrorCodesMaximumNumberOfGuildsReached            RESTJSONErrorCodes = 30001
	RESTJSONErrorCodesMaximumNumberOfFriendsReached           RESTJSONErrorCodes = 30002
	RESTJSONErrorCodesMaximumNumberOfPinsReachedForTheChannel RESTJSONErrorCodes = 30003
	RESTJSONErrorCodesMaximumNumberOfRecipientsReached        RESTJSONErrorCodes = 30004
	RESTJSONErrorCodesMaximumNumberOfGuildRolesReached        RESTJSONErrorCodes = 30005

	RESTJSONErrorCodesMaximumNumberOfWebhooksReached RESTJSONErrorCodes = 30007
	RESTJSONErrorCodesMaximumNumberOfEmojisReached   RESTJSONErrorCodes = 30008

	RESTJSONErrorCodesMaximumNumberOfReactionsReached RESTJSONErrorCodes = 30010

	RESTJSONErrorCodesMaximumNumberOfGuildChannelsReached RESTJSONErrorCodes = 30013

	RESTJSONErrorCodesMaximumNumberOfAttachmentsInAMessageReached RESTJSONErrorCodes = 30015
	RESTJSONErrorCodesMaximumNumberOfInvitesReached               RESTJSONErrorCodes = 30016

	RESTJSONErrorCodesMaximumNumberOfAnimatedEmojisReached RESTJSONErrorCodes = 30018
	RESTJSONErrorCodesMaximumNumberOfServerMembersReached  RESTJSONErrorCodes = 30019

	RESTJSONErrorCodesMaximumNumberOfServerCategoriesReached RESTJSONErrorCodes = 30030

	RESTJSONErrorCodesGuildAlreadyHasTemplate RESTJSONErrorCodes = 30031

	RESTJSONErrorCodesMaximumThreadParticipants RESTJSONErrorCodes = 30033

	RESTJSONErrorCodesMaximumNumberOfNonGuildMemberBansHasBeenExceeded RESTJSONErrorCodes = 30035

	RESTJSONErrorCodesMaximumNumberOfBanFetchesHasBeenReached               RESTJSONErrorCodes = 30037
	RESTJSONErrorCodesMaximumNumberOfUncompletedGuildScheduledEventsReached RESTJSONErrorCodes = 30038

	RESTJSONErrorCodesMaximumNumberOfStickersReached             RESTJSONErrorCodes = 30039
	RESTJSONErrorCodesMaximumNumberOfPruneRequestsHasBeenReached RESTJSONErrorCodes = 30040

	RESTJSONErrorCodesMaximumNumberOfGuildWidgetSettingsUpdatesHasBeenReached RESTJSONErrorCodes = 30042

	RESTJSONErrorCodesMaximumNumberOfEditsToMessagesOlderThanOneHourReached RESTJSONErrorCodes = 30046
	RESTJSONErrorCodesMaximumNumberOfPinnedThreadsInForumHasBeenReached     RESTJSONErrorCodes = 30047
	RESTJSONErrorCodesMaximumNumberOfTagsInForumHasBeenReached              RESTJSONErrorCodes = 30048

	RESTJSONErrorCodesBitrateIsTooHighForChannelOfThisType RESTJSONErrorCodes = 30052

	RESTJSONErrorCodesUnauthorized                           RESTJSONErrorCodes = 40001
	RESTJSONErrorCodesVerifyYourAccount                      RESTJSONErrorCodes = 40002
	RESTJSONErrorCodesOpeningDirectMessagesTooFast           RESTJSONErrorCodes = 40003
	RESTJSONErrorCodesSendMessagesHasBeenTemporarilyDisabled RESTJSONErrorCodes = 40004
	RESTJSONErrorCodesRequestEntityTooLarge                  RESTJSONErrorCodes = 40005
	RESTJSONErrorCodesFeatureTemporarilyDisabledServerSide   RESTJSONErrorCodes = 40006
	RESTJSONErrorCodesUserBannedFromThisGuild                RESTJSONErrorCodes = 40007

	RESTJSONErrorCodesTargetUserIsNotConnectedToVoice  RESTJSONErrorCodes = 40032
	RESTJSONErrorCodesThisMessageWasAlreadyCrossposted RESTJSONErrorCodes = 40033

	RESTJSONErrorCodesApplicationCommandWithThatNameAlreadyExists RESTJSONErrorCodes = 40041

	RESTJSONErrorCodesInteractionHasAlreadyBeenAcknowledged RESTJSONErrorCodes = 40060
	RESTJSONErrorCodesTagNamesMustBeUnique                  RESTJSONErrorCodes = 40061

	RESTJSONErrorCodesMissingAccess                                     RESTJSONErrorCodes = 50001
	RESTJSONErrorCodesInvalidAccountType                                RESTJSONErrorCodes = 50002
	RESTJSONErrorCodesCannotExecuteActionOnDMChannel                    RESTJSONErrorCodes = 50003
	RESTJSONErrorCodesGuildWidgetDisabled                               RESTJSONErrorCodes = 50004
	RESTJSONErrorCodesCannotEditMessageAuthoredByAnotherUser            RESTJSONErrorCodes = 50005
	RESTJSONErrorCodesCannotSendAnEmptyMessage                          RESTJSONErrorCodes = 50006
	RESTJSONErrorCodesCannotSendMessagesToThisUser                      RESTJSONErrorCodes = 50007
	RESTJSONErrorCodesCannotSendMessagesInNonTextChannel                RESTJSONErrorCodes = 50008
	RESTJSONErrorCodesChannelVerificationLevelTooHighForYouToGainAccess RESTJSONErrorCodes = 50009
	RESTJSONErrorCodesOAuth2ApplicationDoesNotHaveBot                   RESTJSONErrorCodes = 50010
	RESTJSONErrorCodesOAuth2ApplicationLimitReached                     RESTJSONErrorCodes = 50011
	RESTJSONErrorCodesInvalidOAuth2State                                RESTJSONErrorCodes = 50012
	RESTJSONErrorCodesMissingPermissions                                RESTJSONErrorCodes = 50013
	RESTJSONErrorCodesInvalidToken                                      RESTJSONErrorCodes = 50014
	RESTJSONErrorCodesNoteWasTooLong                                    RESTJSONErrorCodes = 50015
	RESTJSONErrorCodesProvidedTooFewOrTooManyMessagesToDelete           RESTJSONErrorCodes = 50016
	RESTJSONErrorCodesInvalidMFALevel                                   RESTJSONErrorCodes = 50017

	RESTJSONErrorCodesMessageCanOnlyBePinnedInTheChannelItWasSentIn RESTJSONErrorCodes = 50019
	RESTJSONErrorCodesInviteCodeInvalidOrTaken                      RESTJSONErrorCodes = 50020
	RESTJSONErrorCodesCannotExecuteActionOnSystemMessage            RESTJSONErrorCodes = 50021

	RESTJSONErrorCodesCannotExecuteActionOnThisChannelType RESTJSONErrorCodes = 50024
	RESTJSONErrorCodesInvalidOAuth2AccessToken             RESTJSONErrorCodes = 50025
	RESTJSONErrorCodesMissingRequiredOAuth2Scope           RESTJSONErrorCodes = 50026

	RESTJSONErrorCodesInvalidWebhookToken RESTJSONErrorCodes = 50027
	RESTJSONErrorCodesInvalidRole         RESTJSONErrorCodes = 50028

	RESTJSONErrorCodesInvalidRecipients                              RESTJSONErrorCodes = 50033
	RESTJSONErrorCodesOneOfTheMessagesProvidedWasTooOldForBulkDelete RESTJSONErrorCodes = 50034
	RESTJSONErrorCodesInvalidFormBodyOrContentType                   RESTJSONErrorCodes = 50035
	RESTJSONErrorCodesInviteAcceptedToGuildWithoutTheBotBeingIn      RESTJSONErrorCodes = 50036

	RESTJSONErrorCodesInvalidAPIVersion RESTJSONErrorCodes = 50041

	RESTJSONErrorCodesFileUploadedExceedsMaximumSize RESTJSONErrorCodes = 50045
	RESTJSONErrorCodesInvalidFileUploaded            RESTJSONErrorCodes = 50046

	RESTJSONErrorCodesCannotSelfRedeemThisGift RESTJSONErrorCodes = 50054
	RESTJSONErrorCodesInvalidGuild             RESTJSONErrorCodes = 50055

	RESTJSONErrorCodesInvalidMessageType RESTJSONErrorCodes = 50068

	RESTJSONErrorCodesPaymentSourceRequiredToRedeemGift RESTJSONErrorCodes = 50070

	RESTJSONErrorCodesCannotDeleteChannelRequiredForCommunityGuilds RESTJSONErrorCodes = 50074

	RESTJSONErrorCodesCannotEditStickersWithinMessage RESTJSONErrorCodes = 50080
	RESTJSONErrorCodesInvalidStickerSent              RESTJSONErrorCodes = 50081

	RESTJSONErrorCodesInvalidActionOnArchivedThread             RESTJSONErrorCodes = 50083
	RESTJSONErrorCodesInvalidThreadNotificationSettings         RESTJSONErrorCodes = 50084
	RESTJSONErrorCodesParameterEarlierThanCreation              RESTJSONErrorCodes = 50085
	RESTJSONErrorCodesCommunityServerChannelsMustBeTextChannels RESTJSONErrorCodes = 50086

	RESTJSONErrorCodesServerNotAvailableInYourLocation RESTJSONErrorCodes = 50095

	RESTJSONErrorCodesServerNeedsMonetizationEnabledToPerformThisAction RESTJSONErrorCodes = 50097

	RESTJSONErrorCodesServerNeedsMoreBoostsToPerformThisAction RESTJSONErrorCodes = 50101

	RESTJSONErrorCodesRequestBodyContainsInvalidJSON RESTJSONErrorCodes = 50109

	RESTJSONErrorCodesYouDoNotHavePermissionToSendThisSticker RESTJSONErrorCodes = 50600

	RESTJSONErrorCodesTwoFactorAuthenticationIsRequired RESTJSONErrorCodes = 60003

	RESTJSONErrorCodesNoUsersWithDiscordTagExist RESTJSONErrorCodes = 80004

	RESTJSONErrorCodesReactionWasBlocked RESTJSONErrorCodes = 90001

	RESTJSONErrorCodesAPIResourceOverloaded RESTJSONErrorCodes = 130000

	RESTJSONErrorCodesTheStageIsAlreadyOpen RESTJSONErrorCodes = 150006

	RESTJSONErrorCodesCannotReplyWithoutPermissionToReadMessageHistory RESTJSONErrorCodes = 160002

	RESTJSONErrorCodesThreadAlreadyCreatedForMessage   RESTJSONErrorCodes = 160004
	RESTJSONErrorCodesThreadLocked                     RESTJSONErrorCodes = 160005
	RESTJSONErrorCodesMaximumActiveThreads             RESTJSONErrorCodes = 160006
	RESTJSONErrorCodesMaximumActiveAnnouncementThreads RESTJSONErrorCodes = 160007

	RESTJSONErrorCodesInvalidJSONForUploadedLottieFile                 RESTJSONErrorCodes = 170001
	RESTJSONErrorCodesUploadedLottiesCannotContainRasterizedImages     RESTJSONErrorCodes = 170002
	RESTJSONErrorCodesStickerMaximumFramerateExceeded                  RESTJSONErrorCodes = 170003
	RESTJSONErrorCodesStickerFrameCountExceedsMaximumOf1000Frames      RESTJSONErrorCodes = 170004
	RESTJSONErrorCodesLottieAnimationMaximumDimensionsExceeded         RESTJSONErrorCodes = 170005
	RESTJSONErrorCodesStickerFramerateIsTooSmallOrTooLarge             RESTJSONErrorCodes = 170006
	RESTJSONErrorCodesStickerAnimationDurationExceedsMaximumOf5Seconds RESTJSONErrorCodes = 170007

	RESTJSONErrorCodesCannotUpdateAFinishedEvent RESTJSONErrorCodes = 180000

	RESTJSONErrorCodesFailedToCreateStageNeededForStageEvent RESTJSONErrorCodes = 180002

	RESTJSONErrorCodesWebhooksCanOnlyCreateThreadsInForumChannels RESTJSONErrorCodes = 220003
)
