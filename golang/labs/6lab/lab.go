package lab6

import (
 "fmt"
)

type Film struct {
 Title  string
 Genre  string
 Length int 
}

func NewFilm(title, genre string, length int) Film {
 return Film{
  Title:  title,
  Genre:  genre,
  Length: length,
 }
}

func (f Film) Play() {
 fmt.Printf("Идет просмотр фильма: %s (Жанр: %s, Длина: %d минут)\n", f.Title, f.Genre, f.Length)
}

func (f Film) Info() string {
 return fmt.Sprintf("Фильм: %s, Жанр: %s, Длина: %d минут", f.Title, f.Genre, f.Length)
}

func (f *Film) UpdateGenre(newGenre string) {
 f.Genre = newGenre
}

func Executelab6() {
 film := NewFilm("Властелин колец", "Фэнтези", 178)
 film.Play()
 fmt.Println(film.Info())
 film.UpdateGenre("Приключения")
 fmt.Println(film.Info())
}