package main

import (
	"interface_lesson/internal/database"
	"interface_lesson/internal/routes"
	"interface_lesson/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		fx.Provide(
			database.NewPool,
			routes.NewGin,
		),
		services.Module,
		routes.Module,

		fx.Invoke(func(*pgxpool.Pool, *gin.Engine) {

		}),
	).Run()
}
