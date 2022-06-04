package rpc

// https://discord.com/developers/docs/topics/opcodes-and-status-codes#rpc-rpc-error-codes
type RPCError uint16

// https://discord.com/developers/docs/topics/opcodes-and-status-codes#rpc-rpc-error-codes
const (
	RPCErrorrUnknownError                   RPCError = 1000
	RPCErrorInvalidPayload                  RPCError = 4000
	RPCErrorInvalidCommand                  RPCError = 4002
	RPCErrorInvalidGuild                    RPCError = 4003
	RPCErrorInvalidEvent                    RPCError = 4004
	RPCErrorInvalidChannel                  RPCError = 4005
	RPCErrorInvalidPermissions              RPCError = 4006
	RPCErrorInvalidClientId                 RPCError = 4007
	RPCErrorInvalidOrigin                   RPCError = 4008
	RPCErrorInvalidToken                    RPCError = 4009
	RPCErrorInvalidUser                     RPCError = 4010
	RPCErrorOAuth2Error                     RPCError = 5000
	RPCErrorSelectChannelTimedOut           RPCError = 5001
	RPCErrorGetGuildTimedOut                RPCError = 5002
	RPCErrorSelectVoiceForceRequired        RPCError = 5003
	RPCErrorCaptureShortcutAlreadyListening RPCError = 5004
)

// https://discord.com/developers/docs/topics/opcodes-and-status-codes#rpc-rpc-error-codes
type RPCCloseEvent uint16

// https://discord.com/developers/docs/topics/opcodes-and-status-codes#rpc-rpc-error-codes
const (
	RPCCloseEventInvalidClientId RPCCloseEvent = 4000
	RPCCloseEventInvalidOrigin   RPCCloseEvent = 4001
	RPCCloseEventRateLimited     RPCCloseEvent = 4002
	RPCCloseEventTokenRevoked    RPCCloseEvent = 4003
	RPCCloseEventInvalidVersion  RPCCloseEvent = 4004
	RPCCloseEventInvalidEncoding RPCCloseEvent = 4005
)
