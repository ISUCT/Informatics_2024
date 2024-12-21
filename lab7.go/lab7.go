package lab7

import (
  "fmt"
)

type Product interface {
  GetPrice() float32
  Discount(float32)
  ChangePrice(float32)
  changeCharacteristic(string, string)
}
func CalculateTotalCost(products []Product) float32 {
  var total float32 = 0.0
  for _, product := range products {
    total += product.GetPrice()
  }
  return total
}

func RunLab7A() {
  product1 := &Highlighter{Color: "Хайлайтер", Price: 1000.0}
  product2 := &Eyeliner{Colot: "Подводка", Price: 1500.0}
  product3 := &Mascara{Color: "Тушь", Price: 3000.0}

  product1.ApplyDiscount(100)
  product2.ApplyDiscount(50)

  products := []Product{product1, product2, product3}

  fmt.Printf("Общая стоимость без учёта скидок: %.2f\n", CalculateTotalCost([]Product{
    &Highlighter{Color: "Хайлайтер", Price: 2000.0},
    &Eyeliner{Material: "Подводка", Price: 3000.0},
    &Mascara{Color: "Тушь", Price: 4000.0},
  }))

  totalAfterDiscounts := CalculateTotalCost(products)
  fmt.Printf("Общая стоимость с учётом скидок: %.2f\n", totalAfterDiscounts)
}


