package api

import (
	"fmt"
	"internal/database"
	"internal/helpers"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (cf *ApiConfig) handlerUnfollowFeed(w http.ResponseWriter, r *http.Request, user database.User){
	defer r.Body.Close()

	feedPathId := chi.URLParam(r, "id")
	if len(feedPathId) == 0 {
		log.Printf("Feed ID required")
		helpers.RespondWithError(w, http.StatusBadRequest, "Feed ID required")
		return
	}

	userID := uuid.NullUUID{UUID: user.ID, Valid: true}
	feedId := uuid.MustParse(feedPathId)

	err := cf.DB.DeleteFollow(r.Context(), database.DeleteFollowParams{
		FeedID: uuid.NullUUID{UUID: feedId, Valid: true},
		UserID: userID,
	})

	if err != nil {
		log.Printf("Error deleting follow %s", err)
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error deleting follow")
		return
	}

	// respond with id and cleaned body
	helpers.RespondWithJSON(w, http.StatusCreated, map[string]string{
		"message": fmt.Sprintf("Unfollowed feed %s", feedId),
	})
}