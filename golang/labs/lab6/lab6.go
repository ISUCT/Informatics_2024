package lab6

import "strconv"

type Cat struct {
	Name string
	Age int
	Breed string
}

func (cat *Cat) SetAge(uage int) {
	(*cat).Age = uage
}

func (cat *Cat) SetName(uname string) {
	(*cat).Name = uname
}

func (cat *Cat) SetBreed(ubreed string) {
	(*cat).Breed = ubreed
}

func (cat Cat) GetStatus () (string){
	return_str := "Имя кота: " + cat.Name + "\nВозраст кота: " + strconv.Itoa(cat.Age) + "\nПорода кота: " + cat.Breed
	return return_str
}