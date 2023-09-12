package slice

import (
	"math/rand"
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

// 压测表明 V2 效果更好, 故采用 V2
func BenchmarkMapAnyV1_MapAnyV2(b *testing.B) {
	src := make([]uint8, 1000000)
	for j := 0; j < 1000000; j++ {
		src[j] = uint8(rand.Uint32())
	}

	b.ResetTimer() // 重置计时器

	b.Run("MapAnyV1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MapAnyV1[int, uint8](src, func(src uint8) int {
				return int(src)
			})
		}
	})

	b.Run("MapAnyV2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MapAnyV2[int, uint8](src, func(src uint8) int {
				return int(src)
			})
		}
	})
}
