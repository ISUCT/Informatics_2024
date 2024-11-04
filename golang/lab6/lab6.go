package lab6

import (
  "fmt"
)

type Dish struct {
	Name string
	Price float64
	Description string
}

func NewDish(name string, price float64, description string) *Dish {
  d := new(Dish)
  d.Name = name
  d.Price = price
  d.Description = description
  return d
}

func (d *Dish) SetPrice(price float64) { d.Price = price }
func (d Dish) GetPrice() float64       { return d.Price }
func (d *Dish) SetDescription(description string) { d.Description = description }
func (d Dish) GetDescription() string            { return d.Description }
func (d Dish) GetName() string                 { return d.Name }

func CompleteLab6() {
  pasta := NewDish("Паста Карбонара", 450, "Спагетти с мелкими кусочками гуанчиале или панчетты, смешанные с соусом из яиц, сыра пекорино романо, соли и свежемолотого чёрного перца.")
  fmt.Println("Название:", pasta.GetName())
  fmt.Println("Цена:", pasta.GetPrice())
  fmt.Println("Описание:", pasta.GetDescription())
}
