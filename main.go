package main

import (
	"internal/api"
	"internal/scraper"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)



func main(){
	godotenv.Load()
	port := os.Getenv("PORT")
	go scraper.ScraperWorker()

	// Start the server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: api.NewRouter(),
	}
	log.Print("Server started at port " + port)
	log.Fatal(srv.ListenAndServe())

}