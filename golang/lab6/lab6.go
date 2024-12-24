package lab6

import "fmt"

type Phone struct {
	Name   string
	Op     string
	Number string
}

func NewPhone(name, op, number string) *Phone {
	p := new(Phone)
	p.Name = name
	p.Op = op
	p.Number = number
	return p
}

func (p *Phone) SetNumber(number string) { p.Number = number }
func (p Phone) GetNumber() string        { return p.Number }
func (p Phone) GetOp() string            { return p.Op }
func (p Phone) GetName() string          { return p.Name }

func Runlab6() {
	phone := NewPhone("Влад", "Билайн", "89621658549")
	phone.SetNumber("89012863969")
	fmt.Println("Имя владельца:", phone.GetName())
	fmt.Println("Оператор сотовой связи:", phone.GetOp())
	fmt.Println("Номер телефона пользователя:", phone.GetNumber())
}
