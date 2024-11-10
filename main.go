package main

import (
	"github.com/mikowitz/trace/pkg/trace"
)

func main() {
	groundMat := trace.Lambertian{Albedo: trace.NewColor(0.8, 0.8, 0)}
	centerMat := trace.Lambertian{Albedo: trace.NewColor(0.1, 0.2, 0.5)}
	leftMat := trace.Dielectric{RefractionIndex: 1.5}
	bubbleMat := trace.Dielectric{RefractionIndex: 1.0 / 1.5}
	rightMat := trace.Metal{Albedo: trace.NewColor(0.8, 0.6, 0.2), Fuzz: 1.0}
	world := trace.HittableList{}
	world.Add(trace.NewSphere(trace.NewVec(0, -100.5, -1), 100, &groundMat))
	world.Add(trace.NewSphere(trace.NewVec(0, 0, -1.2), 0.5, &centerMat))
	world.Add(trace.NewSphere(trace.NewVec(-1, 0, -1), 0.5, &leftMat))
	world.Add(trace.NewSphere(trace.NewVec(-1, 0, -1), 0.4, &bubbleMat))
	world.Add(trace.NewSphere(trace.NewVec(1, 0, -1), 0.5, &rightMat))

	c := trace.NewCamera()
	c.AspectRatio(16.0 / 9.0)
	c.ImageWidth(400)
	c.SamplesPerPixel(100)
	c.MaxDepth(50)

	c.Render(world)
}
