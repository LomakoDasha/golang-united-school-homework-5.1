package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (square Square) End() Point {
	sideLength := int(square.a)
	square.start.x = square.start.x + sideLength
	square.start.y = square.start.y + sideLength
	return square.start
}

func (square Square) Area() uint {
	return square.a * square.a
}

func (square Square) Perimeter() uint {
	return square.a * 4
}
