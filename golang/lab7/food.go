package lab7

type food struct {
	product
}

func Newfood(id int, name string, price float64) *food {
	f := &food{
		product: newProduct(id, name, price),
	}
	return f
}
