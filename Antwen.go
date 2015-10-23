package main

import (
	"io"
	"net/http"

	"github.com/leowenyang/Antwen/antServer"
)

type HomeController struct {
	antServer.Controller
}

func (this *HomeController) Get(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "say Bye")
}

func main() {
	// route
	antServer.SetRouter("/bye", &HomeController{})
	antServer.StartServer()
}
