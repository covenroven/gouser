package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/covenroven/gouser/internal/api"
)

func Init() (chi.Router, error) {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/users", func(r chi.Router) {
	    r.Get("/", api.IndexUsers);
	    r.Post("/", api.StoreUser);
	    r.Get("/{userID}", api.ShowUser);
	    // r.Put("/{userID}", api.ShowUser);
	    // r.Delete("/{userID}", api.ShowUser);
	})

	return r, nil
}
