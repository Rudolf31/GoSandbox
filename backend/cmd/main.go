package main

import (
	"interface_lesson/internal/config"
	"interface_lesson/internal/database"
	"interface_lesson/internal/routes"
	"interface_lesson/internal/services"
	"interface_lesson/internal/swagger"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"go.uber.org/fx"
	"go.uber.org/zap"
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
			config.NewConfig,
			database.NewPool,
			routes.NewGin,
			zap.NewProduction,
		),
		services.Module,
		routes.Module,
		swagger.Module,

		fx.Invoke(func(pool *pgxpool.Pool, router *gin.Engine, log *zap.Logger) {

		}),
	).Run()
}
