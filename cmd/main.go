package main

import (
	"interface_lesson/internal/models/dto"
	"interface_lesson/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	calc := services.NewCalculatorService()
	router := gin.Default()

	profileService := services.NewProfileService()

	router.GET("/add/:num1/:num2", func(c *gin.Context) {
		num1_str := c.Param("num1")
		num2_str := c.Param("num2")
		num1, err := strconv.Atoi(num1_str)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		num2, err := strconv.Atoi(num2_str)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		result := calc.Addition(num1, num2)
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	router.GET("/sub/:num1/:num2", func(c *gin.Context) {
		num1_str := c.Param("num1")
		num2_str := c.Param("num2")

		num1, err := strconv.Atoi(num1_str)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		num2, err := strconv.Atoi(num2_str)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		result := calc.Subtraction(num1, num2)
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	router.GET("/count", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"count": calc.GetOpperation()})
	})

	router.POST("/profile", func(c *gin.Context) {
		var profile dto.NewProfileDTO

		if err := c.ShouldBindJSON(&profile); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := profileService.CreateProfile(profile)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": *result})
	})

	router.Run()
}
