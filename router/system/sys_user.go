package system

import (
	"api-platform/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userGroup := Router.Group("user")
	userApi := v1.Apis.UserApi
	{
		userGroup.POST("register", userApi.Register)
		userGroup.POST("login", userApi.Login)
	}
}
