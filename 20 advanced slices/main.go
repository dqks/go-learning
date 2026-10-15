package main

import (
	"fmt"
)

func main() {
	// при передаче слайса по значению копируется header,
	// который представляет из себя len, cap, array pointer
	// при изменении элемента backing array в add изменения slice видны снаружи
	// но при добавлении в add меняется lenтолько внутри add
	// поэтому изменения снаружи len не видно, но по факту backing array изменился и снаружи
	// просто из-за "старого" значения len нам этих изменений не видно
	func() {
		slice := make([]int, 3, 10)
		add(slice)

		fmt.Println(slice, len(slice), "from main")   // [1000 0 0] 3 from main
		fmt.Println(slice[:4], "from main slice[:4]") // [1000 0 0 99] from main slice[:4]
		fmt.Println()

		// Из-за нового backing array изменения снаружи видны не будут
		slice2 := make([]int, 3, 3)
		add(slice2)
		fmt.Println(slice2, len(slice2), "from main") // [0 0 0] 3 from main
		// fmt.Println(slice2[:4], "from main slice[:4]") // ошибка
		fmt.Println()
	}()

	//
	func() {

		a := []int{1, 2, 3, 4}
		b := a[:2] // скопируется тот же backing array
		b = append(b, 99)
		fmt.Println(a, b) // [1 2 99 4] [1 2 99]
		c := a[:4]
		c = append(c, 99)
		fmt.Println(a, b, c) // [1 2 99 4] [1 2 99] [1 2 99 4 99]

		// в c при копировании len = 4 и cap = 4
		// поэтому при добавлении backing array поменяется
	}()

	// для копирования слайса slices.Clone
	// для удаления элемента

}
