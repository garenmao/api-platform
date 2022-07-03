package system

import (
	"api-platform/api/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (r *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("base")
	baseApi := v1.Apis.BaseApi
	{
		baseRouter.GET("/captcha", baseApi.Captcha)
	}
}
