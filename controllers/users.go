package controllers

import (
	"net/http"
)

type Users struct {
	Templates struct {
		// New views.Template
		New TemplateExecuter
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}
