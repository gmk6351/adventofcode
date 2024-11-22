package main

import (
	"fmt"
	"os"
	"bufio"
)

const inputFile = "../input/3"

const north = '^'
const south = 'v'
const east = '>'
const west = '<'

func main() {
	var posA int
	var posB int

	var m map[string]int
	m = make(map[string]int)
	m[fmt.Sprintf("%d_%d", posA, posB)] = 1

	f, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		c, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		switch c {
		case north:
			posA++
		case south:
			posA--
		case east:
			posB++
		case west:
			posB--
		}
		m[fmt.Sprintf("%d_%d", posA, posB)] = 1
	}
	fmt.Println("Houses visited:", len(m))
}
