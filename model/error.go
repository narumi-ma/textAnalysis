package model

type ErrorResponse struct {
	ID      interface{} `json:"id"`      // Can be nil or another type like string, hence interface{}
	JSONRPC string      `json:"jsonrpc"` // The fixed value returned is "2.0"
	Error   ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
