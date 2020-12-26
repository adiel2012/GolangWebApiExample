package main

type IRouter interface {
}

var routerInstance IRouter = nil

func GetRouter() IRouter {
	return routerInstance
}

func SetRouter(r IRouter) {
	routerInstance = r
}
