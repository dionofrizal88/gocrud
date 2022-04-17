package entity

//Person object for REST(CRUD)
type Book struct {
	ID        int    `json:"id"`
	Name 	  string `json:"name"`
	Isbn  	  string `json:"isbn"`
	Page      int    `json:"page"`
}