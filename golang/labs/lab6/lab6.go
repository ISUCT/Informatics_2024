package lab6

import (
	"fmt"
)

type Car struct {
	speed float64
	color string
	move  string
}

func (c *Car) setSpeed(newSpeed float64) {
	c.speed = newSpeed
}

func (c Car) getSpeed() float64 {
	return c.speed
}

func (c *Car) setColor(newColor string) {
	c.color = newColor
}

func (c Car) getColor() string {
	return c.color
}

func (c *Car) setMove(newMove string) {
	c.move = newMove
}

func (c Car) getMove() string {
	return c.move
}

func RunLab6Task() {
	car := Car{speed: 60.1, color: "Red", move: "Forward"}

	fmt.Printf("Скорость движения автомобиля %f км/ч\n", car.getSpeed())
	car.setSpeed(70.4)
	fmt.Printf("Скорость движения автомобиля %f км/ч\n", car.getSpeed())

	fmt.Printf("Цвет автомобиля %s\n", car.getColor())
	car.setColor("Green")
	fmt.Printf("Цвет автомобиля %s\n", car.getColor())

	fmt.Printf("Направление движения автомобиля %s\n", car.getMove())
	car.setMove("Back")
	fmt.Printf("Цвет автомобиля %s\n", car.getMove())
}
