package main

import "fmt"

// Интерфейс для всех животных
type Animal interface {
	MakeSound() string
	Move() string
	GetInfo() string
	GetType() string
}

// Интерфейс для животных, которые умеют плавать
type Swimmer interface {
	Swim() string
	CanSwim() bool
}

// Лиса
type Fox struct {
	Name string
}

func (f Fox) MakeSound() string {
	return "Тяф-тяф!"
}

func (f Fox) Move() string {
	return "Лиса хитро бежит"
}

func (f Fox) GetInfo() string {
	return fmt.Sprintf("Лиса %s - хитрая и умная!", f.Name)
}

func (f Fox) GetType() string {
	return "Млекопитающее"
}

func (f Fox) Bark() string {
	return "Тяф-тяф-тяф!"
}

func (f Fox) CanSwim() bool {
	return false
}

// Енот
type Raccoon struct {
	Name string
}

func (r Raccoon) MakeSound() string {
	return "Кырк-кырк!"
}

func (r Raccoon) Move() string {
	return "Енот лазает по деревьям"
}

func (r Raccoon) GetInfo() string {
	return fmt.Sprintf("Енот %s - милый и проказливый!", r.Name)
}

func (r Raccoon) GetType() string {
	return "Млекопитающее"
}

func (r Raccoon) Chatter() string {
	return "Кырк-кырк-кырк!"
}

func (r Raccoon) CanSwim() bool {
	return true
}

func (r Raccoon) Swim() string {
	return "Енот плавает, чтобы перейти реку"
}

// Слон
type Elephant struct {
	Name string
}

func (e Elephant) MakeSound() string {
	return "Трамп-трамп-трамп!"
}

func (e Elephant) Move() string {
	return "Слон ходит медленно и мощно"
}

func (e Elephant) GetInfo() string {
	return fmt.Sprintf("Слон %s - самый большой наземный млекопитающий!", e.Name)
}

func (e Elephant) GetType() string {
	return "Млекопитающее"
}

func (e Elephant) Trumpet() string {
	return "Громкая труба!"
}

func (e Elephant) CanSwim() bool {
	return false
}

// Лягушка
type Frog struct {
	Name string
}

func (f Frog) MakeSound() string {
	return "Ква-ква!"
}

func (f Frog) Move() string {
	return "Лягушка прыгает"
}

func (f Frog) GetInfo() string {
	return fmt.Sprintf("Лягушка %s - земноводное!", f.Name)
}

func (f Frog) GetType() string {
	return "Земноводное"
}

func (f Frog) Ribbit() string {
	return "Ква-ква-ква!"
}

func (f Frog) CanSwim() bool {
	return true
}

func (f Frog) Swim() string {
	return "Лягушка плавает"
}

// Заяц
type Hare struct {
	Name string
}

func (h Hare) MakeSound() string {
	return "Пш-пш-пш!"
}

func (h Hare) Move() string {
	return "Заяц быстро бегает"
}

func (h Hare) GetInfo() string {
	return fmt.Sprintf("Заяц %s - пушистый и быстрый!", h.Name)
}

func (h Hare) GetType() string {
	return "Млекопитающее"
}

func (h Hare) Hop() string {
	return "Заяц прыгает"
}

func (h Hare) CanSwim() bool {
	return false
}

func main() {
	lisa := Fox{"Лиса"}
	enot := Raccoon{"Енот"}
	slon := Elephant{"Слон"}
	lyagushka := Frog{"Лягушка"}
	zayac := Hare{"Заяц"}

	animals := []Animal{lisa, enot, slon, lyagushka, zayac}

	for _, animal := range animals {
		fmt.Println(animal.GetInfo())
		fmt.Println("Звук:", animal.MakeSound())
		fmt.Println("Движение:", animal.Move())
		fmt.Println("Тип:", animal.GetType())
		if swimmer, ok := animal.(Swimmer); ok {
			fmt.Println("Умеет плавать:", swimmer.CanSwim())
			fmt.Println("Способ плавания:", swimmer.Swim())
		}
		fmt.Println("----")
	}
}
