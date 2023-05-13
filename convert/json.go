package convert

import (
	"encoding/json"
	"fmt"
)

func FromJSON[T any](b []byte) (*T, error) {
	var v T
	if err := json.Unmarshal(b, &v); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return &v, nil
}
