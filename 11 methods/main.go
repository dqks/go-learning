package main

import "fmt"

type names []string

// (n names) - ресивер, получатель
// n - имя ресивера
// names - тип ресивера, к которому нужно прицепить метод
func (n names) print() {
	for _, v := range n {
		fmt.Print(v, "\t")
	}
}

// В n копируется значение, не по ссылке передается
// Т.е. вне этого метода срез namesSlice не поменяется
func (n names) add(name string) {
	n = append(n, name)
	fmt.Println(n)
}

type user struct {
	firstName, lastName, patronymic string
	age                             uint
}

// Испольщуем указатель по причине описанной ниже
func (u *user) getFullName() string {
	return u.lastName + " " + u.firstName + " " + u.patronymic
}

// Когда мы в методе структуры мы принимаем указатель на структуру
// При вызовые Go сам подставляет адрес
// Тут IDE жалуется на то, что есть метод с user а есть метод *user
// Если есть хотя бы метод *user, то стоит для всех методов и дальше прописывать *user
// Если даже это нигде не используется
func (u *user) setAge(age uint) {
	u.age = age
}

func main() {
	// Метод для псевдонима
	namesSlice := names{"Ivan", "Sergey", "Alex"}
	namesSlice.print()
	fmt.Println()
	namesSlice.add("Vanya")
	namesSlice.print()
	fmt.Println()

	// Метод для структуры
	user := user{firstName: "Максим", lastName: "Штин", patronymic: "Олегович"}
	fmt.Println(user.getFullName())
	fmt.Println(user.age) // 0
	user.setAge(20)
	fmt.Println(user.age) // 20
}
