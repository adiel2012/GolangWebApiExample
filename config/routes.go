package config

import (
	"gowebapi/controllers"
	"net/http"
)

func Routes(routes *map[string]func(http.ResponseWriter, *http.Request)) {
	(*routes)["/test"] = controllers.Test
	(*routes)["/product/get"] = controllers.ListProducts
}
