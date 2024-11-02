package lab6

import (
	"fmt"
)

type Fox struct {
	name       string
	age        int
	wool_color string
}

func NewFox(name string, age int, wool_color string) *Fox {
	f := &Fox{
		name:       name,
		age:        age,
		wool_color: wool_color,
	}
	return f
}

func (f *Fox) GetView() string {
	a := "  /\\   /\\  \n" + " {  ---'  } \n" + " {  O   O  } \n" + " >  V  < \n" + "  \\ \\|/ \n" + "   -----'____ \n" + "   /     \\    \\_ \n" + "  {       }\\  )_ \\_   _ \n" + "  |  \\_/  |/ /   \\_\\_/ \n" + "   \\__/  /(_/ \n" + "     (__/ \n"
	return a
}

func (f Fox) DisplayInfo() (string, int, string) {
	return f.name, f.age, f.wool_color
}

func (f Fox) Speak(message string) string {
	return message
}

func Laba6() {
	myFox := NewFox("Жорик", 3, "Рыжий")

	Name, Age, Whool_color := myFox.DisplayInfo()
	fmt.Printf("Имя: %v Возраст: %v Цвет: %v \n", Name, Age, Whool_color)

	fmt.Println(myFox.GetView())

	fmt.Println(myFox.Speak("Wow-wow"))
}
