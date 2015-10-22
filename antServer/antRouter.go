package antServer

import (
	"io"
	"net/http"
)

type AntRouter struct {
	// route
	Item map[string]Controller
}

func NewAntRouter() *AntRouter {
	item := make(map[string]Controller)
	return &AntRouter{Item: item}
}

func (this *AntRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle rounte function
	if v, ok := this.Item[r.URL.String()]; ok {
		v(w, r)
	} else {
		io.WriteString(w, "404 not found")
	}
}
