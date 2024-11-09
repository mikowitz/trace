package trace

type HitRecord struct {
	P         Point
	Normal    Vec
	T         float64
	FrontFace bool
	Material  Material
}

func (rec *HitRecord) SetFaceNormal(r Ray, outwardNormal Vec) {
	rec.FrontFace = r.Direction.Dot(outwardNormal) < 0.0
	if rec.FrontFace {
		rec.Normal = outwardNormal
	} else {
		rec.Normal = outwardNormal.Neg()
	}
}

type Hittable interface {
	Hit(r Ray, i Interval) (bool, HitRecord)
}
