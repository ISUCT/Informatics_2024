package lab7

type ElectronicDevice struct {
	Name        string
	Price       float64
	OldPrice    float64
	Description string
	Brand       string
	Model       string
}

func (e *ElectronicDevice) GetName() string {
	return e.Name
}

func (e *ElectronicDevice) GetPrice() float64 {
	return e.Price
}

func (e *ElectronicDevice) GetDescription() string {
	return e.Description
}

func (e *ElectronicDevice) ApplyDiscount(discount float64) {
	e.OldPrice = e.Price
	e.Price = e.Price - (e.Price * discount)
}

func (e *ElectronicDevice) SetPrice(newPrice float64) {
	e.OldPrice = e.Price
	e.Price = newPrice
}

func (e *ElectronicDevice) SetDescription(newDescription string) {
	e.Description = newDescription
}

func (e *ElectronicDevice) GetOldPrice() float64 {
	return e.OldPrice
}
