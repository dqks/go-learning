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

// Принимаем срез []employee
// Срез состоит из 3 элементов: указатель (ссылается на массив), длина и емкость
// Копируется и указатель и длина и емкость
// Но скопированный указатель указывает на все тотже массив, что и тот, который мы передаем в функцию
// Т.е. чаще всего нет смысла передавать в функцию указатель на сам срез, а лучше просто передавать срез
func findEmployeeByName(employees []employee, firstName string) *employee {
	for i := range employees {
		// Обращаемся по индексу к элементу среза
		// А не через отдельную переменную
		if employees[i].firstName == firstName {
			// Возвращаем ссылку на элемент ВСЕ ТОГО ЖЕ СРЕЗА
			// Который мы передаем в ЭТУ функцию
			return &employees[i]
		}
	}
	return nil
}

func getAverageSalary(employees []employee) float64 {
	if len(employees) == 0 {
		return 0
	}
	var sum float64 = 0
	for _, e := range employees {
		sum += float64(e.salary)
	}
	return sum / float64(len(employees))
}

func getOldestEmployee(employees []employee) *employee {
	if len(employees) == 0 {
		return nil
	}
	oldestEmployee := &employees[0]
	for i := 1; i < len(employees); i++ {
		if employees[i].age > oldestEmployee.age {
			oldestEmployee = &employees[i]
		}
	}
	return oldestEmployee
}

func main() {
	alex := employee{firstName: "Alex", lastName: "Shmitd", age: 0, salary: 11}
	sam := employee{firstName: "Sam", lastName: "Sulek", age: 10, salary: 10}
	bob := employee{firstName: "Robert", lastName: "Pollson", age: 0, salary: 10}

	employees := []employee{alex, sam, bob}
	fmt.Println(employees[0].getFullName())
	fmt.Println()

	employees[0].increaseSalary(10)
	fmt.Println(employees[0])
	fmt.Println()

	foundEmployee := findEmployeeByName(employees, " ")
	if foundEmployee != nil {
		fmt.Printf("%p\n", foundEmployee)
		fmt.Printf("%p\n", &employees[1])
		fmt.Println("*foundEmployee == sam ", foundEmployee == &employees[1])
		foundEmployee.salary = 0
	}
	fmt.Println()

	printEmployees(employees)
	fmt.Println(getAverageSalary(employees), "average salary")

	fmt.Println()
	foundOldestEmployee := getOldestEmployee(employees)
	fmt.Println(foundOldestEmployee)
}
