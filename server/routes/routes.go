package routes

import (
	"github.com/manjeshpv/qgotify/server/api/todo/routes"
	"github.com/manjeshpv/qgotify/server/common/static"
	"github.com/julienschmidt/httprouter"
)

func Init(router *httprouter.Router) {
	todoroutes.Init(router)
	static.Init(router)
}
