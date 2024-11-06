package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/mikowitz/trace/pkg/trace"
	"github.com/schollz/progressbar/v3"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	aspectRatio := 16.0 / 9.0
	imageWidth := 400

	imageHeight := int(float64(imageWidth) / aspectRatio)
	if imageHeight < 1 {
		imageHeight = 1
	}

	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))
	cameraCenter := trace.NewVec(0, 0, 0)

	viewportU := trace.NewVec(viewportWidth, 0, 0)
	viewportV := trace.NewVec(0, -viewportHeight, 0)

	pixelDeltaU := viewportU.Div(float64(imageWidth))
	pixelDeltaV := viewportV.Div(float64(imageHeight))

	viewportUpperLeft := cameraCenter.Sub(trace.NewVec(0, 0, focalLength)).Sub(viewportU.Div(2)).Sub(viewportV.Div(2))
	pixel00Loc := viewportUpperLeft.Add(pixelDeltaU.Add(pixelDeltaV).Mul(0.5))

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

	semaphore := make(chan struct{}, 100)
	var wg sync.WaitGroup

	pixels := make([]string, imageWidth*imageHeight)

	for y := range imageHeight {
		for x := range imageWidth {
			wg.Add(1)
			go func(x, y int) {
				defer wg.Done()
				semaphore <- struct{}{}
				defer func() { <-semaphore }()

				pixelCenter := pixel00Loc.Add(pixelDeltaU.Mul(float64(x))).Add(pixelDeltaV.Mul(float64(y)))
				rayDirection := pixelCenter.Sub(cameraCenter)
				ray := trace.NewRay(cameraCenter, rayDirection)

				pixelColor := ray.Color()
				pixels[y*imageWidth+x] = pixelColor.ToPpm()
				handle(err)
			}(x, y)
		}
		wg.Wait()

		err = bar.Add(1)
		handle(err)
	}

	pixelString := strings.Join(pixels, "")
	_, err = f.WriteString(pixelString)
	handle(err)
}
