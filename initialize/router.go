package initialize

import (
	"api-platform/router/system"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	routers := system.Routers
	PublicGroup := Router.Group("")
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "0k")
		})
	}
	{
		routers.InitUserRouter(PublicGroup)
		routers.InitBaseRouter(PublicGroup)
	}
	return Router
}
