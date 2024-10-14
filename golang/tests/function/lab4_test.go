package function_test

import (
	"math"
	"testing"

	"isuct.ru/informatics2022/function"
)

type Test struct {
	x, a, b float64
	out     float64
}

var tests = []Test{
	{0.33, 0.06, 0.05, 1.527163095495038},
	{1.05, 0.06, 0.05, math.NaN()},
	{0.95, 0.06, 0.05, 1.0078412577756953},
	{0.24, 0.05, 0.07, 1.633123935319537e+16},
}

func TestLab4(t *testing.T) {
	for _, test := range tests {
		got := function.RunLab4Task(test.x, test.a, test.b)
		if math.IsNaN(test.out) {
			if !math.IsNaN(got) {
				t.Errorf("F(%f, %f, %f) = %f, want NaN", test.x, test.a, test.b, got)
			}
		} else {
			if got != test.out {
				t.Errorf("F(%f, %f, %f) = %f, want %f", test.x, test.a, test.b, got, test.out)
			}
		}
	}
}
