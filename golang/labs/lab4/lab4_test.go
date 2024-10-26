package lab4

import (
	"math"
	"testing"
)

// Тест для функции calculate_y
func TestCalculateY(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{0.5, math.Acos(0.5)},
		{1.5, math.Pow(1.2, 1.5) - math.Pow(1.5, 1.2)},
		{-0.9, math.Acos(-0.9)},
		{-1.5, math.Pow(1.2, -1.5) - math.Pow(-1.5, 1.2)},
	}

	for _, tt := range tests {
		result := calculate_y(tt.input)

		if math.IsNaN(tt.expected) {
			if !math.IsNaN(result) {
				t.Errorf("calculate_y(%v) = %v, expected NaN", tt.input, result)
			}
		} else if result != tt.expected {
			t.Errorf("calculate_y(%v) = %v, expected %v", tt.input, result, tt.expected)
		}
	}
}

// Тест для функции Task_A
func TestTask_A(t *testing.T) {
	epsilon := 1e-9
	result := Task_A(0.2, 2.2, 0.4)
	expected := []float64{
		calculate_y(0.2),
		calculate_y(0.6),
		calculate_y(1.0),
		calculate_y(1.4),
		calculate_y(1.8),
		calculate_y(2.2)}

	for i, value := range result {
		if math.Abs(value-expected[i]) > epsilon {
			t.Errorf("Task_A(%v) = %v, expected %v", i, value, expected[i])
		}
	}
}

// Тест для функции Task_B
func TestTask_B(t *testing.T) {
	epsilon := 1e-9
	arr := []float64{0.1, 0.9, 1.2, 1.5, 2.3}
	result := Task_B(arr)
	expected := []float64{
		calculate_y(0.1),
		calculate_y(0.9),
		calculate_y(1.2),
		calculate_y(1.5),
		calculate_y(2.3),
	}

	for i, value := range result {
		if math.Abs(value-expected[i]) > epsilon {
			t.Errorf("Task_B(%v) = %v, expected %v", i, value, expected[i])
		}
	}
}
