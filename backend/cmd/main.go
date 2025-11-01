package main

import (
	"interface_lesson/internal/routes"
	"interface_lesson/internal/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		fx.Provide(
			routes.NewGin,
		),
		services.Module,
		routes.Module,

		fx.Invoke(func(*gin.Engine) {

		}),
	).Run()
}
