package api

import (
	"github.com/go-chi/chi"
)

func feedFollowsRouter(cf *ApiConfig) *chi.Mux {
	feedFollows := chi.NewRouter()

	feedFollows.Post("/", cf.middlewareAuth(cf.handlerdFollowFeed))
	feedFollows.Delete("/{id}", cf.middlewareAuth(cf.handlerUnfollowFeed))

	return feedFollows
}