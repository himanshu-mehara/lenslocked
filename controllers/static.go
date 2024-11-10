package controllers

import (
	"html/template"
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "is there are any free version",
			Answer:   "yes! we are offering a free trial for 30 days",
		},
		{
			Question: "what are your support hours ?",
			Answer:   "we have support staff answering emails",
		},
		{
			Question: "how do i contact support",
			Answer:   "email us at ffdsaf;ds",
		},
		{
			Question: "Where is your office located?",
			Answer:   "our entire team is remote !",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
