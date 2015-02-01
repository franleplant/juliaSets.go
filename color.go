package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

func main() {

	out, err := os.Create("./output.jpg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	img := image.NewRGBA(image.Rect(0, 0, 300, 300))
	fmt.Printf("%T", img)
	draw.Draw(img, image.Rect(0, 0, 10, 10), &image.Uniform{color.RGBA{0, 0, 255, 255}}, image.ZP, draw.Src)
	draw.Draw(img, image.Rect(20, 20, 10, 10), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.ZP, draw.Src)

	var opt jpeg.Options

	opt.Quality = 80
	// ok, write out the data into the new JPEG file

	err = jpeg.Encode(out, img, &opt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Generated image to output.jpg \n")
}
