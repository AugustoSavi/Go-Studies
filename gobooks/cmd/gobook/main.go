package main

import (
	"database/sql"
	"net/http"
	"os"

	"gobooks/internal/cli"
	"gobooks/internal/service"
	"gobooks/internal/web"

	_ "github.com/mattn/go-sqlite3"
)



func main(){
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = service.CreateTable(db)
	if err != nil {
		panic(err)
	}

	bookService := service.NewBookService(db)
	bookHandler := web.NewBookHandlers(bookService)

	if len(os.Args) > 1 && (os.Args[1] == "search" || os.Args[1] == "simulate") {
		bookCLI := cli.NewBookCli(bookService)
		bookCLI.Run()
		return
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /books", bookHandler.GetBooks)
	router.HandleFunc("POST /books", bookHandler.CreateBook)
	router.HandleFunc("GET /books/{id}", bookHandler.GetBookById)
	router.HandleFunc("PUT /books/{id}", bookHandler.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", bookHandler.DeleteBook)

	http.ListenAndServe(":8080", router)
}