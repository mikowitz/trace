package trace

import "math"

type Ray struct {
	Origin    Point
	Direction Vec
	Time      float64
}

func NewRay(origin Point, direction Vec, time float64) Ray {
	return Ray{Origin: origin, Direction: direction, Time: time}
}

func (r Ray) At(t float64) Point {
	return r.Origin.Add(r.Direction.Mul(t))
}

func (r Ray) Color(world Hittable, depth int) Color {
	if depth <= 0 {
		return NewColor(0, 0, 0)
	}

	isHit, hitRec := world.Hit(r, NewInterval(0.001, math.Inf(1)))
	if isHit {
		scatters, scatterRec := hitRec.Material.Scatter(r, hitRec)
		if scatters {
			return scatterRec.Attenuation.Prod(scatterRec.Scattered.Color(world, depth-1))
		}
		return NewColor(0, 0, 0)
	}

	unitDirection := r.Direction.Normalize()
	a := 0.5 * (unitDirection[1] + 1.0)
	return NewColor(1, 1, 1).Mul(1.0 - a).Add(NewColor(0.5, 0.7, 1).Mul(a))
}

func (r Ray) HitSphere(center Point, radius float64) float64 {
	oc := center.Sub(r.Origin)
	a := r.Direction.LengthSquared()
	h := r.Direction.Dot(oc)
	c := oc.LengthSquared() - radius*radius

	discrimant := h*h - a*c

	if discrimant <= 0.0 {
		return -1.0
	}
	return (h - math.Sqrt(discrimant)) / a
}
