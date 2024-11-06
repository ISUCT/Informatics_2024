package lab7

import "errors"

type furniture struct {
	product
	material string
}

func (f *furniture) SetMaterial(material string) error {
	for _, standardmaterial := range []string{"кожа", "дерево", "Фанера", "дсп"} {
		if material == standardmaterial {
			f.material = material
			return nil
		}
	}
	return errors.New("invalid material")
}

func NewFurniture(id int, name string, price float64, material string) *furniture {
	f := &furniture{
		product: newProduct(id, name, price),
	}
	f.SetMaterial(material)
	return f
}
