package main

import (
	"github.com/RegalOctopus/go-react/models"
	// "github.com/gin-gonic/contrib/cors"
	// "github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func validateToken(c *gin.Context) {
	token := c.Request.Header.Get(("X-API-Key"))

	if token == "" {
		c.AbortWithStatus(401)
	} else if token == "handshake" {
		c.Next()
	} else {
		c.AbortWithStatus(401)
	}
}

func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        validateToken(c)
        c.Next()
    }
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-API-Key")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main()	{
	router := gin.Default()
	router.Use(CORS())
	// router.Use(static.Serve("/", static.LocalFile("./frontend/dist", true)))
	
	api := router.Group("/api", authMiddleware())
	{
		api.GET("/books", models.GetBooks)
		api.GET("/books/:id", models.BookById)
		api.POST("/books", models.AddBook)
		api.PATCH("/checkout", models.CheckOutBook)
		api.PATCH("/checkin", models.CheckInBook)
	}
	
	router.Run("localhost:8080")
}