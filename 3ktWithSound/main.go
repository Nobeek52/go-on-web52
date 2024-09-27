package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Интерфейс для всех животных
type Animal interface {
	MakeSound() (string, error) // Возвращаем ошибку для обработки
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

func (f Fox) MakeSound() (string, error) {
	return "Тяф-тяф!", nil
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

func (r Raccoon) MakeSound() (string, error) {
	return "Кырк-кырк!", nil
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

func (e Elephant) MakeSound() (string, error) {
	return "Трамп-трамп-трамп!", nil
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

func (f Frog) MakeSound() (string, error) {
	return "Ква-ква!", nil
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

func (h Hare) MakeSound() (string, error) {
	return "Пш-пш-пш!", nil
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

// Неизвестное животное -  генерирует ошибку при вызове MakeSound
type UnknownAnimal struct {
	Name string
}

func (u UnknownAnimal) MakeSound() (string, error) {
	return "", fmt.Errorf("Неизвестное животное не может издавать звуки")
}

func (u UnknownAnimal) Move() string {
	return "Движется неизвестным образом"
}

func (u UnknownAnimal) GetInfo() string {
	return fmt.Sprintf("Неизвестное животное %s", u.Name)
}

func (u UnknownAnimal) GetType() string {
	return "Неизвестный тип"
}

func (u UnknownAnimal) CanSwim() bool {
	return false
}

// Запись ошибок в файл логов
func logError(err error) {
	f, err := os.OpenFile("errors.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия лог-файла:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%sn", err))
	if err != nil {
		fmt.Println("Ошибка записи в лог-файл:", err)
		return
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите тип животного (лиса, енот, слон, лягушка, заяц, неизвестное) или 'exit' для выхода:")
		input, _ := reader.ReadString('\n') // Читаем строку, включая \n
		input = strings.TrimSpace(input)    // Удаляем пробелы и \n

		if input == "exit" {
			break
		}

		var animal Animal

		switch input {
		case "лиса":
			fmt.Println("Введите имя лисы:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name) // Удаляем пробелы и \n
			animal = Fox{Name: name}
		case "енот":
			fmt.Println("Введите имя енота:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name) // Удаляем пробелы и \n
			animal = Raccoon{Name: name}
		case "слон":
			fmt.Println("Введите имя слона:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name) // Удаляем пробелы и \n
			animal = Elephant{Name: name}
		case "лягушка":
			fmt.Println("Введите имя лягушки:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name) // Удаляем пробелы и \n
			animal = Frog{Name: name}
		case "заяц":
			fmt.Println("Введите имя зайца:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name) // Удаляем пробелы и \n
			animal = Hare{Name: name}
		case "неизвестное":
			fmt.Println("Введите имя неизвестного животного:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name) // Удаляем пробелы и \n
			animal = UnknownAnimal{Name: name}
		default:
			fmt.Println("Неверный тип животного. Попробуйте снова.")
			continue
		}

		// Вывод информации о животном
		fmt.Println(animal.GetInfo())

		// Вывод звука
		sound, err := animal.MakeSound()
		if err != nil { // Обработка ошибок при вызове MakeSound
			fmt.Println("Ошибка:", err)
			logError(err)
			continue
		}
		fmt.Println("Звук:", sound)

		// Вывод движения
		fmt.Println("Движение:", animal.Move())

		// Вывод типа
		fmt.Println("Тип:", animal.GetType())

		// Проверка умения плавать
		if swimmer, ok := animal.(Swimmer); ok {
			fmt.Println("Умеет плавать:", swimmer.CanSwim())
			fmt.Println("Способ плавания:", swimmer.Swim())
		}

		fmt.Println("----")
	}
}
