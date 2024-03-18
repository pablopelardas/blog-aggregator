package api

import (
	"github.com/go-chi/chi"
)

func userRouter(cf *ApiConfig) *chi.Mux {
	userRouter := chi.NewRouter()

	userRouter.Get("/", cf.handlerGetUser)
	userRouter.Post("/", cf.handlerCreateUser)

	return userRouter
}