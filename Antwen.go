package main

import (
	"io"

	"github.com/leowenyang/Antwen/antServer"
)

type HomeController struct {
	antServer.Controller
}

func (this *HomeController) Get() {
	io.WriteString(this.Ctx.W, "say Home")
}

type MainController struct {
	antServer.Controller
}

func (this *MainController) Get() {
	io.WriteString(this.Ctx.W, "say Main")
}

func main() {
	// route
	antServer.SetRouter("/home", &HomeController{})
	antServer.SetRouter("/main", &MainController{})
	antServer.StartServer()
}
