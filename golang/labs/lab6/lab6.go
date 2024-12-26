package lab6

import (
	"fmt"
)

type Mouse struct {
	Name   string
	Age    int
	Weight float64
}

func NewMouse(name string, age int, weight float64) Mouse {
	return Mouse{Name: name, Age: age, Weight: weight}
}

func (r *Mouse) GetAge() int {
	return r.Age
}

func (r *Mouse) SetAge(age int) {
	if age > 0 {
		r.Age = age
	}
}

func (r Mouse) Info() string {
	return fmt.Sprintf("Имя: %s, Возраст: %d, Вес: %.2f кг", r.Name, r.Age, r.Weight)
}

func RunLab6() {
	mouse := NewMouse("Гена", 2, 1.5)

	fmt.Println(mouse.Info())

	mouse.SetAge(3)

	fmt.Printf("Обновленный возраст: %d\n", mouse.GetAge())

	fmt.Println(mouse.Info())
}
