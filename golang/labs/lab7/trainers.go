package lab7

import "fmt"

type Shoes struct {
    name  string
    price float64
    brand string
    size  string
}

func (s *Shoes) getPrice() float64 {
    return s.price
}

func (s *Shoes) setPrice(newPrice float64) {
    s.price = newPrice
}

func (s *Shoes) applyDiscount(discount float64) {
    if discount < 0 || discount > 100 {
        fmt.Println("Скидка должна быть в диапазоне от 0 до 100%")
        return
    }
    s.price = s.price * (1 - discount/100)
}

func (s *Shoes) getName() string {
    return s.name
}

func (s *Shoes) setName(newName string) {
    s.name = newName
}

func (s *Shoes) setBrand(newBrand string) {
    s.brand = newBrand
}