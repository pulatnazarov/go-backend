package main

func Division(x, y float64) interface{} {
	if y == 0.0 {
		return "can not divide by 0"
	}

	if x == 0.0 {
		return 0.0
	}

	return x / y
}

func main() {

}
