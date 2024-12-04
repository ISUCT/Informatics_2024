package lab7

type vkusnyshki struct {
	product
}

func Newvkusnyshki(id int, name string, price float64) *vkusnyshki {
	f := &vkusnyshki{
		product: newProduct(id, name, price),
	}
	return f
}
