package interfaces

type Area interface {
	Area()
	Parameter()
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

func (r Circle) Area() float64 {
	return 3.14 * r.Radius * r.Radius
}

func (l Rectangle) Area() float64 {
	return l.Length * l.Breadth
}

func (l Rectangle) Parameter() float64 {
	return 2.0 * (l.Length + l.Breadth)
}

func (r Circle) Parameter() float64 {
	return 2.0 * 3.14 * r.Radius
}
