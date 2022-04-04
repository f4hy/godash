// Functional Tools for Golang useing generics inspired by lodash
package godash

// All returns true if all elements match the predicate
func All[T any](values []T, predicate func(T) bool) bool {
	for _, v := range values {
		if !predicate(v) {
			return false
		}
	}
	return true
}
