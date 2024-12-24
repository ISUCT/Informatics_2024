package labs

import "fmt"

type Media struct {
  name     string
  price    float64
  m.genre    string
  m.duration string
}

func (m *Media) getPrice() float64 {
  return m.price
}

func (m *Media) setPrice(newPrice float64) {
  m.price = newPrice
}

func (m *Media) applyDiscount(discount float64) {
  if discount < 0 || discount > 100 {
    fmt.Println("Скидка на товар в диапазоне от 0 - 100%")
    return
  }
  m.price = m.price * (1 - discount/100)
}

func (m *Media) getName() string {
  return m.name
}

func (m *Media) setName(newName string) {
  m.name = newName
}

func (m *Media) setGenre(newGenre string) {
  m.genre = newGenre
}
func (m *Media) setDuration(newDuration string) {
  m.duration = newDuration
}
