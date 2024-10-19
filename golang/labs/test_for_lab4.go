package labs

import (
	"math"
	"testing"
)

func testing_outputY(t *testing.T) {
	const const_a = 2
	const const_b = 3
	result_y := Return_value_y(0.11, const_a, const_b)
	true_res := 1.581565621572023
	if math.IsNaN(result_y) {
		t.Errorf("The value is a not a namber!")
	} else if result_y != true_res {
		t.Errorf("Returns %v, expected %v", result_y, true_res)
	}

}
func testing_taskA(t *testing.T) {
	y_list_true := []float64{1.581565621572023,
		1.5923031114990622,
		1.6056494993987385,
		1.620871010842551,
		1.6372494546297607,
		1.6540889552598357}
	y_list_output := Task_A(2, 3)
	if len(y_list_output) != 6 {
		t.Errorf("Returns %v y, expected %v y", len(y_list_output), len(y_list_true))
	}
	for i, y := range y_list_output {
		if y != y_list_true[i] {
			t.Errorf("The function was calculated incorrectly")
		}
	}
}
func testing_taskB(t *testing.T) {
	y_list_true := []float64{1.5766843704639988,
		1.6208710138786773,
		1.650716645162541,
		1.6707225698630865,
		1.7060955449869923}
	y_list_output := Task_B(2, 3)
	for i, y := range y_list_output {
		if y != y_list_true[i] {
			t.Errorf("The function was calculated incorrectly")
		}
	}
}
