package laba4_test

import (
	"math"
	"testing"

	"isuct.ru/informatics2022/laba4"
)

type Test struct {
	in  float64
	out float64
}

func TestColcul(t *testing.T) {
	var tests []Test = []Test{
		{1.15, 0.6097818153301673},
		{1.5299999999999998, math.NaN()},
		{1.9099999999999997, math.NaN()},
		{2.2899999999999996, 0.042702414782365784},
		{2.6699999999999995, 0.3282127650567373},
		{3.0499999999999994, math.NaN()},

		{1.2, 0.588690605680308},
		{1.36, 0.11819766614205533},
		{1.57, math.NaN()},
		{1.93, math.NaN()},
		{2.25, 0.008019238155234967},
	}

	for _, test := range tests {
		var result float64 = laba4.Function(2.5, 4.6, test.in)

		if math.IsNaN(test.out) {
			if !math.IsNaN(result) {
				t.Errorf("Function(2.5, 4.6, %v) = %v, want %v, error", test.in, result, test.out)
			}
		} else {
			if result != test.out {
				t.Errorf("Function(2.5, 4.6, %v) = %v, want %v, error", test.in, result, test.out)
			}
		}
	}
}
