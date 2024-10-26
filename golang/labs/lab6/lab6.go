package lab6

import (
	"strconv"
	"fmt"
)
type Cat struct {
	Name string
	Age int
	Breed string
}

func (cat *Cat) SetAge(age int) {
	cat.Age = age
}

func (cat *Cat) SetName(name string) {
	cat.Name = name
}

func (cat *Cat) SetBreed(breed string) {
	cat.Breed = breed
}

func (cat Cat) GetInfo () (string){
	return_str := "Имя кота: " + cat.Name + "\nВозраст кота: " + strconv.Itoa(cat.Age) + "\nПорода кота: " + cat.Breed
	return return_str
}

func AnsLab6(){
	var cat = Cat{Name: "Имя", Age: 18, Breed: "Парода"}

	fmt.Println(cat.GetInfo())
	
	cat.SetName("Мурзик")
	cat.SetAge(10)
	cat.SetBreed("Сфинкс")

	fmt.Println(cat.GetInfo())
}