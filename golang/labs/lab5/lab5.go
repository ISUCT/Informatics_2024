package lab5

import "fmt"

type Computer struct {
	Hdd      int
	HddBrand string
	Ram      int
	RamBrand string
}

func NewComputer(hdd int, hddBrand string, ram int, ramBrand string) Computer {
	return Computer{
		Hdd:      hdd,
		HddBrand: hddBrand,
		Ram:      ram,
		RamBrand: ramBrand,
	}
}

func (c Computer) GetHdd() int {
	return c.Hdd
}

func (c *Computer) SetHdd(newHdd int) {
	c.Hdd = newHdd
}

func (c Computer) GetHddBrand() string {
	return c.HddBrand
}

func (c Computer) GetRam() int {
	return c.Ram
}

func (c *Computer) SetRam(newRam int) {
	c.Ram = newRam
}

func (c Computer) GetRamBrand() string {
	return c.RamBrand
}

func RunLab5() {
	myPc := Computer{Hdd: 512, HddBrand: "Asus", Ram: 16, RamBrand: "Corsair"}
	fmt.Printf("Объем оперативной памяти: %d Gb\n", myPc.GetRam())
	myPc.SetRam(32)
	fmt.Printf("Объем оперативной памяти: %d Gb\n", myPc.GetRam())

	fmt.Printf("Объем жесткого диска: %d Gb\n", myPc.GetHdd())
	myPc.SetHdd(1024)
	fmt.Printf("Объем жесткого диска: %d Gb\n", myPc.GetHdd())

	fmt.Printf("Бренд жесткого диска: %s\nБренд оперативной памяти: %s\n", myPc.HddBrand, myPc.RamBrand)

	fmt.Println()

	newPc := NewComputer(2048, "WD Blue", 64, "HyperX")
	fmt.Printf("Объем жесткого диска нового компьютера: %d Gb\n", newPc.GetHdd())
	newPc.SetHdd(1024)
	fmt.Printf("Объем жесткого диска нового компьютера после изменений: %d Gb\n", newPc.GetHdd())
	fmt.Printf("Объем оперативной памяти нового компьютера: %d Gb бренда %s\n", newPc.GetRam(), newPc.GetRamBrand())
}
