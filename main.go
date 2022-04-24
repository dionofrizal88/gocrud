package main

import (
	"log"
	"net/http"
	"CRUDGolang/controllers"
	"CRUDGolang/database"
	"CRUDGolang/entity"

	"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/person/create", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/person/get", controllers.GetAllPerson).Methods("GET")
	// router.HandleFunc("/get/{id}", controllers.GetPersonByID).Methods("GET")
	// router.HandleFunc("/update/{id}", controllers.UpdatePersonByID).Methods("PUT")
	// router.HandleFunc("/delete/{id}", controllers.DeletPersonByID).Methods("DELETE")

	// get all book
	router.HandleFunc("/book/get", controllers.GetAllBook).Methods("GET")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "",
			DB:         "golearn",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	// add table person into database golearn
	database.Migrate(&entity.Person{})

	// add table book into database golearn
	database.MigrateBook(&entity.Book{})
}