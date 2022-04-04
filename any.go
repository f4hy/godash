package godash

// Any returns true if any of the elements match the predicate
func Any[T any](values []T, predicate func(T) bool) bool {
	for _, v := range values {
		if predicate(v) {
			return true
		}
	}
	return false
}
