package slice

import "github.com/mikluke/co-pilot/convert"

func Contains[T comparable](s []T, v T) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}

	return false
}

func Unique[T comparable](s []T) []T {
	m := make(map[T]struct{}, len(s))

	for _, v := range s {
		m[v] = struct{}{}
	}

	return convert.Map2Slice(m, func(k T, _ struct{}) T {
		return k
	})
}
