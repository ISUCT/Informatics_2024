package lab6

import (
	"fmt"
	"strconv"
)

const maxPowerUsage = 800

type PC struct {
	Brand      string
	GPU        string
	CPU        string
	PowerUsage int
	Storage    int
}

func (pc *PC) SetBrand(brand string) {
	pc.Brand = brand
}

func (pc *PC) SetPowerUsage(powerUsage int) {
	if powerUsage > maxPowerUsage {
		fmt.Println("Ошибка: превышение максимальной мощности потребления!")
	} else {
		pc.PowerUsage = powerUsage
	}
}

func (pc *PC) SetGPU(gpu string) {
	pc.GPU = gpu
}

func (pc *PC) SetCPU(cpu string) {
	pc.CPU = cpu
}

func (pc *PC) SetStorage(storage int) {
	pc.Storage = storage
}

func (pc *PC) GetInfo() string {
	return "Бренд: " + pc.Brand + "\nПроцессор: " + pc.CPU + "\nВидеокарта: " + pc.GPU + "\nЖесткий диск: " + strconv.Itoa(pc.Storage) + " Гб" + "\nЭнергопотребление: " + strconv.Itoa(pc.PowerUsage) + " Вт"
}

func RunLab6() {
	pc := PC{
		Brand:      "LG",
		CPU:        "Intel Xeon 231",
		GPU:        "Nvidia 1050 GTX",
		Storage:    256,
		PowerUsage: 300}
	fmt.Println(pc.GetInfo())
	pc.SetBrand("BMW")
	pc.SetCPU("Razen Radeon 5632")
	pc.SetGPU("Invoker RTX SunsrikeSuper 2090")
	pc.SetStorage(512)
	pc.SetPowerUsage(2000)
	fmt.Println(pc.GetInfo())
}
