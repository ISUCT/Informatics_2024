package internal_test

import (
	"testing"

	"isuct.ru/informatics2022/lab"
)

func TestTaskA(t *testing.T) {
	var values = []struct {
		text    string
		b       float64
		xn      float64
		xk      float64
		delx    float64
		answers []float64
	}{
		{"TaskA(2.5, 1.28, 2.08, 0.4)", 2.5, 1.28, 2.08, 0.4, []float64{0.6965184026921305, 0.7313025427889659, 0.42530098762123236}},
		{"TaskA(2.5, 2.48, 3.28, 0.4)", 2.5, 2.48, 3.28, 0.4, []float64{0.4024271469763357, 0.5702037834473751, 0.3678671257606618}},
		{"TaskA(2, 1.04, 2.04, 0.5)", 2, 1.04, 2.04, 0.5, []float64{0.520316235595818, 0.7177865337405192, 0.5878711701956065}},
		{"TaskA(4, 9.034, 9.039, 0.02)", 4, 9.034, 9.039, 0.02, []float64{0.1113898676606079, 0.1498764947240627, 0.19861893241865275}},
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
		text    string
		b       float64
		slice   []float64
		answers []float64
	}{
		{"TaskB(2.5, []float64{1.1, 2.4})", 2.5, []float64{1.1, 2.4}, []float64{0.7393337604443017, 0.599437811110944}},
		{"TaskB(2.5, []float64{3.6, 1.7, 3.9})", 2.5, []float64{3.6, 1.7, 3.9}, []float64{0.32139427020687567, 0.7252847867828259, 0.2828468929960873}},
		{"TaskB(1.321, []float64{4.5, 3.222, 5.00})", 1.321, []float64{4.5, 3.222, 5.00}, []float64{0.3376564479793544, 0.5664507473034334, 0.3965790442838673}},
		{"TaskB(15, []float64{0, 3.7})", 15, []float64{0, 3.7}, []float64{0.10947371185997536, 0.12849119680379323}},
	}
	for _, value := range values {
		t.Run(value.text, func(t *testing.T) {
			results := lab.TaskB(value.b, value.slice)
			for index, result := range results {
				if result != value.answers[index] {
					t.Errorf(`TaskB(%e, %e) = %e error`, value.b, value.slice[index], result)
				}
			}
		})
	}
}
