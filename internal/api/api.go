package api

import (
	"database/sql"
	"internal/database"
	"log"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type ApiConfig struct{
	DB *database.Queries
}

func GetAPIConfig() *ApiConfig {
	db, err := sql.Open("postgres", os.Getenv("DB_CONN_STRING"))
	if err != nil {
		log.Fatal(err)
	}
	return &ApiConfig{
		DB: database.New(db),
	}
}

func NewRouter() *chi.Mux {
	apiConfig := GetAPIConfig()
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	router.Mount("/v1", v1Router(apiConfig))

	return router
}