package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}

	//

	fmt.Println()

	// Перебор элементов массива
	var array = [3]int{1, 2, 3}
	for i, val := range array {
		fmt.Println(i, val)
	}

	//

	fmt.Println()

	// Перебор элементов строки
	// будет выводить код символа
	hello := "Hello"
	for i, c := range hello {
		fmt.Println(i, c)
	}

	fmt.Println()

	for i, c := range hello {
		fmt.Printf("%d, %c\n", i, c)
	}

	//

	fmt.Println()

	// В Go также есть операторы break и continue

	array1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	for _, n := range array1 {
		if n == 5 {
			break
		}

		fmt.Println(n)
	}

}
