package lab6

type Hero struct {
	name  string
	class string
	lvl   float64
}

func (h *Hero) SetName(name string) {
	h.name = name
}

func (h *Hero) GetName() string {
	return h.name
}

func (h *Hero) SetClass(class string) {
	h.class = class
}

func (h *Hero) GetClass() string {
	return h.class
}

func (h *Hero) SetLvl(lvl float64) {
	h.lvl = lvl
}

func (h *Hero) GetLvl() float64 {
	return h.lvl
}
