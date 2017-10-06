package main

import (
	"fmt"

	"github.com/2ptO/golang/playground/bet_example/dog"
)

type canine struct {
	name string
	age  int
}

/*
 * Example derived from Todd's exercise code at
 * https://github.com/GoesToEleven/go-programming/tree/master/code_samples/010-ninja-level-thirteen/01/starting-code
 */

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
