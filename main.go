package main

import (
	"github.com/mikowitz/trace/pkg/trace"
)

func main() {
	world := trace.HittableList{}
	world.Add(trace.NewSphere(trace.NewVec(0, 0, -1), 0.5))
	world.Add(trace.NewSphere(trace.NewVec(0, -100.5, -1), 100))

	c := trace.NewCamera()
	c.AspectRatio(16.0 / 9.0)
	c.ImageWidth(400)

	c.Render(world)
}
