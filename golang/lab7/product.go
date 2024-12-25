package lab7

type Product interface {
	Sale(discount float64) error
	Change_price(new_price float64)
	Change_specifications(s1, s2, s3 string)
	Get_price() float64
	Get_info() [3]string
}