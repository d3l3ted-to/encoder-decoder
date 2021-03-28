package main

import (
	"encoderReplacer"
	"flag"
	"fmt"
)

var (
	lowcase     = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	capitalcase = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	input = flag.String("t", "", "text to encode")
	key   = flag.Int("k", 0, "key for encoding")
)

func init() {
	flag.Parse()
}

func main() {
	replacer := encoderReplacer.GenerateReplacer(*key, lowcase, capitalcase)
	fmt.Printf("Result: %s\n", encoderReplacer.Replace(*input, replacer))
}
