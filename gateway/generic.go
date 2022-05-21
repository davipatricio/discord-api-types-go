package gateway

type GatewayURLQuery struct {
	// The gateway version to use
	V string `json:"v"`
	// The encoding to use. Can be either "json" or "etf"
	Encoding string `json:"encoding"`
	// Which compression to use. Can empty or "zlib-stream"
	Compress string `json:"compress"`
}
