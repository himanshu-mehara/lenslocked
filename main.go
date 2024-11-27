package main

import (
	"fmt"
	"net/http"
	"os"
	"webdev/controllers"
	"webdev/migrations"
	"webdev/models"
	"webdev/views"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
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
	cfg := models.DefaultPostgresConfig()
	fmt.Println(cfg.String())
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	userService := models.UserService{
		DB: db,
	}
	sessionService := models.SessionService{
		DB: db,
	}

	umw := controllers.UserMiddlefware{
		SessionService: &sessionService,
	}

	fs := os.DirFS("templates")
	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect([]byte(csrfKey), csrf.Secure(false),)

	usersC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(fs, "signup.gohtml", "tailwind.gohtml"))
	usersC.Templates.SignIn = views.Must(views.ParseFS(fs, "signin.gohtml", "tailwind.gohtml"))
	usersC.Templates.ForgotPassword = views.Must(views.ParseFS(fs, "forgot-pw.gohtml", "tailwind.gohtml"))


	r := chi.NewRouter()
	r.Use(csrfMw)
	r.Use(umw.SetUser)
	// tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	// if err != nil {
	// 	panic(err)
	// }
	// fs := os.DirFS("templates")
	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(fs, "home.gohtml", "tailwind.gohtml"))))
	// tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	// if err != nil {
	// 	panic(err)
	// }
	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(fs, "contact.gohtml", "tailwind.gohtml"))))
	// tpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
	// if err != nil {
	// 	panic(err)
	// }
	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(fs, "faq.gohtml", "tailwind.gohtml"))))
	// cfg := models.DefaultPostgresConfig()
	// fmt.Println(cfg.String())
	// db, err := models.Open(cfg)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// err = models.MigrateFS(db, migrations.FS, ".")
	// if err != nil {
	// 	panic(err)
	// }

	// userService := models.UserService{
	// 	DB: db,
	// }
	// sessionService := models.SessionService{
	// 	DB: db,
	// }
	// usersC := controllers.Users{
	// 	UserService:    &userService,
	// 	SessionService: &sessionService,
	// }
	// usersC.Templates.New = views.Must(views.ParseFS(fs, "signup.gohtml", "tailwind.gohtml"))
	// usersC.Templates.SignIn = views.Must(views.ParseFS(fs, "signin.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Get("/forgot-pw",usersC.ForgotPassword)
	r.Post("/forgot-pw",usersC.ProcessForgotPassword)
	r.Route("/users/me",func(r chi.Router) {
		r.Use(umw.RequireUser)
		r.Get("/",usersC.CurrentUser)
	})
	// r.Get("/users/me", usersC.CurrentUser)
	r.Post("/signout", usersC.ProcessSignOut)
	// r.Get("/signup", controllers.FAQ(
	// 	views.Must(views.ParseFS(fs, "signup.gohtml", "tailwind.gohtml"))))
	r.Get("/contact/{user-id}", MyRequestHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	// fmt.Println("starting the server on :3000...")
	// umw := controllers.UserMiddlefware{
	// 	SessionService: &sessionService,
	// }

	// csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	// csrfMw := csrf.Protect([]byte(csrfKey), csrf.Secure(false))
	fmt.Println("starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}

// func TimerMiddleWare(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		h(w,r)
// 		fmt.Println("request time :",time.Since(start))
// 	}
// }
