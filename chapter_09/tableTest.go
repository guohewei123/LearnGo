package main

import (
	"fmt"
	"math"
)

func add(a, b int32) int32 {
	return a + b
}

func main() {

	tests := []struct {
		a, b, c int32
	}{
		{1, 2, 3},
		{0, 2, 2},
		{0, 0, 0},
		{-1, 1, 0},
		{math.MaxInt32, 1, math.MinInt32},
	}

	for _, test := range tests {
		if actual := add(test.a, test.b); actual != test.c {
			fmt.Printf("%d + %d != %d\n", test.a, test.b, test.c)
		}
	}
}
