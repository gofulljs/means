package slice

func MapUint8ToInt(src []uint8, m func(src uint8) int) []int {
	return MapUint8ToIntV2(src, m)
}

func MapUint8ToIntV1(src []uint8, m func(src uint8) int) []int {
	res := make([]int, 0, len(src))
	for i := 0; i < len(src); i++ {
		res = append(res, m(src[i]))
	}

	return res
}

func MapUint8ToIntV2(src []uint8, m func(src uint8) int) []int {
	res := make([]int, len(src))
	for i := 0; i < len(src); i++ {
		res[i] = m(src[i])
	}

	return res
}
