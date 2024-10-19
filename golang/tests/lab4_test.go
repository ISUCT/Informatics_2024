package labs_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/labs"
)

var a float64 = 0.05
var b float64 = 0.06

func TestEquation(t *testing.T) {

	tests := []struct {
		x    float64
		want float64
	}{
		{0.20, 40.91},
		{0.35, 12.07},
		{0.50, 5.29},
		{0.65, 2.63},
		{0.80, 1.27},
		{0.15, 77.59},
		{0.26, 23.13},
		{0.37, 10.66},
		{0.48, 5.84},
		{0.56, 3.97},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Euation(%f)", tt.x), func(t *testing.T) {
			result := labs.Calculate_y(tt.x, a, b)
			assert.InDelta(t, tt.want, result, 0.01, "Expected %f but got %f", tt.want, result)
		})
	}
}
