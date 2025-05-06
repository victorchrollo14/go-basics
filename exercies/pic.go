package main

import (
"fmt"
"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
  fmt.Println(dx, dy);
  	
  // pic_slice := make([][]uint8, dy)
  pic_slice := [][]uint8{}
  for i := 0; i < dy; i++ {
     // pic_slice[i] = make([]uint8, dx)
	 row := []uint8{}
     for j := 0; j < dx; j++ {
	     // pic_slice[i][j] = uint8((i+j)/2)
		 row = append(row, uint8((i+j)/2) )
	 }
	 pic_slice = append(pic_slice, row)
  }
  
  fmt.Println(pic_slice)
  return pic_slice;
	 
}

func main(){
  pic.Show(Pic)
}
