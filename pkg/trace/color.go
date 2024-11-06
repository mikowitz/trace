package trace

import "fmt"

type Color = Vec

func NewColor(r, g, b float64) Color {
	return Color{r, g, b}
}

func (c Color) ToPpm() string {
	r := int(255.999 * c[0])
	g := int(255.999 * c[1])
	b := int(255.999 * c[2])

	return fmt.Sprintf("%d %d %d\n", r, g, b)
}
