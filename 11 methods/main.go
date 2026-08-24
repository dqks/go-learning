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

func main() {
	namesSlice := names{"Ivan", "Sergey", "Alex"}
	namesSlice.print()
	fmt.Println()
	namesSlice.add("Vanya")
	namesSlice.print()

	//

}
