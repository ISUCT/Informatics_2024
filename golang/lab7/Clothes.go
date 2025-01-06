package lab7

import "fmt"

type Dress struct {
	Description string
	Fabric string
	Brand string
	Price float64
}

func (f Dress) GetDescription() {
	fmt.Println("Это", f.Description, "из", f.Fabric, "от бренда", f.Brand ,"стоимостью", f.Price)
}

func (f Dress) GetPrice() float64 {
	return f.Price
}

func (f *Dress) ApplyDiscount(x float64) {
	(*f).Price = (f.Price) * (x)
}

func (f *Dress) SetNewPrice(x float64) {
	(*f).Price = x
}

func (f *Dress) SetNewDescription(x string) {
	(*f).Fabric = x
}
