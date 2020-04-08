package opticalcheckbox

import "image"

type CheckBox struct {
	Vanilla bool
	Bounds  image.Rectangle
}

func checkImageFile(inputPath string) ([]bool, error) {
	return []bool{false}, nil

}

//see https://stackoverflow.com/questions/16072910/trouble-getting-a-subimage-of-an-image-in-go
type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func checkBoxDebug(im image.Image, box *CheckBox) (bool, image.Image, float64) {

	checkImage := im.(SubImager).SubImage(box.Bounds)
	cum := uint32(0)
	bounds := checkImage.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := checkImage.At(x, y).RGBA()
			cum = cum + r>>11 + g>>11 + b>>11
		}
	}

	colourCount := 3
	pixelCount := colourCount * (bounds.Max.X - bounds.Min.X) * (bounds.Max.Y - bounds.Min.Y)
	averagePixelValue := float64(cum / uint32(pixelCount))

	if box.Vanilla {
		return averagePixelValue > 30.0, checkImage, averagePixelValue
	} else {
		return averagePixelValue < 1.0, checkImage, averagePixelValue
	}
}

func checkBox(im image.Image, box *CheckBox) bool {

	checkImage := im.(SubImager).SubImage(box.Bounds)
	cum := uint32(0)
	bounds := checkImage.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := checkImage.At(x, y).RGBA()
			cum = cum + r>>11 + g>>11 + b>>11
		}
	}

	colourCount := 3
	pixelCount := colourCount * (bounds.Max.X - bounds.Min.X) * (bounds.Max.Y - bounds.Min.Y)
	averagePixelValue := float64(cum / uint32(pixelCount))

	if box.Vanilla {
		return averagePixelValue > 30.0
	} else {
		return averagePixelValue < 1.0
	}
}
