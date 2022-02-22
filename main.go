package main

import (
	"github.com/RegalOctopus/go-react/models"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main()	{
	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api")
	{
		api.GET("/books", models.GetBooks)
		api.GET("/books/:id", models.BookById)
		api.POST("/books", models.AddBook)
		api.PATCH("/checkout", models.CheckOutBook)
		api.PATCH("/checkin", models.CheckInBook)
	}
	
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", true)))
	
	router.Run("localhost:5000")
}