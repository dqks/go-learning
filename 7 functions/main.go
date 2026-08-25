package main

import "fmt"

func main() {
	var numMain int8 = 10
	// Передается по значению
	inc(numMain)
	fmt.Println(numMain, " after inc func")

	// Можно использовать неограниченное количество аргументов
	fmt.Println(sumNums(1, 2, 3, 4, 5, 6, 7, 8), " after sum")
	// Но при этом если хотим передать массив, то нужно в его конец добавлять ...
	// Для того, чтобы раскрыть сам массив
	fmt.Println(sumNums([]int8{1, 2, 3, 10, 10, 10}...), " after sum")

	fmt.Println(arrSum([3]int8{10, 20, 30}), " sum of array")
	// Даже если мы передад
	fmt.Println(arrSum([...]int8{10, 20, 30}), " sum of array")

	fmt.Println(concat("Hello", "World"), " after string concat")

	// Из функции, которая возвращает несколько значений, можно через запятую записать значения в переменную
	var num, name = add(5, 10, "Maxim", "Shtin")
	fmt.Println(num, " ", name)

	// У функции также есть тип
	// Например, у фукнции inc ниже тип такой: func(int8)
	// А у sum - func(...int8) int8
	// В переменную можно записать функцию
	var funcInc func(int8) = inc
	funcInc(28)

	//Также функцию можно передавать в качестве параметра в другую функцию
	// При этом, логично, мы не сможем записаь функцию ДРУГОГО типа в эту переменную
	var multFunc = selectFunc(2)
	fmt.Println(multFunc(2, 10000))

	numInc := outerFunc()
	fmt.Println(numInc())
	fmt.Println(numInc())
	fmt.Println(numInc())

	certainMultiplier := multiplyByNum(5)
	fmt.Println(certainMultiplier(1))
	fmt.Println(certainMultiplier(2))
	fmt.Println(certainMultiplier(3))
	fmt.Println(certainMultiplier(4))

	fmt.Println()
	fmt.Println(factorial(5))
}

func inc(num int8) {
	num += 20
	fmt.Println(num, " inside inc func")
}

func sumNums(nums ...int8) int8 {
	var sum int8 = 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func arrSum(num [3]int8) int8 {
	var sum int8 = 0
	for _, n := range num {
		sum += n
	}
	return sum
}

func concat(s1, s2 string) (s string) {
	s = s1 + s2
	// Вернет s
	return
}

// Можно вернуть сразу несколько значений из функции
func add(x, y int, firstName, lastName string) (num int, fullName string) {
	num = x + y
	fullName = firstName + " " + lastName
	return
}

func sum(x, y int) int {
	return x + y
}

func mult(x, y int) int {
	return x * y
}

// Функцию можно возвращать из другой функции
func selectFunc(n int8) func(int, int) int {
	if n == 1 {
		return sum
	} else {
		return mult
	}
}

// Также можно возвращать анонимные функции / лямбда функции
// Также их можно присваивать в переменные и передавать в качестве аргумента
// Как и все обычные функции
func selectFuncAnon(n int8) func(int, int) int {
	if n == 1 {
		return func(x, y int) int { return x + y }
	} else {
		return func(x, y int) int { return x * y }
	}
}

func outerFunc() func() int {
	num := 10
	return func() int {
		num++
		return num
	}
}

func multiplyByNum(num1 int) func(int) int {
	return func(num2 int) int { return num1 * num2 }
}

func factorial(num int) int {
	if num == 1 {
		return 1
	}

	return num * factorial(num-1)
}
