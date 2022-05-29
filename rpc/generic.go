package rpc

// https://discord.com/developers/docs/topics/opcodes-and-status-codes#rpc-rpc-error-codes
const (
	RPCErrorrUnknownError                   int = 1000
	RPCErrorInvalidPayload                  int = 4000
	RPCErrorInvalidCommand                  int = 4002
	RPCErrorInvalidGuild                    int = 4003
	RPCErrorInvalidEvent                    int = 4004
	RPCErrorInvalidChannel                  int = 4005
	RPCErrorInvalidPermissions              int = 4006
	RPCErrorInvalidClientId                 int = 4007
	RPCErrorInvalidOrigin                   int = 4008
	RPCErrorInvalidToken                    int = 4009
	RPCErrorInvalidUser                     int = 4010
	RPCErrorOAuth2Error                     int = 5000
	RPCErrorSelectChannelTimedOut           int = 5001
	RPCErrorGetGuildTimedOut                int = 5002
	RPCErrorSelectVoiceForceRequired        int = 5003
	RPCErrorCaptureShortcutAlreadyListening int = 5004
)

// https://discord.com/developers/docs/topics/opcodes-and-status-codes#rpc-rpc-error-codes
const (
	RPCCloseEventInvalidClientId int = 4000
	RPCCloseEventInvalidOrigin   int = 4001
	RPCCloseEventRateLimited     int = 4002
	RPCCloseEventTokenRevoked    int = 4003
	RPCCloseEventInvalidVersion  int = 4004
	RPCCloseEventInvalidEncoding int = 4005
)
