package helper

import (
	"github.com/goccy/go-json"
)

// JSONMarshal is to stringify struct/map into json string.
func JSONMarshal(p any) string {
	b, err := json.Marshal(p)
	if err != nil {
		return "{}"
	}

	return string(b)
}

// JSONUnmarshal is to parse string into given struct.
func JSONUnmarshal[T any](s string, data *T) T {
	if data == nil {
		var res T

		if err := json.Unmarshal([]byte(s), &res); err != nil {
			return res
		}

		return res
	}

	if err := json.Unmarshal([]byte(s), data); err != nil {
		return *data
	}

	return *data
}
