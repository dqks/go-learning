package main

import "fmt"

func main() {
	var1 := 5
	var2 := 3

	if var1 == var2 {
		fmt.Println(true, "var1 == var2")
	} else {
		fmt.Println(false, "var1 != var2")
	}

	array1 := [3]int{1, 2, 3}
	var array2 = [3]int{1, 2, 3}

	if array1 == array2 {
		fmt.Println(true, "array1 == array2")
	} else {
		fmt.Println(false, "array1 != array2")
	}

	// var num int8 =  43; можно намеренно указать тип
	// var num = 43 можно автоматически задать тип
	num := 43 // только автоматический тип

	switch num {
	case 43:
		fmt.Println(43)
		fallthrough // с помощью него можно выполниь текущий блок условия а также спуститься на 1 ниже
	case 44:
		fmt.Println(44)
	case 45:
		fmt.Println(45)
	}

}
