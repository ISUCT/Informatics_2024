package todo

type Task struct {
	Name string `json:"Name"`
	Date int    `json:"Date"`
	Time int    `json:"Time"`
	Teg  string `json:"Teg"`
}

type Todo struct {
	List []Task `json:"list"`
}
