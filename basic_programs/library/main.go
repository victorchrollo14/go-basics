package main

import (
	"fmt"
)

// Problem 5: Nested Structs and Maps
// Design a program that models a library with books and authors using structs.
// Create a map where keys are author names and values are slices of book structs.
// Implement functions to add books and search for books by author or title.
func main() {
  // initialAuthor := AuthorList;
  // initialBooks := BookList;

  fmt.Println("Books library - lets you view, add, remove books")
 
  options := map[int]string{
    1: "show book list",
    2: "add new book",
    3: "remove a book",
    4: "search books based on title or author",
  }
 
  for true {
    for key, val := range options {
      fmt.Println(key, val)
    }
 
    var optNum int;
    fmt.Print("choose your options, enter the number for the operation you want to perform: ")
    fmt.Scanln(&optNum)

    fmt.Println("the option you have chosen is: ", options[optNum]);
    break

  }


}
