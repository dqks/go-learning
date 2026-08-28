package main

import (
	"fmt"
	"strings"
)

func main() {
	// В Go строка - это слайс байтов
	myString := "Строка"
	fmt.Println(myString)
	fmt.Println(len(myString))

	binary_str := []byte{208, 159, 209, 128, 208, 184, 208, 178, 208, 181, 209, 130, 32, 208, 188, 208, 184, 209, 128}
	str := string(binary_str)

	fmt.Println(str) // Привет мир

	// Руна - это псевдоним для int32
	var myRune = 'd'
	myRune1 := 0x41f
	myRunes := []rune{0x41f, 0x41f, 0x41f}
	fmt.Println(myRune)
	fmt.Println(myRune1)
	fmt.Println(string(myRunes))

	// Выдаст ошибку
	//myString[1] = 'a'
	// Но если нужно поменять строку то можно ее изменить
	// Переведя ее в срез рун и поменять по индексам
	stringToChange := "Мир"
	runes := []rune(stringToChange)
	runes[0] = 'П'
	stringToChange = string(runes)
	fmt.Println(stringToChange)

	helloWorld := "Hello World"

	fmt.Println(strings.ToUpper(helloWorld))
	fmt.Println(strings.ToLower(helloWorld))
	fmt.Println(strings.HasSuffix(helloWorld, "Hell"))
	fmt.Println(strings.HasSuffix(helloWorld, "ld"))
	fmt.Println(strings.Contains(helloWorld, " "))
	fmt.Println(strings.Count(helloWorld, "l"))

	sliceString := []string{"Hello", "World"}
	fmt.Println(strings.Join(sliceString, " "))
}
