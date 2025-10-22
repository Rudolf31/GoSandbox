package main

import (
	customeerrors "interface_lesson/internal/customeErrors"
	"interface_lesson/internal/models/dto"
	"interface_lesson/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	calculatorService := services.NewCalculatorService()
	profileService := services.NewProfileService()

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

	router.GET("/profile/:id", func(c *gin.Context) {

		id_str := c.Param("id")
		id, err := strconv.Atoi(id_str)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		result, serviceErr := profileService.GetProfile(id)
		if serviceErr != nil {
			switch serviceErr.Error {
			case customeerrors.ErrNotFound:
				c.JSON(http.StatusNotFound, serviceErr)
			default:
				c.Status(http.StatusInternalServerError)
			}
			return
		}

		c.JSON(http.StatusOK, *result)
	})

	router.PUT("/profile/:id", func(c *gin.Context) {

		var profile dto.NewProfileDTO

		idStr := c.Param("id")

		id, errId := strconv.Atoi(idStr)
		if errId != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		if err := c.ShouldBindJSON(&profile); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, serviceErr := profileService.UpdateProfile(id, profile)
		if serviceErr != nil {
			switch serviceErr.Error {
			case customeerrors.ErrNotFound:
				c.JSON(http.StatusNotFound, serviceErr)
			default:
				c.Status(http.StatusInternalServerError)
			}
			return
		}

		c.JSON(http.StatusOK, *result)
	})

	router.Run()
}
