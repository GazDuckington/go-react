package main

import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

// capital case so it can be exported
// `json:field` so it can be read as json
type book struct{
	ID			string 	`json:"id"`
	Title		string 	`json:"title"`
	Author		string 	`json:"author"`
	Quantity	int		`json:"quantity"`
}

// slices of book struct
var books = []book{
	{ID: "1", Title: "Babylon's Ash", Author: "James S.A. Corey", Quantity: 5},
	{ID: "2", Title: "God's Demons", Author: "Wayne Barlowe", Quantity: 3},
	{ID: "3", Title: "Roadside Picnic", Author: "Arkady and Boris Strugatsky", Quantity: 6},
}

// convert books slice as json
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func checkOutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
	}

	book, err :=  getBookById(id)

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

func getBookById(id string) (*book, error) {
	for i, b := range books{
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookByTitle(t string) error {
	for _, b := range books{
		if b.Title == t {
			return errors.New("book already registered")
		}
	}
	return nil
}

func addBook(c *gin.Context) {
	var newBook book

	// &var get var's memory address
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	err := getBookByTitle(newBook.Title)

	if err != nil {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

func checkInBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
	}
	
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)


}

func main()	{
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", addBook)
	router.PATCH("/checkout", checkOutBook)
	router.PATCH("/checkin", checkInBook)
	router.Run("localhost:8080")
}