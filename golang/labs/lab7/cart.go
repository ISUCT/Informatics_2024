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

func (c cart) itogSell() string {
	var sellWithoutDiscass uint
	var sellWithDiscass float64

	var sellWithoutDiscassString string
	var sellWithDiscassString string
	var saveMoney string

	for _, products := range c.products {
		_, price, discass, _ := products.product.getInfo()
		sellWithoutDiscass += price * products.count_product
		sellWithDiscass += float64(price) * float64(products.count_product) * float64((1-float64(discass)/100.0))
	}
	
	sellWithoutDiscassString = fmt.Sprintf("Конечная ценна без скидки: %d\n", sellWithoutDiscass)
	sellWithDiscassString = fmt.Sprintf("Со скидкой: %.0f\n", sellWithDiscass)
	if sellWithDiscass == float64(sellWithoutDiscass) {
		saveMoney = "Скидки применены не были"
	} else {
		saveMoney = fmt.Sprintf("Вы сэкономили %.0f!", float64(sellWithoutDiscass)-sellWithDiscass)
	}

	return sellWithoutDiscassString+sellWithDiscassString+saveMoney
}
