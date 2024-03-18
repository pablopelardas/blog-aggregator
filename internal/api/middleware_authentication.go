package api

import (
	"internal/database"
	"internal/helpers"
	"net/http"

	"github.com/google/uuid"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cf *ApiConfig) middlewareAuth(next authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKeyHeader := r.Header.Get("Authorization")
		if len(apiKeyHeader) == 0 {
			helpers.RespondWithError(w, http.StatusUnauthorized, "API Key required")
			return
		}
		apiKey := apiKeyHeader[7:]
		user, err := cf.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "User not found")
			return
		}
		if user.ID == uuid.Nil {
			helpers.RespondWithError(w, http.StatusNotFound, "User not found")
			return
		}
		next(w, r, user)
	}
}