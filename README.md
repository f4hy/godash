[![Go Reference](https://pkg.go.dev/badge/github.com/f4hy/godash.svg)](https://pkg.go.dev/github.com/f4hy/godash)
# godash
Functional programming methods for golang inspired by [lodash](https://lodash.com/). Now possible using generics from go 1.18


Usage

```go
package main

import (
	"fmt"
	__ "github.com/f4hy/godash"
)

func main() {
	example := []int{1, 2, 4, 3}
	mapped := __.Map(example, func(i int) int { return i + 1 })
	fmt.Printf("mapped: %v \n", mapped)
	// [2 3 5 4]
	greaterThan2 := func(i int) bool { return i > 2 }
	filtered := __.Filter(example, greaterThan2)
	fmt.Printf("filtered: %v \n", filtered)
	// [4 3]
	alllargerthan2 := __.All(example, greaterThan2)
	anylargerthan2 := __.Any(example, greaterThan2)
	fmt.Printf("all larger than 2: %t \n", alllargerthan2) // false
	fmt.Printf("any larger than 2: %t \n", anylargerthan2) // true

	sum := __.Reduce(example, func(acc int, i int) int { return acc + i }, 0)
	fmt.Printf("sum: %d \n", sum)
	// 10
}
```

Fucntions
```
func All[T any](values []T, predicate func(T) bool) bool
func Any[T any](values []T, predicate func(T) bool) bool
func Filter[T any](values []T, predicate func(T) bool) []T
func Find[T any](values []T, predicate func(T) bool) *T
func ForEach[T any](values []T, fun func(T))
func Map[T any, R any](values []T, mapper func(T) R) []R
func Reduce[T any, R any](values []T, reducer func(R, T) R, accumulator R) R
```

