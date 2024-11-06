package trace

import "math"

type Vec [3]float64
type Point = Vec

func NewVec(x, y, z float64) Vec {
	return Vec{x, y, z}
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
