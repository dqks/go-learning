package main

import "fmt"

func add(s []int) {
	s = append(s, 99)
	s[0] = 1000
	fmt.Println(s, len(s), "from add") // [1000 0 0 99] 4 from add
}
