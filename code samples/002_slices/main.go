package main

import (
	"fmt"
)

func main() {
	var lastOver []int
	lastOver = append(lastOver, 4, 6, 4, 6, 6, 6)
	inningsOvers := make([][]int, 50)
	const overPerInnings = 50
	fmt.Println(lastOver)
	inningsOvers[overPerInnings-1] = lastOver
	fmt.Println(len(inningsOvers[49]),
					cap(inningsOvers[49]),
					len(inningsOvers[48]),
					cap(inningsOvers[48]))
}
