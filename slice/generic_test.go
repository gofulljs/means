package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapAny(t *testing.T) {
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
				return MapAny[int, uint8](src.([]uint8), func(src uint8) int {
					return int(src)
				})
			},
		},
		{
			name:    "[]byte to string",
			src:     [][]byte{[]byte("hi"), []byte("Hello world!")},
			wantRes: []string{"hi", "Hello world!"},
			fn: func(src any) any {
				return MapAny[string, []byte](src.([][]byte), func(src []byte) string {
					return string(src)
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
