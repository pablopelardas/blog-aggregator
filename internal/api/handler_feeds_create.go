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

func (cf *ApiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User){
	defer r.Body.Close()
	type requestBody struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}
	type returnBody struct {
		Id uuid.UUID `json:"id"`
		Name string `json:"name"`
		Url string `json:"url"`
		UserID uuid.UUID `json:"user_id"`
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
	hasUrl := len(rBody.Url) > 0
	if !hasName || !hasUrl {
		log.Printf("Name and URL required")
		helpers.RespondWithError(w, http.StatusBadRequest, "Name and URL required")
		return
	}
	id := uuid.New()
	createdAt := time.Now()
	updatedAt := time.Now()

	userID := uuid.NullUUID{UUID: user.ID, Valid: true}
	// insert into database
	_, err = cf.DB.CrateFeed(r.Context(), database.CrateFeedParams{
		ID: id,
		Name: rBody.Name,
		Url: rBody.Url,
		UserID: userID,
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
		Url: rBody.Url,
		UserID: userID.UUID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
}