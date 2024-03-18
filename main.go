package main

import (
	"database/sql"
	"internal/helpers"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pablopelardas/blog-aggregator/internal/database"
)

type apiConfig struct{
	DB *database.Queries
}

func main(){
	godotenv.Load()
	port := os.Getenv("PORT")
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	apiConfig := &apiConfig{
		DB: dbQueries,
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	v1Router.Get("/err", func(w http.ResponseWriter, r *http.Request) {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	})

	router.Mount("/v1", v1Router)

	// Start the server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Print("Server started at port " + port)
	log.Fatal(srv.ListenAndServe())

}