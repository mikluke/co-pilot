package convert

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompactJSON(t *testing.T) {
	t.Parallel()

	for _, v := range []struct {
		name string
		inp  []byte
		exp  []byte
	}{
		{
			name: "nil",
			inp:  nil,
			exp:  nil,
		},
		{
			name: "min",
			inp:  json.RawMessage(`{"a":["b","c"]}`),
			exp:  json.RawMessage(`{"a":["b","c"]}`),
		},
		{
			name: "case 1",
			inp:  json.RawMessage(`{"a" : ["b", "c"]}`),
			exp:  json.RawMessage(`{"a":["b","c"]}`),
		},
	} {
		tc := v
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			act, err := CompactJSON(tc.inp)
			require.NoError(t, err)
			require.Equal(t, tc.exp, act)
		})
	}
}
