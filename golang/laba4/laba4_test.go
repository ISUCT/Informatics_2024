package laba4_test

import (
	"testing"

	"isuct.ru/informatics2022/laba4"
)

func TestColcul(t *testing.T) {
	var result float64 = laba4.Colcul(1, 1, 1)
	if result != 0.36670521511960164 {
		t.Fatalf("Colcul(2.5, 4.6, 1.2) = %f, want 0.36670521511960164, error", result)
	}
}
