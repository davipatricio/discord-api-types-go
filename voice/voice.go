package voice

const VoiceGatewayVersion string = "4"

// https://discord.com/developers/docs/topics/opcodes-and-status-codes#voice-voice-opcodes
type VoiceOpcodes uint8

// https://discord.com/developers/docs/topics/opcodes-and-status-codes#voice-voice-opcodes
const (
	//Begin a voice websocket connection
	VoiceOpcodesIdentify VoiceOpcodes = iota
	// Select the voice protocol
	VoiceOpcodesSelectProtocol
	// Complete the websocket handshake
	VoiceOpcodesReady
	// Keep the websocket connection alive
	VoiceOpcodesHeartbeat
	// Describe the session
	VoiceOpcodesSessionDescription
	// Indicate which users are speaking
	VoiceOpcodesSpeaking
	// Sent to acknowledge a received client heartbeat
	VoiceOpcodesHeartbeatAck
	// Resume a connection
	VoiceOpcodesResume
	// Time to wait between sending heartbeats in milliseconds
	VoiceOpcodesHello
	// Acknowledge a successful session resume
	VoiceOpcodesResumed
	// A client has connected to the voice channel
	VoiceOpcodesClientConnect VoiceOpcodes = 12
	// A client has disconnected from the voice channel
	VoiceOpcodesClientDisconnect VoiceOpcodes = 13
)

// https://discord.com/developers/docs/topics/opcodes-and-status-codes#voice-voice-close-event-codes
type VoiceCloseCodes uint16

// https://discord.com/developers/docs/topics/opcodes-and-status-codes#voice-voice-close-event-codes
const (
	// You sent an invalid opcode
	VoiceCloseCodesUnknownOpcode VoiceCloseCodes = iota + 4001
	// You sent a invalid payload in your identifying to the Gateway
	VoiceCloseCodesFailedToDecode
	// You sent a payload before identifying with the Gateway
	VoiceCloseCodesNotAuthenticated
	// The token you sent in your identify payload is incorrect
	VoiceCloseCodesAuthenticationFailed
	// You sent more than one identify payload. Stahp
	VoiceCloseCodesAlreadyAuthenticated
	// Your session is no longer valid
	VoiceCloseCodesSessionNoLongerValid
	// Your session has timed out
	VoiceCloseCodesSessionTimeout VoiceCloseCodes = 4009
	// We can't find the server you're trying to connect to
	VoiceCloseCodesServerNotFound VoiceCloseCodes = 4011
	// We didn't recognize the protocol you sent
	VoiceCloseCodesUnknownProtocol VoiceCloseCodes = 4012
	// Either the channel was deleted you were kicked or the main gateway session was dropped. Should not reconnect
	VoiceCloseCodesDisconnected VoiceCloseCodes = 4014
	// The server crashed. Our bad! Try resuming
	VoiceCloseCodesVoiceServerCrashed VoiceCloseCodes = 4015
	// We didn't recognize your encryption
	VoiceCloseCodesUnknownEncryptionMode VoiceCloseCodes = 4016
)
