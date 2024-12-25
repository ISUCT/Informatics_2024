package lab6

import "fmt"

func Lab6() {
	Hero1 := Hero{
		name:  "Иван",
		class: "Человек",
		lvl:   12,
	}
	fmt.Println("Старое имя:", Hero1.GetName())
	Hero1.SetName("Вася")
	fmt.Println("Новое имя:", Hero1.GetName())
	fmt.Println("Старый класс", Hero1.GetClass())
	Hero1.SetClass("Робот")
	fmt.Println("Новый класс:", Hero1.GetClass())
	fmt.Println("Старый уровент:", Hero1.GetLvl())
	Hero1.SetLvl(15)
	fmt.Println("Новый уровень:", Hero1.GetLvl())
}
