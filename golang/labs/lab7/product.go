package lab7

type product interface{
	setName(string)
	setPrice(uint)
	setDiscount(uint)
	setCatigories(string)
	getInfo() (string, uint, uint, string)
}