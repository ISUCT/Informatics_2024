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

func (p *pig) ColorPopularity(color string) string {
	var a string
	switch color {
	case "Розовый":
		a = "Высокая распространенность"
	case "Черный":
		a = "Умеренная распространенность"
	case "Серый":
		a = "Распространенность ниже среднего"
	case "Коричневый":
		a = "Низкая распространенность"
	default:
		a = "Очень низкая распространенность"
	}
	return a

}

func (p *pig) InformationAboutPig(color string) {
	colorPopularity := p.ColorPopularity(color)
	fmt.Printf("Свинка %s породы %s весит %.2f кг, её цвет %s\nУ цвета %s %s", p.Name, p.Breed, p.Weight, p.Color, p.Color, colorPopularity)
}

func Runlab6() {
	pig := NewPig("Нюша", "Йоркшир", 100, "Розовый")
	pig.InformationAboutPig(pig.Color)
	pig.Eat(0.7)

}
