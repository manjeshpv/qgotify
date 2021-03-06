package main

import (
	"fmt"
	"github.com/manjeshpv/qgotify/server/routes"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const port string = ":3333"

func main() {
	fmt.Printf("Running at %v\n", port)

	r := httprouter.New()
	fmt.Println("main")
	routes.Init(r)

	panic(http.ListenAndServe(port, r))
}
