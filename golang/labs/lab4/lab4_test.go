package lab4_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/labs/lab4"
)

var a float64 = 0.1
var b float64 = 0.5

type Test struct {
	name string
	a, b float64
	in   float64
	out  float64
}

var tests = []Test{
	{"NaN answer", a, b, 3.21, math.NaN()},
	{"0.33 test", a, b, 0.33, 4.6487441844155634e-05},
}

func TestLab4(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := lab4.Calculate(test.a, test.b, test.in)
			if math.IsNaN(test.out) {
				assert.True(t, math.IsNaN(got), "Expected NaN but got %f", got)
			} else {
				assert.Equal(t, test.out, got, "For input %f, expected %f but got %f", test.in, test.out, got)
			}
		})
	}
}
