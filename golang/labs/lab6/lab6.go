package labs

import "fmt"

type Table struct {
	length, width, height float64
}

func NewTable(length, width, height float64) Table {
	return Table{
		length: length,
		width:  width,
		height: height,
	}
}

func (t *Table) SetDimensions(length, width, height float64) {
	t.length = length
	t.width = width
	t.height = height
}

func (t *Table) GetLength() float64 {
	return t.length
}

func (t *Table) GetWidth() float64 {
	return t.width
}

func (t *Table) GetHeight() float64 {
	return t.height
}

func RunLab6Task() {
	myTable := NewTable(120, 60, 75)

	fmt.Printf("Начальные размеры стола: Длина = %.2f, Ширина = %.2f, Высота = %.2f\n", myTable.GetLength(), myTable.GetWidth(), myTable.GetHeight())

	myTable.SetDimensions(150, 70, 80)

	fmt.Printf("Обновленные размеры стола: Длина = %.2f, Ширина = %.2f, Высота = %.2f\n", myTable.GetLength(), myTable.GetWidth(), myTable.GetHeight())
}
