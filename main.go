package main

import (
	"github.com/wicoady1/redissimulation/handler"
	"github.com/wicoady1/redissimulation/router"
)

func main() {
	module := handler.InitModule()
	router.NewRouter(module)
}
