package godash

// Find returns the first element which matches the predicate.
func Find[T any](values []T, predicate func(T) bool) *T {
	for _, v := range values {
		if predicate(v) {
			return &v
		}
	}
	return nil
}
