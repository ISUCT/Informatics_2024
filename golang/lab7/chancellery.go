package lab7
 
type Chancellery struct { 
 name     string 
 price    float64 
 color    string 
 quantity int 
} 
 
func (c *Chancellery) GetName() string { 
 return c.name 
} 
 
func (c *Chancellery) GetPrice() float64 { 
 return c.price 
} 
 
func (c *Chancellery) SetPrice(price float64) { 
 c.price = price 
} 
 
func (c *Chancellery) ApplyDiscount(discount float64) { 
 c.price -= c.price * discount / 100 
}