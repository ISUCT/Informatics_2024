package labs

import "fmt"

type Sydno struct {
	width, height, tosadka, polnota float64
}

func (t Sydno) Size() (float64, float64, float64, float64) {
	return t.width, t.height, t.tosadka, t.polnota
}

func (t Sydno) SetSize(width, height, tosadka, polnota float64) {
	t.width = width
	t.height = height
	t.tosadka = tosadka
	t.polnota = polnota
}

func (t Sydno) Calculate_Area() float64 {
	return t.height * t.width
}

//водоизмещение
func (t Sydno) Water() float64 {
	return t.height * t.width * t.polnota
}

func PrintSydnoInfo(sydno Sydno) {
	width, height, tosadka, polnota := sydno.Size()
	fmt.Printf("Размеры судна: Ширина = %.2f, Высота = %.2f, Осадка судна = %.2f, Коэффициент полноты = %.2f\n", width, height, tosadka, polnota)
	fmt.Printf("Площадь судна: %.2f м²\n", sydno.Calculate_Area())
	fmt.Printf("Водоизмещение судна: %.2f м³\n", sydno.Water())
}
func FinishSudno() {
	sydno := Sydno{}
	sydno.SetSize(1.25, 3.4, 2.3, 5.65)
	PrintSydnoInfo(sydno)
}
