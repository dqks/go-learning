package main

import "fmt"

func main() {
	var stringVar = "Maxim"
	var intNum int8 = 13
	var booleanVar bool = true

	fmt.Printf("stringVar type - %T, value - %s\n", stringVar, stringVar)
	fmt.Printf("intNum type - %T, value - %d\n", intNum, intNum)
	fmt.Printf("booleanVar type - %T, value - %t", booleanVar, booleanVar)
}
