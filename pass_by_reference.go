package main

import (
	"fmt"
)

func add2(a *int) {
	*a += 2
}

func main() {
	var a int
	a = 2
	add2(&a)
	fmt.Println(a)
}
