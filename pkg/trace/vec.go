package trace

import (
	"math"
	"math/rand"
)

type Vec [3]float64
type Point = Vec

func NewVec(x, y, z float64) Vec {
	return Vec{x, y, z}
}

func RandomVec() Vec {
	return Vec{
		rand.Float64(),
		rand.Float64(),
		rand.Float64(),
	}
}

func RandomVecIn(min, max float64) Vec {
	return Vec{
		randFloat64In(min, max),
		randFloat64In(min, max),
		randFloat64In(min, max),
	}
}

func RandomUnitVector() Vec {
	for {
		p := RandomVecIn(-1, 1)
		lensq := p.LengthSquared()
		if 1e-160 < lensq && lensq <= 1.0 {
			return p.Normalize()
		}
	}
}

func RandomVecInUnitDisk() Vec {
	for {
		p := Vec{
			randFloat64In(-1, 1),
			randFloat64In(-1, 1),
			0,
		}
		if p.LengthSquared() < 1.0 {
			return p
		}
	}
}

func (u Vec) Neg() Vec {
	return Vec{-u[0], -u[1], -u[2]}
}

func (u Vec) Add(v Vec) Vec {
	return Vec{u[0] + v[0], u[1] + v[1], u[2] + v[2]}
}

func (u Vec) Sub(v Vec) Vec {
	return Vec{u[0] - v[0], u[1] - v[1], u[2] - v[2]}
}

func (u Vec) Prod(v Vec) Vec {
	return Vec{u[0] * v[0], u[1] * v[1], u[2] * v[2]}
}

func (u Vec) Mul(t float64) Vec {
	return Vec{u[0] * t, u[1] * t, u[2] * t}
}

func (u Vec) Div(t float64) Vec {
	return u.Mul(1.0 / t)
}

func (u Vec) Length() float64 {
	return math.Sqrt(u.LengthSquared())
}

func (u Vec) LengthSquared() float64 {
	return u[0]*u[0] + u[1]*u[1] + u[2]*u[2]
}

func (u Vec) Dot(v Vec) float64 {
	return u[0]*v[0] + u[1]*v[1] + u[2]*v[2]
}

func (u Vec) Cross(v Vec) Vec {
	return Vec{
		u[1]*v[2] - u[2]*v[1],
		u[2]*v[0] - u[0]*v[2],
		u[0]*v[1] - u[1]*v[0],
	}
}

func (u Vec) Normalize() Vec {
	return u.Div(u.Length())
}

func (u Vec) IsNearZero() bool {
	s := 1e-8
	return math.Abs(u[0]) < s && math.Abs(u[1]) < s && math.Abs(u[2]) < s
}

func (u Vec) Reflect(n Vec) Vec {
	return u.Sub(n.Mul(2.0 * u.Dot(n)))
}

func (u Vec) Refract(n Vec, eta float64) Vec {
	cosTheta := math.Min(u.Neg().Dot(n), 1.0)
	rOutPerp := u.Add(n.Mul(cosTheta)).Mul(eta)
	rOutParallel := n.Mul(-math.Sqrt(math.Abs(1.0 - rOutPerp.LengthSquared())))

	return rOutPerp.Add(rOutParallel)
}

func randFloat64In(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}
