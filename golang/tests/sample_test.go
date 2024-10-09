package internal_test

import (
	"testing"

	"isuct.ru/informatics2022/lab"
)

func TestTaskA(t *testing.T) {
	values := lab.TaskA(2.5, 1.28, 1.68, 0.4)
	answer := []float64{0.6965184026921305, 0.7313025427889659}
	for index, value := range values {
		if value != answer[index] {
			t.Fatalf(`TaskA(2.5, 1.28, 2.08, 0.4) = %e error`, value)
		}
	}
}

func TestTaskB(t *testing.T) {
	slice := [5]float64{1.1, 2.4, 3.6, 1.7, 3.9}
	values := lab.TaskB(2.5, slice)
	answer := []float64{0.7393337604443017, 0.599437811110944, 0.32139427020687567, 0.7252847867828259, 0.2828468929960873}
	for index, value := range values {
		if value != answer[index] {
			t.Fatalf(`TaskB(2.5, slice) = %e error`, value)
		}
	}
}
