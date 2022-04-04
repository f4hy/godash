package godash

func Map[T any, R any](values []T, mapper func(T) R) []R {
	mapped := make([]R, 0)
	for _, v := range values {
		mapped = append(mapped, mapper(v))
	}
	return mapped
}
