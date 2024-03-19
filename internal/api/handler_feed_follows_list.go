package api

import (
	"internal/database"
	"internal/helpers"
	"net/http"

	"github.com/google/uuid"
)

func (cf *ApiConfig) handlerGetAllFeedFollows(w http.ResponseWriter, r *http.Request, user database.User){
	defer r.Body.Close()
	type returnBody struct {
		Following []database.UsersFeed `json:"follows"`
	}
	follows, err := cf.DB.GetFollowsByUser(r.Context(), uuid.NullUUID{UUID: user.ID, Valid: true})
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error reading follows")
		return
	}
	helpers.RespondWithJSON(w, http.StatusOK, returnBody{
		Following: follows,
	})


}