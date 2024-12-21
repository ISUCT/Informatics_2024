package lab7

type Product interface {
    getPrice() float64
    setPrice(newPrice float64)
    applyDiscount(discount float64)
    getName() string
    setName(newName string)
}