package views

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"webdev/context"
	"webdev/models"

	"github.com/gorilla/csrf"
)

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	tpl := template.New(patterns[0])
	// tpl, err := template.ParseFS(fs, patterns...)
	// if err != nil {
	// 	return Template{}, fmt.Errorf("parsing template: %w", err)
	// }
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() (template.HTML, error) {
				return "", fmt.Errorf("csrfFIELD NOT IMPLEMENTED")
			},
			"currentUser": func() (template.HTML, error) {
				return "", fmt.Errorf("currentUser not implemented")
			},

			"errors": func() []string {
				return nil 
					// "don't do this",
					// "the email address you provided is a already associated with an account",
					// "something went wrong",

				},
			},
	)
	tpl, err := tpl.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}

// func Parse(filepath string) (Template, error) {
// 	htmltpl, err := template.ParseFiles(filepath)
// 	if err != nil {
// 		return Template{}, fmt.Errorf("parsing template: %w", err)
// 	}
// 	return Template{
// 		htmlTpl: htmltpl,
// 	}, nil
// }

type Template struct {
	htmlTpl *template.Template
}

// type Data struct {
// 	Yield interface{} 
// 	Alerts []Alert 
// }

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}, errs ...error) {
	tpl, err := t.htmlTpl.Clone()
	if err != nil {
		log.Printf("cloning template : %v", err)
		http.Error(w, "there was an error rendering the page.", http.StatusInternalServerError)
	}
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return csrf.TemplateField(r)
			},
			"currentUser": func() *models.User {
				return context.User(r.Context())
			},
			"errors": func() []string {
				var errMessages []string 
				for _,err := range errs {
					errMessages= append(errMessages,err.Error())
				}
				return errMessages
			},
		},
	)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "there was an error executing the template.", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}
