package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	msg := flag.String("msg", "sup", "the message to be deisplayed")
	num := flag.Int("num", 1, "how much times to print the message")
	caps := flag.Bool("u", false, "uppercase")

	flag.Parse()

	if *caps {
		*msg = strings.ToUpper(*msg)
	}

	for i := 0; i < *num; i++ {
		fmt.Println(*msg)
	}
}
