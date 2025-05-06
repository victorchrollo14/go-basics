
package main

import (
   "fmt"
   "strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
    parts := strings.Fields(s)
    fmt.Println(parts)
	wordMap := make(map[string]int)
	
	for _, v := range parts {
	   wordCount := 0
	   
	   for _, val := range parts {
	       if(v == val){
		     wordCount++
		   }
	   }
	   wordMap[v] = wordCount
	}
	
	fmt.Println(wordMap)
	
	return wordMap;
}

// better wordCount
func wordCountFast = (s string) map[string]int {
   parts := strings.Fields(s);

   wordCountMap = make(map[string]int)
   for _, word := range parts {
      wordCountMap[word]++ 
   }

   return workCountMap

}


func main() {
	wc.Test(WordCount)
}
