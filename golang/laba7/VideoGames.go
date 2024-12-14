package laba7

import "fmt"

type VideoGames struct {
	Name    string
	Size    float64
	Price   float64
	Company string
}

func (v *VideoGames) ApplyDiscount(discount float64) {
	v.Price = v.Price * (1 - discount/100)
}

func (v *VideoGames) GetPrice() float64 {
	return v.Price
}
func (v *VideoGames) SetPrice(newPrice float64) {
	v.Price = newPrice
}

func (v *VideoGames) UpdateCompany(newCompany string) {
	v.Company = newCompany
}

func (v *VideoGames) getInfo() string {
	return fmt.Sprintf("Name: %s, Size: %.2f , Price: %.2f,Company: %s", v.Name, v.Size, v.Price, v.Company)
}
