package main

import (
	"CRUD_GO/internal/service"
	"CRUD_GO/internal/store"
	"CRUD_GO/internal/transport"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {

	//conectar SQLLite
	db, err := sql.Open("sqlite", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//create table if not exist
	q := `
	CREATE TABLE IF NOT EXISTS books (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	author TEXT NOT NULL
	)
`
	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}
	//inyectar dependencias
	bookStore := store.New(db)
	bookService := service.New(bookStore)
	bookHandler := transport.New(bookService)
	//configuración de rutas
	http.HandleFunc("/books", bookHandler.HandleBooks)
	http.HandleFunc("/books/", bookHandler.HandleBookByID)

	fmt.Println("Servidor ejecutandose en http://localhost:8081")
	fmt.Println("API Endpoints: ")
	fmt.Println("GET       /books            -Obtener todos los libros")
	fmt.Println("POST      /books           -Crear un nuevo libro")
	fmt.Println("GET       /books/{id}       -Obtener un libro especifico")
	fmt.Println("PUT       /books{id}         -Actualizar un libro")
	fmt.Println("DELETE    /books/{id}    -Eliminar un libro")

	//comenzar y escuchar al servidor
	log.Fatal(http.ListenAndServe(":8081", nil))

}
