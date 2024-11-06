package main

import (
	"fmt"
	"os"

	"github.com/mikowitz/trace/pkg/trace"
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
			pixelColor := trace.NewColor(
				0.0,
				float64(y)/float64(imageHeight-1),
				float64(x)/float64(imageHeight-1),
			)

			_, err = f.WriteString(pixelColor.ToPpm())
			handle(err)
		}

		err = bar.Add(1)
		handle(err)
	}
}
