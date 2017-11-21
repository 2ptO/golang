package main

import (
	"fmt"

	"github.com/2ptO/golang/playground/bet_example2/02-code-complete/quote"
	"github.com/2ptO/golang/playground/bet_example2/02-code-complete/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
