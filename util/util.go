package util

import "image"

func SizeTo(original image.Point, target image.Point) (float64, float64) {
	return float64(target.X) / float64(original.X),
		float64(target.Y) / float64(original.Y)
}
