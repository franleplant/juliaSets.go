//TODO: improve the code

package main

import (
	"fmt"
	"github.com/franleplant/juliaSets/utils"
	"image"
	"image/color"
	"image/draw"
	"math/cmplx"
)

const (
	ER             = 1
	N              = 500
	MIN_x          = float64(-1.0)
	c              = c7
	c1             = -0.4 + 0.6i //nice
	c2             = -0.285 + 0.1i
	c3             = (1 - 1.618033987) + 0.0i //Golden ratio
	c4             = -0.8 + 0.156i            //nice
	c5             = 0.279 + 0.0i             //really nice
	c6             = 0 + 0i                   //circle
	c7             = -2 + 0i
	MAX_ITERATIONS = 3000
	PHASE          = 150.0
	K_COLOR        = 5.0
)

func fz1(z complex128) complex128 {
	return cmplx.Pow(z, 2) + 0.279 + 0.0i
}
func main() {
	getColor := utils.GetColorFactory(K_COLOR, PHASE)
	escapeTime(fz1, getColor, N, MIN_x)
}

// Sufficiently big enough matrix
type CMatrix [5000][5000]complex128

type fzType func(z complex128) complex128

var fzi = fz1

var m CMatrix
var z *complex128

func escapeTime(fz fzType, getColor utils.GetColorType, N int, MIN_x float64) {

	var count int
	var r, g, b uint8
	var (
		M      = N
		MAX_x  = -MIN_x
		MIN_y  = MIN_x
		MAX_y  = MAX_x
		STEP_x = (MAX_x - MIN_y) / float64(M)
		STEP_y = (MAX_y - MIN_y) / float64(N)
	)

	img := image.NewRGBA(image.Rect(0, 0, M, N))

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			z = &m[i][j]
			*z = complex(MIN_x+float64(j)*STEP_x, MAX_y-float64(i)*STEP_y)
			//fmt.Println(*z)

			for count = 0; count < MAX_ITERATIONS; count++ {
				//*z = cmplx.Pow(*z, 2) + c
				*z = fzi(*z)
				if cmplx.Abs(*z) > ER {
					//fmt.Printf("count %v color %v ", count, (math.Mod(PHASE*float64(count), 255.0))/255.0)
					break
				}
			}

			if count == MAX_ITERATIONS {
				fmt.Println("Max Iterations reached")
			}

			r, g, b = getColor(count)
			//fmt.Printf("r %v g %v b %v \n", r, g, b)

			draw.Draw(img, image.Rect(j, i, j+1, i+1), &image.Uniform{color.RGBA{r, g, b, 255}}, image.ZP, draw.Src)
		}
	}

	utils.SaveImg(img)
}
