package godash

// Reduce uses an intial value, apply a binary function to the running value and the next element of the given slice.
func Reduce[T any, R any](values []T, reducer func(R, T) R, accumulator R) R {
	for _, v := range values {
		accumulator = reducer(accumulator, v)
	}
	return accumulator
}
