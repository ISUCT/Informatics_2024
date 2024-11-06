package labs_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/labs"
)

func TestEquation(t *testing.T) {
	var a = 0.05
	var b = 0.06
	tests := []struct {
		name string
		x    float64
		want float64
	}{
		{"случай 1", 0.20, 40.91},
		{"случай 2", 0.35, 12.07},
		{"случай 3", 0.50, 5.29},
		{"случай 4", 0.65, 2.63},
		{"случай 5", 0.80, 1.27},
		{"Граничный случай 1", 0.15, 77.59},
		{"Граничный случай 2", 0.26, 23.13},
		{"Граничный случай 3", 0.37, 10.66},
		{"Граничный случай 4", 0.48, 5.84},
		{"Граничный случай 5", 0.56, 3.97},
		{"Граничный случай: NaN", math.NaN(), 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := labs.CalculateY(tt.x, a, b)
			if math.IsNaN(tt.x) {
				assert.True(t, math.IsNaN(result), "ожидалось NaN %f", result)
			} else {
				assert.InDelta(t, tt.want, result, 0.01, "Ожидалось %f, но получено %f", tt.want, result)
			}
		})
	}
}
