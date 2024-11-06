package lab7

func Lab7() {
	var ProductList []Product = []Product{
		NewCar(0, "бентли", 100000, "красный"),
		NewFurniture(1, "диван", 1000, "кожа"),
	}

	for _, product := range ProductList {
		err := product.SetDiscount(10)
		if err != nil {
			panic(err)
		}
	}
	Write(ProductList)
}
