package trace

type Ray struct {
	Origin    Point
	Direction Vec
}

func NewRay(origin Point, direction Vec) Ray {
	return Ray{Origin: origin, Direction: direction}
}

func (r Ray) At(t float64) Point {
	return r.Origin.Add(r.Direction.Mul(t))
}

func (r Ray) Color() Color {
	if r.HitSphere(NewVec(0, 0, -1), 0.5) {
		return NewColor(1, 0, 0)
	}

	unitDirection := r.Direction.Normalize()
	a := 0.5 * (unitDirection[1] + 1.0)
	return NewColor(1, 1, 1).Mul(1.0 - a).Add(NewColor(0.5, 0.7, 1).Mul(a))
}

func (r Ray) HitSphere(center Point, radius float64) bool {
	oc := center.Sub(r.Origin)
	a := r.Direction.Dot(r.Direction)
	b := -2.0 * r.Direction.Dot(oc)
	c := oc.Dot(oc) - radius*radius

	discrimant := b*b - 4.0*a*c

	return discrimant >= 0.0
}
