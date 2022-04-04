package godash

func Reduce[T any, R any](values []T, reducer func(R, T) R, accumulator R) R {
	for _, v := range values {
		accumulator = reducer(accumulator, v)
	}
	return accumulator
}
