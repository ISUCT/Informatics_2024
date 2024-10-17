package labs_test

import (
	"math"
	"testing"

	"isuct.ru/informatics2022/labs"
)

type Test struct {
	in  float64
	out float64
}

var tests = []Test{
	{1.2, 95.65974212841763},
	{1.28, 105.92978633506368},
	{1.36, 116.8010604328461},
	{-1, math.NaN()},
}

func TestLab4(t *testing.T) {
	for _, test := range tests {
		got := labs.Calculate_y(test.in, 2.5, 4.6)
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
