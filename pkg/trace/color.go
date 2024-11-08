package trace

import "fmt"

type Color = Vec

func NewColor(r, g, b float64) Color {
	return Color{r, g, b}
}

func (c Color) ToPpm() string {
	intensity := NewInterval(0.000, 0.999)
	r := int(256 * intensity.Clamp(c[0]))
	g := int(256 * intensity.Clamp(c[1]))
	b := int(256 * intensity.Clamp(c[2]))

	return fmt.Sprintf("%d %d %d\n", r, g, b)
}
