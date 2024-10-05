package function_test

import (
	"testing"

	"isuct.ru/informatics2022/function"
)

func TestFunction(t *testing.T) {
	test_y := function.Uravnenie(1, 1)
	if test_y != 1.826821810431806 {
		t.Fatalf(`Фунция с b = 1, x = 1 выдаёт неверное значение %f не 1.826821810431806`, test_y)
	}
}