package function

import (
	"fmt"
	"math"
	"testing"
)

var tests = []struct {
	x, a, b float64
	out     float64
}{
	{0.33, 0.06, 0.05, 1.527163},
	{0.95, 0.06, 0.05, 0.403693},
	{0.24, 0.05, 0.07, 1.007841},
}

func TestCalculateYWithNaN(t *testing.T) {
	for _, test := range tests {
		if math.IsNaN(test.out) {
			got := CalculateY(test.x, test.a, test.b)
			if !math.IsNaN(got) {
				t.Errorf("Test failed for NaN case x=%f, a=%f, b=%f: got %f, want NaN", test.x, test.a, test.b, got)
			}
		}
	}
}

func TestCalculateYWithoutNaN(t *testing.T) {
	for _, test := range tests {
		if !math.IsNaN(test.out) {
			got := CalculateY(test.x, test.a, test.b)
			if math.Abs(got-test.out) > 1e-9 { // small tolerance for floating point comparison
				t.Errorf("Test failed for x=%f, a=%f, b=%f: got %f, want %f", test.x, test.a, test.b, got, test.out)
			}
