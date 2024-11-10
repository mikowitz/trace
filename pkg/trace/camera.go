package trace

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"sync"

	"github.com/schollz/progressbar/v3"
)

type Camera struct {
	aspectRatio             float64
	imageWidth, imageHeight int
	samplesPerPixel         int
	maxDepth                int
	vfov                    float64
	lookfrom, lookat        Point
	vup                     Vec

	center, pixel00Loc       Point
	pixelDeltaU, pixelDeltaV Vec
	pixelSampleScale         float64
}

func NewCamera() Camera {
	return Camera{}
}

func (c *Camera) AspectRatio(aspectRatio float64) {
	c.aspectRatio = aspectRatio
}

func (c *Camera) ImageWidth(imageWidth int) {
	c.imageWidth = imageWidth
}

func (c *Camera) SamplesPerPixel(samplesPerPixel int) {
	c.samplesPerPixel = samplesPerPixel
}

func (c *Camera) MaxDepth(maxDepth int) {
	c.maxDepth = maxDepth
}

func (c *Camera) Vfov(vfov float64) {
	c.vfov = vfov
}

func (c *Camera) Lookfrom(lookfrom Point) {
	c.lookfrom = lookfrom
}

func (c *Camera) Lookat(lookat Point) {
	c.lookat = lookat
}

func (c *Camera) Vup(vup Vec) {
	c.vup = vup
}

func (c *Camera) Render(world Hittable) {
	c.initialize()

	f, err := os.Create("image.ppm")
	handle(err)
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", c.imageWidth, c.imageHeight))
	handle(err)

	bar := progressbar.NewOptions(c.imageWidth*c.imageHeight,
		progressbar.OptionSetWidth(50),
	)

	semaphore := make(chan struct{}, 10)
	var wg sync.WaitGroup

	pixels := make([]string, c.imageWidth*c.imageHeight)

	for y := range c.imageHeight {
		for x := range c.imageWidth {
			wg.Add(1)
			go func(x, y int) {
				defer wg.Done()
				defer func() {
					err = bar.Add(1)
					handle(err)
				}()
				semaphore <- struct{}{}
				defer func() { <-semaphore }()

				pixelColor := NewColor(0, 0, 0)
				for _ = range c.samplesPerPixel {
					r := c.getRay(x, y)
					pixelColor = pixelColor.Add(r.Color(world, c.maxDepth))
				}
				pixels[y*c.imageWidth+x] = pixelColor.Mul(c.pixelSampleScale).ToPpm()
			}(x, y)
		}
		wg.Wait()
	}

	pixelString := strings.Join(pixels, "")
	_, err = f.WriteString(pixelString)
	handle(err)
}

func (c *Camera) getRay(x, y int) Ray {
	xOffset := rand.Float64() - 0.5
	yOffset := rand.Float64() - 0.5
	pixelSample := c.pixel00Loc.Add(c.pixelDeltaU.Mul(float64(x) + xOffset)).
		Add(c.pixelDeltaV.Mul(float64(y) + yOffset))

	origin := c.center
	direction := pixelSample.Sub(origin)
	return NewRay(origin, direction)
}

func (c *Camera) initialize() {
	c.imageHeight = int(float64(c.imageWidth) / c.aspectRatio)
	if c.imageHeight < 1 {
		c.imageHeight = 1
	}

	c.pixelSampleScale = 1.0 / float64(c.samplesPerPixel)

	c.center = c.lookfrom

	focalLength := c.lookfrom.Sub(c.lookat).Length()
	theta := c.vfov * math.Pi / 180.0
	h := math.Tan(theta / 2.0)
	viewportHeight := 2.0 * h * focalLength
	viewportWidth := viewportHeight * (float64(c.imageWidth) / float64(c.imageHeight))

	w := c.lookfrom.Sub(c.lookat).Normalize()
	u := c.vup.Cross(w).Normalize()
	v := w.Cross(u)

	viewportU := u.Mul(viewportWidth)
	viewportV := v.Neg().Mul(viewportHeight)

	c.pixelDeltaU = viewportU.Div(float64(c.imageWidth))
	c.pixelDeltaV = viewportV.Div(float64(c.imageHeight))

	viewportUpperLeft := c.center.Sub(w.Mul(focalLength)).Sub(viewportU.Div(2)).Sub(viewportV.Div(2))
	c.pixel00Loc = viewportUpperLeft.Add(c.pixelDeltaU.Add(c.pixelDeltaV).Mul(0.5))
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}
