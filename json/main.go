package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Animal struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Sound string `json:"sound"`
}

func main() {
	animals := []Animal{}

	file, err := os.Open("file.json")
	if err != nil {
		panic(err)
	}

	if err := json.NewDecoder(file).Decode(&animals); err != nil {
		panic(err)
	}

	fmt.Println(animals)
}
