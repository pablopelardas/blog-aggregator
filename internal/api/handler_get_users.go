package api

import (
	"internal/helpers"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cf *ApiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request){
	apiKeyHeader := r.Header.Get("Authorization")
	if len(apiKeyHeader) == 0 {
		log.Printf("API Key required")
		helpers.RespondWithError(w, http.StatusUnauthorized, "API Key required")
		return
	}
	apiKey := apiKeyHeader[7:]
	user, err := cf.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		log.Printf("Error getting user %s", err)
		helpers.RespondWithError(w, http.StatusInternalServerError, "User not found")
		return
	}	
	if user.ID == uuid.Nil {
		log.Printf("User not found")
		helpers.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}
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