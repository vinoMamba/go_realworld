package utils

import "encoding/json"

func JsonMarshal(v any) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bytes)
}
