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
		authService services.AuthService,
	) {
		profileRoutes(router, profileService, authService)
		calculatorRoutes(router, calculatorService)
	}),
)
