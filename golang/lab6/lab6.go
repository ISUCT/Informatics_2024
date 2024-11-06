package lab6

import "fmt"

type Document struct {
	Type   string
	Size   int
	Author string
}

func NewDocument(tYpe string, size int, author string) *Document {
	p := new(Document)
	p.Type = tYpe
	p.Size = size
	p.Author = author
	return p
}

func (p *Document) SetType(tYpe string) { p.Type = tYpe }

func (p *Document) UpdateStruct(new_type string, new_size int, new_author string) {
	p.Type = new_type
	p.Size = new_size
	p.Author = new_author
}

func (p Document) GetType() string   { return p.Type }
func (p Document) GetSize() int      { return p.Size }
func (p Document) GetAuthor() string { return p.Author }

func RunLab6Tasks() {
	doc := NewDocument("PDF", 120, "Максим Александров")
	doc.SetType("txt")
	fmt.Println(doc.GetType())
	fmt.Println(doc.GetSize())
	fmt.Println(doc.GetAuthor())
	doc.UpdateStruct("ZIP", 80, "Антон Иванов")
	fmt.Println(doc.GetType())
	fmt.Println(doc.GetSize())
	fmt.Println(doc.GetAuthor())
}
