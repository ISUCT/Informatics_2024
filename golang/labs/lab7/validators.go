package lab7

import "fmt"

func validatePrice(newPrice, currentPrice float64) error {
	if newPrice < 0 {
		return fmt.Errorf("заданная цена меньше нуля")
	} else if newPrice == currentPrice {
		return fmt.Errorf("заданное значение равно исходному")
	}
	return nil
}

func validateDiscount(discount float64) error {
	if discount < 0 {
		return fmt.Errorf("заданная скидка меньше нуля")
	} else if discount > 100 {
		return fmt.Errorf("заданная скидка больше ста")
	}
	return nil
}
