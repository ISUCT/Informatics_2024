package lab6

import (
	"fmt"
)

type Sydno struct {
	width, height, osadka, polnota float64
}

func (t Sydno) Size() (float64, float64, float64, float64) {
	return t.width, t.height, t.osadka, t.polnota
}

func (t Sydno) CalculateArea() float64 {
	return t.height * t.width
}

func (t Sydno) Displacement() float64 {
	return t.height * t.width * t.polnota
}

func RunLab6() {
	sydno := Sydno{2.45, 5.23, 7.34, 1.65}
	width, height, osadka, polnota := sydno.Size()
	fmt.Printf("Размеры судна: Ширина = %.2f, Высота = %.2f, Осадка судна = %.2f, Коэффициент полноты = %.2f\n", width, height, osadka, polnota)
	fmt.Printf("Площадь судна: %.2f м²\n", sydno.CalculateArea())
	fmt.Printf("Водоизмещение судна: %.2f м³\n", sydno.Displacement())
}
