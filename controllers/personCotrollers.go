package controllers

import (
	"encoding/json"
	// "io/ioutil"
	"net/http"
	"CRUDGolang/database"
	"CRUDGolang/entity"
	// "strconv"

	// "github.com/gorilla/mux"
)

//GetAllPerson get all person data
func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	var persons []entity.Person
	database.Connector.Find(&persons)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(persons)
}