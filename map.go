package godash

// Map returns a new slice from an exsting slice applying a function to each element.
func Map[T any, R any](values []T, mapper func(T) R) []R {
	mapped := make([]R, 0)
	for _, v := range values {
		mapped = append(mapped, mapper(v))
	}
	return mapped
}
