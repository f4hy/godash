package godash

import (
	"reflect"
	"testing"
)

func TestReduceSum(t *testing.T) {
	reducer := func(acc int, i int) int { return acc + i }
	sum := Reduce([]int{10, 200, 7}, reducer, 0)
	if sum != 217 {
		t.Fatalf("Reduce summing ints failed got %v expected %v", sum, 217)
	}
}

func TestReduceConcat(t *testing.T) {
	reducer := func(acc string, i string) string { return acc + i }
	concatted := Reduce([]string{"hello", " ", "world", "!"}, reducer, "start:")
	if concatted != "start:hello world!" {
		t.Fatalf("Reduce concat string failed got %v expected %v", concatted, "hello world!")
	}
}

func TestReduceAppendDouble(t *testing.T) {
	reducer := func(acc []int, i int) []int { return append(acc, i*2) }
	doubled := Reduce([]int{1, 5, 3}, reducer, []int{})
	expected := []int{2, 10, 6}
	if !reflect.DeepEqual(doubled, expected) {
		t.Fatalf("Reduce as doubling map failed got %v expected %v", doubled, expected)
	}
}
