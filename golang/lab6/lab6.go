package lab6

import (
	"fmt"
)

type pig struct {
	Name   string
	Breed  string
	Weight float64
	Color  string
}

func NewPig(name string, breed string, weight float64, color string) *pig {
	return &pig{
		Name:   name,
		Breed:  breed,
		Weight: weight,
		Color:  color,
	}
}

func (p *pig) Eat(WeightFood float64) {
	p.Weight += WeightFood

}

func (p *pig) ColorPopularity() string {
	var a string
	switch p.Color {
	case "Розовый":
		a = "высокая распространенность"
	case "Черный":
		a = "умеренная распространенность"
	case "Серый":
		a = "распространенность ниже среднего"
	case "Коричневый":
		a = "низкая распространенность"
	default:
		a = "очень низкая распространенность"
	}
	return a

}

func (p *pig) InformationAboutPig(weightFood float64, colorPopularity string) {
	fmt.Printf("Свинка %s породы %s весит %.2f кг, её цвет %s\nУ цвета %s %s\nОна съела %.2f кг корма\nПосле того как %s покушает, она будет весить %.2f кг", p.Name, p.Breed, p.Weight, p.Color, p.Color, colorPopularity, weightFood, p.Name, p.Weight+weightFood)
}

func Runlab6() {
	weightFood := 8.50
	pig := NewPig("Нюша", "Йоркшир", 100, "Розовый")
	colorPopularity := pig.ColorPopularity()
	pig.InformationAboutPig(weightFood, colorPopularity)
}
