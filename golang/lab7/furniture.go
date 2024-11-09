package lab7

import "errors"

var ErrMaterial error = errors.New("invalid material")

type Material string

const (
	Skin    Material = "кожа"
	Wood    Material = "дерево"
	Plywood Material = "Фанера"
	Dsp     Material = "дсп"
)

type furniture struct {
	product
	material Material
}

func (f *furniture) SetMaterial(material Material) error {
	for _, standardmaterial := range []Material{Skin, Wood, Plywood, Dsp} {
		if material == standardmaterial {
			f.material = material
			return nil
		}
	}
	return ErrMaterial
}

func NewFurniture(id int, name string, price float64, material Material) *furniture {
	f := &furniture{
		product: newProduct(id, name, price),
	}
	f.SetMaterial(material)
	return f
}
