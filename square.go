package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (point Point) End() Point {
	// implement me
	return point
}

func (square Square) Area() uint {
	return square.a * square.a
}

func (square Square) Perimeter() uint {
	return square.a * 4
}
