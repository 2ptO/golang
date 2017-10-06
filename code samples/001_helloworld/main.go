package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, world")
	s := "qiuck brown jumps over lazy dog"
	s = strings.Join(strings.Split(s, " "), " ")
	fmt.Println(s)
}