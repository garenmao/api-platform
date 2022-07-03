package v1

import (
	"api-platform/service"
)

type apiGroup struct {
	UserApi
	BaseApi
}

var Apis = new(apiGroup)

var (
	userService = service.Group.UserService
)
