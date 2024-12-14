package lab4

import "math"

a float64 = 1.6
func Lab4(a, x float64) float64 {
	return (math.Pow(a, math.Pow(x, 2)-1) - math.Log10(math.Pow(x, 2)-1) + math.Cbrt(math.Pow(x, 2)-1))
}
func lab4A(imin, imax, razn) {
	for imin; imin < imax; imin += razn {
		fmt.Println(lab4.Lab4(a, i))
	}
}
func lab4B() {
	var x [5]float64 = [5]float64{1.28, 1.36, 2.47, 3.68, 4.56}
		for index , value := range x {
			fmt.Println(lab4.Lab4(a, x[i]))
		}
}
func Lab4AB(){
	fmt.Println(lab4.lab4A(1.2, 3.7, 0.5))
	fmt.Println(lab4.lab4B())
}
