package main

import (
	"fmt"
	"math"
)

func beyondHello() {
	var x int // manual declaring the type
	x = 3
	y := 4 // automatically infers type

	sum, prod := learnMultiple(x, y)
	fmt.Println("prod =", prod)
	fmt.Println("sum =", sum)

	learnTypes()
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}

func learnTypes() {
	// string types
	str := "learning go"
	str_literal := `A "raw
  " string literal ${str}`

	fmt.Println(str, str_literal)

	// adding a variable inside a string
	var var_in_string string
	var_in_string = fmt.Sprintf(`A "raw" string literal %s`, str) // %s - strings, %d - int, %f - float, %t - boolean,
	fmt.Println(var_in_string)

	g := 'Σ'    // rune type, int32
	f := 3.3123 // floating point number
	c := 3 + 5i // complex128 type
	fmt.Println(g, f, c)

	var u uint = 10 // unsigned integer
	var pi float32 = 22. / 7
	fmt.Println(u, pi)

	n := byte('\n')
	fmt.Println(n)

	var a4 [4]int              // an array of size 4, where all values are initialized to 0
	a5 := []int{3, 4, 4, 2, 3} // array initialized with fixed elements
	a5 = append(a5, 7)
	a4 = append(a4, 6)

	fmt.Println(a4, a5)
}

func main() {
	fmt.Println("Hello world")
	fmt.Println(math.Sqrt(100))

	beyondHello()
}
