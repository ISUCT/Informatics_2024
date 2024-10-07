package main

import (
	"fmt"

	"isuct.ru/informatics2022/labs/lab5"
)

func main() {
	myPc := lab5.Computer{Hdd: 512, HddBrand: "Asus", Ram: 16, RamBrand: "Corsair"}
	fmt.Printf("Объем оперативной памяти: %d Gb\n", myPc.GetRam())
	myPc.SetRam(32)
	fmt.Printf("Объем оперативной памяти: %d Gb\n", myPc.GetRam())

	fmt.Printf("Объем жесткого диска: %d Gb\n", myPc.GetHdd())
	myPc.SetHdd(1024)
	fmt.Printf("Объем жесткого диска: %d Gb\n", myPc.GetHdd())

	fmt.Printf("Бренд жесткого диска: %s\nБренд оперативной памяти: %s\n", myPc.HddBrand, myPc.RamBrand)

	fmt.Println()

	newPc := lab5.NewComputer(2048, "WD Blue", 64, "HyperX")
	fmt.Printf("Объем жесткого диска нового компьютера: %d Gb\n", newPc.GetHdd())
	newPc.SetHdd(1024)
	fmt.Printf("Объем жесткого диска нового компьютера после изменений: %d Gb\n", newPc.GetHdd())
	fmt.Printf("Объем оперативной памяти нового компьютера: %d Gb бренда %s\n", newPc.GetRam(), newPc.GetRamBrand())
}
