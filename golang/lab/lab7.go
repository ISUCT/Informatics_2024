package lab

import "fmt"

type Product interface {
	GetInfo()
	GetPrice() float32
	Sale(float32)
	ChangePrice(float32)
	ChangeChar(string)
}

type Fruit struct {
	Name      string
	Freshness string
	Price     float32
}

func (f Fruit) GetInfo() {
	fmt.Println("У нас есть в наличии", f.Name, "оно", f.Freshness, "и стоит", f.Price)
}

func (f Fruit) GetPrice() float32 {
	return f.Price
}

func (f *Fruit) Sale(x float32) {
	(*f).Price = (f.Price / 100) * (100 - x)
}

func (f *Fruit) ChangePrice(x float32) {
	(*f).Price = x
}

func (f *Fruit) ChangeChar(x string) {
	(*f).Freshness = x
}

type Book struct {
	Name   string
	Format string
	Price  float32
}

func (b Book) GetInfo() {
	fmt.Println("У нас есть в наличии", b.Name, "в", b.Format, "формате и стоит", b.Price)
}

func (b Book) GetPrice() float32 {
	return b.Price
}

func (b *Book) Sale(x float32) {
	(*b).Price = (b.Price / 100) * (100 - x)
}

func (b *Book) ChangePrice(x float32) {
	(*b).Price = x
}

func (b *Book) ChangeChar(x string) {
	(*b).Format = x
}

type Clothes struct {
	Name  string
	Size  string
	Price float32
}

func (c Clothes) GetInfo() {
	fmt.Println("У нас есть в наличии", c.Name, "размера", c.Size, "стоимостью", c.Price)
}

func (c Clothes) GetPrice() float32 {
	return c.Price
}

func (c *Clothes) Sale(x float32) {
	(*c).Price = (c.Price / 100) * (100 - x)
}

func (c *Clothes) ChangePrice(x float32) {
	(*c).Price = x
}

func (c *Clothes) ChangeChar(x string) {
	(*c).Size = x
}

func CalculatePok(list []Product) {
	var sum float32 = 0
	for _, price := range list {
		sum += price.GetPrice()
	}
	fmt.Println("Общая стоимость товаров:", sum)
}

func RunLab7Tasks() {
	var interface_fruit Product = &Fruit{"яблоко", "свежее", 40}
	var interface_book Product = &Book{"Война и мир", "бумажном", 500}
	var interface_clothes Product = &Clothes{"худи", "M", 3500}
	interface_fruit.GetInfo()
	interface_book.GetInfo()
	interface_clothes.GetInfo()
	var pok []Product = []Product{interface_fruit, interface_book, interface_clothes}
	CalculatePok(pok)

	interface_fruit.ChangePrice(60)
	interface_book.ChangeChar("электронном")
	interface_book.Sale(40)
	interface_clothes.Sale(60)
	interface_clothes.ChangeChar("L")
	interface_fruit.GetInfo()
	interface_book.GetInfo()
	interface_clothes.GetInfo()

	CalculatePok(pok)
}
