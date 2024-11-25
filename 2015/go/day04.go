package main

import (
	"fmt"
	"crypto/md5"
	"strings"
	"io"
)

const input = "iwrupvqb"
func findHashWithPrefix(prefix string) {
	var number = 1
	for {
		h := md5.New()
		io.WriteString(h, input)
		io.WriteString(h, fmt.Sprintf("%d", number))

		if strings.HasPrefix(fmt.Sprintf("%x", h.Sum(nil)), prefix) {
			fmt.Printf("%x\n", h.Sum(nil))
			fmt.Println("Number:", number)
			break
		}
		number++
	}

}
func main() {
	findHashWithPrefix("00000")
	findHashWithPrefix("000000")
}

