package lab6

import (
	"fmt"
)

type Fox struct {
	Name  string
	Color string
	Age   int
}

func (f *Fox) Speak() string {
	return "Rrrr-rrrr-rrrr-rrrr-rraavv!"
}

func (f *Fox) GetDescription() string {
	return fmt.Sprintf("Это %s, ей %d года %s лиса.", f.Name, f.Age, f.Color)
}

func (f *Fox) GetView() string {
	return ` 
         \\ // 
        ( o.o )
         > ^ < 
      `
}

func (f *Fox) IsAdult() bool {
	return f.Age >= 2
}

func RunLab6() {
	fox := Fox{
		Name:  "Foxy",
		Age:   3,
		Color: "красная",
	}

	fmt.Println(fox.GetDescription())

	if fox.IsAdult() {
		fmt.Println("Эта лиса - взрослая.")
	} else {
		fmt.Println("Эта лиса - не взрослая.")
	}

	fmt.Println(fox.Speak())
	fmt.Println(fox.GetDescription())
	fmt.Println(fox.GetView())
}
