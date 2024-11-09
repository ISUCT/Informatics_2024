package labs

type Product interface {
	discount(discount float64)
	setPrice(newPrice float64)
	getPrice() float64
	changeCharacteristic(string, string)
}
