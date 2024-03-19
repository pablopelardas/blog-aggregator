package api

import (
	"github.com/go-chi/chi"
)

func feedsRouter(cf *ApiConfig) *chi.Mux {
	feedsRouter := chi.NewRouter()

	feedsRouter.Post("/", cf.middlewareAuth(cf.handlerCreateFeed))
	feedsRouter.Get("/", cf.handlerListFeeds)

	return feedsRouter
}