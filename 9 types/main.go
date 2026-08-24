package main

import "fmt"

// Определяем новый тип, фактически это и есть uint, но с новым названием
// Тип mile "играет" по правилам uint
type mile uint

// funcInt название с маленько потому что внутри 1 пакета
// Если бы мы его экспортировали, то было бы FuncInt
type funcInt func(int, int) int

func action(n1, n2 int, operation funcInt) int {
	fmt.Println("Вывод из функции action")
	return operation(n1, n2)
}

func add(x, y int) int {
	return x + y
}

func incAge(user *person) {
	user.age += 1
}

// Можно определяеть структуру
// Если тип у полей структуры одинаковый, то их можно перечислить через запятую и в конце написать их тип
type person struct {
	firstName, lastName, patronymic string
	age                             uint8
	company                         string
}

func main() {
	// Этот новый тип можно использовать также как и все обычные типы
	var milesCount mile = 302
	fmt.Println(milesCount)

	// Вместо того, чтобы писать func(int, int) int
	// При определении функции и в типе переменной мы пишем funcInt
	var sumOperation funcInt = add
	result := action(10, 20, sumOperation)
	fmt.Println(result)

	// Определим переменную структуры
	var maxim person
	fmt.Println(maxim)
	maxim.firstName = "Maxim"
	maxim.lastName = "Brady"
	maxim.patronymic = "Adolfovich"
	maxim.company = "Google"
	maxim.age = 10
	fmt.Println(maxim)

	var edmon person = person{"Edmon", "Narmetov", "Adolfovich", 100, "Amazon"}
	fmt.Println(edmon)

	artur := person{}
	fmt.Println(artur)

	// Можно также определить анонимную структуру
	// И тут же ее инициализировать
	tom := struct {
		name string
		age  uint
	}{
		name: "Tom",
		age:  20,
	}
	fmt.Println(tom)

	//
	var bob struct {
		name string
		age  uint
	}

	var p_bob = &bob
	fmt.Println(p_bob.age, "возраст Боба без разыменовывания")
	fmt.Println((*p_bob).age, "возраст Боба с разыменовыванием")

	var p_age *uint = &p_bob.age
	fmt.Println(p_age, "адрес на возраст Боба")
	fmt.Println(*p_age, "возраст Боба ")

	//
	robin := new(person)
	robin.firstName = "Robin"

	// Оба вывода покажут одно и тоже
	// Потому что robin.age - это синтаксический сахар, компилятор сам прописывает (*robin).age
	fmt.Println(robin.age)
	fmt.Println((*robin).age)

	fmt.Println(robin, "адрес Робина")
	fmt.Println(*robin, "значение Робина")

	// Копирование структуры
	alice := person{age: 40}
	dasha := alice
	dasha.age = 20
	// Таким образом, копирование происходит не по ссылке, а по значению
	fmt.Println(alice, "alice") // 40 alice
	fmt.Println(dasha, "dasha") // 20 dasha

	//
	user := person{age: 10}
	// Передаем адрес, чтобы вне фукнции поменялась структура
	// Иначе передается копия структуры
	incAge(&user)
	fmt.Println(user)

	// Сравнение структуры
	user1 := person{}
	user2 := person{}
	fmt.Println(user1 == user2) // true
	user1.age = 10
	fmt.Println(user1 == user2) // false

	type p1 struct {
		name string
		age  uint
	}

	type p2 struct {
		name string
		age  uint
	}

	// user3 := p1{}
	// user4 := p2{}
	// fmt.Println(user3 == user4) // mismatched types
}
