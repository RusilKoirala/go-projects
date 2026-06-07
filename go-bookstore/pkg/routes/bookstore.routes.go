package routes

import (
	"github.com/gorilla/mux"
	"github.com/rusilkoirala/go-projects/go-bookstore/pkg/controllers"
)

// All major routes :)
var ResgisterBookStoreRoutes = func(router *mux.Router) {
	// Create Book
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	// Get All books
	router.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	// Get Books by [id]
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	// edit (update ) book {both are same btw}
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	// erasing data ;)
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
