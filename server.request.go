package server

import (
	"encoding/json"
)

// Base struct describing server request
type ServerRequest struct {
	Method string `json:"method"`
	Params json.RawMessage `json:"params,string"`
}