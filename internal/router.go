package router

import (
	"github.com/go-chi/chi"
)

func Init() (chi.Router, error) {
	r := chi.newRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.write([]byte("welcome"))
	})

	return r, nil
}
