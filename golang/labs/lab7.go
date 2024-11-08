package labs

import "fmt"

type Product interface {
	discount(discount float64)
	setPrice(newPrice float64)
	getPrice() float64
	changeCharacteristic(string, string)
}

type Food struct {
	price    float64
	foodtype string
	brand    string
}

type Cosmetics struct {
	price         float64
	brand         string
	cosmeticstype string
}

type Clothes struct {
	price       float64
	clothestype string
	material    string
}

func (f *Food) getPrice() float64 {
	return f.price
}

func (f *Food) discount(discount float64) {
	f.price -= f.price * discount / 100
}

func (f *Food) setPrice(newPrice float64) {
	f.price = newPrice
}

func (f *Food) changeCharacteristic(newFoodtype string, newBrand string) {
	f.foodtype = newFoodtype
	f.brand = newBrand
}

func (c *Cosmetics) getPrice() float64 {
	return c.price
}

func (c *Cosmetics) discount(discount float64) {
	c.price -= c.price * discount / 100
}

func (c *Cosmetics) setPrice(newPrice float64) {
	c.price = newPrice
}

func (c *Cosmetics) changeCharacteristic(newBrand string, newCosmeticstype string) {
	c.brand = newBrand
	c.cosmeticstype = newCosmeticstype
}

func (c *Clothes) getPrice() float64 {
	return c.price
}

func (c *Clothes) discount(discount float64) {
	c.price -= c.price * discount / 100
}

func (c *Clothes) setPrice(newPrice float64) {
	c.price = newPrice
}

func (c *Clothes) changeCharacteristic(newClothestype string, newMaterial string) {
	c.clothestype = newClothestype
	c.material = newMaterial
}

func CalculateProductsSum(products []Product) float64 {
	var sum float64 = 0
	for _, product := range products {
		sum += product.getPrice()
	}
	return sum
}

func RunLab7() {
	product1 := &Food{50.15, "картошка", "пятерочка"}
	product2 := &Cosmetics{1700, "Chanel", "помада"}
	product3 := &Clothes{3000, "толстовка", "хлопок"}

	products := []Product{product1, product2, product3}
	fmt.Println("Общая стоимость товаров:", CalculateProductsSum(products), "рублей")
	product1.discount(20)
	product2.discount(10)
	product3.discount(15)
	fmt.Println("Общая стоимость товаров после применения скидок:", CalculateProductsSum(products), "рублей")
}
