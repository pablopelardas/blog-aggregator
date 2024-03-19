package api

import (
	"internal/database"
	"internal/helpers"
	"net/http"

	"github.com/google/uuid"
)

func (cf *ApiConfig) handlerListPostsByUser(w http.ResponseWriter, r *http.Request, user database.User){
	defer r.Body.Close()
	type returnBody struct {
		Posts []database.Post `json:"posts"`
	}
	posts, err := cf.DB.ListPostsByUser(r.Context(), database.ListPostsByUserParams{
		UserID: uuid.NullUUID{UUID: user.ID, Valid: true},
		Limit: 100,
		Offset: 0,
	})
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error reading follows")
		return
	}
	helpers.RespondWithJSON(w, http.StatusOK, returnBody{
		Posts: posts,
	})


}