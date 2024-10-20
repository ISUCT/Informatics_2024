package lab4

import ( 
 "math"
)

func Colcul (a, b, x float64) float64 {
  if x == 0 {
    return 0
  }
 var y = math.Sqrt(float64(math.Abs(a-b*x)) / math.Pow(math.Log10(float64(x)), 3))

 return y;
}