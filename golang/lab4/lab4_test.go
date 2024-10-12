package lab4_test

import (
	"testing"

	"isuct.ru/informatics2022/lab4"
)

func TestFunction(t *testing.T) {
	result := lab4.Function(1.28)
	if result != 0.69 {
		t.Fatalf(`Function(1.28) = %f, want 0.69, error`, result)
	}
}
