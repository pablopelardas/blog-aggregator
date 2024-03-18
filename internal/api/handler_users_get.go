package api

import (
	"internal/database"
	"internal/helpers"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cf *ApiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User){
	type returnBody struct {
		Id uuid.UUID `json:"id"`
		Name string `json:"name"`
		ApiKey string
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	helpers.RespondWithJSON(w, http.StatusOK, returnBody{
		Id: user.ID,
		Name: user.Name,
		ApiKey: user.ApiKey,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}