// arrays, maps and structs programs

package main

import (
	"fmt"
	"slices"
	"strings"
)

// Problem 1: Custom Struct with Methods
// Create a `Person` struct with fields for name, age, and address.
// Then implement a method called `Birthday()` that increments the person's age and
// returns a string announcing their new age.
type Person struct {
	name    string
	age     int
	address string
}

func (p *Person) Birthday() string {
	p.age = p.age + 1

	msg := fmt.Sprintf("the new age of %v after Birthday is %v", p.name, p.age)
	return msg
}

// Problem 2: Slice Operations
// Write a function that takes two slices of integers and returns
// a new slice containing only the elements that appear in both input slices (the intersection).
func MergeSlice(s1 []int, s2 []int) []int {
	merged_slice := []int{}

	// sorting the slices
	slices.Sort(s1)
	slices.Sort(s2)

	for _, val := range s1 {
		_, is_common := slices.BinarySearch(s2, val)
		_, is_added := slices.BinarySearch(merged_slice, val)

		if is_common == true && is_added == false {
			merged_slice = append(merged_slice, val)
		}
	}

	return merged_slice

}

// gpt code
func testMergeSlice() {
	tests := []struct {
		s1 []int
		s2 []int
	}{
		{[]int{1, 2, 3}, []int{2, 3, 4}},
		{[]int{5, 6, 7}, []int{1, 2, 3}},
		{[]int{1, 1, 2, 2}, []int{2, 2, 3}},
		{[]int{}, []int{1, 2}},
		{[]int{0, -1, -2}, []int{-2, 0, 1}},
	}

	for _, t := range tests {
		fmt.Println("s1 =", t.s1)
		fmt.Println("s2 =", t.s2)
		fmt.Println("intersection =", MergeSlice(t.s1, t.s2))
		fmt.Println("-----")
	}
}

// Problem 3: Map as a Frequency Counter
// Create a function that takes a string and returns a map where keys are individual
// characters and values are the number of times each character appears in the string.
func FrequencyCounter(text string) map[string]int {
	freq_map := map[string]int{}

	// remove unnecessary char
	has_space := strings.Contains(text, "")
	filtered_text := text
	if has_space == true {
		filtered_text = strings.ReplaceAll(text, " ", "")
	}

	// iterates on index, rune
	for _, val := range filtered_text {
		// _, ok := freq_map[string(val)]
		// if it's not present then its the first iteration of that char
		// if(ok == false){
		//   freq_map[string(val)] = 1
		// } else {
		//   freq_map[string(val)] = freq_map[string(val)] + 1
		// }

		// no need to do above part, since if a key is not present, its val would be 0
		freq_map[string(val)]++

	}

	return freq_map
}

// Problem 4: Dynamic Multidimensional Slice
// Write a program that creates a "matrix" (a slice of slices of integers) of a size specified by the user,
// populates it with a simple pattern, and prints it in a formatted way.
func Matrix(row int, col int) {
	matrix := make([][]int, row)

	for i := 0; i < row; i++ {
		matrix[i] = make([]int, col)

		for j := 0; j < col; j++ {
			matrix[i][j] = (i+1)*10 + (j + 1)
		}
	}

	for i, _ := range matrix {
		for _, val := range matrix[i] {
			fmt.Print(val, "\t")
		}
		fmt.Println()
	}
}

func main() {
	victor := Person{"victor", 22, "bangalore"}
	fmt.Println("init values = ", victor)
	msg := victor.Birthday()
	fmt.Println(msg)

	s1 := []int{23, 12, -10, 0, 5, 7, 7, 8, 8, 3}
	s2 := []int{22, 12, -90, 0, 5, 7, 8, 8, 8, 4}
	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s1)
	intersection := MergeSlice(s1, s2)
	fmt.Println("intersection =", intersection)
	testMergeSlice()
	FrequencyCounter("hello hello, what are you doing")

	Matrix(3, 3)

}
