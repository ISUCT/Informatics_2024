package main

import "fmt"

type PC struct {
	HDD      int    // объем жесткого диска в ГБ
	RAM      int    // объем оперативной памяти в ГБ
	CPUModel string // модель процессора
}

//структура PC.
func NewPC(hdd int, ram int, cpuModel string) *PC {
	return &PC{
		HDD:      hdd,
		RAM:      ram,
		CPUModel: cpuModel,
	}
}

func (c *PC) GetHDD() int {
	return c.HDD
}

func (c *PC) SetHDD(newHDD int) {
	c.HDD = newHDD
}

func (c *PC) Info() {
	fmt.Printf("Компьютер с процессором %s, %d ГБ ОЗУ, жесткий диск: %d ГБ\n", c.CPUModel, c.RAM, c.HDD) //вывод инфы о компьютере
}

func main() {

	comp := NewPC(500, 16, "Intel Core i5-9600f")

	comp.Info()

	comp.SetHDD(1000)

	comp.Info()
}
