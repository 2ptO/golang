package word_test

import (
	"testing"

	"github.com/2ptO/golang/playground/bet_example2/02-code-complete/quote"
	"github.com/2ptO/golang/playground/bet_example2/02-code-complete/word"
)

func TestCount(t *testing.T) {
	type testInput struct {
		input     string
		wordCount int
	}
	testInputs := []testInput{
		{"quick brown fox", 3},
		{"", 0},
		{"one", 1},
	}
	for _, ti := range testInputs {
		if wc := word.Count(ti.input); wc != ti.wordCount {
			t.Fatal("Expected", ti.wordCount, "for ", ti.input, "Got", wc)
		}
	}
}

func TestUseCount(t *testing.T) {
	input := "one two two three two one"
	wordCount := word.UseCount(input)
	for k, v := range wordCount {
		switch k {
		case "one":
			if v != 2 {
				t.Fatal("Expected", 2, "Got", v)
			}
		case "two":
			if v != 3 {
				t.Fatal("Expected", 3, "Got", v)
			}
		case "three":
			if v != 1 {
				t.Fatal("Expected", 1, "Got, v")
			}
		}
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		word.Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		word.UseCount(quote.SunAlso)
	}
}
