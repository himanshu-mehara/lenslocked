package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"webdev/controllers"
	"webdev/views"

	"github.com/go-chi/chi/v5"
)

// func executeTemplate(w http.ResponseWriter, filepath string) {
// 	t, err := views.Parse((filepath))
// 	if err != nil {
// 		log.Printf("parsing template: %v", err)
// 		http.Error(w, "there was an error parsing the template.", http.StatusInternalServerError)
// 		return
// 	}
// 	t.Execute(w, nil)
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
	tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.Get("/contact/{user-id}", MyRequestHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	fmt.Println("starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
