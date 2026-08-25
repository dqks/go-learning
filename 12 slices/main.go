package main

import "fmt"

func main() {
	// Самый простой срез
	var people []string = []string{"Ivan", "Sergey", "Sasha"}
	fmt.Println(people)

	people2 := []string{"Ivan", "Sergey", "Sasha"}
	fmt.Println(people2)

	// Срез можно создавать из массива
	numbersArr := [8]string{"1", "2", "3", "4", "5", "6", "7", "8"}

	numbersSlice1 := numbersArr[2:6] // отбираем элементы МЕЖДУ индексами 2 и 6
	fmt.Println(numbersSlice1)       // [3 4 5 6]
	numbersSlice2 := numbersArr[:6]
	fmt.Println(numbersSlice2) // [1, 2, 3, 4, 5, 6]
	numbersSlice3 := numbersArr[4:]
	fmt.Println(numbersSlice3) // [5, 6, 7, 8]

	// Срез - это указатель, он ссылается на массив или раздел базового массива
	// В таком случае, если мы изменим базовый массив, то и срез тоже поменяется
	numbersArr[3] = "453"
	fmt.Println(numbersSlice1, "numbersSlice1 after change") // [3 453 5 6]

	// Срез можно создавать из строки
	hello := "Hello World!"
	helloSlice := hello[5:]
	fmt.Println(helloSlice) // " World"
	//hello = "Hello World Again"
	//fmt.Println(helloSlice) // " World"

	// Срез можно создавать из другого среза
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := slice1[2:5] // 3 4 5
	fmt.Println(slice2, "slice2")

	// Перебираем и обращаемся к элементам среза также как и в массиве
	fmt.Println(slice1[0], " slice1[0]")
	for _, n := range slice2 {
		fmt.Println(n, " inside of loop")
	}

	// Для получения длины среза используем len()
	fmt.Println(len(slice2), " len(slice2)")

	// Емкость среза
	// Это то, на сколько можно расширить срез
	// В данном случае этго 4
	// Т.к. до максимальной длинмы базового массива не хватает 1 элемента
	fmt.Println(cap(slice2), "cap(slice2)") // 4

	// Срез можно создать с помощью функции make()
	sliceMake := make([]int, 9)
	sliceMake[0] = 1
	sliceMake[2] = 3
	sliceMake[8] = 2

	// Добавление в срез
	sliceMake = append(sliceMake, 10)
	fmt.Println(sliceMake)

	// Удаление элемента из среза
	sliceMake = append(sliceMake[:2], sliceMake[3:]...) // Удалили 3
	fmt.Println(sliceMake)

	// Копирование элементов
	sliceCopy := []int{1, 2, 3, 4, 5}
	sliceToCopy := []int{}
	copy(sliceToCopy, sliceCopy)
	fmt.Println(sliceToCopy, "sliceToCopy") // Пустой

	// Копирование происходит по заполнению элементами среза
	// Мы не копируем длину исходного среза
	// Если длина исходного среза больше, то срез в который мы копируем элементы заполнится полностью
	sliceToCopy2 := make([]int, 3)
	copy(sliceToCopy2, sliceCopy)
	fmt.Println(sliceToCopy2, "sliceToCopy2")

	sliceToCopy3 := make([]int, 10)
	copy(sliceToCopy3, sliceCopy)
	fmt.Println(sliceToCopy3, "sliceToCopy3") //  Значения, которые не копируются из другого среза остаются как были

}
