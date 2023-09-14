package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapUint8ToInt(t *testing.T) {
	tests := []struct {
		name    string
		src     any
		wantRes any
		fn      func(src any) any
	}{
		{
			name:    "uint8 to int",
			src:     []uint8{1, 2, 3},
			wantRes: []int{1, 2, 3},
			fn: func(src any) any {
				return MapUint8ToIntV1(src.([]uint8), func(src uint8) int {
					return int(src)
				})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.fn(tt.src)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
