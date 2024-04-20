package routers

import (
	"context"

	"github.com/gin-gonic/gin"
)

func InitRouter(ctx context.Context) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		InitOrderRoutes(v1, "/orders", ctx)
		InitOrderItemRoutes(v1, "/order-items", ctx)
	}

	return router
}
