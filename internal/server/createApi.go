package server

import (
	"context"
	"keyify/internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (r *CreateApiRequest) Valid(ctx context.Context) (problems map[string]string) {
	problems = make(map[string]string)

	if r.ApiName == "" {
		problems["apiName"] = "apiName is required"
	}

	return
}

func (s *Server) CreateApiHandler(w http.ResponseWriter, r *http.Request) {
	// Get rootKey from context
	session := r.Context().Value("session").(Session)

	// Validate the request body
	decodedJson, problems, err := decodeValid[*CreateApiRequest](r)

	if err != nil {
		encode(w, r, http.StatusBadRequest, problems)
		return
	}

	var apiRow = &database.Api{
		ID:          uuid.New(),
		WorkspaceId: session.WorkspaceId,
		ApiName:     decodedJson.ApiName,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, createErr := s.Db.CreateApi(apiRow)

	if createErr != nil {
		encode(w, r, http.StatusInternalServerError, "Failed to create api")
		return
	}

	respBody := CreateApiResponse{
		ApiId: apiRow.ID,
	}

	encode(w, r, http.StatusCreated, respBody)
}
