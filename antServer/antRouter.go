package antServer

import (
	"io"
	"net/http"
	"reflect"
)

type RouterInfo struct {
	pattern        string
	controllerType reflect.Value
}

type AntRouter struct {
	// route
	Item []*RouterInfo
}

func NewAntRouter() *AntRouter {
	item := make([]*RouterInfo, 0)
	return &AntRouter{Item: item}
}

func (this *AntRouter) Add(r string, f ControlInfter) {
	controllerType := reflect.ValueOf(f)
	router := &RouterInfo{pattern: r, controllerType: controllerType}
	this.Item = append(this.Item, router)
}

func (this *AntRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// find router
	var findRouter bool
	var runRouter *RouterInfo

	runURL := r.URL.String()
	for _, route := range this.Item {
		if runURL == route.pattern {
			findRouter = true
			runRouter = route
			break
		}
	}
	if !findRouter {
		io.WriteString(w, "404 not found")
		return
	}

	// handle route function
	method := runRouter.controllerType.MethodByName("Init")
	in := make([]reflect.Value, 1)
	ctx := &Context{w, r}
	in[0] = reflect.ValueOf(ctx)
	method.Call(in)

	in = make([]reflect.Value, 0)
	switch r.Method {
	case "GET":
		method := runRouter.controllerType.MethodByName("Get")
		method.Call(in)
	case "POST":
		method := runRouter.controllerType.MethodByName("Post")
		method.Call(in)
	case "PUT":
		method := runRouter.controllerType.MethodByName("Put")
		method.Call(in)
	case "DELETE":
		method := runRouter.controllerType.MethodByName("Delete")
		method.Call(in)
	default:
		io.WriteString(w, "404 not found")
	}
}
