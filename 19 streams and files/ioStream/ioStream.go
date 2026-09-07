package ioStream

import (
	"fmt"
	"os"
)

func IoStream() {
	file, err := os.OpenFile("hello.txt", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Ошибка открытия файла")
		return
	}

	defer file.Close()

	// Запись можно осуществлять с помощью пакета fmt
	fmt.Fprintln(file, "Строка на запись")

	person := struct {
		name string
		age  int
	}{
		name: "Ivan",
		age:  20,
	}

	// Запись может быть форматированной
	fmt.Fprintf(file, "Человек %s, возраст %d", person.name, person.age)

}
