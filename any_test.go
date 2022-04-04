package godash

import (
	"testing"
)

func TestAny(t *testing.T) {
	GreaterThan5 := func(i int) bool { return i > 5 }
	if Any([]int{1, 10, 2}, GreaterThan5) != true {
		t.Fatalf("Any({1,10,2}, GreaterThan5) should be true")
	}
	if Any([]int{1, 4, 2, 3}, GreaterThan5) != false {
		t.Fatalf("Any({1,4,2,3}, GreaterThan5) should be True")
	}
	LengthUnder5 := func(s string) bool { return len(s) < 5 }
	if Any([]string{"three", "two"}, LengthUnder5) != true {
		t.Fatalf("Any({'three','two'}, LengthUnder5) should be True")
	}
	if Any([]string{"three", "hundred"}, LengthUnder5) != false {
		t.Fatalf("All({'three','hundred'}, LengthUnder5) should be False")
	}
}
