package godash

func ForEach[T any](values []T, fun func(T)) {
	for _, v := range values {
		fun(v)
	}
}
