package main

import (
	"github.com/franleplant/juliaSets/alg"
	"github.com/franleplant/juliaSets/utils"
	"math/cmplx"
)

const (
	N       = 1000
	MIN_x   = float64(-1.0)
	PHASE   = 150.0 //Start from Blue
	K_COLOR = 10.0
)

func main() {
	getColor := utils.GetColorFactory(K_COLOR, PHASE)
	alg.EscapeTime(fz1, getColor, N, MIN_x)
}

// Functions available
//

func fz1(z complex128) complex128 {
	return cmplx.Pow(z, 2) + 0.279 + 0.0i
}

// Golden ratio
func fz2(z complex128) complex128 {
	return cmplx.Pow(z, 2) + (1 - 1.618033987) + 0.0i
}

func fz3(z complex128) complex128 {
	return cmplx.Pow(z, 2) - 0.8 + 0.156i
}

func fz4(z complex128) complex128 {
	return cmplx.Pow(z, 2) - 0.4 + 0.6i
}

// 3rd degree
func fz5(z complex128) complex128 {
	return cmplx.Pow(z, 3) + 0.4 + 0.0i
}

// 4rd degree
func fz6(z complex128) complex128 {
	return cmplx.Pow(z, 4) + 0.484 + 0.0i
}

// 7th degree
func fz7(z complex128) complex128 {
	return cmplx.Pow(z, 7) + 0.626 + 0.0i
}

func fz8(z complex128) complex128 {
	return cmplx.Exp(cmplx.Pow(z, 3)) - 0.59 + 0.0i
}

func fz9(z complex128) complex128 {
	return z*cmplx.Exp(z) + 0.04 + 0.0i
}

func fz10(z complex128) complex128 {
	return cmplx.Pow(z, 2)*cmplx.Exp(z) + 0.21 + 0.0i
}

// Crazy ones
func fz11(z complex128) complex128 {
	return cmplx.Sqrt(cmplx.Sinh(cmplx.Pow(z, 2))) + 0.065 + 0.122i
}

func fz12(z complex128) complex128 {
	return (cmplx.Pow(z, 2)+z)/cmplx.Log(z) + 0.268 + 0.06i
}
