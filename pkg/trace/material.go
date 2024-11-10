package trace

import (
	"math"
	"math/rand"
)

type ScatterRecord struct {
	Attenuation Color
	Scattered   Ray
}

type Material interface {
	Scatter(r Ray, hitRec HitRecord) (bool, ScatterRecord)
}

type Lambertian struct {
	Albedo Color
}

func (m *Lambertian) Scatter(r Ray, hitRec HitRecord) (bool, ScatterRecord) {
	scatterDirection := hitRec.Normal.Add(RandomUnitVector())
	if scatterDirection.IsNearZero() {
		scatterDirection = hitRec.Normal
	}
	scattered := NewRay(hitRec.P, scatterDirection)
	return true, ScatterRecord{Attenuation: m.Albedo, Scattered: scattered}
}

type Metal struct {
	Albedo Color
	Fuzz   float64
}

func (m *Metal) Scatter(r Ray, hitRec HitRecord) (bool, ScatterRecord) {
	reflected := r.Direction.Reflect(hitRec.Normal)
	reflected = reflected.Normalize().Add(RandomUnitVector().Mul(m.Fuzz))
	scattered := NewRay(hitRec.P, reflected)
	scatters := scattered.Direction.Dot(hitRec.Normal) > 0.0
	return scatters, ScatterRecord{Attenuation: m.Albedo, Scattered: scattered}
}

type Dielectric struct {
	RefractionIndex float64
}

func (m *Dielectric) Scatter(r Ray, hitRec HitRecord) (bool, ScatterRecord) {
	var ri float64
	if hitRec.FrontFace {
		ri = 1.0 / m.RefractionIndex
	} else {
		ri = m.RefractionIndex
	}

	unitDirection := r.Direction.Normalize()
	cosTheta := math.Min(unitDirection.Neg().Dot(hitRec.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	cannotRefract := ri*sinTheta > 1.0
	var direction Vec

	if cannotRefract || reflectance(cosTheta, ri) > rand.Float64() {
		direction = unitDirection.Reflect(hitRec.Normal)
	} else {
		direction = unitDirection.Refract(hitRec.Normal, ri)
	}

	scattered := NewRay(hitRec.P, direction)

	return true, ScatterRecord{Attenuation: NewColor(1, 1, 1), Scattered: scattered}
}

func reflectance(cosine, refractionIndex float64) float64 {
	r0 := (1.0 - refractionIndex) / (1.0 + refractionIndex)
	r0 = r0 * r0
	return r0 + (1.0-r0)*math.Pow(1-cosine, 5)
}
