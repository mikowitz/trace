package trace

type HittableList struct {
	Objects []Hittable
}

func (hl *HittableList) Add(object Hittable) {
	hl.Objects = append(hl.Objects, object)
}

func (hl HittableList) Hit(r Ray, i Interval) (bool, HitRecord) {
	rec := HitRecord{}
	hitAnything := false
	closest := i.Max

	for _, object := range hl.Objects {
		isHit, hitRec := object.Hit(r, NewInterval(i.Min, closest))
		if isHit {
			rec = hitRec
			hitAnything = true
			closest = rec.T
		}
	}

	return hitAnything, rec
}
