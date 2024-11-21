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
	var firstNegative int

	// Read file
	data, e := os.ReadFile(inputFile)
	check(e)
	str := string(data)

	// Loop through input runes
	for pos, char := range str {
		switch char {
		case up:
			floor++
		case down:
			floor--
		}
		if firstNegative == 0 && floor < 0 {
			firstNegative = pos+1
		}
	}
	fmt.Println("Result:", floor)
	fmt.Println("First Negative Position:", firstNegative)
}
