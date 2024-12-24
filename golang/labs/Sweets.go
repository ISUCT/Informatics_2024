package labs

import "fmt"

type Sweets struct {
  name  string
  price float64
}

func (s *Sweets) getPrice() float64 {
  return s.price
}

func (s *Sweets) setPrice(newPrice float64) {
  s.price = newPrice
}

func (s *Sweets) applyDiscount(discount float64) {
  if discount < 0 || discount > 100 {
    fmt.Println("Скидка на сладости в диапазоне от 0 - 100%")
    return
  }
  s.price = s.price * (1 - discount/100)
}

func (s *Sweets) getName() string {
  return s.name
}

func (s *Sweets) setName(newName string) {
  s.name = newName
}