package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	ageOfTimmy := 4
	ageOfTimmysDog := Years(ageOfTimmy)
	if ageOfTimmysDog != 28 {
		t.Fatal("Expected:", 28, "Got:", ageOfTimmysDog)
	}
}

func TestYearsTwo(t *testing.T) {
	ageOfTimmy := 4
	ageOfTimmysDog := YearsTwo(ageOfTimmy)
	if ageOfTimmysDog != 28 {
		t.Fatal("Expected:", 28, "Got:", ageOfTimmysDog)
	}
}

func ExampleYears() {
	ageOfTimmy := 4
	fmt.Println(Years(ageOfTimmy))
	// Output:
	// 28
}

func ExampleYearsTwo() {
	ageOfTimmy := 4
	fmt.Println(YearsTwo(ageOfTimmy))
	// Output:
	// 28
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(8)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(16)
	}
}
