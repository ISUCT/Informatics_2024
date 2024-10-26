package laba6

import "fmt"

type Movie struct {
	Title    string
	Director string
	Duration int
}

func (m *Movie) setTitle(newTitle string) {
	m.Title = newTitle
}

func (m Movie) getTitle() string {
	return m.Title
}

func (m *Movie) setDirector(newDirector string) {
	m.Director = newDirector
}

func (m Movie) getDirector() string {
	return m.Director
}

func (m *Movie) setDuration(newDuration int) {
	m.Duration = newDuration
}

func (m Movie) getDuration() int {
	return m.Duration
}

func Runlab6() {
	movie := Movie{Title: "Гарри Поттер и Филасовский камень", Director: "Джоан Роулинг", Duration: 120}

	fmt.Printf("Название фильма %s\n", movie.getTitle())
	movie.setTitle("Гарри Поттер и Узник Аскабана")
	fmt.Printf("Название фильма %s\n", movie.getTitle())

	fmt.Printf("Автор Фильма %s\n", movie.getDirector())
	movie.setDirector("Джон Роулинг")
	fmt.Printf("Автор Фильма %s\n", movie.getDirector())

	fmt.Printf("Длительность Фильма %d\n", movie.getDuration())
	movie.setDuration(130)
	fmt.Printf("Длительность Фильма %d\n", movie.getDuration())
}
