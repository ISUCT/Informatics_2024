package lab7

type Vkusnyshki struct {
	product
}

func NewVkusnyshki(id int, name string, price float64) *Vkusnyshki {
	f := &Vkusnyshki{
		product: NewProduct(id, name, price),
	}
	return f
}
