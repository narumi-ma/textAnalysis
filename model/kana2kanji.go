package model

// Request parameters (POST)

type Request_K struct {
	ID      interface{} `json:"id"`      // integer or string
	JSONRPC string      `json:"jsonrpc"` // The value should be "2.0".
	Method  string      `json:"method"`  // The value should be "jlp.jimservice.conversion"
	Params  Params_K    `json:"params"`
}

type Params_K struct {
	Q          string   `json:"q"` // The text to be converted
	Format     string   `json:"format"`
	Mode       string   `json:"mode"`
	Option     []string `json:"option"`
	Dictionary []string `json:"dictionary"`
	Results    int      `json:"results"`
}
