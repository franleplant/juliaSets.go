package main

import (
	"github.com/franleplant/juliaSets/alg"
	"github.com/franleplant/juliaSets/utils"
	"math/cmplx"
)

const (
	N       = 1000
	MIN_x   = float64(-1.0)
	c2      = -0.285 + 0.1i
	PHASE   = 150.0
	K_COLOR = 4.0
)

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

func fz5(z complex128) complex128 {
	return cmplx.Pow(z, 2) - 0.285 + 0.1i
}

func main() {
	getColor := utils.GetColorFactory(K_COLOR, PHASE)
	alg.EscapeTime(fz1, getColor, N, MIN_x)
}
