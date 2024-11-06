package lab7

type furniture struct {
	product
	material string
}

func NewFurniture(id int, name string, price float64, material string) *furniture {
	return &furniture{
		product:  newProduct(id, name, price),
		material: material,
	}
}
