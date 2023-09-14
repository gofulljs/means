package slice

import (
	"math/rand"
	"testing"
)

// 压测表明 V2 效果更好, 故采用 V2
// result:
// goos: linux
// goarch: amd64
// pkg: github.com/gofulljs/means/slice
// cpu: 11th Gen Intel(R) Core(TM) i5-11300H @ 3.10GHz
// BenchmarkMapUint8ToIntV1_MapUint8ToIntV2
// BenchmarkMapUint8ToIntV1_MapUint8ToIntV2/MapAnyV1
// BenchmarkMapUint8ToIntV1_MapUint8ToIntV2/MapAnyV1-8         	     346	   3669220 ns/op	 8003591 B/op	       1 allocs/op
// BenchmarkMapUint8ToIntV1_MapUint8ToIntV2/MapAnyV2
// BenchmarkMapUint8ToIntV1_MapUint8ToIntV2/MapAnyV2-8         	     523	   2397103 ns/op	 8003586 B/op	       1 allocs/op
// BenchmarkMapUint8ToIntV1_MapUint8ToIntV2/MapUint8ToIntV1
// BenchmarkMapUint8ToIntV1_MapUint8ToIntV2/MapUint8ToIntV1-8  	     453	   2614571 ns/op	 8003588 B/op	       1 allocs/op
// BenchmarkMapUint8ToIntV1_MapUint8ToIntV2/MapUint8ToIntV2
// BenchmarkMapUint8ToIntV1_MapUint8ToIntV2/MapUint8ToIntV2-8  	     499	   2370396 ns/op	 8003588 B/op	       1 allocs/op
// 可以看出 泛型版本 V2 与非泛型版本差距不大
func BenchmarkMapUint8ToIntV1_MapUint8ToIntV2(b *testing.B) {
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

	b.Run("MapUint8ToIntV1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MapUint8ToIntV1(src, func(src uint8) int {
				return int(src)
			})
		}
	})

	b.Run("MapUint8ToIntV2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MapUint8ToIntV2(src, func(src uint8) int {
				return int(src)
			})
		}
	})
}
