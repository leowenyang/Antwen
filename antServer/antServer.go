// Copyright 2015 Antwen Author. All Rights Reserved.
//
// Usage:
//
// import "github.com/leowenyang/Antwen/antServer"
//
// antServer.StartServer()
//

package antServer

import "net/http"

// Server will use the blow var
var (
	AntApp *AntServer
	IPAddr string
	IPort  string
)

// AntServer is a Server abstract
type AntServer struct {
	Router *AntRouter
}

func NewAntServer() *AntServer {
	router := NewAntRouter()
	return &AntServer{Router: router}
}

// init Server var
func init() {
	AntApp = NewAntServer()
	IPort = ":8080"
}

// the below function will call by outer
func StartServer() {
	http.ListenAndServe(IPort, AntApp.Router)
}

func SetRouter(r string, f Controller) {
	AntApp.Router.Item[r] = f
}
