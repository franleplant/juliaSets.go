package main

import "fmt"

type f func(y int) int

func a(x int) f {

	return func(y int) int {
		return x + y
	}
}

func main() {
	b := a(1)
	fmt.Println(b(2))
}
