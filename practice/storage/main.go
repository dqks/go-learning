package main

import "fmt"

// Если меняется длина слайса, его капасити, то тогда принимаем в функцию указатель на слайс
// Иначе просто слайс

type product struct {
	name       string
	price      float64
	amountLeft int
}

func printProduct(p product, i int) {
	fmt.Printf("%d. Название: %s, цена: %.2f, остаток на складе: %d\n", i+1, p.name, p.price, p.amountLeft)
}

func printAllProducts(products []product) {
	for i, p := range products {
		printProduct(p, i)
	}
}

func getProductByName(products []product, name string) *product {
	for i := range products {
		if products[i].name == name {
			return &products[i]
		}
	}
	return nil
}

func getTheMostExpensiveProduct(products []product) *product {
	if len(products) == 0 {
		return nil
	}
	expensiveProduct := &products[0]
	for i := 1; i < len(products); i++ {
		if expensiveProduct.price < products[i].price {
			expensiveProduct = &products[i]
		}
	}
	return expensiveProduct
}

func countFullProductPrice(products []product) float64 {
	if len(products) == 0 {
		return 0
	}
	var sum float64 = 0
	for _, p := range products {
		sum += p.price * float64(p.amountLeft)
	}
	return sum
}

func getEndingProducts(products []product) {
	fmt.Println("Осталось меньше 5 товаров:")
	for i, p := range products {
		if p.amountLeft < 5 {
			printProduct(p, i)
		}
	}
}

func (p *product) addAmountLeft(amountToAdd int) {
	if amountToAdd > 0 {
		p.amountLeft += amountToAdd
	}
}

func (p *product) sellProduct(amountToSell int) {
	if p.amountLeft-amountToSell >= 0 && amountToSell > 0 {
		p.amountLeft = p.amountLeft - amountToSell
	}
}

func (p *product) changeProductPrice(price float64) {
	if price > 0 {
		p.price = price
	}
}

func main() {
	milk := product{name: "Milk", price: 100, amountLeft: 0}
	cookie := product{name: "Cookie", price: 0, amountLeft: 15}
	bread := product{name: "Bread", price: 0, amountLeft: 3}
	rice := product{name: "Rice", price: 0, amountLeft: 8}
	pizza := product{name: "Pizza", price: 0, amountLeft: 10}

	products := []product{milk, cookie, bread, rice, pizza}

	foundProduct := getProductByName(products, "Milk")
	if foundProduct != nil {
		foundProduct.amountLeft = 1000
		fmt.Println(products[0])
	}
	fmt.Println()

	products[1].sellProduct(15)
	products[0].addAmountLeft(5)
	products[2].changeProductPrice(100000)
	printAllProducts(products)

	fmt.Println(countFullProductPrice(products))
	getEndingProducts(products)
	expensiveProduct := getTheMostExpensiveProduct(products)
	expensiveProduct.name = "gdfhfghfghgf"
	fmt.Println()
	fmt.Println()
	fmt.Println()
}
