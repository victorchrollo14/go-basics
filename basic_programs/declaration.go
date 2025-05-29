// bunch of basic programs
package main

import (
	"fmt"
	"math"
)

// program 1 that imports "fmt" and "math" packages and print area of circle
func circle_area() {
	var radius float64 = 7
	area := math.Pi * (radius * radius)
	fmt.Printf("The area of circle of radius %v = %v", radius, area)
}

// function that takes in single string and returns 2 strings by splitting into 2
func split_string(word string) (string, string) {
	word_length := len(word)
	mid_term := (word_length / 2)

	fmt.Println("the first split string = ", word[:mid_term])
	fmt.Println("the second split string = ", word[mid_term:])

	return word[:mid_term], word[mid_term:]
}

// function that returns max, min and avg of an array or slice of integers
func stats(nums []int) (maxi, mini, avg int) {
	mini = nums[0]
	maxi = nums[0]
	avg = 0

  sum := 0
	for _, val := range nums { 
    if(mini > val){
      mini = val 
    } 
    if(maxi < val){
      maxi = val 
    }
    sum += val
	}

  avg = sum / len(nums)

 fmt.Printf("the min value in slice %v = %v\n", nums, mini)
 fmt.Printf("the max value in slice %v = %v\n", nums, maxi)
 fmt.Printf("the average value in slice %v = %v\n", nums, avg)
 
 return maxi, mini, avg
}

// closures functions
func fibonacci() func() int {
  first, second := 0, 1  
  iteration := 0 

  return func() int {
    result := first
    first, second = second, (first + second)
    fmt.Printf("the index %v value in fibonacci series = %v\n",iteration, result)
    iteration++
    return result 
  }

} 


func main() {
	circle_area()
	split_string("hellohai")
	stats([]int{23, 12, 14, 13, 87, 8, 62, 23})
  fib_val := fibonacci()
  array := [10]int{} 
  for _ = range array{ 
    fib_val()
  }
}
