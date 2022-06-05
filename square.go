package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (receiver) End() Point {
	// implement me
}

func (square Square) Area() uint {
	return square.a * square.a
}

func (square Square) Perimeter() uint {
	return square.a * 4
}
