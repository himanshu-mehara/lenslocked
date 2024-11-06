package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "there was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "there was an error executing the template.", http.StatusInternalServerError)
		return
	}
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplpath := filepath.Join("templates", "home.gothml")
	executeTemplate(w, tplpath)
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplpath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplpath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tplpath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, tplpath)
}

// func faqHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "text/html; charset=utf-8")
// 	fmt.Fprint(w, `<h1>FAQ PAGE </h1>
// 	<ul>
// 	<li> <b>is there any free course ?</b> yes there is free trial of the course .</li>
// 	<li> <b> what are your support hours ? </b> we offer a 30 days free trial </li>
// 	<li> <b> is there any free version ? </b> yes there are two free versions </li>
// 	</ul>

// 	`)
// }

func MyRequestHandler(w http.ResponseWriter, r *http.Request) {
	// fetch the url parameter `"userID"` from the request of a matching
	// routing pattern. An example routing pattern could be: /users/{userID}
	userID := chi.URLParam(r, "userID")
	// fetch `"key"` from the request context
	ctx := r.Context()
	key := ctx.Value("key").(string)
	// respond to the client
	w.Write([]byte(fmt.Sprintf("hi %v, %v", userID, key)))
}
func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/contact/{user-id}", MyRequestHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	fmt.Println("starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
