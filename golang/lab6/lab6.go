package lab6

import (
	"fmt"
)

type Gun struct {
	Model        string
	Caliber      float32
	AmmoCapacity int
}

func NewGun(model string, caliber float32, ammoCapacity int) Gun {
	return Gun{
		Model:        model,
		Caliber:      caliber,
		AmmoCapacity: ammoCapacity,
	}
}

func (g Gun) GetInfo() string {
	return fmt.Sprintf(g.Model, g.Caliber, g.AmmoCapacity)
}

func (g *Gun) SetCaliber(newCaliber float32) {
	g.Caliber = newCaliber
}

func (g *Gun) SetAmmoCapacity(newCapacity int) {
	g.AmmoCapacity = newCapacity
}

func Lab6() {
	gun := NewGun("Glock 19", 9.0, 15)
	fmt.Println(gun.GetInfo())

	gun.SetCaliber(10.0)
	fmt.Println("После изменения калибра:", gun.GetInfo())

	gun.SetAmmoCapacity(17)
	fmt.Println("После изменения вместимости магазина:", gun.GetInfo())
}
