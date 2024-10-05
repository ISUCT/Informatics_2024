package laba4_test

import (
	"testing"

	"isuct.ru/informatics2022/laba4"
)

func TestColcul(t *testing.T) {
	result := laba4.Colcul(2.5, 4.6, 1.2)
	if result != 0.588691 {
		t.Fatalf("Colcul(2.5, 4.6, 1.2) = %f, want 0.588691, error", result)
	}
}
