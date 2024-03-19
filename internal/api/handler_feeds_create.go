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
		Feed database.Feed `json:"feed"`
		FeedFollow database.UsersFeed `json:"follow"`
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
	feedFollowId := uuid.New()
	// create feedFollow
	_, err = cf.DB.CreateFollow(r.Context(), database.CreateFollowParams{
		ID: feedFollowId,
		FeedID: uuid.NullUUID{UUID: id, Valid: true},
		UserID: userID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})

	// respond with id and cleaned body
	helpers.RespondWithJSON(w, http.StatusCreated, returnBody{
		Feed: database.Feed{
			ID: id,
			Name: rBody.Name,
			Url: rBody.Url,
			UserID: uuid.NullUUID{UUID: user.ID, Valid: true},
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
		FeedFollow: database.UsersFeed{
			ID: uuid.New(),
			UserID: uuid.NullUUID{UUID: user.ID, Valid: true},
			FeedID: uuid.NullUUID{UUID: feedFollowId, Valid: true},
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	})
}