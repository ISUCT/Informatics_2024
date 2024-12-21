package lab6

import "fmt"

type City struct {
	Country    string
	Population int
}

func NewCity(country string, population int) *City {
	return &City{
		Country:    country,
		Population: population,
	}
}
func (d *City) Hymn() string {
	return "Russland veri well!"
}

func (d *City) GetCountry() string {
	return d.Country
}

func (d *City) SetCountry(country string) {
	d.Country = country
}

func RunLab6() {
	city := NewCity("Reih", 865345)
	fmt.Println(city.Hymn())
	fmt.Println(city.GetCountry())
	city.SetCountry("Russland")
	fmt.Println(city.GetCountry())
}
