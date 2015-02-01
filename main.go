//TODO: improve the code

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"math"
	"math/cmplx"
	"os"

	sadboxColor "code.google.com/p/sadbox/color"
)

const (
	ER             = 1
	N              = 500
	M              = N
	MIN_x          = float64(-1.0)
	MAX_x          = -MIN_x
	MIN_y          = MIN_x
	MAX_y          = MAX_x
	STEP_x         = (MAX_x - MIN_y) / M
	STEP_y         = (MAX_y - MIN_y) / N
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

type CMatrix [N][M]complex128

type fz func(z complex128) complex128

func fz1(z complex128) complex128 {
	return cmplx.Pow(z, 2) + 0.279 + 0.0i
}

var m CMatrix
var z *complex128

func main() {

	var count int

	img := image.NewRGBA(image.Rect(0, 0, M, N))

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			z = &m[i][j]
			*z = complex(MIN_x+float64(j)*STEP_x, MAX_y-float64(i)*STEP_y)
			//fmt.Println(*z)

			for count = 0; count < MAX_ITERATIONS; count++ {
				//*z = cmplx.Pow(*z, 2) + c
				*z = fz1(*z)
				if cmplx.Abs(*z) > ER {
					//fmt.Printf("count %v color %v ", count, (math.Mod(PHASE*float64(count), 255.0))/255.0)
					break
				}
			}

			if count == MAX_ITERATIONS {
				fmt.Println("Max Iterations reached")
			}

			//TODO: if count is 0 then return black and if is MAX_ITER then return white (or viceversa)
			r, g, b := sadboxColor.HSVToRGB((math.Mod(PHASE+K_COLOR*float64(count), 255.0))/255.0, 1.0, 1.0)
			//fmt.Printf("r %v g %v b %v \n", r, g, b)

			draw.Draw(img, image.Rect(j, i, j+1, i+1), &image.Uniform{color.RGBA{r, g, b, 255}}, image.ZP, draw.Src)
		}
	}

	//Image saving
	out, err := os.Create("./output.jpg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var opt jpeg.Options

	opt.Quality = 100
	// ok, write out the data into the new JPEG file

	// TODO: try png or other type of img encoding to see if it improves the quality
	err = jpeg.Encode(out, img, &opt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
