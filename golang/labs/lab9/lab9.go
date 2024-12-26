package lab9

type Task struct {
	Notion      string `json:"notion"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

func RunLab9() {
	ChooseTypeNotion()
}
