// Вариант 5: Первое уравнение.
// Функция для вычисления значения
func CalculateY(a, elements float64) float64 {
  numerator := math.Pow(math.Log10(math.Pow(a, 2)+elements), 2)
  denominator := (a + elements) * (a + elements)
  y := numerator / denominator
  return y
}
func TaskA(a, xn, xk, xdel float64) []float64 {
  var res []float64
  for i := xn; i < xk; i += xdel {
    res = append(res, CalculateY(a, i))
}
  return res
}
func TaskB(a float64, x []float64) []float64 {
  var res []float64
  for _, i := range x {
    res = append(res, CalculateY(a, i))
  return res
}
