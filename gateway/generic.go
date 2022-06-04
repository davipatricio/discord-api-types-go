package gateway

type CompressionType string

const (
	CompressionTypeNone       CompressionType = ""
	CompressionTypeZLibStream CompressionType = "zlib-stream"
)

// When initially creating and handshaking connections to the Gateway, a user can choose whether they wish to communicate over plain-text JSON or binary ETF.
type Encoding string

const (
	EncodingNone Encoding = ""
	EncodingJSON Encoding = "json"
	EncodingETF  Encoding = "etf"
)

type GatewayURLQuery struct {
	// The gateway version to use
	V string `json:"v"`
	// The encoding of received gateway packets.
	Encoding Encoding `json:"encoding"`
	// The (optional) compression of gateway packets.
	Compress CompressionType `json:"compress"`
}
