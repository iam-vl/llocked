package controllers

import "net/http"

type TemplateExecuter interface {
	Execute(w http.ResponseWriter, data interface{})
}
type Couponer interface {
	Execute(w http.ResponseWriter, data interface{})
}
