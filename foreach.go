package godash

// ForEach calls a function on each elemnt. Returns nothing.
func ForEach[T any](values []T, fun func(T)) {
	for _, v := range values {
		fun(v)
	}
}
