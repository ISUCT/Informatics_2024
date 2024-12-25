package lab7

import "errors"

var ErrMaterial error = errors.New("invalid material")

type Material string

const (
	Skin    Material = "кожа"
	Wood    Material = "дерево"
	Diamonds Material = "бриллианты"
	Dsp     Material = "дсп"
)

type mebel struct {
	product
	material Material
}

func (f *mebel) SetMaterial(material Material) error {
	for _, standardMaterial := range []Material{Skin, Wood, Diamonds, Dsp} {
		if material == standardMaterial {
			f.material = material
			return nil
		}
	}
	return ErrMaterial
}

func NewMebel(id int, name string, price float64, material Material) *mebel {
	f := &Mebel{
		product: newProduct(id, name, price),
	}
	f.SetMaterial(material)
	return f
}
