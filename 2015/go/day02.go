package main

import (
	"os"
	"fmt"
	"bufio"
	"slices"
)

const inputFile = "../input/2"

var requiredPaper int
var requiredRibbon int

func addRequiredPaper(size string) {
	var dim [3]int // l, w, h
	fmt.Sscanf(size, "%dx%dx%d", &dim[0], &dim[1], &dim[2])
	slices.Sort(dim[:])

	requiredPaper += 2*dim[0]*dim[1] + 2*dim[1]*dim[2] + 2*dim[2]*dim[0] + dim[0]*dim[1]
}

func addRequiredRibbon(size string) {
	var dim [3]int // l, w, h
	fmt.Sscanf(size, "%dx%dx%d", &dim[0], &dim[1], &dim[2])
	slices.Sort(dim[:])

	requiredRibbon += 2*dim[0] + 2*dim[1] + dim[0]*dim[1]*dim[2]
}

func main() {
	f, e := os.Open(inputFile)
	if e != nil {
		panic(e)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// "<length>x<width>x<height>
		addRequiredPaper(scanner.Text())
		addRequiredRibbon(scanner.Text())
	}

	fmt.Println("Required Paper:", requiredPaper)
	fmt.Println("Required Ribbon:", requiredRibbon)
}
