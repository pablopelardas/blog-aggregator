package api

import (
	"internal/helpers"
	"net/http"

	"github.com/go-chi/chi"
)

func v1Router(cf *ApiConfig) *chi.Mux {
	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	v1Router.Get("/err", func(w http.ResponseWriter, r *http.Request) {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	})
	v1Router.Mount("/users", userRouter(cf))

	return v1Router
}