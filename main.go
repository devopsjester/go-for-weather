package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Serve static files for Tailwind CSS
	router.Static("/static", "./static")

	// Define a route to display the current weather
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Current Weather",
			"weather": "Sunny", // Placeholder for actual weather data
		})
	})

	router.LoadHTMLGlob("templates/*")
	router.Run(":8080")
}
