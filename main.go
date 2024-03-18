package main

import (
	"log"
	"net/http"
	"os"

	"internal/helpers"
	"internal/middlewares"

	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	port := os.Getenv("PORT")
	r := http.NewServeMux()
	corsMux := middlewares.Cors(r)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Hello, World!"})
	})

	// Start the server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}
	log.Print("Server started at port " + port)
	log.Fatal(srv.ListenAndServe())

}