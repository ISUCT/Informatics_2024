package labs

import "fmt"

func CalculationSumProduct(listproducts []Product) float64 {
  var sum float64 = 0
  for _, product := range listproducts {
    sum += product.getPrice()
  }
  return sum
}

func labs() {
  film := &Media{price: 350.00, name: "Субстанция", genre: "Horror", duration: "98 minutes"}
  waffle := &Sweets{price: 300.00, name: "Венские вафли"}
  smartphone := &Phones{price: 39000.00, name: "Huawei 16 Pro'", brand: "Huawei", color: "White"}

  film.setPrice(450.00)
  film.setName("Red One")
  film.setGenre("Comedy")
  film.setDuration("130 minutes")
  smartphone.setBrand("Samsung")
  smartphone.setPrice(79000.99)

  listproducts := []Product{film, waffle, smartphone}

  fmt.Printf("Сумма товаров, без учёта скидки, равна: %.2f рублей \n", CalculationSumProduct(listproducts))

  discounts := []float64{5, 25, 15}
  for i, product := range listproducts {
    product.applyDiscount(discounts[i])
  }

  fmt.Printf("Сумма товаров, с учётом скидки, равна: %.2f рублей \n", CalculationSumProduct(listproducts))
}