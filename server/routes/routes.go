package routes

import (
	"github.com/manjeshpv/qgotify/server/api/todo/routes"
//	"github.com/manjeshpv/qgotify/server/common/static"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func Init(router *httprouter.Router) {
	fmt.Println("routes")
	todoroutes.Init(router)
//	static.Init(router)
}
