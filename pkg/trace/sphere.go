package trace

import "math"

type Sphere struct {
	Center   Ray
	Radius   float64
	Material Material
}

func NewSphere(center Point, radius float64, material Material) Sphere {
	return Sphere{
		Center:   NewRay(center, NewVec(0, 0, 0), 0),
		Radius:   radius,
		Material: material,
	}
}

func NewMovingSphere(center1, center2 Point, radius float64, material Material) Sphere {
	return Sphere{
		Center:   NewRay(center1, center2.Sub(center1), 0),
		Radius:   radius,
		Material: material,
	}

}

func (s Sphere) Hit(r Ray, i Interval) (bool, HitRecord) {
	rec := HitRecord{}

	center := s.Center.At(r.Time)

	oc := center.Sub(r.Origin)
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
	outwardNormal := rec.P.Sub(center).Div(s.Radius)
	rec.SetFaceNormal(r, outwardNormal)
	rec.Material = s.Material

	return true, rec
}
