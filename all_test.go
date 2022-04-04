package godash

import (
	"testing"
)

func TestAll(t *testing.T) {
	GreaterThan5 := func(i int) bool { return i > 5 }
	if All([]int{1, 10, 2, 20}, GreaterThan5) != false {
		t.Fatalf("All({1,10,2,20}, GreaterThan5) should be false")
	}
	if All([]int{100, 20, 111, 10}, GreaterThan5) != true {
		t.Fatalf("All({1,2,1,0}, GreaterThan5) should be True")
	}
	LengthUnder5 := func(s string) bool { return len(s) < 5 }
	if All([]string{"one", "two"}, LengthUnder5) != true {
		t.Fatalf("All({'one','two'}, LengthUnder5) should be True")
	}
	if All([]string{"one", "two", "three"}, LengthUnder5) != false {
		t.Fatalf("All({'one','two','three'}, LengthUnder5) should be False")
	}
}
