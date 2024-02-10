package main

type Square struct {
	X int
}

type Rectangle struct {
	X int
	Y int
}

func PerimeterSquare(s Square) interface{} {
	if s.X < 0 {
		return "X can not negative"
	} else if s.X == 0 {
		return "X can not be 0"
	}

	perimeter := 4 * s.X
	return perimeter
}

func PerimeterRectangle(r Rectangle) interface{} {
	if r.X == 0 {
		return "X can not be 0"
	}

	if r.Y == 0 {
		return "Y can not be 0"
	}

	if r.X < 0 || r.Y < 0 {
		return "X or Y can not be negative"
	}

	perimeter := 2 * (r.Y + r.X)

	return perimeter
}
