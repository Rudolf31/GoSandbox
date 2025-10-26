package routes

import (
	"interface_lesson/internal/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Module("routes",
	fx.Invoke(func(
		router *gin.Engine,
		profileService services.ProfileService,
		calculatorService services.CalculatorService,
	) {
		profileRoutes(router, profileService)
		calculatorRoutes(router, calculatorService)
	}),
)
