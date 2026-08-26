package main

import (
	"fmt"
)

type employee struct {
	firstName, lastName string
	age, salary         uint
}

func (e employee) getFullName() string {
	return e.firstName + " " + e.lastName
}

func (e *employee) increaseSalary(percent uint) {
	e.salary = e.salary + (e.salary * percent / 100)
}

func printEmployees(employees []employee) {
	fmt.Println("Сотрудники")
	for i, e := range employees {
		num := i + 1
		fmt.Printf("%d. %s, возраст: %d, зарплата: %d\n", num, e.getFullName(), e.age, e.salary)
	}
}

// Принимаем указатель на срез []employee
func findEmployeeByName(employees *[]employee, firstName string) *employee {
	for i := range *employees {
		// Обращаемся по индексу к элементу среза
		// А не через отдельную переменную
		if (*employees)[i].firstName == firstName {
			// Возврщааем ссылку на элемент ВСЕ ТОГО ЖЕ СРЕЗА
			// Который мы передаем в ЭТУ функцию
			return &(*employees)[i]
		}
	}
	return nil
}

func increaseSalary(e *employee, percent uint) {
	e.salary = e.salary + (e.salary * percent / 100)
}

func getAverageSalary(employees []employee) int {
	if len(employees) == 0 {
		return 0
	}
	sum := 0
	for _, e := range employees {
		sum += int(e.salary)
	}
	return sum / len(employees)
}

func getOldestEmployee(employees *[]employee) *employee {
	var oldestEmployee *employee = &(*employees)[0]
	for i := range *employees {
		if (*employees)[i].age > oldestEmployee.age {
			oldestEmployee = &(*employees)[i]
		}
	}
	return oldestEmployee
}

func main() {
	alex := employee{"Alex", "Shmitd", 0, 20000}
	sam := employee{"Sam", "Sulek", 10, 50000}
	bob := employee{"Robert", "Pollson", 0, 5000}

	employees := []employee{alex, sam, bob}
	fmt.Println(employees[0].getFullName())
	fmt.Println()

	employees[0].increaseSalary(10)
	fmt.Println(employees[0])
	fmt.Println()

	foundEmployee := findEmployeeByName(&employees, "Sam")
	if foundEmployee != nil {
		fmt.Printf("%p\n", foundEmployee)
		fmt.Printf("%p\n", &employees[1])
		fmt.Println("*foundEmployee == sam ", foundEmployee == &employees[1])
	}
	foundEmployee.salary = 0
	fmt.Println()

	increaseSalary(&employees[1], 20)
	fmt.Println(employees[1], " increased Sam's salary")
	fmt.Println()

	printEmployees(employees)
	fmt.Println(getAverageSalary(employees), "average salary")

	fmt.Println()
	foundOldestEmployee := getOldestEmployee(&employees)
	fmt.Println(foundOldestEmployee)
}
