package main

import (
	"fmt"
	"interface_lesson/internal/services"
)

func main() {
	a := services.NewCalculatorService()
	num1, num2 := 1, 2
	fmt.Print(a.Addition(num1, num2), " ", a.Subtraction(num1, num2), " ", a.GetOpperation())
}
