package lab7

import "fmt"

func Start7lab() {
	product1 := &Clothes{Name: "Шорты", Price: 1000, Size: "S", Color: "White"}
	product1.make_sale(500)
	product2 := &Vegetables{Name: "Перец", Weight: 100, Price: 100}
	product2.make_sale(50)
	product3 := &Book{Name: "Алиса в стране чудес", Format: "бумажный", Price: 450}
	product3.make_sale(100)

	products := []Product{product1, product2, product3}

	for _, product := range products {
		fmt.Println(product.get_productInfo())
	}
	totalPrice := get_priceAllProducts(products)
	fmt.Printf("Общая стоимость: %.2f\n", totalPrice)
}
