package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//DummyWebHandler parse dummy webpage for redis test
func (m *Module) DummyWebHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t := template.New("index.html")
	t = template.Must(t.ParseFiles("template/index.html"))
	t.Execute(w, map[string]interface{}{"test": "test"})
}

//DummyRequest handle request which is come from ajax (webpage)
func (m *Module) DummyRequest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response := map[string]string{
		"status": "OK!",
	}
	responseCode := http.StatusOK

	err := m.Response("test")
	if err != nil {
		response["status"] = "Error"
		response["message"] = fmt.Sprintf("%+v", err)
		responseCode = http.StatusNotAcceptable
	}

	res, _ := json.Marshal(response)
	NewAjax(w, r, res, responseCode)
}

//NewAjax throw ajax a response with its response code
func NewAjax(w http.ResponseWriter, r *http.Request, js []byte, code int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(js))
}
