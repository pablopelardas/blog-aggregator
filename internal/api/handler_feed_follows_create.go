package api

import (
	"encoding/json"
	"fmt"
	"internal/database"
	"internal/helpers"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cf *ApiConfig) handlerdFollowFeed(w http.ResponseWriter, r *http.Request, user database.User){
	defer r.Body.Close()
	type requestBody struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	type returnBody struct {
		Id uuid.UUID `json:"id"`
		FeedId uuid.UUID `json:"feed_id"`
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
	hasFeed := len(rBody.FeedId) > 0
	if !hasFeed {
		log.Printf("Feed ID required")
		helpers.RespondWithError(w, http.StatusBadRequest, "Feed ID required")
		return
	}
	id := uuid.New()
	createdAt := time.Now()
	updatedAt := time.Now()

	userID := uuid.NullUUID{UUID: user.ID, Valid: true}
	feedID := uuid.NullUUID{UUID: rBody.FeedId, Valid: true}
	// insert into database
	_, err = cf.DB.CreateFollow(r.Context(), database.CreateFollowParams{
		ID: id,
		FeedID: feedID,
		UserID: userID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
	if err != nil {
		log.Printf("Error creating follow %s", err)
		helpers.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating follow from user %s to feed %s", userID, feedID))
		return
	}

	// respond with id and cleaned body
	helpers.RespondWithJSON(w, http.StatusCreated, returnBody{
		Id: id,
		UserID: userID.UUID,
		FeedId: feedID.UUID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
}