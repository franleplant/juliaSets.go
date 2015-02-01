package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func SaveImg(img *image.RGBA) {
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
