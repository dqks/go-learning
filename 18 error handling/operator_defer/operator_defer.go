package operator_defer

import (
	"fmt"
	"os"
)

func finish() {
	fmt.Println("Finish")
}

func PrintLines() {
	// Выведется последней
	defer finish()
	// Выведется предпоследней
	defer fmt.Println("The last line")
	fmt.Println("First Line")
	fmt.Println("Second Line")
	// Если в функции несколько defer,
	// То последним выведется тот, который определен первее всех
}

// Реальный пример использования
// В конце выполнения функции закрываем файл main.go
func OpenFile() {
	file, err := os.Open("./main.go")
	if err == nil {
		defer file.Close()
	}
	fmt.Println(file)
	fmt.Println(err)
}
