package lab4_test

import (
	"math"
	"testing"

	"isuct.ru/informatics2022/lab4"
)

func TestFunction(t *testing.T) {
	result := lab4.CalculateFunction(1.28, 2.5)
	expected := 0.6965184026921305
	if math.IsNaN(result) {
		t.Errorf("Returns the value NaN")
	} else if result != expected {
		t.Errorf("Returns %v, wanted to %v", result, expected)
	}
}
func TestTaskA(t *testing.T) {
	resultA := lab4.TaskA(1.28, 3.28, 0.4)
	if len(resultA) != 6 {
		t.Errorf("problems with the function step")
	}
}

func TestTaskB(t *testing.T) {
	resultB := lab4.TaskB(1.1, 2.4, 3.6, 1.7, 3.9)
	yb := []float64{0.7393337604443017, 0.599437811110944, 0.32139427020687567, 0.7252847867828259, 0.2828468929960873}
	for i, v := range resultB {
		if v != yb[i] {
			t.Errorf("The function was calculated incorrectly")
		}
	}
}
