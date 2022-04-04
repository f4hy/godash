package godash

import (
	"fmt"
	"testing"
)

func TestForEachSum(t *testing.T) {
	sum := 0
	summingClosure := func(i int) { sum = sum + i }
	ForEach([]int{10, 200, 7}, summingClosure)
	if sum != 217 {
		t.Fatalf("ForEach summing ints closure failed got %v expected %v", sum, 217)
	}
}

func ExampleForEachPrint() {
	printer := func(s string) { fmt.Println(s) }
	ForEach([]string{"hello", "world"}, printer)
	// Output:
	// hello
	// world
}
