package models

import (
	"github.com/gin-gonic/gin"
	"errors"
	"net/http"
)

// slices of book struct
var books = []book{
	{ID: "1", Title: "Babylon's Ash", Author: "James S.A. Corey", Quantity: 5},
	{ID: "2", Title: "God's Demons", Author: "Wayne Barlowe", Quantity: 3},
	{ID: "3", Title: "Roadside Picnic", Author: "Arkady and Boris Strugatsky", Quantity: 6},
}

// convert books slice as json
func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func CheckOutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
	}

	book, err :=  GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func GetBookById(id string) (*book, error) {
	for i, b := range books{
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func BookById(c *gin.Context) {
	id := c.Param("id")
	book, err := GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func GetBookByTitle(t string) error {
	for _, b := range books{
		if b.Title == t {
			return errors.New("book already registered")
		}
	}
	return nil
}

func AddBook(c *gin.Context) {
	var newBook book

	// &var get var's memory address
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	err := GetBookByTitle(newBook.Title)

	if err != nil {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

func CheckInBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
	}
	
	book, err := GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}