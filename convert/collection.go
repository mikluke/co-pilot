package convert

func Slice[from any, to any](s []from, f func(v from) to) []to {
	if s == nil {
		return nil
	}

	out := make([]to, len(s))
	for i, v := range s {
		out[i] = f(v)
	}

	return out
}

func Map[FromKey comparable, FromValue any, ToKey comparable, ToValue any](
	m map[FromKey]FromValue,
	f func(k FromKey, v FromValue) (ToKey, ToValue)) map[ToKey]ToValue {
	if m == nil {
		return nil
	}

	out := make(map[ToKey]ToValue, len(m))
	for k, v := range m {
		ck, cv := f(k, v)
		out[ck] = cv
	}

	return out
}
