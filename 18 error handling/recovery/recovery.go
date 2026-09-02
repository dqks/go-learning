package recovery

import "fmt"

// Функция recover нужна для обработки ошибок, вызванных panic
// И используется только внутри defer
// recover() возвращает значение, переданное через panic

func zeroDivisionHandler() {
	// С помощью recover получаем значение из panic()
	// И проверяем что ошибка есть
	if result := recover(); result != nil {
		fmt.Println("Error:", result)
	}
}

func Divide(n1, n2 float64) float64 {
	defer zeroDivisionHandler()
	if n2 == 0 {
		panic("Division by zero")
	}
	return n1 / n2
}
