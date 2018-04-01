package router

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wicoady1/redissimulation/handler"
)

func NewRouter(m *handler.Module) {
	router := httprouter.New()

	router.GET("/", m.DummyWebHandler)
	router.POST("/req", m.DummyRequest)
	router.ServeFiles("/script/*filepath", http.Dir("template/script"))

	log.Println("Router Initialization Complete on port 5555")
	http.ListenAndServe(":5555", router)
}
