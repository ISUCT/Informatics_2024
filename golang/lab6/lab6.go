package lab6

import (
	"fmt"
)

type Plane struct {
	name            string
	flight_number   int
	flight_distance float32
	passenger_mass  int
}

func (c Plane) Speed_calculation(time float32) float32 {
	speed := c.flight_distance / time
	return speed
}

func (c Plane) Get_routes() string {
	routes := []string{1: "Москва--Санкт-Петербург", 2: "Москва-Сочи", 3: "Мюнхен-Берлин", 4: "Нью-Йорк--Лос-Анджелес"}
	number := routes[c.flight_number]
	return number
}

func (c Plane) Load_capacity_calculation(luggage_mass int) int {
	load_capacity := luggage_mass + c.passenger_mass
	return load_capacity
}

func RunLab6() {
	Plane := Plane{"Super Jet", 4, 3937.6, 10}
	fmt.Printf("Скорость самолета %v равна %.1f км/ч\n", Plane.name, Plane.Speed_calculation(4.4))
	fmt.Printf("Самолет %v преодолеет расстояние в %.0f км\n", Plane.name, Plane.flight_distance)
	fmt.Printf("Полезная нагрузка самолета %d тонн \n", Plane.Load_capacity_calculation(27))
	fmt.Printf("Самолет %v летит по маршруту номер %v - %v\n", Plane.name, Plane.flight_number, Plane.Get_routes())
}
