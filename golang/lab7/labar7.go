package main

import "fmt"

type Product interface {
	ApplyDiscount(discount float64) // Применить скидку
	GetDetails() string             // Получить детали товара
	GetPrice() float64              // Получить текущую цену
}

// Структура Линейка
type Ruler struct {
	Length int     // Длина в см
	Type   string  // Тип линейки (пластиковая, металлическая и т.п)
	Price  float64 // Цена
}

// Реализация интерфейса Product для Ruler
func (r *Ruler) ApplyDiscount(discount float64) { r.Price -= r.Price * discount / 100 }
func (r *Ruler) GetDetails() string {
	return fmt.Sprintf("Длина: %d см, Тип: %s, Цена: %.2f руб.", r.Length, r.Type, r.Price)
}
func (r *Ruler) GetPrice() float64 { return r.Price }

func CreateRuler(length int, rulerType string, price float64) *Ruler { // Функция для создания структуры
	return &Ruler{
		Length: length,
		Type:   rulerType,
		Price:  price,
	}
}

func CalculateTotalPrice(products []Product) float64 { // расчёт общей стоимости
	total := 0.0
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func main() {
	// создание товаров через функцию
	ruler1 := CreateRuler(30, "Пластиковая", 150.0)
	ruler2 := CreateRuler(50, "Металлическая", 250.0)

	//применяется скидка на товары
	ruler1.ApplyDiscount(10) // 10% скидка на первую линейку
	ruler2.ApplyDiscount(15) // 15% скидка на вторую линейку

	// Срез продуктов
	products := []Product{ruler1, ruler2}

	//Выводим информацию о каждом товаре
	for _, product := range products {
		fmt.Println(product.GetDetails())
	}
	// Рассчитываем общую стоимость
	total := CalculateTotalPrice(products)
	fmt.Printf("\nОбщая стоимость товаров: %.2f руб.\n", total)
}
