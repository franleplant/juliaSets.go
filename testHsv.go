package main

import (
	"fmt"

	"code.google.com/p/sadbox/color"
)

func main() {
	r, g, b := color.HSVToRGB(0.6, 1.0, 1.0)
	fmt.Println(r)
	fmt.Println(g)
	fmt.Println(b)
}
