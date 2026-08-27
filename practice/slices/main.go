package main

import "fmt"

func getSliceSumAndMax(slice []int) (int, int) {
	if len(slice) == 0 {
		return 0, 0
	}
	maxElem := slice[0]
	sum := 0
	for _, n := range slice {
		sum += n
		if n > maxElem {
			maxElem = n
		}
	}

	return sum, maxElem
}

func getEvenNumsSlice(slice []int) []int {
	newSlice := make([]int, 0, len(slice))
	for _, n := range slice {
		if n%2 == 0 {
			newSlice = append(newSlice, n)
		}
	}
	return newSlice
}

func deleteSliceElem(slice *[]int, i int) {
	if i >= 0 && i <= len(*slice)-1 {
		*slice = append((*slice)[:i], (*slice)[i+1:]...)
	}
}

type product struct {
	name  string
	price float64
}

// Возвращаем указатель
func getProductByName(products []product, name string) *product {
	for i := range products {
		if products[i].name == name {
			return &products[i]
		}
	}
	return nil
}

func main() {
	// Task 1
	fmt.Println("Task 1")
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	fmt.Println(len(nums), "before len(nums)")
	fmt.Println(cap(nums), "before cap(nums)")

	nums = append(nums, 6, 7, 8)

	fmt.Println(len(nums), "after len(nums)")
	// При увеличении длины среза его емкость увеличивается в соответсвии
	// С алгоритмом Go
	fmt.Println(cap(nums), "after cap(nums)")

	fmt.Println(nums[0], "first elem")
	fmt.Println(nums[len(nums)-1], "last elem")
	fmt.Println("")

	// Task 2
	fmt.Println("Task 2")
	fmt.Println(getSliceSumAndMax(nums))
	fmt.Println("")

	// Task 3
	fmt.Println("Task 3")
	numsTask3 := []int{5, 8, 13, 4, 7, 10, 2}

	fmt.Println(getEvenNumsSlice(numsTask3))
	fmt.Println("")

	// Task 4
	fmt.Println("Task 4")
	numsTask4 := []int{10, 20, 30, 40, 50}
	deleteSliceElem(&numsTask4, 2)
	fmt.Println(numsTask4)
	fmt.Println("")

	// Task 5
	fmt.Println("Task 5")
	milk := product{name: "milk", price: 0}
	bread := product{name: "bread", price: 0}
	cucumber := product{name: "cucumber", price: 0}
	products := []product{milk, bread, cucumber}
	foundMilk := getProductByName(products, "milk")
	if foundMilk != nil {
		foundMilk.price = 1000
		fmt.Println(products[0], "products[0]")
	}
	fmt.Println("")

	// Task 6
	//
	fmt.Println("Task 6")
	// Срезы ссылаются на один и тот же underlying array
	slice1 := []int{200, 300, 400, 500}
	slice2 := slice1[:2]
	fmt.Printf("%p - slice1\n", slice1)
	fmt.Printf("%p - slice2\n", slice2)
	fmt.Println()

	numbers := make([]int, 3, 5)
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	second := numbers[:2]
	fmt.Println(numbers, len(numbers), cap(numbers))
	fmt.Println(second, len(second), cap(second))
	fmt.Println()

	second = append(second, 20, 30, 40, 50)
	fmt.Println(numbers, len(numbers), cap(numbers))
	fmt.Println(second, len(second), cap(second))
	fmt.Println()

	second[0] = 100
	fmt.Println(numbers, len(numbers), cap(numbers))
	fmt.Println(second, len(second), cap(second))
	fmt.Println()

	// Тут уже не будут ссылаться на один underlying array
	// Поскольку для слайса second из-за выхода из capacity
	// Создался новый underlying array с бОльшим capacity
	fmt.Printf("%p - numbers\n", numbers)
	fmt.Printf("%p - second\n", second)
	fmt.Println()

}
