package lab7
import "fmt"
type Technic struct {
	Name      string
	Brend     string
	Price     float32
}

func (t Technic) GetInformation() {
	fmt.Println("У нас есть в наличии", t.Name, "от компании", t.Brend, "и стоит", t.Price)
}

func (t Technic) GetPrice() float32 {
	return t.Price
}

func (t *Technic) ApplyDiscount(x float32) {
	(*t).Price = (t.Price / 100) * (100 - x)
}

func (t *Technic) ChangePrice(x float32) {
	(*t).Price = x
}

func (t *Technic) ChangeDescription(x string) {
	(*t).Brend = x
}
