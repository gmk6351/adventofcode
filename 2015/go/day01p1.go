package main

import (
	"fmt"
	"os"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	const inputFile = "../input/1"
	var floor int = 0 // Start at ground (0) floor
	const up = '('
	const down = ')'

	// Read file
	data, e := os.ReadFile(inputFile)
	check(e)
	str := string(data)

	// Loop through input runes
	for _, char := range str {
		switch char {
		case up:
			floor++
		case down:
			floor--
		}
	}
	fmt.Println("Result:", floor)
}
