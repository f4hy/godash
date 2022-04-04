package godash

func Find[T any](values []T, predicate func(T) bool) *T {
	for _, v := range values {
		if predicate(v) {
			return &v
		}
	}
	return nil
}
