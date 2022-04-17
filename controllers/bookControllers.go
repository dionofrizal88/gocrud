package controllers

import (
	"encoding/json"
	"net/http"
	"CRUDGolang/database"
	"CRUDGolang/entity"
)

//GetAllBook get all book data
func GetAllBook(w http.ResponseWriter, r *http.Request) {
	var books []entity.Book
	database.Connector.Find(&books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}