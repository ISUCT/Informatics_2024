package labs

import (
	"fmt"
)

type Document struct {
	name         string
	type_doc     string
	wight, lenth float64
}

func (d *Document) set_name(new_name string) {
	d.name = new_name
}
func (d *Document) set_size(new_wight, new_lenth float64) {
	d.wight = new_wight
	d.lenth = new_lenth
}
func (d *Document) set_type(new_type string) {
	d.type_doc = new_type
}
func (d Document) get_name() string {
	return d.name
}
func (d Document) get_aboutdoc() [3]string {
	return [3]string{d.type_doc, (fmt.Sprintf("%g", d.wight)), (fmt.Sprintf("%g", d.lenth))}
}
func Go_lab6() {
	var doc = new(Document)
	doc.set_name("фото")
	doc.set_size(3, 4)
	doc.set_type("png")
	name := doc.get_name()
	var aboutdoc = doc.get_aboutdoc()
	var document = map[string]interface{}{name: aboutdoc}
	fmt.Println("\n1:", document)
	fmt.Println("2:", document[name])
	fmt.Println("3:\nТип файла: " + aboutdoc[0] + "\nРазмер файла: " + aboutdoc[1] + "x" + aboutdoc[2]+"\n")
}
