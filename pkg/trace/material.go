package trace

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
