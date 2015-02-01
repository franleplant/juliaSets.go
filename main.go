package main

import (
	"github.com/franleplant/juliaSets/alg"
	"github.com/franleplant/juliaSets/utils"
	"math/cmplx"
)

const (
	N       = 500
	MIN_x   = float64(-1.0)
	c1      = -0.4 + 0.6i //nice
	c2      = -0.285 + 0.1i
	c3      = (1 - 1.618033987) + 0.0i //Golden ratio
	c4      = -0.8 + 0.156i            //nice
	c5      = 0.279 + 0.0i             //really nice
	c6      = 0 + 0i                   //circle
	c7      = -2 + 0i
	PHASE   = 150.0
	K_COLOR = 5.0
)

func fz1(z complex128) complex128 {
	return cmplx.Pow(z, 2) + 0.279 + 0.0i
}
func main() {
	getColor := utils.GetColorFactory(K_COLOR, PHASE)
	alg.EscapeTime(fz1, getColor, N, MIN_x)
}
