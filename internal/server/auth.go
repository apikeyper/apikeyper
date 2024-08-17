package server

import (
	"context"
	"keyify/internal/database"
	"keyify/internal/database/utils"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type Session struct {
	WorkspaceId uuid.UUID
}

func Auth(dbService database.Service, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Verify the request
		authHeader := r.Header["Authorization"]

		if len(authHeader) == 0 {
			encode(w, r, http.StatusUnauthorized, "Authorization header is missing")
			return
		}

		// split the header into parts
		authHeaderRaw := strings.Split(authHeader[0], " ")
		scheme := authHeaderRaw[0]
		if scheme != "Bearer" {
			encode(w, r, http.StatusUnauthorized, "Invalid Authorization header scheme")
			return
		}

		rootKey := authHeaderRaw[1]

		hashedRootKey := utils.HashString(rootKey)

		// Verify the rootKey exists
		rootKeyRow, err := dbService.FetchRootKey(hashedRootKey)

		if err != nil {
			encode(w, r, http.StatusUnauthorized, "Invalid root key")
			return
		}

		session := Session{
			WorkspaceId: rootKeyRow.WorkspaceId,
		}

		// Set values on request context
		r = r.WithContext(context.WithValue(r.Context(), "session", session))
		handler(w, r)
	}
}
