package laba6

import "fmt"

type Dish struct {
	Name  string
	Price int
	Grams int
}

func (d *Dish) UpdateGrams(newGrams int) {
	d.Grams = newGrams
}

func (d *Dish) Applydiscount(newPrice int) {
	d.Price = newPrice
}
func (d *Dish) ChangeDish(newName string) {
	d.Name = newName
}

func Runlaba6() {
	var Borsh = Dish{Name: "Borsh", Price: 51, Grams: 400}
	fmt.Println(Borsh.Grams)
	Borsh.UpdateGrams(1000)
	fmt.Println(Borsh.Grams)
	var ChocolateCake = Dish{Name: "ChocolateCake", Price: 125, Grams: 250}
	fmt.Println(ChocolateCake.Price)
	ChocolateCake.Applydiscount(89)
	fmt.Println(ChocolateCake.Price)
	var Shrimps = Dish{Name: "Shrimps", Price: 100, Grams: 50}
	fmt.Println(Shrimps.Name)
	Shrimps.ChangeDish("FriedChiken")
	fmt.Println(Shrimps.Name)
}
