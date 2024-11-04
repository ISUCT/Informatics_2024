package main

import (
 "fmt"
)

type Fox struct {
 name   string
 age    int
 color  string
}

func NewFox(name string, age int, color string) Fox {
 return Fox{name: name, age: age, color: color}
}

func (f Fox) GetView() string {
 return 
       /   /
      {   + f.color +   }
       + f.name +  is  + fmt.Sprint(f.age) +  years old.
        __/
    
}

func (f *Fox) SetAge(age int) {
 f.age = age
}

func (f *Fox) SetColor(color string) {
 f.color = color
}

func (f Fox) GetInfo() string {
 return fmt.Sprintf("Name: %s, Age: %d, Color: %s", f.name, f.age, f.color)
}

func Runlab6() {
 fox := NewFox("Foxy", 3, "red")

 fmt.Println(fox.GetView())

 fmt.Println(fox.GetInfo())

 fox.SetAge(4)
 fox.SetColor("white")

 fmt.Println(fox.GetInfo())
}