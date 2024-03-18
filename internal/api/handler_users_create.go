package api

import (
	"encoding/json"
	"internal/database"
	"internal/helpers"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cf *ApiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	type requestBody struct {
		Name string `json:"name"`
	}
	type returnBody struct {
		Id uuid.UUID `json:"id"`
		Name string `json:"name"`
		ApiKey string `json:"api_key"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`

	}
	dat, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body %s", err)
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error reading body")
		return
	}
	rBody := requestBody{}
	err = json.Unmarshal(dat, &rBody)
	if err != nil {
		log.Printf("Error unmarshalling JSON %s", err)
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error unmarshalling JSON")
		return	
	}
	hasName := len(rBody.Name) > 0

	if !hasName {
		log.Printf("Name required")
		helpers.RespondWithError(w, http.StatusBadRequest, "Name required")
		return
	}

	id := uuid.New()
	createdAt := time.Now()
	updatedAt := time.Now()
	// generate random sha256
	apiKey := helpers.GenerateRandomString(64)

	// insert into database
	_, err = cf.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: id,
		Name: rBody.Name,
		ApiKey: apiKey,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
	if err != nil {
		log.Printf("Error creating user %s", err)
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	// respond with id and cleaned body
	helpers.RespondWithJSON(w, http.StatusCreated, returnBody{
		Id: id,
		Name: rBody.Name,
		ApiKey: apiKey,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
}