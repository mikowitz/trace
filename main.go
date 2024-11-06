package main

import (
	"fmt"
	"os"

	"github.com/schollz/progressbar/v3"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	imageWidth := 256
	imageHeight := 256

	f, err := os.Create("image.ppm")
	handle(err)
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", imageWidth, imageHeight))
	handle(err)

	bar := progressbar.NewOptions(imageHeight,
		progressbar.OptionSetWidth(50),
		progressbar.OptionSetPredictTime(true),
		progressbar.OptionSetElapsedTime(true),
	)

	for y := range imageHeight {
		for x := range imageWidth {
			r := 0.0
			g := float32(y) / float32(imageHeight-1)
			b := float32(x) / float32(imageHeight-1)

			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			_, err = f.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
			handle(err)
		}

		err = bar.Add(1)
		handle(err)
	}
}
