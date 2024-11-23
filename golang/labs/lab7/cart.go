package lab7

import "fmt"

type cart struct {
	products []cart_product
}

type cart_product struct {
	product
	count_product uint
}

func (c *cart) add(product product) {
	check := 0
	for index, products := range c.products {
		if products.product == product {
			c.products[index].count_product++
			check++
		}
	}
	if check == 0 {
		c.products = append(c.products, cart_product{product, 1})
	}
}

func (c *cart) purge(product product) {
	for index, products := range c.products {
		if products.product == product {
			if products.count_product == 1 {
				c.products = append(c.products[:index], c.products[index+1:]...)
			} else {
				c.products[index].count_product--
			}
		}
	}
}

func (c cart) getTotalSell() string {
	var sellWithoutDiscount uint
	var sellWithDiscount float64

	var sellWithoutDiscountString string
	var sellWithDiscountString string
	var saveMoney string

	for _, products := range c.products {
		_, price, discount, _ := products.product.getInfo()
		sellWithoutDiscount += price * products.count_product
		sellWithDiscount += float64(price * products.count_product) * (1-float64(discount)/100.0)
	}
	
	sellWithoutDiscountString = fmt.Sprintf("Конечная ценна без скидки: %d\n", sellWithoutDiscount)
	sellWithDiscountString = fmt.Sprintf("Со скидкой: %.0f\n", sellWithDiscount)
	if sellWithDiscount == float64(sellWithoutDiscount) {
		saveMoney = "Скидки применены не были"
	} else {
		saveMoney = fmt.Sprintf("Вы сэкономили %.0f!", float64(sellWithoutDiscount)-sellWithDiscount)
	}

	return sellWithoutDiscountString+sellWithDiscountString+saveMoney
}
