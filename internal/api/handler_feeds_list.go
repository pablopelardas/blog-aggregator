package api

import (
	"internal/database"
	"internal/helpers"
	"net/http"
)

func (cf *ApiConfig) handlerListFeeds(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	type returnBody struct {
		Feeds []database.Feed `json:"feeds"`
	}

	feeds, err := cf.DB.ListFeeds(r.Context())

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error reading feeds")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, returnBody{
		Feeds: feeds,
	})

}