package lab7

type motorbike struct {
	name string
	price uint
	discount uint
	catigories string
}

func (m *motorbike) setName(newName string) {
	m.name = newName
}

func (m *motorbike) setPrice(newPrice uint) {
	m.price = newPrice
}

func (m *motorbike) setDiscount(newDiscount uint) {
	m.discount = newDiscount
}

func (m *motorbike) setCatigories(newCatigories string) {
	m.catigories = newCatigories
}

func (m motorbike) getInfo() (string, uint, uint, string) {
	return m.name, m.price, m.discount, m.catigories
}