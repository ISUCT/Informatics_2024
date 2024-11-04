package lab6

import "fmt"

type Person struct {
	Class  string
	Lvl    int
	Weapon string
}

func NewPerson(class string, lvl int, weapon string) *Person {
	pers := new(Person)
	pers.Class = class
	pers.Lvl = lvl
	pers.Weapon = weapon
	return pers
}

func (pers *Person) SetLvl(lvl int) {
	pers.Lvl = lvl
}

func (pers Person) GetLvl() int {
	return pers.Lvl
}

func (pers Person) GetWeapon() string {
	return pers.Weapon
}

func Lab6() {
	vedmak := NewPerson("Ведьмак", 10, "Меч")
	vedmak.SetLvl(15)
	fmt.Println(vedmak.GetLvl())
	fmt.Println(vedmak.GetWeapon())
}
