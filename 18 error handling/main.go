package main

import (
	"error_handling/error_type"
	"error_handling/iso_handling"
	"error_handling/operator_defer"
	"error_handling/operator_panic"
	"error_handling/recovery"
	"fmt"
)

func main() {
	// -------------------------------------------------------------------
	// Из функций можно возвращать ошибки вторым возвращаемым значением
	fact1, err := iso_handling.Factorial(5)
	fmt.Println(fact1, err)

	fact2, err2 := iso_handling.Factorial(0)
	fmt.Println(fact2, err2)

	// -------------------------------------------------------------------
	fmt.Println()
	fmt.Println()
	// Тут мы вместо пустого интерфейса используем структуру, которая реализует
	// Интерфейс error1
	error_type.ShowParamError()
	fact3, err3 := error_type.Factorial(5)
	fmt.Println(fact3, err3)
	fact4, err4 := error_type.Factorial(0)
	fmt.Println(fact4, err4)

	// -------------------------------------------------------------------
	// Оператор defer позволяет отложить выполнение фукнции
	// В самый конец оборачивающей функции
	fmt.Println()
	fmt.Println()
	operator_defer.PrintLines()
	operator_defer.OpenFile()

	// -------------------------------------------------------------------
	fmt.Println()
	fmt.Println()
	fmt.Println(operator_panic.Divide(2, 5))
	// fmt.Println(operator_panic.Divide(2, 0))

	// -------------------------------------------------------------------
	fmt.Println()
	fmt.Println()
	fmt.Println(recovery.Divide(5, 2))
	fmt.Println(recovery.Divide(5, 0))
}
