package main

import (
	"flag"
	"fmt"
	"strings"
)

func reverse(char string) string {
	var result string
	for _, value := range char {
		result = string(value) + result
	}

	return result
}

func main() {
	msg := flag.String("msg", "sup", "the message to be deisplayed")
	num := flag.Int("num", 1, "how much times to print the message")
	caps := flag.Bool("u", false, "uppercase")
	revrse := flag.Bool("r", false, "reverse")
	flag.Parse()

	if *caps {
		*msg = strings.ToUpper(*msg)
	}
	if *revrse {
		*msg = reverse(*msg)
	}

	for i := 0; i < *num; i++ {
		fmt.Println(*msg)
	}
}
