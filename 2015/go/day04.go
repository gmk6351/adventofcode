package main

import (
	"fmt"
	"crypto/md5"
	"strings"
	"io"
)

const input = "iwrupvqb"
func main() {
	var number = 1
	for {
		h := md5.New()
		io.WriteString(h, input)
		io.WriteString(h, fmt.Sprintf("%d", number))

		if strings.HasPrefix(fmt.Sprintf("%x", h.Sum(nil)), "00000") {
			fmt.Printf("%x\n", h.Sum(nil))
			fmt.Println("Number:", number)
			break
		}
		number++
	}
}

