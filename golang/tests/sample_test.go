package internal_test

import (
	"testing"

	"isuct.ru/informatics2022/lab"
)

// func TestTaskA(t *testing.T) {
// 	values := lab.TaskA(2.5, 1.28, 2.08, 0.4)
// 	answer := []float64{0.6965184026921305, 0.7313025427889659, 0.42530098762123236}
// 	for index, value := range values {
// 		if value != answer[index] {
// 			t.Fatalf(`TaskA(2.5, 1.28, 2.08, 0.4) = %e error`, value)
// 		}
// 	}
// }

// func TestTaskB(t *testing.T) {
// 	slice := [5]float64{1.1, 2.4, 3.6, 1.7, 3.9}
// 	values := lab.TaskB(2.5, slice)
// 	answer := []float64{0.7393337604443017, 0.599437811110944, 0.32139427020687567, 0.7252847867828259, 0.2828468929960873}
// 	for index, value := range values {
// 		if value != answer[index] {
// 			t.Fatalf(`TaskB(2.5, slice) = %e error`, value)
// 		}
// 	}
// }

//Table тест
func TestTaskA(t *testing.T) {
	var values = []struct {
		text  string
		b float64
		xn float64
		xk float64
		delx float64
		answers []float64
	}{
		{"Верно", 2.5, 1.28, 2.08, 0.4, []float64 {0.6965184026921305, 0.7313025427889659, 0.42530098762123236}},
		{"Верно", 2.5, 2.48, 3.28, 0.4, []float64{0.4024271469763357, 0.5702037834473751, 0.3678671257606618}},
	}
	for _, value := range values {
		t.Run(value.text, func(t *testing.T) {
			results := lab.TaskA(value.b, value.xn, value.xk, value.delx)
			for index, result := range results {
				if result != value.answers[index] {
					t.Errorf(`TaskA(2.5, 1.28, 2.08, 0.4) = %e error`, result)
				}
			}
		})
	}
}

func TestTaskB(t *testing.T) {
	var values = []struct {
		text  string
		b float64
		slice []float64
		answers []float64
	}{
		{"Верно", 2.5, []float64{1.1, 2.4}, []float64{0.7393337604443017, 0.599437811110944}},
		{"Верно", 2.5, []float64{3.6, 1.7, 3.9}, []float64{0.32139427020687567, 0.7252847867828259, 0.2828468929960873}},
	}
	for _, value := range values {
		t.Run(value.text, func(t *testing.T) {
			results := lab.TaskB(value.b, value.slice)
			for index, result := range results {
				if result != value.answers[index] {
					t.Errorf(`TaskB(2.5, slice) = %e error`, result)
				}
			}
		})
	}
}