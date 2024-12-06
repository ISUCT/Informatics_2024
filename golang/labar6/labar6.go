package laba6

import "fmt"

type PC struct {
	HDD      int
	RAM      int
	CPUModel string
}

func NewPC(hdd int, ram int, cpuModel string) *PC {
	return &PC{HDD: hdd, RAM: ram, CPUModel: cpuModel}
}

func (c *PC) Info() {
	fmt.Printf("Компьютер с процессором %s, %d ГБ ОЗУ, жесткий диск: %d ГБ\n", c.CPUModel, c.RAM, c.HDD)
}

func RunLab6() {
	comp := NewPC(500, 16, "Intel Core i5-9600f")
	comp.Info()
	comp.HDD = 1000
	comp.Info()
}
