package lab7

import (
	"fmt"
)

func RunLab7() {

	phone := &Electronic{name: "iPhone 14", price: 1000, description: "New iPhone"}
	tshirt := &Clothing{name: "T-Shirt", price: 20, description: "Cotton T-Shirt", size: "M"}
	phone.SetDescription("New iPhone 14 Pro Max")
	tshirt.SetDescription("Cotton T-Shirt with white logo")

	products := []Product{phone, tshirt}

	fmt.Println("Total price before discount:", CalculateTotalPrice(products))

	phone.ApplyDiscount(0.1)
	tshirt.ApplyDiscount(0.05)

	fmt.Println("Total price after discount:", CalculateTotalPrice(products))

	tshirt.SetName("New T-Shirt")
	tshirt.SetPrice(25)
	fmt.Println("Price of T-Shirt after changes:", tshirt.GetPrice())
	fmt.Println("Name of T-Shirt after changes:", tshirt.GetName())
	fmt.Println("Description of T-Shirt after changes:", tshirt.GetDescription())
	fmt.Println("Phone Characreristics after changes:", phone.GetDescription())

}
