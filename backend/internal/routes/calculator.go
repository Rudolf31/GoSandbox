package routes

import (
	"interface_lesson/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func calculatorRoutes(router *gin.Engine, calculatorService services.CalculatorService) {

	router.GET("/add/:num1/:num2", func(c *gin.Context) {
		num1_str := c.Param("num1")
		num2_str := c.Param("num2")
		num1, err := strconv.Atoi(num1_str)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid num1"})
			return
		}

		num2, err := strconv.Atoi(num2_str)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid num2"})
			return
		}

		result := calculatorService.Addition(num1, num2)
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	router.GET("/sub/:num1/:num2", func(c *gin.Context) {
		num1_str := c.Param("num1")
		num2_str := c.Param("num2")

		num1, err := strconv.Atoi(num1_str)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid num1"})
			return
		}
		num2, err := strconv.Atoi(num2_str)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid num2"})
			return
		}

		result := calculatorService.Subtraction(num1, num2)
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	router.GET("/count", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"count": calculatorService.GetOpperation()})
	})
}
