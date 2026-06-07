package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rusilkoirala/go-projects/go-bookstore/pkg/config"
	"github.com/rusilkoirala/go-projects/go-bookstore/pkg/models"
	"github.com/rusilkoirala/go-projects/go-bookstore/pkg/utils"
	"gorm.io/gorm"
)

var NewBook models.Book
var db *gorm.DB

func init() {
	db = config.GetDb()
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	books = models.GetAllBooks(db)
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	book, err := models.GetBookbyId(db, id)

	if err != nil {
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	err := utils.ParseBody(r, &book)
	if err != nil {
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = models.CreateBook(db, &book)
	if err != nil {
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, book, http.StatusOK)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

}
func DeleteBook(w http.ResponseWriter, r *http.Request) {

}
