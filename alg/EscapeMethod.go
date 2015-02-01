package alg

import (
	"fmt"
	"github.com/franleplant/juliaSets/utils"
	"image"
	"image/color"
	"image/draw"
	"math/cmplx"
)

const (
	MAX_ITERATIONS = 3000
	ER             = 1
)

// Sufficiently big enough matrix
type CMatrix [5000][5000]complex128
type fzType func(z complex128) complex128

var m CMatrix
var z *complex128

// fz is the f(z) that will be used to iterate over
// getColor is a function that takes a count int and returns an RGB color
// N is the size of the image to be generated (the image will be a square), 5000 is the max
// MIX_x is the range from which the complex plane will be calculated and graphicated
func EscapeTime(fz fzType, getColor utils.GetColorType, N int, MIN_x float64) {

	var (
		M       = N
		MAX_x   = -MIN_x
		MIN_y   = MIN_x
		MAX_y   = MAX_x
		STEP_x  = (MAX_x - MIN_y) / float64(M)
		STEP_y  = (MAX_y - MIN_y) / float64(N)
		r, g, b uint8
		count   int
		img     = image.NewRGBA(image.Rect(0, 0, M, N))
	)

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			z = &m[i][j]
			*z = complex(MIN_x+float64(j)*STEP_x, MAX_y-float64(i)*STEP_y)
			//fmt.Println(*z)

			for count = 0; count < MAX_ITERATIONS; count++ {
				*z = fz(*z)
				if cmplx.Abs(*z) > ER {
					break
				}
			}

			if count == MAX_ITERATIONS {
				fmt.Println("Max Iterations reached")
			}

			r, g, b = getColor(count)
			draw.Draw(img, image.Rect(j, i, j+1, i+1), &image.Uniform{color.RGBA{r, g, b, 255}}, image.ZP, draw.Src)
		}
	}

	utils.SaveImg(img)
}
