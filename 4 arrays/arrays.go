package main

import "fmt"

func main() {
	array1 := [...]int8{1, 2, 3}
	array2 := [...]int8{2, 2, 3}

	fmt.Println(array1 == array2) // Сравниваются по значениям внутри самого массива, порядок важен

	//Массив можно также обозначить так, при этом ключом может быть только int:
	stringArray := [3]string{2: "first element", 0: "second element", 1: "thirnd element"}
	fmt.Println(stringArray[2])

	twoDimensionNums := [3][2]int{
		{1, 2},
		{4, 5},
		{7, 8},
	}
	fmt.Println(twoDimensionNums)

	// Для получения длины массива используется фукнция len()
	fmt.Println(len(twoDimensionNums))

	// В Go массив - это значение, а не ссылка
	// В таком случае, если 1 массиву присвоить другой массив,
	// То в него скопируются значения

	fmt.Println("")

	array3 := array1

	array3[0] = 50

	fmt.Println(array1)
	fmt.Println(array3)

}
