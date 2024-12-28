package lab7

import "fmt"

type Jeans struct {
	name     string
	price    float64
	brand    string
	material string
	season   string
}

func (j *Jeans) getName() string {
	return j.name
}

func (j *Jeans) getPrice() float64 {
	return j.price
}

func (j *Jeans) setdiscount(discount float64) {
	j.price -= j.price * discount / 100
}

func (j *Jeans) setPrice(newPrice float64) {
	j.price = newPrice
}

func (j *Jeans) getBrand() string {
	return j.brand
}
func (j *Jeans) getSeason() string {
	return j.season
}
func (j *Jeans) getMaterial() string {
	return j.material
}
func (j *Jeans) GetInfo() string {
	return fmt.Sprintf("Продается товар: %s, ценой: %.2f, бренда: %s, сделанной из: %s,для сезона: %s", j.name, j.price, j.brand, j.material, j.season)
}
