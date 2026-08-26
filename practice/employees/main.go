package main

import (
	"fmt"
	"strconv"
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
		fmt.Println(strconv.Itoa(num) + ". " + e.getFullName() + ", возраст: " + strconv.Itoa(int(e.age)) + ", зарплата: " + strconv.Itoa(int(e.salary)))
	}
}

func findEmployeeByName(employees *[]employee, firstName string) *employee {
	for i := range *employees {
		if (*employees)[i].firstName == firstName {
			return &(*employees)[i]
		}
	}
	return nil
}

func increaseSalary(e *employee, percent uint) {
	e.salary = e.salary + (e.salary * percent / 100)
}

func getAverageSalary(employees []employee) int {
	sum := 0
	for _, e := range employees {
		sum += int(e.salary)
	}
	return sum / len(employees)
}

func main() {

	alex := employee{"Alex", "Shmitd", 20, 20000}
	sam := employee{"Sam", "Sulek", 30, 50000}
	bob := employee{"Robert", "Pollson", 26, 5000}

	employees := []employee{alex, sam, bob}
	fmt.Println(employees[0].getFullName())
	fmt.Println()

	employees[0].increaseSalary(10)
	fmt.Println(employees[0])
	fmt.Println()

	foundEmployee := findEmployeeByName(&employees, "Sam")
	if foundEmployee != nil {
		// fmt.Println(foundEmployee, " adress of found employee")
		fmt.Printf("%p\n", foundEmployee)
		fmt.Printf("%p\n", &employees[1])
		fmt.Println("*foundEmployee == sam ", foundEmployee == &employees[1])
	}
	fmt.Println()

	increaseSalary(&employees[1], 20)
	fmt.Println(employees[1], " increased Sam's salary")
	fmt.Println()

	printEmployees(employees)
	fmt.Println(getAverageSalary(employees), "average salary")

}
