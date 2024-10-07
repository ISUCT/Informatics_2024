package labs_test

import (
  "math"
  "testing"

  "isuct.ru/informatics2022/labs"
)

func TestF(t *testing.T) {
  a := 0.05
  b := 0.06
  testValues := []float64{0.2, 0.2345, 1.23, 0.132}

  for _, x := range testValues {
    result := labs.F(x, a, b)

    resulta := labs.F(math.Acos(math.Pow(x, 2)-math.Pow(b, 2)), a, b)
    resultb := labs.F(math.Asin(math.Pow(x, 2)-math.Pow(a, 2)), a, b)

    if result < 0 {
      t.Fatalf(`F(%f, %f, %f) = %f; ожидается положительное значение`, x, a, b, result)
    }

    if resultb != 0 && resulta/resultb == 0 {
      t.Fatalf(`F(%f, %f, %f) = %f; ожидается NaN`, x, a, b, result)
    }
  }

}
