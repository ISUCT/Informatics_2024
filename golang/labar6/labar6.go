package laba6

import (
	"fmt"
)

type PC struct {
	Name     string
	Color    string
	CPU      float64
	HardDisk float64
}

func NewPC(name, color string, cpu, hardDisk float64) *PC {
	c := new(PC)
	c.Name = name
	c.Color = color
	c.CPU = cpu
	c.HardDisk = hardDisk
	return c
}

func (c *PC) SetCPU(cpu float64)           { c.CPU = cpu }
func (c PC) GetCPU() float64               { return c.CPU }
func (c *PC) SetHardDisk(hardDisk float64) { c.HardDisk = hardDisk }
func (c PC) GetHardDisk() float64          { return c.HardDisk }
func (c PC) GetColor() string              { return c.Color }

func RunLab6() {
	newPC := NewPC("ПК", "Черный", 3.4, 500.0)
	newPC.SetCPU(3.8)
	newPC.SetHardDisk(1000.0)
	fmt.Println("скорость процессора:", newPC.GetCPU())
	fmt.Println("объем жесткого диска в GB:", newPC.GetHardDisk())
	fmt.Println("Цвет сборки пк:", newPC.GetColor())
}
