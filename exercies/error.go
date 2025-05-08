package main

import (
	"fmt"
)

// this type implements the Error interface 
// type Error interface {
//    Error() string
// }
type ErrNegativeSquare float64

func (e ErrNegativeSquare) Error() string {
  return fmt.Sprintf("cannot sqrt negative number: %v", float64(e))
}

// function will return the value followed by error
func Sqrt(x float64) (float64, error) {
  if(x < 0) {
    return 0, ErrNegativeSquare(x) 
  }

  z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(8))
	fmt.Println(Sqrt(-8))
}
