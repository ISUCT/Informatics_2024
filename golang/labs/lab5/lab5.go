package lab5

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
