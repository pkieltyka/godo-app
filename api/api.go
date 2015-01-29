package api

import (
	"net/http"

	"github.com/pkieltyka/godo-app/lib/ws"

	"github.com/pkieltyka/godo-app"
	"github.com/pkieltyka/godo-app/api/sessions"
	"github.com/pkieltyka/godo-app/api/todos"
	"github.com/pkieltyka/godo-app/api/users"

	"github.com/zenazn/goji/web/middleware"

	"github.com/zenazn/goji/web"
)

func New() http.Handler {
	r := web.New()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", http.RedirectHandler("/app/", 301))
	r.Get("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ws.Text(w, 200, ".")
	}))

	r.Post("/signup", users.SignUp)
	r.Post("/login", sessions.Login)
	r.Get("/login", sessions.Login) // for demo purposes

	// Static files
	r.Handle("/app/*", http.StripPrefix("/app/", http.FileServer(http.Dir(godo.App.Config.Webapp.Path))))

	// Authed router
	a := web.New()
	a.Use(middleware.SubRouter)
	a.Use(godo.App.TokenAuth.Handler)
	a.Use(sessions.UserContext)

	a.Get("/todos", todos.Index)
	a.Post("/todos", todos.Create)
	a.Get("/todos/:id", todos.Ctx.On(todos.Read))
	a.Put("/todos/:id", todos.Ctx.On(todos.Update))
	a.Delete("/todos/:id", todos.Ctx.On(todos.Delete))

	r.Handle("/*", a)

	return r
}
