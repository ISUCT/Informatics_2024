package lab7

type Phone struct {
	Size  float64
	Model string
	Price float64
}

type PC struct {
	Model string
	Price float64
	RAM   float64
}

type Laptop struct {
	Model string
	Price float64
	ROM   float64
}

func (l *Laptop) SetModel(Model string) {
	l.Model = Model
}
func (l *Laptop) GetModel() string {
	return l.Model
}
func (l *Laptop) SetPrice(Price float64) {
	l.Price = Price
}
func (l *Laptop) GetPrice() float64 {
	return l.Price
}

func (l *Laptop) SetROM(ROM float64) {
	l.ROM = ROM
}

func (l *Laptop) SetDiscount(percent float64) {
	l.Price -= (l.Price / 100) * percent
}

func (p *Phone) SetSize(Size float64) {
	p.Size = Size
}

func (p *Phone) GetSize() float64 {
	return p.Size
}
func (p *Phone) SetModel(Model string) {
	p.Model = Model
}

func (p *Phone) SetPrice(Price float64) {
	p.Price = Price
}

func (p *Phone) GetPrice() float64 {
	return p.Price
}

func (p *Phone) SetDiscount(percent float64) {
	p.Price -= (p.Price / 100) * percent
}

func (c *PC) SetModel(Model string) {
	c.Model = Model
}

func (c *PC) SetPrice(Price float64) {
	c.Price = Price
}

func (c *PC) GetPrice() float64 {
	return c.Price
}

func (c *PC) GetModel() string {
	return c.Model
}

func (c *PC) SetDiscount(percent float64) {
	c.Price -= (c.Price / 100) * percent
}

func (c *PC) SetRAM(RAM float64) {
	c.RAM = RAM
}

func (c *PC) GetRAM() float64 {
	return c.RAM
}
