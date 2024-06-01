package controllers

import (
	"net/http"

	"github.com/iam-vl/llocked/views"
)

type Static struct {
	Template views.Template
}

func (s Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Template.Execute(w, nil)
}

// Closure approach
func HandleStatic(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
