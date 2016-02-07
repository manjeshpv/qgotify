package todoroutes

import (
	"github.com/manjeshpv/qgotify/server/api/todo/controller"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func Init(router *httprouter.Router) {
	fmt.Println("todoroutes")
	router.GET("/api/todos", todocontroller.GetAll)
	router.POST("/api/todos", todocontroller.NewTodo)
	router.DELETE("/api/todos/:id", todocontroller.RemoveTodo)
}
