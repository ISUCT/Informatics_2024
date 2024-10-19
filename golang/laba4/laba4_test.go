package laba4_test

import (
	"math"
	"testing"

	"isuct.ru/informatics2022/laba4"
)

type Test struct {
	name string
	in   float64
	out  float64
}

func TestFunction(t *testing.T) {
	var tests []Test = []Test{
		{
			name: "проверка задачи А",
			in:   1.15,
			out:  0.6097818153301673,
		},
		{
			name: "проверка задачи А с NaN результатом",
			in:   1.5299999999999998,
			out:  math.NaN(),
		},
		{
			name: "проверка задачи А с большим значением после запятой",
			in:   2.2899999999999996,
			out:  0.042702414782365784,
		},
		{
			name: "проверка задачи B",
			in:   1.2,
			out:  0.588690605680308,
		},
		{
			name: "проверка задачи B с большим значением после запятой",
			in:   1.93,
			out:  math.NaN(),
		},
	}

	for _, test := range tests {
		var result float64 = laba4.СalculateFunction(2.5, 4.6, test.in)

		if (math.IsNaN(test.out)) && (!math.IsNaN(result)) {
			t.Errorf("Function(2.5, 4.6, %v) = %v, want %v, error", test.in, result, test.out)
		}
		if (!math.IsNaN(test.out)) && (result != test.out) {
			t.Errorf("Function(2.5, 4.6, %v) = %v, want %v, error", test.in, result, test.out)
		}
	}
}
