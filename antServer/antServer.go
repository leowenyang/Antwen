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

func StartServer() {
	http.ListenAndServe(":8888", &Router{})
}

func SetRouter(r string, f Controler) {
	route[r] = f
}
