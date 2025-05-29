// Flow Control programs
package main

import (
	"fmt"
	"os"
	"unicode"
)

// Problem 1: FizzBuzz with For Loop
// Implement the classic FizzBuzz problem using a for loop: Print numbers from 1 to 100,
// but for multiples of 3 print "Fizz", for multiples of 5 print "Buzz", and for multiples of both 3 and 5 print "FizzBuzz".

func fizzbuzz() {
	fizzbuzz_map := map[int]string{}

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fizzbuzz_map[i] = "FizzBuzz"
		} else if i%3 == 0 {
			fizzbuzz_map[i] = "Fizz"
		} else if i%5 == 0 {
			fizzbuzz_map[i] = "Buzz"
		}
	}

	fmt.Println(fizzbuzz_map)
}

//  Problem 2: If-Else Number Classifier
//  Write a function that takes an integer and returns a string describing the number: "negative" if less than 0,
//  "zero" if equal to 0, "small positive" if between 1 and 10, "medium positive" if between 11 and 100, and
//  "large positive" if greater than 100.

func num_classifier(val int) string {
	if val == 0 {
		return "zero"
	} else if val >= 1 && val <= 10 {
		return "small positive"
	} else if val >= 11 && val <= 100 {
		return "medium positive"
	} else if val > 100 {
		return "large positive"
	}

	return "negative num not allowed"
}

//	Problem 3: Advanced Switch Case
//
// Create a function that takes a single rune (character) and uses a switch statement
// to determine what category it falls into: uppercase letter, lowercase letter, digit, or other symbol.
func check_char(char rune) string {

	switch {
	case unicode.IsDigit(char):
		return "Digit"
	case unicode.IsUpper(char):
		return "uppercase-letter"
	case unicode.IsLower(char):
		return "lowercase-letter"
	default:
		return "other-symbol"
	}

}

// Problem 4: Defer Function Execution
// Write a program that opens a file, writes some text to it,
// and ensures the file is properly closed using defer. Make sure your program handles potential errors.
func WriteToFile(line string) (bool, error) {

	// checking if file exists
	_, err := os.Stat("block.txt")
	if os.IsNotExist(err) {
		fmt.Println("file doesn't exist, creating one...")
		newFile, err := os.Create("block.txt")

		if err != nil {
			return false, err
		}
    newFile.Close()
	}

	file, err := os.OpenFile("block.txt", os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return false, err
	}

  defer func(){
    fmt.Println("closing the file...")
    file.Close()
  }()

	val, err := file.WriteString(line + "\n")
	if err != nil {
		return false, err
	}
  fmt.Printf("wrote %v lines to block.txt\n", val)


	return true, err

}

func main() {
	// fizzbuzz()

	// a := num_classifier(3)
	// b := num_classifier(0)
	// c := num_classifier(18)
	// fmt.Println(a, b, c)

	// digit := check_char('3')
	// upper := check_char('A')
	// lower := check_char('a')
	// special := check_char('$')
	// fmt.Println(digit, upper, lower, special)
  success,_ := WriteToFile("pinterest.com")
  fmt.Println(success)
	WriteToFile("youtube.com")

}
