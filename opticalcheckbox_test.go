package opticalcheckbox

import (
	"fmt"
	"image"

	//_ "image/jpeg"
	jpeg "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"testing"
)

func TestHist(t *testing.T) {

	reader, err := os.Open("test.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()

	// Calculate a 16-bin histogram for m's red, green, blue and alpha components.
	//
	// An image's bounds do not necessarily start at (0, 0), so the two loops start
	// at bounds.Min.Y and bounds.Min.X. Looping over Y first and X second is more
	// likely to result in better memory access patterns than X first and Y second.
	var histogram [16][4]int
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			// A color's RGBA method returns values in the range [0, 65535].
			// Shifting by 12 reduces this to the range [0, 15].
			histogram[r>>12][0]++
			histogram[g>>12][1]++
			histogram[b>>12][2]++
			histogram[a>>12][3]++
		}
	}

	// Print the results.
	fmt.Printf("%-14s %6s %6s %6s %6s\n", "bin", "red", "green", "blue", "alpha")
	for i, x := range histogram {
		fmt.Printf("0x%04x-0x%04x: %6d %6d %6d %6d\n", i<<12, (i+1)<<12-1, x[0], x[1], x[2], x[3])
	}
}
func TestSub(t *testing.T) {

	reader, err := os.Open("test.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	r0 := image.Rect(10, 10, 40, 40)

	c0 := m.(SubImager).SubImage(r0)

	var cum uint32

	cum = 0

	bounds := c0.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := c0.At(x, y).RGBA()
			cum = cum + r + g + b
		}
	}

	avg := float64(cum) / float64((bounds.Max.X-bounds.Min.X)*(bounds.Max.Y-bounds.Min.Y))

	fmt.Printf("%v avg: %v\n", c0.Bounds(), avg)

	r0 = image.Rect(60, 60, 90, 90)

	c0 = m.(SubImager).SubImage(r0)

	cum = 0

	bounds = c0.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := c0.At(x, y).RGBA()
			cum = cum + r + g + b
		}
	}

	of, err := os.Create("c0.jpg")
	if err != nil {
		panic(err)
	}
	defer of.Close()
	err = jpeg.Encode(of, c0, nil)
	if err != nil {
		t.Errorf("writing file %v\n", err)
	}

	avg = float64(cum) / float64((bounds.Max.X-bounds.Min.X)*(bounds.Max.Y-bounds.Min.Y))

	fmt.Printf("%v avg: %v\n", c0.Bounds(), avg)

	r0 = image.Rect(10, 60, 40, 90)

	c0 = m.(SubImager).SubImage(r0)

	cum = 0

	bounds = c0.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := c0.At(x, y).RGBA()
			cum = cum + r + g + b
		}
	}

	avg = float64(cum) / float64((bounds.Max.X-bounds.Min.X)*(bounds.Max.Y-bounds.Min.Y))

	fmt.Printf("%v avg: %v\n", c0.Bounds(), avg)

	r0 = image.Rect(70, 20, 80, 30)

	c0 = m.(SubImager).SubImage(r0)

	cum = 0

	bounds = c0.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := c0.At(x, y).RGBA()
			cum = cum + r + g + b
		}
	}

	avg = float64(cum) / float64((bounds.Max.X-bounds.Min.X)*(bounds.Max.Y-bounds.Min.Y))

	fmt.Printf("%v avg: %v\n", c0.Bounds(), avg)
}

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}
