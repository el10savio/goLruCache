package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func factorial(number float64) float64 {
	if number == 0 || number == 1 {
		return 1
	} else {
		return number * factorial(number-1)
	}
}

func apiHandler(c *gin.Context) {
	numberParam := c.Param("number")
	number, _ := strconv.ParseFloat(numberParam, 64)
	c.JSON(http.StatusOK, gin.H{"result": factorial(number)})
}

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello"})
}

func main() {

	// Use Gin Router
	router := gin.Default()

	// Set routes
	router.GET("/", indexHandler)
	router.GET("/api/:number", apiHandler)

	router.Run(":8085")

}
