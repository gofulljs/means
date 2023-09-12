package slice

func MapAny[D, S any](src []S, m func(src S) D) []D {
	return MapAnyV2[D, S](src, m)
}

func MapAnyV1[D, S any](src []S, m func(src S) D) []D {
	res := make([]D, 0, len(src))
	for i := 0; i < len(src); i++ {
		res = append(res, m(src[i]))
	}

	return res
}

func MapAnyV2[D, S any](src []S, m func(src S) D) []D {
	res := make([]D, len(src))
	for i := 0; i < len(src); i++ {
		res[i] = m(src[i])
	}

	return res
}
