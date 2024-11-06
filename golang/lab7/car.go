package lab7

type car struct {
	product
	color string
}

func (c *car) SetColor() error {
	return nil
}

func NewCar(id int, name string, price float64, color string) *car {
	return &car{
		product: newProduct(id, name, price),
		color:   color,
	}
}
