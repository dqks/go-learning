package iso_handling

// Используем interface{} т.к. им может выступать любой тип
func Factorial(num int) (int, interface{}) {

	if num <= 0 {
		// Если выполняется услоие, то возвращаем 0 и строку с ошибкой
		return 0, "Ошибка. Факториала меньше нуля не существует"
	}

	result := 1

	for i := 1; i <= num; i++ {
		result *= i
	}

	return result, nil

}
