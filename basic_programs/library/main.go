package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Problem 5: Nested Structs and Maps
// Design a program that models a library with books and authors using structs.
// Create a map where keys are author names and values are slices of book structs.
// Implement functions to add books and search for books by author or title.
func SearchList(bookList []Book, id string) (Book, bool) {
	for _, val := range bookList {
		if val.id == id {
			return val, true
		}
	}

	return Book{}, false
}

// returns the index of the search item
func SearchIndex(bookList []Book, id string) int {
	for i, val := range bookList {
		if val.id == id {
			return i
		}
	}

	return -1
}

func LogBooks(bookList []Book) {
	for _, val := range bookList {
		fmt.Printf("{id: '%v', author: '%v', title: '%v', genre: '%v' }\n", val.id, val.author, val.title, val.genre)
	}
	fmt.Println("\n\n")

}

func main() {
	// initialAuthor := AuthorList;
	initialBooks := BookList

	fmt.Println("Books library - lets you view, add, remove books")

	options := map[int]string{
		1: "show book list",
		2: "add new book",
		3: "remove a book",
		4: "search books based on title or author",
		5: "Quit",
	}

	for true {
		for key, val := range options {
			fmt.Println(key, val)
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("choose your options, enter the number for the operation you want to perform: ")
		optStr, _ := reader.ReadString('\n')
		optStr = strings.TrimSpace(optStr)

		optNum, err := strconv.Atoi(optStr)
		if err != nil {
			fmt.Println("Doesn't accept alphabets and special characters, only digits in the range 1 - 5 are accepted")
			fmt.Println(err)
			os.Exit(1)
		}

		switch optNum {
		case 5:
			fmt.Println("quiting...")
			os.Exit(0)

		case 1:
			fmt.Println("\nThe Books list: ")
      LogBooks(initialBooks)

		case 2:
			fmt.Println("Creating an new book...")

			id := fmt.Sprintf("B00%v", len(initialBooks)+1)

			fmt.Println("Enter book details")
			fmt.Print("Name of the book: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("Author of the book: ")
			author, _ := reader.ReadString('\n')
			author = strings.TrimSpace(author)

			fmt.Print("Genre of the book: ")
			genre, _ := reader.ReadString('\n')
			genre = strings.TrimSpace(genre)

			initialBooks = append(initialBooks, Book{id, author, title, genre})	
      LogBooks(initialBooks)

		case 3:
			fmt.Println("Deleting book...")
			fmt.Println("Enter the book id: ")
			delId, _ := reader.ReadString('\n')
			delId = strings.TrimSpace(delId)

			book, found := SearchList(initialBooks, delId)
			if found == false {
				fmt.Print("Didn't find any book with the id: %v, try other id\n", delId)
				continue
			}

			fmt.Println("the book to be deleted is ", book)
			index := SearchIndex(initialBooks, delId)
			fmt.Println("the index of book to be deleted is ", index)
			initialBooks = append(initialBooks[:index], initialBooks[index+1:]...)
			fmt.Println("the library after deleting a book")
      LogBooks(initialBooks)

		default:
			fmt.Printf("Option %v doesn't exist\n", optNum)

		}

	}

}
