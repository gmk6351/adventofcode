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

func part2() {
	var santaPosA int
	var santaPosB int
	var robosantaPosA int
	var robosantaPosB int

	var m map[string]int
	m = make(map[string]int)
	m[fmt.Sprintf("%d_%d", 0, 0)] = 1

	f, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var count int
	reader := bufio.NewReader(f)
	for {
		c, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		switch count%2 {
		case 0:
			switch c {
			case north:
				santaPosA++
			case south:
				santaPosA--
			case east:
				santaPosB++
			case west:
				santaPosB--
			}
			m[fmt.Sprintf("%d_%d", santaPosA, santaPosB)] = 1
		case 1:
			switch c {
			case north:
				robosantaPosA++
			case south:
				robosantaPosA--
			case east:
				robosantaPosB++
			case west:
				robosantaPosB--
			}
			m[fmt.Sprintf("%d_%d", robosantaPosA, robosantaPosB)] = 1
		}
		count++
	}
	fmt.Println("Houses visited (year++):", len(m))
}

func part1() {
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

func main() {
	part1()
	part2()
}
