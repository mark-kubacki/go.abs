package abs // import "plugin.hosting/go/abs"

import (
	"fmt"
	"math"
	"testing"
)

func ExampleAbs64() {
	fmt.Printf("%d", Abs64(-1)) // 😃
	// Output: 1
}

var fromTo64 = []struct {
	given    int64
	expected uint64
}{
	{0, 0},
	{-0, 0},
	{1, 1},
	{-1, 1},
	{math.MaxInt64, math.MaxInt64},
	{-math.MaxInt64, math.MaxInt64},
	{-math.MaxInt64 - 1, math.MaxInt64 + 1},
}

func TestAbs64Inductive(t *testing.T) {
	for _, pair := range fromTo64 {
		got := Abs64(pair.given)
		if got != pair.expected {
			t.Errorf("Abs64(%v) = %v, expected %v", pair.given, got, pair.expected)
		}
	}
}

var fromTo32 = []struct {
	given    int32
	expected uint32
}{
	{0, 0},
	{-0, 0},
	{1, 1},
	{-1, 1},
	{math.MaxInt32, math.MaxInt32},
	{-math.MaxInt32, math.MaxInt32},
	{-math.MaxInt32 - 1, math.MaxInt32 + 1},
}

func TestAbs32Inductive(t *testing.T) {
	for _, pair := range fromTo32 {
		got := Abs32(pair.given)
		if got != pair.expected {
			t.Errorf("Abs32(%v) = %v, expected %v", pair.given, got, pair.expected)
		}
	}
}
