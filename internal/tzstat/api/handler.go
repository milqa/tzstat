package api

import "net/http"

type Methods interface {
	Handlers() []Method
}

type Method struct {
	Name    string
	Path    string
	Handler http.Handler
}
