package lab6

import "fmt"

// Person структура описывает человека.
type Person struct {
	Name   string
	Age    int
	Height float64
}

// NewPerson конструктор для создания нового объекта типа Person.
func NewPerson(name string, age int, height float64) *Person {
	p := new(Person)
	p.Name = name
	p.Age = age
	p.Height = height
	return p
}

// SetAge устанавливает возраст человека.
func (p *Person) SetAge(age int) {
	p.Age = age
}

// GetAge возвращает возраст человека.
func (p Person) GetAge() int {
	return p.Age
}

// GetHeight возвращает рост человека.
func (p Person) GetHeight() float64 {
	return p.Height
}

// CompleteLab6 реализует логику лабораторной работы №6.
func CompleteLab6() {
	ivan := NewPerson("Дима", 20, 1.78)
	ivan.SetAge(20)

	// Выводим имя, возраст и рост
	fmt.Printf("Имя: %s\nВозраст: %d\nРост: %.2f\n", ivan.Name, ivan.GetAge(), ivan.GetHeight())
}
