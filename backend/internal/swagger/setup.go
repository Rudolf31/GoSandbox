package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

var Module = fx.Module("swagger",
	fx.Invoke(func(
		router *gin.Engine,

	) {
		SwaggerInfo.Schemes = []string{"http"}
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}),
)
