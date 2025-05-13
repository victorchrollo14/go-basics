package main

import (
	"io"
	"os"
	"strings"
  // "fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) Read(r1 []byte) (n int, err error){
   // A - Z ascii values (65, 90)
   // a - z ascii values (97, 122) 
   
   for {
     buffer := make([]byte, 8)
     n, err = r13.r.Read(buffer);

     if err == io.EOF {
       // fmt.Println("end of file reached") 
       return 0, io.EOF 
     }

     // fmt.Println("original buffer = ", buffer[:n]) 
     for i, val := range buffer[:n]{
       modVal := 0
       
       if val >= 65 && val <= 90 {
         delta := int(val) + 13
         
         if(delta > 90){
            overflow := delta - 90
            modVal = 64 + overflow 
            buffer[i] = byte(modVal)
            continue 
         }
 
         buffer[i] = byte(delta) 
       } else if val >= 97 && val <= 122 {
         delta := int(val) + 13
 
         if(delta > 122){
           overflow := delta - 122
           modVal = 96 + overflow 
           buffer[i] = byte(modVal)
           continue
         }
 
         buffer[i] = byte(delta)
       }
     }

     copy(r1, buffer[:n])
     return n, nil 
  }  

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
