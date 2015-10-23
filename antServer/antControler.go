package antServer

import (
	"io"
	"net/http"
)

type Controller struct{}

type ControlInfter interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	Put(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func (this *Controller) Get(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "404 not found")
}

func (this *Controller) Post(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "404 not found")
}

func (this *Controller) Put(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "404 not found")
}

func (this *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "404 not found")
}
