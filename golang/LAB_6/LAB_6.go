package LAB_6

import (
	"fmt"
	"strconv"
)

type Car struct {
	Model    string
	Speed    int
	MaxSpeed int
}

func (car *Car) SetModel(model string) {
	car.Model = model
}

func (car *Car) SetSpeed(speed int) {
	if speed <= car.MaxSpeed {
		car.Speed = speed
	} else {
		fmt.Println("Ошибка: превышена максимальная скорость!")
	}
}

func (car *Car) SetMaxSpeed(maxSpeed int) {
	car.MaxSpeed = maxSpeed
}

func (car Car) GetInfo() string {
	return "Модель: " + car.Model + "\nТекущая скорость: " + strconv.Itoa(car.Speed) + " км/ч\nМаксимальная скорость: " + strconv.Itoa(car.MaxSpeed) + " км/ч"
}

func Lab6() {
	car := Car{Model: "Toyota Corolla", Speed: 0, MaxSpeed: 180}

	fmt.Println(car.GetInfo())

	car.SetModel("Mazda rx-7")
	car.SetSpeed(120)
	car.SetMaxSpeed(250)

	fmt.Println(car.GetInfo())
}
