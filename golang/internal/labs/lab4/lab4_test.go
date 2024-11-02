package lab4

import (
	"math"
	"testing"
)

func TestTask_A(t *testing.T) {
	epsilon := 1e-9
	result := Task_A(0.22, 0.92, 0.14)
	expected := []float64{
		Calculate(0.22),
		Calculate(0.36),
		Calculate(0.50),
		Calculate(0.64),
		Calculate(0.78),
		Calculate(0.92),
	}

	for i, value := range result {
		if math.Abs(value-expected[i]) > epsilon {
			t.Errorf("Task_A(%v) = %v, expected %v", i, value, expected[i])
		}
	}
}

// Тест для функции Task_B
func TestTask_B(t *testing.T) {
	epsilon := 1e-9
	args := []float64{0.1, 0.35, 0.4, 0.55, 0.6}
	result := Task_B(args)
	expected := []float64{
		Calculate(args[0]),
		Calculate(args[1]),
		Calculate(args[2]),
		Calculate(args[3]),
		Calculate(args[4])}

	for i, value := range result {
		if math.Abs(value-expected[i]) > epsilon {
			t.Errorf("Task_A(%v) = %v, expected %v", i, value, expected[i])
		}
	}
}
