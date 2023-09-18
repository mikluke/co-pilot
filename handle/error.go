package handle

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}

func IgnoreError[T any](v T, _ error) T {
	return v
}
