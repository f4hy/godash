package godash

func All[T any](values []T, predicate func(T) bool) bool {
	for _, v := range values {
		if !predicate(v) {
			return false
		}
	}
	return true
}
