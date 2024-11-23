package lab6

import (
	"errors"
	"fmt"
)

type GameCharacter struct {
	Name      string
	Class     string
	Level     int
	Health    int
	MaxHealth int
}

func NewGameCharacter(name, class string, level, health, maxHealth int) (*GameCharacter, error) {
	if name == "" || class == "" || level < 1 || health < 1 || maxHealth < 1 {
		return nil, errors.New("некорректные параметры")
	}

	return &GameCharacter{
		Name:      name,
		Class:     class,
		Level:     level,
		Health:    health,
		MaxHealth: maxHealth,
	}, nil
}

func (gc *GameCharacter) LevelUp() {
	gc.Level++
	gc.MaxHealth += 10
	gc.Health = gc.MaxHealth
}

func (gc *GameCharacter) TakeDamage(damage int) {
	if damage >= gc.Health {
		gc.Health = 0
	} else {
		gc.Health -= damage
	}
}

func (gc *GameCharacter) Heal(healing int) {
	if gc.Health+healing > gc.MaxHealth {
		gc.Health = gc.MaxHealth
	} else {
		gc.Health += healing
	}
}

func Completelaba() {
	character, err := NewGameCharacter("Артас", "Рыцарь смерти", 50, 100, 200)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Имя: %s, Класс: %s, Уровень: %d, Здоровье: %d/%d\n",
		character.Name, character.Class, character.Level, character.Health, character.MaxHealth)

	character.TakeDamage(80)
	fmt.Printf("После получения урона: Здоровье: %d/%d\n",
		character.Health, character.MaxHealth)

	character.Heal(30)
	fmt.Printf("После лечения: Здоровье: %d/%d\n",
		character.Health, character.MaxHealth)

	character.LevelUp()
	fmt.Printf("После повышения уровня: Имя: %s, Класс: %s, Уровень: %d, Здоровье: %d/%d\n",
		character.Name, character.Class, character.Level, character.Health, character.MaxHealth)
}
