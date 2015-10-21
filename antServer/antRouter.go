package antServer

import "net/http"

type Router struct{}

// route
var route map[string]func(http.ResponseWriter, *http.Request)

func (*Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle rounte function
	if v, ok := route[r.URL.String()]; ok {
		v(w, r)
		return
	}
}

func init() {
	route = make(map[string]func(http.ResponseWriter, *http.Request), 2)

}
