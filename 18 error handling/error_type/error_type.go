package error_type

import (
	"fmt"
)

// В Go есть тип error
// Он представляет собой интерфейс с методом Error()
// Как и для всех интерфейсов, если тип реалзиует его
// Метод, то тогда он сможет использоваться вместо него
type paramError struct{}

// Реализуем для структуры метод Error()
func (p paramError) Error() string {
	return "Invalid parameter"
}

func ShowParamError() {
	obj := paramError{}
	// Можем вызвать так ошибку
	fmt.Println(obj.Error())
	// А можем и через сокращенную форму
	fmt.Println(paramError{})
}

// Теперь мы можем определить функцию факториала
// Используя новый тип paramError
func Factorial(num int) (int, error) {

	if num <= 0 {
		// Если выполняется условие, то возвращаем 0
		// и вызов Error() у paramError
		return 0, paramError{}

		// Либо можно воспользоваться встроенной библиотекой Go
		// errors.New вернет ошибку типа интерфейса error
		// С сообщением об ошибке
		// return 0, errors.New("Invalid parameter")
	}

	result := 1

	for i := 1; i <= num; i++ {
		result *= i
	}

	return result, nil

}
