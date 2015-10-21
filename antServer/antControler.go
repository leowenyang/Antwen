package antServer

import "net/http"

type Controler func(http.ResponseWriter, *http.Request)
