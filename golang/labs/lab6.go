package labs

import (
	"fmt"
)

type Sydno struct {
	width   float64
	height  float64
	tosadka float64
	polnota float64
}

// Метод для получения размеров судна
func (t *Sydno) Size() (float64, float64, float64, float64) {
	return t.width, t.height, t.tosadka, t.polnota
}

// Метод для установки размеров судна
func (t *Sydno) SetSize(width, height, tosadka, polnota float64) {
	t.width = width
	t.height = height
	t.tosadka = tosadka
	t.polnota = polnota
}

// Метод для вычисления площади судна
func (t *Sydno) Area() float64 {
	return t.height * t.width
}

// Метод для вычисления водоизмещения
func (t *Sydno) Water() float64 {
	return t.height * t.width * t.polnota
}

// Функция для завершения работы с судном
func FinishSudno() {
	sydno := Sydno{}

	// Устанавливаем размеры судна
	sydno.SetSize(1.25, 3.4, 2.3, 5.65)

	// Получаем размеры судна
	width, height, tosadka, polnota := sydno.Size()
	fmt.Printf("Размеры судна: Ширина = %.2f, Высота = %.2f, Осадка судна = %.2f, Коэффициент полноты = %.2f\n", width, height, tosadka, polnota)

	// Вычисляем и выводим площадь судна
	area := sydno.Area()
	fmt.Printf("Площадь судна: %.2f м²\n", area)

	// Вычисляем и выводим водоизмещение
	water := sydno.Water()
	fmt.Printf("Водоизмещение судна: %.2f м³\n", water)
}
