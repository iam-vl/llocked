package controllers

import (
	"html/template"
	"net/http"
)

type Static struct {
	// Template views.Template
	Template TemplateExecuter
	CouponT  Couponer
}

func (s Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Template.Execute(w, nil)
}

// Closure approach
// func HandleStatic(tpl views.Template) http.HandlerFunc {
func HandleStatic(tpl TemplateExecuter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

// func FAQ(tpl views.Template) http.HandlerFunc {
func FAQ(tpl TemplateExecuter) http.HandlerFunc {
	type QA struct {
		Question string
		Answer   template.HTML
		// Answer   string
	}
	questions := []QA{
		{
			Question: "Is there a free version?",
			Answer:   "Yes, we offer 30 days...",
		},
		{
			Question: "What are your support hours?",
			Answer:   "It's 24/7, though response times may be slower on weekends.",
		},
		{
			Question: "How do I contact support?",
			Answer:   `Just send us an email at <a href+"mailto:support@llocked.com">support@llocked.com</a>`,
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
