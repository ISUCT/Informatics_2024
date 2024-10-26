package lab4_test

import (
	"math"
	"testing"

	"isuct.ru/informatics2022/lab4"
)

func TestFunction(t *testing.T) {
	tests := []struct {
		x, b, expected float64
	}{
		{1.28, 2.5, 0.6965184026921305},
		{1.3, 0.5, 1.1585607285400352},
		{-2.5, 2.5, math.NaN()},
	}
	for _, tt := range tests {
		if math.IsNaN(tt.expected) {
			result := lab4.CalculateFunction(tt.x, tt.b)
			if !math.IsInf(result, 1) {
				t.Errorf("Test failed for +Inf case x = %f, b = %f: result = %f, want +Inf", tt.x, tt.b, result)
			}
		} else {
			result := lab4.CalculateFunction(tt.x, tt.b)
			if result != tt.expected {
				t.Errorf("Test failed for x=%f, b=%f: received %f, want %f", tt.x, tt.b, result, tt.expected)
			}
		}
	}
}
