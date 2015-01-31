package main

import (
	"fmt"
	"math/cmplx"
)

const (
	ER             = 1
	N              = 1000
	M              = N
	RANGE_x        = 2.0
	RANGE_y        = 2.0
	STEP_x         = RANGE_x / M
	STEP_y         = RANGE_y / N
	c              = 0.1 + 0.6i
	MAX_ITERATIONS = 500
)

type CMatrix [N][M]complex128

var m CMatrix
var z *complex128

func main() {
	var count int

	var a = complex(1, 2)
	z = &a
	fmt.Println(*z)

	for i := 0; i < N; i++ {

		for j := 0; j < M; j++ {
			z = &m[i][j]
			*z = complex(-1+float64(j)*STEP_x, 1-float64(i)*STEP_y)
			//fmt.Println(*z)

			for count = 0; count < MAX_ITERATIONS; count++ {
				*z = cmplx.Sqrt(*z) + c
				if cmplx.Abs(*z) > ER {
					break
				}
			}
			//fmt.Printf("%v \n", count)
		}
	}

}
