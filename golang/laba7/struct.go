package laba7

type Electronics struct {
	name  string
	price float64
	model string
}

type Clothes struct {
	name     string
	price    float64
	material string
}

type Books struct {
	name   string
	price  float64
	author string
}

func (E *Electronics) SetModel(model string) {
	E.model = model
}

func (E *Electronics) SetPrice(price float64) {
	E.price = price
}

func (E *Electronics) SetDiscount(percent float64) {
	E.price -= (E.price / 100) * percent
}

func (E *Electronics) GetPrice() float64 {
	return E.price
}

func (C *Clothes) SetMaterial(material string) {
	C.material = material
}

func (C *Clothes) SetPrice(price float64) {
	C.price = price
}

func (C *Clothes) SetDiscount(percent float64) {
	C.price -= (C.price / 100) * percent
}

func (C *Clothes) GetPrice() float64 {
	return C.price
}
func (B *Books) SetAuthor(author string) {
	B.author = author
}

func (B *Books) SetPrice(price float64) {
	B.price = price
}

func (B *Books) SetDiscount(percent float64) {
	B.price -= (B.price / 100) * percent
}

func (B *Books) GetPrice() float64 {
	return B.price
}
