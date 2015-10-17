// Copyright 2015 Antwen Author. All Rights Reserved.
//
// Usage:
//
// import "github.com/leowenyang/Antwen/antServer"
//
// antServer.StartServer()
//

package antServer

import (
	"io"
	"net/http"
	"time"
)

// route
var route map[string]func(http.ResponseWriter, *http.Request)

func StartServer() {
	serve := http.Server{
		Addr:           ":8080",
		Handler:        &myHandler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// route
	route = make(map[string]func(http.ResponseWriter, *http.Request), 2)
	route["/hello"] = sayHello
	route["/bye"] = sayBye
	serve.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle rounte function
	if v, ok := route[r.URL.String()]; ok {
		v(w, r)
		return
	}
}

// route
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "say Hello")
}

// route
func sayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "say Bye")
}
func main() {
}
