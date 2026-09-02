package operator_panic

import "fmt"

// panic принимает произвольный тип данных
// И останавливает выполнение программы

func divideDefer() {
	fmt.Println("Defer divide")
}

func Divide(n1, n2 float64) float64 {
	defer divideDefer()
	if n2 == 0 {
		panic("Division by zero")
	}
	return n1 / n2
}
