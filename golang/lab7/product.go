package lab7

import (
	"errors"
	"fmt"
)

//const err = errors.New("the wrong percentage value")

type Product interface {
	SetDiscount(float64) error
	GetDiscount() float64
	GetPrice() float64
	GetName() string
}

type product struct {
	id       int
	name     string
	price    float64
	discount float64
}

func (p product) GetPrice() float64 { return p.price }

func (p product) GetName() string { return p.name }

func (p product) GetDiscount() float64 { return p.discount }

func (p *product) SetDiscount(percent float64) error {
	if percent < 1 || percent > 100 {
		return errors.New("the wrong percentage value")
	}
	p.discount = percent

	change := (p.price / 100) * percent

	if change > p.price {
		p.price = 0
		return nil
	}
	p.price -= change
	return nil
}

func newProduct(id int, name string, price float64) product {
	return product{
		id:       id,
		name:     name,
		price:    price,
		discount: 0,
	}
}

func Write(PList []Product) {
	fmt.Println("____список товаров_____")
	for _, p := range PList {
		fmt.Printf(
			"________________________\nтовар \"%v\" стоит %v$ благодоря скидке в %v процентов\n",
			p.GetName(), p.GetPrice(), p.GetDiscount())
	}
}

func GetTotalСost(ProductList []Product) float64 {
	var totalСost float64

	for _, product := range ProductList {
		totalСost += product.GetPrice()
	}
	return totalСost
}
