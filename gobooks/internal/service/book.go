package service

import (
	"database/sql"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Genre  string
}

func CreateTable(db *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		genre TEXT NOT NULL
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	return nil
}

func (b Book) GetFullBook() string {
	return b.Title + " by " + b.Author
}

type BookService struct {
	db *sql.DB
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{db: db}
}

func (s *BookService) CreateBook(book *Book) error {
	query := "insert into books (title, author, genre) values(?,?,?)"

	result, err := s.db.Exec(query, book.Title, book.Author, book.Genre)
	if err != nil {
		return err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	book.ID = int(lastInsertId)
	return nil
}

func (s *BookService) GetBooks() ([]Book, error) {
	query := "select id, title, author, genre from books"
	results, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	var books []Book
	for results.Next() {
		var book Book
		err := results.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (s *BookService) GetBookById(id int) (*Book, error) {
	query := "select id, title, author, genre from books where id = ?"
	result := s.db.QueryRow(query, id)
	
	var book Book
	err := result.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *BookService) UpdateBook(book *Book) error {
	query := "update books set title=?, author=?, genre=? where id = ?"
	_, err := s.db.Exec(query, book.Title, book.Author, book.Genre, book.ID)
	
	return err
}


func (s *BookService) DeleteBook(id int) error {
	query := "delete from books where id = ?"
	_, err := s.db.Exec(query, id)

	return err
}