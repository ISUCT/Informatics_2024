package laba4_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"isuct.ru/informatics2022/laba4"
)

func TestFunction(t *testing.T) {
	type Test struct {
		name string
		in   float64
		out  float64
	}

	var result float64
	var tests []Test = []Test{
		{
			name: "проверка с обычным числом",
			in:   1.15,
			out:  0.6097818153301673,
		},
		{
			name: "проверка с большим значением после запятой",
			in:   2.2899999999999996,
			out:  0.042702414782365784,
		},
		{
			name: "проверка с NaN результатом",
			in:   1.93,
			out:  math.NaN(),
		},
		{
			name: "проверка с большим значением после запятой и NaN результатом",
			in:   1.5299999999999998,
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

func TestTaskA(t *testing.T) {
	type Test struct {
		name string
		in   []float64
		out  []float64
	}

	var result []float64
	var test Test = Test{
		name: "проверка задачи А",
		in:   []float64{1.15, 3.05, 0.38},
		out:  []float64{0.6097818153301673, math.NaN(), math.NaN(), 0.042702414782365784, 0.3282127650567373, math.NaN()},
	}

	t.Run(test.name, func(t *testing.T) {
		result = laba4.CompleteTaskA(2.5, 4.6, test.in[0], test.in[1], test.in[2])

		for i := 0; i < len(result); i++ {
			if math.IsNaN(test.out[i]) {
				assert.True(t, math.IsNaN(result[i]))
			} else {
				assert.Equal(
					t,
					result[i],
					test.out[i],
				)
			}
		}

	})
}

func TestTaskB(t *testing.T) {
	type Test struct {
		name string
		in   []float64
		out  []float64
	}

	var result []float64
	var test Test = Test{
		name: "проверка задачи B",
		in:   []float64{1.2, 1.36, 1.57, 1.93, 2.25},
		out:  []float64{0.588690605680308, 0.11819766614205533, math.NaN(), math.NaN(), 0.008019238155234967},
	}

	t.Run(test.name, func(t *testing.T) {
		result = laba4.CompleteTaskB(2.5, 4.6, test.in)

		for i := 0; i < len(result); i++ {
			if math.IsNaN(test.out[i]) {
				assert.True(t, math.IsNaN(result[i]))
			} else {
				assert.Equal(
					t,
					result[i],
					test.out[i],
				)
			}
		}

	})
}
