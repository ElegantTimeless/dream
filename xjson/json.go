package xjson

import (
	"encoding/json"
)

// IsJSON 检查是否是JSON
func IsJSON(s string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(s), &js) == nil
}
