package utils

import (
	"encoding/json"
	"io"
)

// ParseJSONResponse parses the JSON response from an HTTP response body.
// The target interface{} should be a pointer to the structure you want to unmarshal the JSON into.
func ParseJSONResponse(body io.Reader, target interface{}) error {
	decoder := json.NewDecoder(body)
	return decoder.Decode(target)
}
