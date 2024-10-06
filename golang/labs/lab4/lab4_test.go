package lab4_test

import (
	"math"
	"testing"

	"isuct.ru/informatics2022/labs/lab4"
)

type Test struct {
	in  float64
	out float64
}

var tests = []Test{
	{0.33, 4.6487441844155634e-05},
	{1.0499999999999998, 0.0012614954453531414},
	{1.24, 0.002464513104951157},
	{3.21, math.NaN()},
}

func TestLab4(t *testing.T) {
	for _, test := range tests {
		got := lab4.F(test.in)
		if math.IsNaN(test.out) {
			if !math.IsNaN(got) {
				t.Errorf("F(%f) = %f, want NaN", test.in, got)
			}
		} else {
			if got != test.out {
				t.Errorf("F(%f) = %f, want %f", test.in, got, test.out)
			}
		}
	}
}
