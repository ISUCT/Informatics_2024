package Lab7

import ( 
	"fmt"
)

type product struct {
	name     string
	quantity int
	price    float64
}

var t1 = product{name: "book", quantity: 120, price: 70}
var t2 = product{price: 40, quantity: 30, name: "sweater"}
var t3 = product{quantity: 100, name: "boots", price: 110}

func Lab7() {

	fmt.Println()
	fmt.Println()
}
