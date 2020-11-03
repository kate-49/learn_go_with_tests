package structs

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (c Circle) Area() float64 {
	return (3.141592653589793 * (c.Radius * c.Radius))
}

func (r Rectangle) Area() float64 {
	return (r.Width * r.Height)
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
