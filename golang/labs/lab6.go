package labs

import (
	"fmt"
)

// Реализовать следующие структуры (1 класс для 1 варианта)
// Определить минимум 3 переменных класса и 3 метода для класса10. Стол (получить  установить размеры стола)
type Table struct {
	width  float64
	height float64
	length float64
}

func (t *Table) GetSize() (float64, float64, float64) {
	return t.width, t.length, t.height
}

func (t *Table) SetSize(width, length, height float64) {
	t.width = width
	t.length = length
	t.height = height
}

func (t *Table) CalculateArea() float64 {
	return t.length * t.width
}

func RunLab6() {
	myTable := Table{}
	myTable.SetSize(1.28, 2.34, 0.76)
	width, length, height := myTable.GetSize()
	fmt.Printf("Размеры стола: Ширина = %.2f, Длина = %.2f, Высота = %.2f\n", width, length, height)
	stable := myTable.CalculateArea()
	fmt.Printf("Площадь стола: %.2f м²\n", stable)
}
