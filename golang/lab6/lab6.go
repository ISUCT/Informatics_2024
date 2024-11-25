package lab6

import (
	"fmt"
)

type Document struct {
	name     string
	type_doc string
	wight    float64
	lenth    float64
}

func (d *Document) SetName(new_name string) {
	d.name = new_name
}
func (d *Document) SetType(new_type string) {
	d.type_doc = new_type
}
func (d *Document) SetSize(new_wight, new_lenth float64) {
	d.wight = new_wight
	d.lenth = new_lenth
}
func (d Document) GetName() string {
	return d.name
}
func (d Document) get_aboutdoc() [3]string {
	return [3]string{d.type_doc, (fmt.Sprintf("%g", d.wight)), (fmt.Sprintf("%g", d.lenth))}
}
func RunLab6() {
	var doc = new(Document)
	doc.SetName("art")
	doc.SetSize(128, 128)
	doc.SetType("png")

	fmt.Println("Имя: " + doc.GetName() + "\nТип файла: " + doc.get_aboutdoc()[0] + "\nРазмер: " + doc.get_aboutdoc()[1] + "x" + doc.get_aboutdoc()[2] + "\n")
}
