package examples

import (
	"math/rand"

	"github.com/mikowitz/trace/pkg/trace"
)

func WeekendFinalBouncing() {
	groundMat := trace.Lambertian{Albedo: trace.NewColor(0.5, 0.5, 0.5)}

	world := trace.HittableList{}
	world.Add(trace.NewSphere(trace.NewVec(0, -1000, 0), 1000, &groundMat))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := trace.NewVec(float64(a)+0.9*rand.Float64(), 0.2, float64(b)+0.9*rand.Float64())

			if center.Sub(trace.NewVec(4, 0.2, 0)).Length() > 0.9 {
				var material trace.Material

				if chooseMat < 0.9 {
					albedo := trace.RandomVec().Prod(trace.RandomVec())
					material = &trace.Lambertian{Albedo: albedo}
					center2 := center.Add(trace.NewVec(0, rand.Float64()*0.25, 0))
					world.Add(trace.NewMovingSphere(center, center2, 0.2, material))
				} else if chooseMat < 0.95 {
					albedo := trace.RandomVecIn(0.5, 1)
					fuzz := rand.Float64() * 0.5
					material = &trace.Metal{Albedo: albedo, Fuzz: fuzz}
					world.Add(trace.NewSphere(center, 0.2, material))
				} else {
					material = &trace.Dielectric{RefractionIndex: 1.5}
					world.Add(trace.NewSphere(center, 0.2, material))
				}

			}
		}
	}

	mat1 := trace.Dielectric{RefractionIndex: 1.5}
	world.Add(trace.NewSphere(trace.NewVec(0, 1, 0), 1, &mat1))

	mat2 := trace.Lambertian{Albedo: trace.NewColor(0.4, 0.2, 0.1)}
	world.Add(trace.NewSphere(trace.NewVec(-4, 1, 0), 1, &mat2))

	mat3 := trace.Metal{Albedo: trace.NewColor(0.7, 0.6, 0.5), Fuzz: 0.0}
	world.Add(trace.NewSphere(trace.NewVec(4, 1, 0), 1, &mat3))

	c := trace.NewCamera()
	c.AspectRatio(16.0 / 9.0)
	c.ImageWidth(400)
	c.SamplesPerPixel(50)
	c.MaxDepth(10)

	c.Vfov(20.0)
	c.Lookfrom(trace.NewVec(13, 2, 3))
	c.Lookat(trace.NewVec(0, 0, 0))
	c.Vup(trace.NewVec(0, 1, 0))

	c.DefocusAngle(0.5)
	c.FocusDist(10.0)

	c.Render(world)
}
