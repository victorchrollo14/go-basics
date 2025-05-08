package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
  var first, second int; 
  
  order := 0;
  return func() int {
     if(order == 0) {
       order++
       first = 0
       return 0
      }
 
      if(order == 1){
        order++
        second = 1
        return 1
      }

      nextNum := first + second;
      first = second
      second = nextNum
      return nextNum
 
  }
}

// llm suggest stuff
func fibonacci_better() func() int {
  a, b := 0, 1
  return func() int {
    result := a
    a, b = b, a + b
    return result
  }
}


func main() {
	f := fibonacci_better()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
