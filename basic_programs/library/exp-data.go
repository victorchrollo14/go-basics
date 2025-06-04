package main

type Author struct {
	name            string
	age             int
	books_published int
	books_list      []string
}

type Book struct {
	id     string
	author string
	title  string
	genre  string
}

var AuthorList = []Author{
	{
		name:            "J.K. Rowling",
		age:             58,
		books_published: 7,
		books_list:      []string{"Harry Potter and the Philosopher's Stone", "Chamber of Secrets", "Prisoner of Azkaban"},
	},
	{
		name:            "George R.R. Martin",
		age:             75,
		books_published: 5,
		books_list:      []string{"A Game of Thrones", "A Clash of Kings", "A Storm of Swords"},
	},
	{
		name:            "Agatha Christie",
		age:             85,
		books_published: 66,
		books_list:      []string{"Murder on the Orient Express", "And Then There Were None", "The A.B.C. Murders"},
	},
	{
		name:            "Stephen King",
		age:             76,
		books_published: 60,
		books_list:      []string{"The Shining", "It", "Misery"},
	},
	{
		name:            "Haruki Murakami",
		age:             75,
		books_published: 14,
		books_list:      []string{"Norwegian Wood", "Kafka on the Shore", "1Q84"},
	},
}

var BookList = []Book{
	{
		id:     "B001",
		author: "J.K. Rowling",
		title:  "Harry Potter and the Sorcerer's Stone",
		genre:  "Fantasy",
	},
	{
		id:     "B002",
		author: "George Orwell",
		title:  "1984",
		genre:  "Dystopian",
	},
	{
		id:     "B003",
		author: "J.R.R. Tolkien",
		title:  "The Hobbit",
		genre:  "Fantasy",
	},
	{
		id:     "B004",
		author: "Harper Lee",
		title:  "To Kill a Mockingbird",
		genre:  "Classic",
	},
	{
		id:     "B005",
		author: "Dan Brown",
		title:  "The Da Vinci Code",
		genre:  "Thriller",
	},
}

