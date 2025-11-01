package routes

import (
	"interface_lesson/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func calculatorRoutes(router *gin.Engine, calculatorService services.CalculatorService) {

	calculatorGroup := router.Group("/calculator")
	{
		calculatorGroup.GET("/add/:num1/:num2", func(c *gin.Context) { addition(c, calculatorService) })
		calculatorGroup.GET("/sub/:num1/:num2", func(c *gin.Context) { subtraction(c, calculatorService) })
		calculatorGroup.GET("/count", func(c *gin.Context) { count(c, calculatorService) })
	}
}

// @Summary 	Addition of two numbers
// @Tags Calculator
// @Accept		json
// @Produce		json
// @Param		num1	path	int	true	"First number"
// @Param		num2	path	int	true	"Second number"
// @Success	200
// @Failure	400 "Bad Request"
// @Router		/calculator/add/{num1}/{num2} [get]
func addition(c *gin.Context, calculatorService services.CalculatorService) {
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
}

// @Summary 	Subtraction of two numbers
// @Tags Calculator
// @Accept		json
// @Produce		json
// @Param		num1	path	int	true	"First number"
// @Param		num2	path	int	true	"Second number"
// @Success	200
// @Failure	400 "Bad Request"
// @Router		/calculator/sub/{num1}/{num2} [get]
func subtraction(c *gin.Context, calculatorService services.CalculatorService) {
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
}

// @Summary 	Count of operations
// @Tags Calculator
// @Accept		json
// @Produce		json
// @Success		200
// @Failure		400 "Bad Request"
// @Router		/calculator/count [get]
func count(c *gin.Context, calculatorService services.CalculatorService) {
	c.JSON(http.StatusOK, gin.H{"count": calculatorService.GetOperation()})
}
