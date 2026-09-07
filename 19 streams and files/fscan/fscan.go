package fscan

import (
	"fmt"
	"os"
)

func WriteData() {
	name := "Tom"
	age := 23

	file, err := os.OpenFile("hello2.txt", os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println("Ошибка открытия файла")
		return
	}
	defer file.Close()

	fmt.Fprintln(file, name)
	fmt.Fprintln(file, age)
}

func ReadData() {
	var name string
	var age int

	file, err := os.OpenFile("hello2.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Ошибка открытия файла")
		return
	}
	defer file.Close()

	fmt.Fscanln(file, &name)
	fmt.Fscanln(file, &age)

	fmt.Println(name, "name")
	fmt.Println(age, "age")
}
