package main

import "strconv"

// DDD - domain driven design
// BDD - behaviour driven development
// TDD - test driven development

func Hello(name string) string {
	if name == "" {
		return "Hello World!"
	}

	_, err := strconv.Atoi(name)
	if err == nil {
		return "Hello Adam"
	}
	return "Hello " + name
}

//func main() {
//}
