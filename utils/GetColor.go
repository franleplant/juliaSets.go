package utils

import (
	"code.google.com/p/sadbox/color"
	"math"
)

type GetColorType func(count int) (r, g, b uint8)

func GetColorFactory(COLOR_K float64, COLOR_PHASE float64) GetColorType {
	//TODO: if count is 0 then return black and if is MAX_ITER then return white (or viceversa)
	return func(count int) (r, g, b uint8) {
		r, g, b = color.HSVToRGB((math.Mod(COLOR_PHASE+COLOR_K*float64(count), 255.0))/255.0, 1.0, 1.0)
		return
	}
}
