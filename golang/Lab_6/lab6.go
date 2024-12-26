package lab6

import (
  "fmt"
  "time"
)

type Employee struct {
  Name     string
  Position string
  HireDate time.Time
}

func NewEmployee(name, position string, hireDate time.Time) Employee {
  return Employee{
    Name:     name,
    Position: position,
    HireDate: hireDate,
  }
}

func (e Employee) DisplayInfo() {
  fmt.Printf("Имя: %s\nДолжность: %s\nДата принятия на работу: %s\n", e.Name, e.Position, e.HireDate.Format("2006-01-02"))
}

func (e Employee) GetYearsOfService() int {
  return int(time.Since(e.HireDate).Hours() / 24 / 365)
}

func (e *Employee) Promote(newPosition string) {
  e.Position = newPosition
  fmt.Printf("%s был повышен до должности: %s\n", e.Name, e.Position)
}
func RunLab6() {
  hireDate, _ := time.Parse("2006-01-02", "2020-05-15")
  employee := NewEmployee("Алексей Иванов", "Инженер", hireDate)
  employee.DisplayInfo()

  yearsOfService := employee.GetYearsOfService()
  fmt.Printf("Стаж работы: %d лет\n", yearsOfService)

  employee.Promote("Старший инженер")
}
