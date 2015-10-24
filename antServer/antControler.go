package antServer

import (
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

type Controller struct {
	Ctx *Context
}

type ControlInfter interface {
	Init(ctx *Context)
	Get()
	Post()
	Put()
	Delete()
}

func (this *Controller) Init(ctx *Context) {
	this.Ctx = ctx
}

func (this *Controller) Get() {
	io.WriteString(this.Ctx.W, "404 not found")
}

func (this *Controller) Post() {
	io.WriteString(this.Ctx.W, "404 not found")
}

func (this *Controller) Put() {
	io.WriteString(this.Ctx.W, "404 not found")
}

func (this *Controller) Delete() {
	io.WriteString(this.Ctx.W, "404 not found")
}
