package main

import (
	"fmt"
	"myapp/hello_en"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	hello_en.SayHelloEn()
}
