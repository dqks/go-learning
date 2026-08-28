package main

import "fmt"

//

type Vehicle interface {
	move()
}

func drive(vehicle Vehicle) {
	fmt.Println("Vehicle is driving")
}

type Car struct{}

func (c Car) move() {
	fmt.Println("Car moves")
}

type Bike struct{}

func (b Bike) move() {
	fmt.Println("Bike moves")
}

type Aircraft struct {
	name string
}

func (a *Aircraft) move() {
	fmt.Println("Aircraft moves")
}

func (a *Aircraft) setName(name string) {
	a.name = name
}

//

type Empty interface{}

func print(value Empty) {
	fmt.Println(value)
}

//

type Eatable interface {
	eat()
}

type Throwable interface {
	throw()
}

// Product Также можно встраивать в в интерфейс другие интерфейсы
// И чтобы тип данных соответствовал этому интерфейсу
// Он должен также реализовать все методы внутренних интерфейсов
type Product interface {
	Eatable
	Throwable
}

// Также мы можем прописать сигнатуры его методов и это также будет работать
//type Product interface {
//	throw()
//	eat()
//}

type Apple struct {
	appleType string
}

func (a Apple) throw() {
	fmt.Println("Apple is thrown")
}

func (a Apple) eat() {
	fmt.Println("Apple is eaten")
}

//

func main() {
	var car Vehicle
	fmt.Println(car)
	// Хоть это и тип, но от него нельзя создать напрямую объект интерфейса
	//car = vehicle{}

	// Чтобы тип данных соответствовал интерфейсу он должен реализовать все его методы
	// Стурктура Car реализовывает все методы интерфейса Vehicle, поэтому он подходит
	car = Car{}
	bike := Bike{}
	// Это полезно, поскольку можно обобщить передаваемый тип данных
	// Например, есть транеспорт - Vehicle, а есть структуры Car и Bike
	// Есть функция drive, которая принимает транспорт Vehicle
	// Если бы не было интерфейса у drive, то пришлось бы прописывать 2 функции
	// Отдельно для Car и Bike, а так, с интерфейсом, и Car и Bike походят под тип
	drive(car)
	drive(bike)

	// Пустой интерфейс будет означать любой тип данных
	print(1)
	print("dsdafsd")
	print(bike)
	print(car)

	// Можно передавать указатель на структуру там, где требуется интерфейс
	drive(&bike)

	// Но указатель на интерфейс нельзя
	//drive(&car)

	// При этом если метод интерфейса реализуется только для указателя
	// То там, где требуется интерфейс можно передать только указатель
	aircraft := Aircraft{name: "Boeing"}
	// метод move структуры Aircraft реализуется только для указателя
	drive(&aircraft)
	// А так уже нельзя
	//drive(aircraft)

	// Метод для указателя нужен если хотим поменять поля структуры
	aircraft.setName("Boeing 731")
	print(aircraft)

	// При этом структуры могут иметь свои методы, вне интерфейса
	// И реализовывать другие интерфейсы
	// Apple реазилует интерфейс Eatable и Throwable
	var goldenApple Product = Apple{appleType: "golden"}
	print(goldenApple)

	// Проверка, реализует ли структура Apple интрефейс Eatable
	value, ok := goldenApple.(Eatable)
	fmt.Println(value, ok) // true

	value1, ok1 := car.(Eatable)
	fmt.Println(value1, ok1) // false
}
