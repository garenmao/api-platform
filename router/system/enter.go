package system

type RouterGroup struct {
	UserRouter
	BaseRouter
}

var Routers = new(RouterGroup)
