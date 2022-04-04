package godash

// Filter returns a slice with only the elements which match the predicate.
func Filter[T any](values []T, predicate func(T) bool) []T {
	filtered := make([]T, 0)
	for _, v := range values {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
