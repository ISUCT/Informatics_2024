package lab7

type Product interface {
	GetPrice() float64
	ApplyingTheDiscount(float64)
	ChangePrice(float64)
	ChangeCharacteristic(string, string)
}
