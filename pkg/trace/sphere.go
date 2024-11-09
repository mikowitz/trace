package trace

import "math"

type Sphere struct {
	Center   Point
	Radius   float64
	Material Material
}

func NewSphere(center Point, radius float64, material Material) Sphere {
	return Sphere{Center: center, Radius: radius, Material: material}
}

func (s Sphere) Hit(r Ray, i Interval) (bool, HitRecord) {
	rec := HitRecord{}

	oc := s.Center.Sub(r.Origin)
	a := r.Direction.LengthSquared()
	h := r.Direction.Dot(oc)
	c := oc.LengthSquared() - s.Radius*s.Radius

	discrimant := h*h - a*c

	if discrimant <= 0.0 {
		return false, rec
	}
	sqrtd := math.Sqrt(discrimant)

	root := (h - sqrtd) / a
	if !i.Surrounds(root) {
		root = (h + sqrtd) / a
		if !i.Surrounds(root) {
			return false, rec
		}
	}

	rec.T = root
	rec.P = r.At(rec.T)
	outwardNormal := rec.P.Sub(s.Center).Div(s.Radius)
	rec.SetFaceNormal(r, outwardNormal)
	rec.Material = s.Material

	return true, rec
}
