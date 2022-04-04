package godash

import (
	"reflect"
	"testing"
)

func TestFilterInts(t *testing.T) {
	values := []int{1, 10, 2, 20, 3, 30, 4, 4, 6}
	GreaterThan5 := func(i int) bool { return i > 5 }
	filtered := Filter(values, GreaterThan5)
	expected := []int{10, 20, 30, 6}
	if !reflect.DeepEqual(filtered, expected) {
		t.Fatalf("Filter([]int{1,10,2,20,3,30,4,4,6}, GreaterThan5) failed")
	}
}

func TestFilterStrings(t *testing.T) {
	values := []string{"Hello", "World", "!"}
	len5 := func(s string) bool { return len(s) == 5 }
	filtered := Filter(values, len5)
	expected := []string{"Hello", "World"}
	if !reflect.DeepEqual(filtered, expected) {
		t.Fatalf("Filter Strings failed. got %s, expected %s", filtered, expected)
	}
}
