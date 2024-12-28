package lab6

import "fmt"

type Telephone struct {
	Model        string
	Weight       float64
	display_size float64
}

func NewTelephone(Model string, Weight, display_size float64) Telephone {
	return Telephone{
		Model:        Model,
		Weight:       Weight,
		display_size: display_size,
	}
}

func (f Telephone) Info() string {
	return fmt.Sprintf("Модель: %s Вес модели: %f, Размер дисплея модели: %f минут", f.Model, f.Weight, f.display_size)
}
func (f *Telephone) UpdateModel(newModel string) {
	f.Model = newModel
}
func (f Telephone) Buy() {
	fmt.Printf("Вы купили телефон модели:%s веса:%f у которого размер дисплея:%f", f.Model, f.Weight, f.display_size)
}
func ExecuteLab6() {
	Telephone := NewTelephone("Iphone 15 Pro Max", 0.221, 6.7)
	Telephone.Buy()
	fmt.Println(Telephone.Info())
	Telephone.UpdateModel("Iphone 16")
	fmt.Println(Telephone.Info())
}
