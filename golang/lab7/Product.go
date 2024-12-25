package lab7

import "errors"

var ErrPercentageValue error = errors.New("the wrong percentage value")

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
		return ErrPercentageValue
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

func NewProduct(id int, name string, price float64) product {
	return product{
		id:       id,
		name:     name,
		price:    price,
		discount: 0,
	}
}
