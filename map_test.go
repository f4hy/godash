package godash

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapInts(t *testing.T) {
	PlusOne := func(i int) int { return i + 1 }
	mapped := Map([]int{1, 10, 2, 20, 3, 30, 4, 4, 6}, PlusOne)
	expectedP1 := []int{2, 11, 3, 21, 4, 31, 5, 5, 7}
	if !reflect.DeepEqual(mapped, expectedP1) {
		t.Fatalf("Map int->int failed got %v expected %v", mapped, expectedP1)
	}

	PlusOneString := func(i int) string { return fmt.Sprintf("%d + 1", i) }
	mappedToStr := Map([]int{1, 10, 8}, PlusOneString)
	expectedStrs := []string{"1 + 1", "10 + 1", "8 + 1"}
	if !reflect.DeepEqual(mappedToStr, expectedStrs) {
		t.Fatalf("Map int->str failed got %v expected %v", mappedToStr, expectedStrs)
	}
}
