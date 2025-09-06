package main

import (
	"fmt"
)

func SumInts(m map[string]int64) int64 {

	var sum int64
	for _, value := range m {
		sum += value

	}

	return sum
}

func SumFloats(m map[string]float64) float64 {
	var sum float64

	for _, value := range m {
		sum += value
	}

	return sum

}


// new stuff added to interface in go 1.18 where we can define constraint on a type with interface 
type Number interface {
  int64 | float64
}

func SumOrFloats[k comparable, V Number](m map[k]V) V {
	var sum V
	for _, value := range m {
		sum += value
	}

	return sum
}

func main() {

	ints := map[string]int64{
		"hello": 1,
		"hi":    23,
	}

	floats :=
		map[string]float64{
			"sum":   1.2,
			"hello": 8.34,
		}

	fmt.Printf("the non generic function outputs %v and %v\n", SumInts(ints), SumFloats(floats))

	intsGeneric := map[string]int64{
		"hello":  12,
		"victor": 23,
		"matt":   34,
	}

	floatsGeneric := map[int]float64{
		1:  1.3,
		33: 23.78,
		4:  34.723,
	}

	fmt.Println("sum of intsGeneric = ", SumOrFloats(intsGeneric))
	fmt.Println("sum of floatsGeneric = ", SumOrFloats(floatsGeneric))

	fmt.Printf("Sum of generics, ints = %v and floats = %v", SumOrFloats[string, int64](ints), SumOrFloats[string, float64](floats))

}
