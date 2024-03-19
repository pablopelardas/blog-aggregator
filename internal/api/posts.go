package api

import (
	"github.com/go-chi/chi"
)

func postsRouter(cf *ApiConfig) *chi.Mux {
	postRouter := chi.NewRouter()

	postRouter.Get("/", cf.middlewareAuth(cf.handlerListPostsByUser))

	return postRouter
}