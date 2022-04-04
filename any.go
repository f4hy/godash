package godash

func Any[T any](values []T, predicate func(T) bool) bool {
	for _, v := range values {
		if predicate(v) {
			return true
		}
	}
	return false
}
