package main

import (
	"interface_lesson/docs"
	"interface_lesson/internal/database"
	"interface_lesson/internal/routes"
	"interface_lesson/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go.uber.org/fx"
)

// @title          CRUD app
// @version        1.0
// @description    My first API
// @license.name   MIT

// @BasePath       /
// @Schemes        http

// @securityDefinitions.apikey JWT
// @in                         header
// @name                       Authorization

func main() {

	fx.New(
		fx.Provide(
			database.NewPool,
			routes.NewGin,
		),
		services.Module,
		routes.Module,

		fx.Invoke(func(pool *pgxpool.Pool, router *gin.Engine) {
			docs.SwaggerInfo.Schemes = []string{"http"}
			router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		}),
	).Run()
}
