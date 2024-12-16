package lab7

import "fmt"

func AnsLab7() {
	var product1 car = car{"Тойота Камри 2021", 1450000, 4, "Легковые машины"}
	var product2 motorbike = motorbike{"ElectroTown Citycoco X7 2000W PRO", 94500, 10, "Электромотоциклы"}
	var product3 tire = tire{"Зимняя резина", 5000, 0, "Шины для машин"}
	var product4 tire = tire{"Зимняя резина", 2500, 20, "Шины для мотоцикла"}

	var cart cart = cart{}

	cart.add(&product1)
	cart.add(&product2)
	for i := 0; i < 5; i++ {
		cart.add(&product3)
	}
	cart.add(&product4)

	fmt.Println(cart.getTotalSell())

	cart.purge(&product2)
	cart.purge(&product4)
	cart.purge(&product4)

	fmt.Println(cart.getTotalSell())
}
