package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/leowenyang/Antwen/antLog"
	"github.com/leowenyang/Antwen/antServer"
)

func main() {
	// route
	antServer.SetRouter("/login", login)
	antServer.SetRouter("/bye", sayBye)
	antServer.StartServer()
}

// route
func login(w http.ResponseWriter, r *http.Request) {
	//antLog.Info("method:", r.Method)
	if r.Method == "GET" {
		// User login, should show INPUT message
		t, err := template.ParseFiles("./antServer/login.gtpl")
		if err != nil {
			antLog.Info("Parse HTML Template Error")
		}
		t.Execute(w, nil)
	} else {
		// User send form success
		r.ParseForm()
		antLog.Info("username:", r.Form["username"])
		antLog.Info("password:", r.Form["password"])

	}
}

// route
func sayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "say Bye")
}
