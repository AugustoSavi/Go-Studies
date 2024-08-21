package cli

import (
	"fmt"
	"gobooks/internal/service"
	"os"
	"strconv"
	"time"
)

type BookCli struct {
	service *service.BookService
}

func NewBookCli(service *service.BookService) *BookCli {
	return &BookCli{service: service}
}

func (cli *BookCli) Run() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: books <command> [args]")
		return
	}

	command := os.Args[1]
	switch command {
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books search <book title>")
			return
		}
		bookName := os.Args[2]
		cli.searchBooks(bookName)
	case "simulate":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books simulate <book_id> <book_id> <book_id> ....")
			return
		}
		booksIds := os.Args[2:]
		cli.SimulateReading(booksIds)
	}
}

func (cli *BookCli) searchBooks(name string) {
	books, err := cli.service.SearchBooksByName(name)
	if err != nil {
		fmt.Println("Error searching books:", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("no books found:")
		return
	}

	fmt.Printf("%d books found\n", len(books))
	for _, book := range books {
		fmt.Printf("Id: %d, Title: %s, Author: %s, Genre: %s\n", book.ID, book.Title, book.Author, book.Genre)
	}
}

func (cli *BookCli) SimulateReading(booksIds []string) {
	var bookIds []int
	for _, idStr := range booksIds {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Printf("invalid book id: %s", idStr)
			continue
		}
		bookIds = append(bookIds, id)
	}

	responses := cli.service.SimulateMultipleReadings(bookIds, 6 * time.Second)
	for _, v := range responses {
		fmt.Println(v)
	}
}
