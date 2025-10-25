package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewGin(lc fx.Lifecycle) *gin.Engine {
	router := gin.Default()
	server := &http.Server{
		Addr: ":8080",
		Handler: router,
	}
	lc.Append(fx.Hook{
		OnStart: ,
	})

	return router
}