package laba4_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"isuct.ru/informatics2022/laba4"
)

type Test struct {
	name string
	in   float64
	out  float64
}

func TestFunction(t *testing.T) {
	var result float64
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
			name: "проверка задачи B с NaN",
			in:   1.93,
			out:  math.NaN(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result = laba4.CalculateFunction(2.5, 4.6, test.in)

			if math.IsNaN(test.out) {
				assert.True(t, math.IsNaN(result))
			} else {
				assert.Equal(
					t,
					result,
					test.out,
				)
			}
		})
	}
}
