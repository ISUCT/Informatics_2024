package lab7

import "fmt"

type Product interface {
	FindOutPrice() float64
	EditPrice(float64)
	ApplyDiscount(float64)
	EditDescription(string)
	GetInfo()
}

type Car struct {
	Mark     string
	Colour   string
	MaxSpeed int
	Price    float64
}

func (c Car) FindOutPrice() float64 {
	return c.Price
}

func (c *Car) EditPrice(NewPrice float64) {
	(*c).Price = NewPrice
}

func (c *Car) ApplyDiscount(Percent float64) {
	(*c).Price = (c.Price / 100) * (100 - Percent)
}

func (c *Car) EditDescription(NewDescription string) {
	(*c).Colour = NewDescription
}

func (c Car) GetInfo() {
	fmt.Println("Продаётся машина", c.Mark, "цвета", c.Colour, ", которая разгоняется до", c.MaxSpeed, "и стоит в долларах:", c.Price)
}

type Headphones struct {
	Brand     string
	Colour    string
	MaxVolume int
	Price     float64
}

func (h Headphones) FindOutPrice() float64 {
	return h.Price
}

func (h *Headphones) EditPrice(NewPrice float64) {
	(*h).Price = NewPrice
}

func (h *Headphones) ApplyDiscount(Percent float64) {
	(*h).Price = (h.Price / 100) * (100 - Percent)
}

func (h *Headphones) EditDescription(NewDescription string) {
	(*h).Colour = NewDescription
}

func (h Headphones) GetInfo() {
	fmt.Println("В продаже есть наушники бренда", h.Brand, "цвета", h.Colour, ", которые имеют громкость до", h.MaxVolume, "дб и стоят:", h.Price)
}

type Keyboard struct {
	Brand      string
	Format     string
	NumberKeys int
	Price      float64
}

func (k Keyboard) FindOutPrice() float64 {
	return k.Price
}

func (k *Keyboard) EditPrice(NewPrice float64) {
	(*k).Price = NewPrice
}

func (k *Keyboard) ApplyDiscount(Percent float64) {
	(*k).Price = (k.Price / 100) * (100 - Percent)
}

func (k *Keyboard) EditDescription(NewDescription string) {
	(*k).Format = NewDescription
}

func (k Keyboard) GetInfo() {
	fmt.Println("В продаже есть клавиатура, ", k.Format, "бренда", k.Brand, ", у которой", k.NumberKeys, " клавиш и стоит:", k.Price)
}

func CalculatePrice(purchase []Product) float64 {
	var summ float64 = 0
	for _, price := range purchase {
		summ += price.FindOutPrice()
	}
	return summ
}

func RunLab7() {
	var car Product = &Car{"BMW", "чёрная", 300, 10000}
	var headphones Product = &Headphones{"Apple", "белые", 100, 25000}
	var keyboard Product = &Keyboard{"Red Square", "проводная", 86, 4599}

	car.GetInfo()
	headphones.GetInfo()
	keyboard.GetInfo()
	fmt.Println(car.FindOutPrice(), headphones.FindOutPrice(), keyboard.FindOutPrice())
	car.EditPrice(60)

	var purchase []Product = []Product{car, headphones, keyboard}
	summ := CalculatePrice(purchase)
	fmt.Println("Общая стоимость товара:", summ)

	headphones.EditDescription("электронном")
	headphones.ApplyDiscount(40)
	keyboard.ApplyDiscount(60)
	keyboard.EditDescription("L")

	car.GetInfo()
	headphones.GetInfo()
	keyboard.GetInfo()

	summ = CalculatePrice(purchase)
	fmt.Println("Общая стоимость товара:", summ)
}
