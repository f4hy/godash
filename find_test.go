package godash

import (
	"testing"
)

func TestFindInts(t *testing.T) {
	values := []int{1, 10, 2, 20, 3, 30, 4, 4, 6}
	GreaterThan5 := func(i int) bool { return i > 5 }
	found := Find(values, GreaterThan5)
	if *found != 10 {
		t.Fatalf("Find([]int{1,10,2,20,3,30,4,4,6}, GreaterThan5) failed to find 10 found %d", *found)
	}
}

func TestFindStrings(t *testing.T) {
	values := []string{"ok", "Hello", "World", "!"}
	len5 := func(s string) bool { return len(s) == 5 }
	found := Find(values, len5)
	if *found != "Hello" {
		t.Fatalf("Find{'ok', 'Hello', 'World', '!'} failed to find 'Hello' found %s", *found)
	}
}
