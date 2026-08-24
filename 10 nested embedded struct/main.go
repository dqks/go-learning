package main

import "fmt"

type adress struct {
	street, house, apartment_number string
}

// nested structure
type userNested struct {
	login, email, password string
	age                    uint
	adress_info            adress
}

// embedded structure
// Определяем анонимное поле, указывая тип
// Теперь можно обращаться через adress
type userEmbedded struct {
	login, email, password string
	age                    uint
	adress
}

type node struct {
	value int
	next  *node
}

func main() {
	// nested structure
	maxNested := userNested{
		login:    "Max",
		email:    "max@mail.ru",
		password: "1234",
		adress_info: adress{
			street:           "Pushkin",
			house:            "19",
			apartment_number: "2",
		},
	}
	fmt.Printf("%+v", maxNested)

	fmt.Println()

	// embedded structure
	maxEmbedded := userEmbedded{
		login:    "Max",
		email:    "max@mail.ru",
		password: "3123",
		age:      10,
		adress:   adress{"", "", ""},
	}
	fmt.Printf("%+v", maxEmbedded)
	fmt.Println()

	// Структура, которая ссылается на саму себя
	// Делается через указатель (иначе ошибка)
	first := node{value: 1}
	second := node{value: 2}
	third := node{value: 3}

	first.next = &second
	second.next = &third

	var current *node = &first
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}

}
