package lab6

import (
	"fmt"
)

type Fox struct {
	Name       string
	Age        int
	Wool_color string
}

func NewFox(name string, age int, wool_color string) *Fox {
	f := &Fox{
		Name:       name,
		Age:        age,
		Wool_color: wool_color,
	}
	return f
}

func (f Fox) GetView() string {
	a := "  /\\   /\\  \n" + " {  ---'  } \n" + " {  O   O  } \n" + " >  V  < \n" + "  \\ \\|/ \n" + "   -----'____ \n" + "   /     \\    \\_ \n" + "  {       }\\  )_ \\_   _ \n" + "  |  \\_/  |/ /   \\_\\_/ \n" + "   \\__/  /(_/ \n" + "     (__/ \n"
	return a
}

func (f Fox) DisplayInfo() {
	fmt.Printf("Имя: %dnВозраст: %dnЦвет: %sn", f.Name, f.Age, f.Wool_color)
}

func (f Fox) Speak(message string) string {
	return fmt.Sprintf("%s говорит: \"%s\"", f.Name, message)
}

func Laba6() {
	myFox := NewFox("Лиса", 3, "Рыжий")
	myFox.DisplayInfo()
	//fmt.Println(myFox.GetView())
	fmt.Println(myFox.Speak("Wow-wow"))
}
