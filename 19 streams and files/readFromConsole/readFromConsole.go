package readFromConsole

import "fmt"

func ReadFromConsole() {
	var name string
	var age int

	fmt.Println("Введите имя")
	fmt.Scan(&name)

	fmt.Println("Введите возраст")
	fmt.Scan(&age)

	fmt.Println(name, age)
}
