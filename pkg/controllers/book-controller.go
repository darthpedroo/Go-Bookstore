package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/darthpedroo/go-bookstore/pkg/utils"
	"github.com/darthpedroo/go-bookstore/pkg/models"
)

var NewBook models.Book 

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, _ = json.Marshal(newBooks)

	//I could make this a function since it gets repeated in every function
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) //200 status
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r*http.Request){
	vars := mux.Vars(r) //Inside this request, there is the book id.
	bookId := vars["bookId"] //Acces the book Id inside of the request. 
	ID, err := strconv.ParseInt(bookId,0,0)

	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ :=models.GetBookById(ID) //The underscore is because we don't need to use the db here.
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{} //This has to be of type book, so it has acces to the "CreateBook method. Which can only by accesed by book types"
	utils.ParseBody(r, CreateBook)
	b:= CreateBook.CreateBook() //Beware, shouldn't we pass a parameter to CreateBook (?)
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r*http.Request){
	vars := mux.vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)

	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Conrent-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var UpdateBook = &models.Book{} //Here, the pointer is so the ParseBody can modify it later
	utils.ParseBody(r, UpdateBook) 
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	if UpdateBook.Name != ""{
		bookDetails.Name = UpdateBook.Name
	}
	if UpdateBook.Author != ""{
		bookDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != ""{
		bookDetails.Publication = UpdateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	
}








//The user interacts with the routes and the routes send control to the controllers. 

// Here there is all the logic, and the controllers have to perform some operations to the database.
