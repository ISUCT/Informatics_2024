package lab6

import (
	"fmt"
)

type plane struct {
	name            string
	flight_number   int
	power           int
	purchased_seats int
}

func (c plane) speed_calculation(f_air int) int {
	speed := c.power / f_air
	return speed
}

func (c plane) revenue_calculation(price_a_place int) int {
	money := price_a_place * c.purchased_seats
	return money
}
func (c plane) get_routes() string {
	routes := []string{1: "Paris-Moscow", 2: "Moscow-Paris", 3: "Moscow-Seoul", 4: "Seoul-Moscow"}
	nomer := routes[c.flight_number]
	return nomer
}

func RunLab6() {
	Plane := plane{"Airobus999", 3, 200000, 150}
	fmt.Printf("Скорось %v равна %d м/с\n", Plane.name, Plane.speed_calculation(5000))
	fmt.Printf("Прибыль с продажи белетов составляет: %d рублей \n", Plane.revenue_calculation(23000))
	fmt.Printf("%v летит по маршруту %v\n", Plane.name, Plane.get_routes())
}
