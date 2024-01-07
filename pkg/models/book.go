package models 

import (
	"github.com/jinzhu/gorm"
	"github.com/darthpedroo/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.model
	Name string `gorm: ""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init (){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	//b is something that the function recieves, which is of type 'Book'. Also, the function return a type 'Book'
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	//Also, Books is the name of the table
	db.Find(&Books) //&Books: The & operator is used to pass a pointer to the slice, allowing Gorm to modify the underlying slice with the query results.
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}